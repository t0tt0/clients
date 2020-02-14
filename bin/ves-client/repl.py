#!/usr/bin/python3
import code
import copy
import getpass
import json
import readline
import rlcompleter
import urllib.parse

import requests


def t():
    return ves.c_ves.send_op_intents('intent.json')


def create_local():
    t = copy.copy(locals())
    t.update(globals())
    return t


_ = rlcompleter
# tab completion
readline.parse_and_bind('tab: complete')


def interact(*args, **kwargs):
    kwargs['local'] = kwargs.get('local', create_local())
    code.interact(*args, **kwargs)


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
                self.identities = data['identity']
                self.token = data['token']
                self.refresh_token = data['refresh_token']
        return response


class VESRemoteClient(Client):
    def __init__(self, host):
        super().__init__(host)

    def send_op_intents(self, filepath=None, intents=None, dependencies=None):
        if filepath is not None:
            return self.send_op_intents_in_file(filepath)
        intents = intents or []
        dependencies = dependencies or []
        return self.post('/v1/session', {
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
    def __init__(self, host='127.0.0.1:23336'):
        self.c_ves = CVESClient(host)
        self.cli = VESRemoteClient(host)

    @wrap_response
    def ping(self):
        response = self.c_ves.ping()
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
