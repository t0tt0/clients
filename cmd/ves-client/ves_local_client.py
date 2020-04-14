import subprocess
import random


class VESLocalClient:
    def __init__(self, process: subprocess.Popen, port):
        self.process = process
        self.port = port
        self.host = "localhost:" + port

    def terminate(self):
        return self.process.terminate()

    def kill(self):
        return self.process.kill()

    def wait(self):
        return self.process.wait()

    @staticmethod
    def from_role(role, port=None,
                  ves_client_binary='../binary/ves-client', log_file_path=None, log_file=None, stdout=None, stderr=None):
        """
        :param ves_client_binary:
        :param port:
        :param log_file_path:
        :param stderr:
        :param stdout:
        :param role:
        :type playbook.Role
        :return:
        """
        if log_file_path is not None and isinstance(log_file_path, str):
            stdout = stderr = open(log_file_path, 'w')
        port = port or str(random.randint(28000, 40000))
        return VESLocalClient(
            subprocess.Popen(
                [ves_client_binary,
                 '-name', role.name,
                 '-port', ":" + port],
                stdout=stdout or log_file, stderr=stderr or log_file
            ), port=port)
