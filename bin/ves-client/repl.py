#!/usr/bin/python3
import code
import copy
import readline
import rlcompleter
import urllib.parse

import requests

_ = rlcompleter


def create_local():
    t = copy.copy(locals())
    t.update(globals())
    return t


# tab completion
readline.parse_and_bind('tab: complete')


def interact(*args, **kwargs):
    kwargs['local'] = kwargs.get('local', create_local())
    code.interact(*args, **kwargs)


class VESRemoteClient:
    def __init__(self, host='39.10.145.91:26670'):
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

    def ping(self):
        return self.get('ping')

    def list_keys(self):
        return self.get('v1/accounts')


if __name__ == '__main__':
    interact(banner='ves client console',
             exitmsg='exiting ves client console...'
             )
