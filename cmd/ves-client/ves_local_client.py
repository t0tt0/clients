import subprocess
import random
import threading
import json
mutex = threading.Lock()
console_filters = []
# console_filters = [
#     # 'error',
#     '!myriad-!dreamin', 'failed', 'the instance of', 'routing', 'blockHash', 'generate']


class TeeIO(object):
    def __init__(self, n, writer):
        self.writer = writer
        self. n = n

    def fileno(self):
        return self.writer.fileno()

    def write(self, s: bytes):
        s = s.decode('utf-8')
        t = s.strip()
        for console_filter in console_filters:
            if t.find(console_filter) != -1:
                t = ''
                break
        i = t.find('{')
        if i != -1:
            t = t[i:]
        mutex.acquire()
        if t != '':
            print(self.n + ':', t)
        mutex.release()
        self.writer.write(s)
        if t.startswith('exited'):
            return

    def writeline(self, line):
        self.write(line)

    def writelines(self, lines):
        for line in lines:
            self.writeline(line)


class VESLocalClient:
    def __init__(self, process: subprocess.Popen, port, stdout, stderr):
        self.process = process
        self.port = port
        self.host = "localhost:" + port
        self.stdout = stdout
        self.stderr = stderr

    def terminate(self):
        return self.process.terminate()

    def kill(self):
        return self.process.kill()

    def wait(self):
        return self.process.wait()

    @staticmethod
    def from_role(role, port=None,
                  ves_client_binary='../binary/ves-client', log_file_path=None, log_file=None, stdout=None,
                  stderr=None):
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
            stdout = stderr = TeeIO(role.name, open(log_file_path, 'w'))
        if log_file is not None:
            log_file = TeeIO(role.name, log_file)
        port = port or str(random.randint(28000, 40000))
        return VESLocalClient(
            subprocess.Popen(
                [ves_client_binary,
                 '-name', role.name,
                 '-port', ":" + port],
                stdout=subprocess.PIPE, stderr=subprocess.PIPE
            ), port=port, stdout=stdout or log_file, stderr=stderr or log_file)
