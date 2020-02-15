#!/usr/bin/python3
import code
import copy
import getpass
import json
import os.path
import readline
import rlcompleter
import urllib.parse

import requests


def t():
    # return ves.cli.send_op_intents('intent.json')
    pass


# interact #####################################

def create_local():
    _local = copy.copy(locals())
    _local.update(globals())
    return _local


_ = rlcompleter
# tab completion
readline.parse_and_bind('tab: complete')


def interact(*args, **kwargs):
    kwargs['local'] = kwargs.get('local', create_local())
    code.interact(*args, **kwargs)


# unwrap #######################################

magic = '<84f4446f>'


def is_wrap_error(e):
    """
    :param e: {string}
    :return:
    """
    return e.startswith(magic)


class FileLine:

    def __init__(self, file='', line=-1):
        self.file = file
        self.line = line

    def __str__(self):
        return f"{self.file}:{self.line}"

    def rel(self, scope):
        return f"{scope.rel_path(self.file)}:{self.line}"

    @staticmethod
    def unwrap(e):
        c = e.rsplit(':', 1)
        if len(c) < 2:
            return FileLine()
        return FileLine(c[0], int(c[1]))


class Func:
    def __init__(self, name='', fileline=None):
        self.name = name
        self.fileline = fileline

    def __str__(self):
        return f"<name:{self.name},fileline:{self.fileline}>"

    def rel(self, scope):
        return f"<name:{scope.rel_pac(self.name)},fileline:{self.fileline.rel(scope)}>"

    @staticmethod
    def unwrap(e):
        if len(e) >= 2 and e[0] == '<' and e[-1] == '>':
            e = e[1:-1]
        else:
            return Func()
        c = e.split(',', 1)
        return Func(c[0], FileLine.unwrap(c[1]))


class StackPos:
    def __init__(self, fn=None, fileline=None):
        self.fn = fn
        self.fileline = fileline

    def __str__(self):
        return f"<fn:{self.fn},fileline:{self.fileline}>"

    def rel(self, scope):
        return f"<fn:{self.fn.rel(scope)},fileline:{self.fileline.rel(scope)}>"

    @staticmethod
    def unwrap(e):
        if len(e) >= 2 and e[0] == '<' and e[-1] == '>':
            e = e[1:-1]
        else:
            return StackPos()
        if len(e) > 0 and e[0] == '!':
            return StackPos()
        c = e.split('>', 1)
        return StackPos(Func.unwrap(c[0] + '>'), FileLine.unwrap(c[1].lstrip(',')))


class Frame(object):
    def __init__(self, pos, _code, err):
        self.pos = pos
        self.code = _code
        self.err = err

    def __str__(self):
        return f"<pos:{self.pos},code:{self.code},err:{self.err}>"

    def rel(self, scope):
        return f"<pos:{self.pos.rel(scope)},code:{self.code},err:{self.err}>"

    @staticmethod
    def unwrap(e):
        if not is_wrap_error(e):
            return e, False
        c = e.split(magic, 3)
        if len(c) < 4:
            return e, False
        pos, _code, err = \
            StackPos.unwrap(c[1][4:-1]), \
            int(c[2][5:-1]), c[3][4:-1]
        return Frame(pos, _code, err), True


class Scope(object):
    def __init__(self, pn='', wd=''):
        self.pn = pn
        self.wd = wd

    def rel_pac(self, pac):
        if pac.startswith(self.pn):
            return pac[len(self.pn):].lstrip('.')
        return pac

    def rel_path(self, path):
        print(path)
        return os.path.relpath(path, self.wd)


def unwrap(e):
    e, ok = Frame.unwrap(e)
    if ok:
        n, ok = Frame.unwrap(e.err)
        if ok:
            e.err = ''
            return [e] + n
        return [e]
    return []


# client #####################################


