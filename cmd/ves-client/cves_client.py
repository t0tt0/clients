# import base64
# from service_code import Code
from client import Client


# '39.10.145.91:26670'


class CVESClient(Client):
    def __init__(self, host=None):
        if isinstance(host, CVESClient):
            host = host.host

        super().__init__(host or '127.0.0.1:23336')
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
        r = self.post('/v1/login', json={
            'name': name,
            'password': password,
        })
        if r.avail:
            data = r.body
            self.id = data['id']
            self.identities = data['identity']
            self.token = data['token']
            self.refresh_token = data['refresh_token']
        return r

    def post_chain_info(self, chain_id, address, user_id=None):
        user_id = user_id or self.id
        if not isinstance(chain_id, int) or chain_id < 0:
            raise TypeError(f'chain_id type error: {type(chain_id)}')
        response = self.post('/v1/chain_info', json={
            'user_id': user_id,
            'chain_id': chain_id,
            'address': self.encode_address(address),
        })
        return response

    def get_chain_info(self, cid=None, chain_id=None, address=None, user_id=None):
        response = self.get(f'/v1/chain_info/{cid}')
        _ = chain_id
        _ = address
        _ = user_id
        return response

    def put_chain_info(self, chain_id=None, address=None, user_id=None):
        pass

    def delete_chain_info(self, cid=None, chain_id=None, address=None, user_id=None):
        response = self.delete(f'/v1/chain_info/{cid}')
        _ = chain_id
        _ = address
        _ = user_id
        return response

    def list_chain_info(self, chain_id=None, address=None, user_id=None, page=1, page_size=10):
        response = self.get('/v1/chain_info-list', params={
            'page': page,
            'page_size': page_size,
        })
        _ = chain_id
        _ = address
        _ = user_id
        return response

    def list_user(self):
        # todo
        pass
