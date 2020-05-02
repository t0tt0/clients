import yaml
import json
import os.path
import logging

from service_code import Code
from cves_client import CVESClient
from ves_local_client import VESLocalClient
from ves_remote_client import VESRemoteClient


def _load_all(stream):
    """
    :param stream:
    :return: a generator return objects serialized in yaml
    :rtype: object
    """
    return yaml.load_all(stream)


def _load(stream) -> dict:
    """
    :param stream:
    :return: a object serialized in yaml
    :rtype: object
    """
    return yaml.load(stream)


class Role(object):
    def __init__(self, name, password, accounts, central_ves: CVESClient = None):
        """
        :param name:
        :param password:
        :param accounts:
        """

        self.chain_map = dict()

        self.client = None
        self.name = name
        if not isinstance(password, str):
            raise TypeError(type(password))

        self.password = password
        if isinstance(accounts, list):
            self.accounts = accounts
        elif isinstance(accounts, str):
            with open(accounts) as f:
                self.accounts = _load(f)

        # check accounts
        if not isinstance(self.accounts, list):
            raise TypeError(f'self.accounts is not good type: {type(self.accounts)}')

        for account in self.accounts:
            self.chain_map[account['chain_id']] = account['address']

        self.central_ves = CVESClient(central_ves)
        self.process: VESLocalClient
        self.process = None
        self.log_file = open(f'client.{self.name}.out', 'w')

    def try_login(self):
        r = self.central_ves.register(self.name, self.password)
        if r.code != Code.InsertError.value and r.code != Code.OK.value:
            raise r.to_error()
        self.central_ves.login(self.name, self.password).maybe_raise()

    def try_start_local_client(self):

        self.process = VESLocalClient.from_role(self, log_file=self.log_file)
        self.client = VESRemoteClient(self.process.host)
        r = self.client.try_ping()
        if r:
            print(r.body)
        else:
            raise ConnectionError('ping failed')

    def try_register_account(self, account):
        if self.central_ves.id is None:
            raise ValueError(f'{self.name}.central_ves.id is None')
        r = self.central_ves.post_chain_info(account['chain_id'], account['address'])

        if r.code != Code.InsertError.value and r.code != Code.DuplicatePrimaryKey.value \
                and r.code != Code.OK.value:
            raise r.to_error()

        r = self.client.post_account(
            chain_type=account['chain_type'], alias=account['alias'], chain_id=account['chain_id'],
            address=account['address'], addition=account.get('private_address'))

        if r.code != Code.InsertError.value and r.code != Code.DuplicatePrimaryKey.value \
                and r.code != Code.OK.value:
            raise r.to_error()

    def try_close(self):
        if self.log_file is not None:
            self.log_file.close()
        if self.process is not None:
            self.process.kill()


class Account(object):
    # noinspection PyPep8Naming
    def __init__(self, domain, user_name, **kwargs):
        self.domain = domain
        self.user_name = user_name

    def __hash__(self):
        return hash(str(self.domain) + "<84f4446f>" + self.user_name)

    def __eq__(self, other):
        return self.domain == other.domain and self.user_name == other.user_name


class NullPlaybook(object):
    @staticmethod
    def parse_account(acc):
        return acc


def parse_obj_to_intents(op_intents: list or dict, context):
    intents = []
    if not isinstance(op_intents, list):
        op_intents = [op_intents]
    for op_intent in op_intents:
        intents.append(OpIntent(op_intent, context))
    return intents


class OpIntent(object):
    def __init__(self, intent, context):
        self.name = intent.get('name')
        self.type = intent.get('type', 'Payment')
        self.resp_accounts = []
        if self.type == 'Payment':
            self.src = Account(**context.parse_account(intent['src']))
            self.dst = Account(**context.parse_account(intent['dst']))
            self.resp_accounts = [
                self.src, self.dst,
            ]
        elif self.type == 'ContractInvocation':
            self.src = Account(**context.parse_account(intent.get('invoker')))
            # self.contract = Contract(**context.parse_account(intent.get('contract')))
            self.resp_accounts = [
                self.src
            ]
        elif self.type == 'IfStatement':
            self.if_body = parse_obj_to_intents(intent.get('if', []), context)
            self.else_body = parse_obj_to_intents(intent.get('else', []), context)
            self.condition = intent.get('condition')
        elif self.type == 'loopFunction':
            self.body = parse_obj_to_intents(intent.get('loop', []), context)
            self.times = intent.get('loopTime', 0)
        else:
            logging.warning('unknown type ', self.type)