class Client(object):
    def __init__(self, host):
        self.host = host
        if not (self.host.startswith('http://') or self.host.startswith('https://')):
            self.host = 'http://' + host

    def parse_url(self, path):
        """
        :return: :string:`str` object, with base self.host
        :rtype: str
        """
        return urllib.parse.urljoin(self.host, path)

    def get(self, url, *args, **kwargs):
        """
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.get(self.parse_url(url), *args, **kwargs)

    def post(self, url, *args, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.post(self.parse_url(url), *args, **kwargs)


# '39.10.145.91:26670'
class CVESClient(Client):
    def __init__(self, host='127.0.0.1:23336'):
        super().__init__(host)
        self.identities = []
        self.token = None
        self.id = None
        self.refresh_token = None

    def ping(self):
        return self.get('ping')

    def register(self, name, password):
        return self.post('/v1/user', json={
            'name': name,
            'password': password,
        })

    def login(self, name, password):
        response = self.post('/v1/login', json={
            'name': name,
            'password': password,
        })
        if response.status_code == 200:
            data = response.json()
            if data['code'] == 0:
                self.id = data['id']
                self.identities = data['identity']
                self.token = data['token']
                self.refresh_token = data['refresh_token']
        return response

    def post_chain_info(self, chain_id, address, user_id=None):
        user_id = user_id or self.id
        # todo
        pass

    def get_chain_info(self, chain_id=None, address=None, user_id=None):
        # todo
        pass

    def put_chain_info(self, chain_id=None, address=None, user_id=None):
        # todo
        pass

    def delete_chain_info(self, chain_id=None, address=None, user_id=None):
        # todo
        pass

    def list_chain_info(self, chain_id=None, address=None, uesr_id=None):
        # todo
        pass

    def list_user(self):
        # todo
        pass


class VESRemoteClient(Client):
    def __init__(self, host):
        super().__init__(host)

    def ping(self):
        return self.get('ping')

    def send_op_intents(self, filepath=None, intents=None, dependencies=None):
        if filepath is not None:
            return self.send_op_intents_in_file(filepath)
        intents = intents or []
        dependencies = dependencies or []
        return self.post('/v1/session', json={
            'intents': intents,
            'dependencies': dependencies,
        })

    def send_op_intents_in_file(self, filepath):
        _ = self.post
        with open(filepath) as intents_file:
            intents = json.load(intents_file)

        op_intents = intents.get('op-intents', [])
        if not isinstance(op_intents, list):
            raise TypeError(f'field op-intents require list type, {type(op_intents)} got')
        op_intents = list(map(json.dumps, op_intents))

        dependencies = intents.get('dependencies', [])
        if not isinstance(dependencies, list):
            raise TypeError(f'field dependencies require list type, {type(dependencies)} got')
        dependencies = list(map(json.dumps, dependencies))

        return self.send_op_intents(intents=op_intents, dependencies=dependencies)

    def list_keys(self):
        return self.get('v1/accounts')


def print_response(response):
    if response is not None:
        print('status  :', response.status_code,
              '\nresponse:', response.json(), '\n')
    return response


def wrap_response(req_func):
    def wrap(*args, **kwargs):
        return print_response(req_func(*args, **kwargs))

    return wrap


class Console(object):
    def __init__(self, c_ves_host='127.0.0.1:23336', cli_host='localhost:26670'):
        self.c_ves = CVESClient(c_ves_host)
        self.cli = VESRemoteClient(cli_host)

    @wrap_response
    def ping_c_ves(self):
        response = self.c_ves.ping()
        return response

    @wrap_response
    def ping_cli(self):
        response = self.cli.ping()
        return response

    @wrap_response
    def register(self, name=None, password=None):
        if name is None:
            name = input("user_name:")
        if password is None:
            password = getpass.getpass('password:')
            retry = getpass.getpass('re enter:')
            if password != retry:
                print('inconsistent password entered')
                return
        response = self.c_ves.register(name, password)
        return response

    @wrap_response
    def login(self, name=None, password=None):
        if name is None:
            name = input("user_name:")
        if password is None:
            password = getpass.getpass('password:')
        response = self.c_ves.login(name, password)
        return response

    @wrap_response
    def send_op_intents(self, filepath=None, intents=None, dependencies=None):
        response = self.cli.send_op_intents(filepath, intents, dependencies)
        return response

    @wrap_response
    def send_op_intents_in_file(self, filepath=None):
        response = self.cli.send_op_intents_in_file(filepath)
        return response

    def switch(self, name=None, password=None):
        # todo
        pass

    def message_to(self, msg, target_name=None, target_id=None):
        # todo
        pass


def main_load_cfg():
    import os.path
    cfg_path = './config.py'
    cfg = {}
    if os.path.exists(cfg_path):
        with open(cfg_path) as f:
            exec(f.read(), None, cfg)
    user = cfg.get('user', [])
    if not isinstance(user, list):
        user = list(user.items())
    else:
        user = list(map(lambda u: list(u.items())[0], user))
    if len(user) > 0:
        default_user = user[0]
    else:
        default_user = None
    if default_user is not None:
        print('config: use "%s" as default user' % (default_user[0]))
        response = ves.c_ves.login(default_user[0], default_user[1])
        if response.status_code == 200 and response.json()['code'] == 0:
            print('config: login as "%s" successfully' % (default_user[0]))
        else:
            print_response(response)


if __name__ == '__main__':
    class ConsoleDesc(object):
        def __init__(self, _='', ):
            self.content = f"""ves: {Console}
ves.c_ves: {CVESClient}"""

        def __str__(self):
            return self.content

        def __repr__(self):
            return self.content

        def __call__(self, obj):
            return help(obj)


    import sys

    banner = '=' * 30 + '   ves client console   ' + '=' * 30
    banner += f"""
Python {sys.version} on {sys.platform}
Type "desc" for more information"""

    ves = Console()
    desc = ConsoleDesc()
    main_load_cfg()

    x = t()

    print('entering ves console...')
    interact(banner=banner,
             exitmsg='exiting ves client console...'
             )
