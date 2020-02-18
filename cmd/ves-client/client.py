import urllib.parse
import base64
import requests
from service_code import *


class Response(object):
    def __init__(self, response):
        """
        :param response:
        :type response requests.Response
        """
        self.resp = response
        self.avail = True
        if self.resp.status_code != 200:
            self.avail = False
        if self.resp.headers['Content-Type'].startswith('application/json'):
            self.body = self.resp.json()
            self.code = self.body.get('code')
            if self.code != 0:
                self.avail = False

    def fail_string(self):
        return f'<code:{self.code}, err:{self.get_error()}, status_code:{self.resp.status_code}>'

    def get_error(self):
        return self.body.get("error")

    def to_error(self):
        return response_to_error(self)

    def maybe_raise(self):
        if not self.avail:
            raise self.to_error()


def process_response(req_func):
    def wrap(*args, **kwargs):
        return Response(req_func(*args, **kwargs))

    return wrap


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

    @staticmethod
    def encode_address(src):
        if isinstance(src, str):
            src = bytes.fromhex(src)
        if isinstance(src, bytes):
            # print('|', base64.encodebytes(src).decode()[:-1], '|')
            return base64.encodebytes(src).decode()[:-1]
        raise TypeError(f'encode address error: {type(src)}')

    @process_response
    def get(self, url, *args, **kwargs):
        """
        :return: :class:`Response <Response>` object
        :rtype: Response
        """
        return requests.get(self.parse_url(url), *args, **self.parse_kwargs(kwargs))

    @process_response
    def post(self, url, *args, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: Response
        """
        return requests.post(self.parse_url(url), *args, **self.parse_kwargs(kwargs))

    @process_response
    def delete(self, url, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: Response
        """
        return requests.delete(self.parse_url(url), **self.parse_kwargs(kwargs))

    @process_response
    def put(self, url, *args, **kwargs):
        """
        :param url: {string}
        :return: :class:`Response <Response>` object
        :rtype: Response
        """
        return requests.put(self.parse_url(url), *args, **self.parse_kwargs(kwargs))
