#!/usr/bin/python3

from cves_client import CVESClient
from decorator import print_response
from console import Console
from interact import interact
from wrapper import is_wrap_error, unwrap, Frame
from fs import FS


class Utils:
    unwrap = staticmethod(unwrap)
    is_wrap_error = staticmethod(is_wrap_error)
    frame = Frame
    fs = FS()


def t():
    # return ves.cli.send_op_intents('intent.json')
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
    utils = Utils()
    desc = ConsoleDesc()
    main_load_cfg()

    print('entering ves console...')
    interact(banner=banner,
             exitmsg='exiting ves client console...'
             )