class OpIntents(object):
    def __init__(self, d):
        """
        :param d:
        :type d dict
        """
        op_intents = d['op-intents']
        self.dependencies = d.get('dependencies', [])
        self.local_accounts = self.build_local_map(d.get('accounts', []), 'userName', 'user_name')
        self.local_contracts = self.build_local_map(d.get('contracts', []), 'contractName', 'user_name')
        self.intents = parse_obj_to_intents(op_intents, self)

    def parse_account(self, acc):
        if isinstance(acc, str):
            xmp = self.local_accounts.get(acc)
            k = None
            if xmp:
                k = xmp.get(None)
            if k is None:
                raise KeyError('accountName %s is not exists or ambiguous' % acc)
            return k
        return acc

    @staticmethod
    def build_local_map(raw_list, k, rename=None):
        mp = dict()
        info:dict
        for info in raw_list:
            un = info[k]
            if rename is not None:
                info.pop(k)
                info[rename] = un
            xmp = mp.get(un, {
                info['domain']: info,
                None: info,
            })
            if un in mp:
                xmp[None] = None
            mp[un] = xmp

        return mp


class Playbook(object):
    """
    :type intents OpIntents

    good form of Playbook:

    root.name: {string} name of playbook
    root.source: {string} file path of op-intent file
    root.ves-clients: {array} describe ves-clients
    root.accounts: {string} describe resp accounts
    ves-clients.name: {string} open with this name
    ves-clients.password: {string} login central ves by this passphrase
    ves-clients.accounts: {array|string} describe resp accounts
    accounts<array>[].chain_id: {int} which chain the resp account is at
    accounts<array>[].address: {hex string} this account's public address

    accounts[string]: the target file (in yaml) describe the resp's accounts
    """

    def __init__(self, stream=None,
                 file_path='playbook.yaml',
                 central_ves_host=None):
        """
        :type central_ves_host: str
        :param stream:
        :param file_path:
        :type file_path: str
        :type stream TextIO
        """

        self.role_map = dict()
        self.base_path = os.path.abspath(os.path.dirname(file_path))

        if file_path is not None:
            stream = stream or open(file_path)
        if stream is None:
            raise ValueError('stream is None')
        obj = _load(stream)
        self.name = obj.get('name', '<none>')
        self.intent_file_path = self.rel_path(obj.get('source', 'intent.json'))
        self.intents = self.prepare_intent_file(self.intent_file_path)

        self.central_ves = CVESClient(central_ves_host)
        r = self.central_ves.ping()
        if not r.avail:
            raise ConnectionError(f'ping failed: {r.fail_string()}')
        self.roles = self.process_roles(obj.get('ves-clients', []))
        self.build_role_map()

    def process_roles(self, roles):
        parsed_roles = []
        for role in roles:
            parsed_roles.append(Role(
                name=str(role['name']),
                password=str(role['password']),
                accounts=self.rel_path(role['accounts']),
                central_ves=self.central_ves,
            ))
        return parsed_roles

    @staticmethod
    def prepare_intent_file(intent_path):
        with open(intent_path) as f:
            intents = json.load(f)
        return OpIntents(intents)

    def build_role_map(self):
        for role in self.roles:
            role.try_start_local_client()
            role.try_login()
            for account in role.accounts:
                role.try_register_account(account)
                print(account)
                self.role_map[role.name] = role

    def close(self):
        for role in self.roles:
            role.try_close()

    @staticmethod
    def parse_account(acc):
        return acc

    def rel_path(self, fp):
        if isinstance(fp, str):
            return os.path.join(self.base_path, fp)
        return fp


def check_and_get_role(playbook: Playbook, intents: list = None) -> Role or None:
    intents = intents or playbook.intents.intents
    some_role = None
    for intent in intents:
        if intent.type == 'IfStatement':
            some_role = check_and_get_role(playbook, intent.if_body) or some_role
            some_role = check_and_get_role(playbook, intent.else_body) or some_role
        elif intent.type == 'loopFunction':
            some_role = check_and_get_role(playbook, intent.body) or some_role

        for account in intent.resp_accounts:
            if account.domain not in playbook.role_map[account.user_name].chain_map:
                raise ValueError(f'<{account.domain}, {account.user_name}> is not in playbook')
            some_role = playbook.role_map[account.user_name]
    return some_role


import threading

# class IOThread(threading.Thread):
#     def __init__(self, role: Role):
#         super().__init__()
#         self.role = role
#
#     def run(self):
#         while True:
#             self.role.process.stdout.writeline(
#                 self.role.process.process.stdout.readline())

def run_playbook(playbook: Playbook):
    some_role: Role = check_and_get_role(playbook)
    if some_role is not None:
        threads = []

        # for role in playbook.roles:
        #     threads.append(IOThread(role))

        # for thread in threads:
        #     thread.start()

        r = some_role.client.send_op_intents_in_file(playbook.intent_file_path)
        print(r, type(r), getattr(r, 'code', None),
              getattr(getattr(r, 'resp', {}), 'status_code'), getattr(r, 'body', None))

        # for thread in threads:
        #     thread.join()
    else:
        raise ValueError('some_role not found')
    # print(resp_accounts)


if __name__ == '__main__':
    pb = Playbook(file_path='./example/playbook.example.yaml')
    run_playbook(pb)
    input('enter any keys to exit')
    pb.close()
