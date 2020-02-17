import urllib.parse

import requests


class Client(object):
    def __init__(self, host):
        self.host = host
        self.identities = []
        self.token = None
        if not (self.host.startswith('http://') or self.host.startswith('https://')):
            self.host = 'http://' + host

    def parse_url(self, path):
        """
        :return: :string:`str` object, with base self.host
        :rtype: str
        """
        return urllib.parse.urljoin(self.host, path)

    def parse_kwargs(self, kwargs):
        if self.token is not None:
            headers = kwargs.get('headers', dict())
            headers['Authorization'] = 'Bearer ' + self.token
            kwargs['headers'] = headers
        return kwargs

    def get(self, url, *args, **kwargs):
        """
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.get(self.parse_url(url), *args, **self.parse_kwargs(kwargs))

    def post(self, url, *args, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.post(self.parse_url(url), *args, **self.parse_kwargs(kwargs))

    def delete(self, url, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.delete(self.parse_url(url), **self.parse_kwargs(kwargs))

    def put(self, url, *args, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: requests.Response
        """
        return requests.put(self.parse_url(url), *args, **self.parse_kwargs(kwargs))
