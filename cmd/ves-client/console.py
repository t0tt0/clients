import getpass

from cves_client import CVESClient
from ves_remote_client import VESRemoteClient
from decorator import wrap_response


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
    def send_op_intents(self, file_path=None, intents=None, dependencies=None):
        response = self.cli.send_op_intents(file_path, intents, dependencies)
        return response

    @wrap_response
    def send_op_intents_in_file(self, file_path=None):
        response = self.cli.send_op_intents_in_file(file_path)
        return response

    @wrap_response
    def post_chain_info(self, chain_id, address, user_id=None):
        response = self.c_ves.post_chain_info(chain_id, address, user_id)
        return response

    @wrap_response
    def get_chain_info(self, cid=None, chain_id=None, address=None, user_id=None):
        response = self.c_ves.get_chain_info(cid, chain_id, address, user_id)
        return response

    # @wrap_response
    # def put_chain_info(self, chain_id=None, address=None, user_id=None):
    #     response = self.c_ves.put_chain_info(chain_id, address, user_id)
    #     return response

    @wrap_response
    def delete_chain_info(self, cid=None, chain_id=None, address=None, user_id=None):
        response = self.c_ves.delete_chain_info(cid, chain_id, address, user_id)
        return response

    @wrap_response
    def list_chain_info(self, chain_id=None, address=None, user_id=None, page=1, page_size=10):
        response = self.c_ves.list_chain_info(chain_id=chain_id, address=address, user_id=user_id, page=page,
                                              page_size=page_size)
        return response

    def list_user(self):
        # todo
        pass

    def switch(self, name=None, password=None):
        # todo
        pass

    def message_to(self, msg, target_name=None, target_id=None):
        # todo
        pass
