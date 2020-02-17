import yaml
import json


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
    def __init__(self, name, password, accounts):
        """
        :param name:
        :param password:
        :param accounts:
        """
        self.name = name
        self.password = password
        if isinstance(accounts, list):
            self.accounts = accounts
        elif isinstance(accounts, str):
            self.accounts = _load(open(accounts))


class Account(object):
    def __init__(self, domain, user_name):
        self.domain = domain
        self.user_name = user_name

    def __hash__(self):
        return hash(str(self.domain) + "<84f4446f>" + self.user_name)

    def __eq__(self, other):
        return self.domain == other.domain and self.user_name == other.user_name


class OpIntent(object):
    def __init__(self, intent):
        self.name = intent.get('name')
        self.op_type = intent.get('op_type', 'Payment')
        if self.op_type == 'Payment':
            self.src = Account(**intent['src'])
            self.dst = Account(**intent['dst'])
            self.resp_accounts = [
                self.src, self.dst,
            ]
        elif self.op_type == 'ContractInvocation':
            self.src = Account(**intent.get('invoker'))
            self.resp_accounts = [
                self.src
            ]


class OpIntents(object):
    def __init__(self, d):
        """
        :param d:
        :type d dict
        """
        op_intents = d['op-intents']
        self.dependencies = d.get('dependencies', [])

        self.intents = []
        if not isinstance(op_intents, list):
            op_intents = [op_intents]
        for op_intent in op_intents:
            self.intents.append(OpIntent(op_intent))


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
    accounts<array>[].address: {hex string} this account's private address

    accounts[string]: the target file (in yaml) describe the resp's accounts
    """

    def __init__(self, stream=None, file_path='playbook.yaml'):
        """
        :param stream:
        :param file_path:
        :type file_path: str
        :type stream TextIO
        """
        if file_path is not None:
            stream = stream or open(file_path)
        if stream is None:
            raise ValueError('stream is None')
        obj = _load(stream)
        self.name = obj.get('name', '<none>')
        self.intents = self.prepare_intent_file(obj.get('source', 'intent.json'))
        self.roles = self.process_roles(obj.get('ves-clients', []))
        self.build_role_map()

    @staticmethod
    def process_roles(roles):
        parsed_roles = []
        for role in roles:
            parsed_roles.append(Role(
                name=role['name'],
                password=role['password'],
                accounts=role['accounts'],
            ))
        return parsed_roles

    @staticmethod
    def prepare_intent_file(intent_path):
        with open(intent_path) as f:
            intents = json.load(f)
        return OpIntents(intents)

    def build_role_map(self):
        for role in self.roles:
            for account in role.accounts:
                pass
                # print(account)


def run_playbook(playbook: Playbook):
    resp_accounts = set()
    for intents in playbook.intents.intents:
        for account in intents.resp_accounts:
            resp_accounts.add(account)
    # print(resp_accounts)


if __name__ == '__main__':
    pb = Playbook(file_path='playbook.example.yaml')
    run_playbook(pb)
