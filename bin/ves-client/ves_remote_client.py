import json

from client import Client


class VESRemoteClient(Client):
    def __init__(self, host):
        super().__init__(host)

    def ping(self):
        return self.get('ping')

    def send_op_intents(self, file_path=None, intents=None, dependencies=None):
        if file_path is not None:
            return self.send_op_intents_in_file(file_path)
        intents = intents or []
        dependencies = dependencies or []
        return self.post('/v1/session', json={
            'intents': intents,
            'dependencies': dependencies,
        })

    def send_op_intents_in_file(self, file_path):
        _ = self.post
        with open(file_path) as intents_file:
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
