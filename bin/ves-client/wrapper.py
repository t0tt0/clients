magic = '<84f4446f>'


def is_wrap_error(e):
    """
    :param e: {string}
    :return:
    """
    return e.startswith(magic)


class FileLine:

    def __init__(self, file='', line=-1):
        self.file = file
        self.line = line

    def __str__(self):
        return f"{self.file}:{self.line}"

    def rel(self, scope):
        return f"{scope.rel_path(self.file)}:{self.line}"

    @staticmethod
    def unwrap(e):
        c = e.rsplit(':', 1)
        if len(c) < 2:
            return FileLine()
        return FileLine(c[0], int(c[1]))


class Func:
    def __init__(self, name='', fileline=None):
        self.name = name
        self.fileline = fileline

    def __str__(self):
        return f"<name:{self.name},fileline:{self.fileline}>"

    def rel(self, scope):
        return f"<name:{scope.rel_pac(self.name)},fileline:{self.fileline.rel(scope)}>"

    @staticmethod
    def unwrap(e):
        if len(e) >= 2 and e[0] == '<' and e[-1] == '>':
            e = e[1:-1]
        else:
            return Func()
        c = e.split(',', 1)
        return Func(c[0], FileLine.unwrap(c[1]))


class StackPos:
    def __init__(self, fn=None, fileline=None):
        self.fn = fn
        self.fileline = fileline

    def __str__(self):
        return f"<fn:{self.fn},fileline:{self.fileline}>"

    def rel(self, scope):
        return f"<fn:{self.fn.rel(scope)},fileline:{self.fileline.rel(scope)}>"

    @staticmethod
    def unwrap(e):
        if len(e) >= 2 and e[0] == '<' and e[-1] == '>':
            e = e[1:-1]
        else:
            return StackPos()
        if len(e) > 0 and e[0] == '!':
            return StackPos()
        c = e.split('>', 1)
        return StackPos(Func.unwrap(c[0] + '>'), FileLine.unwrap(c[1].lstrip(',')))


class Frame(object):
    def __init__(self, pos, code, err):
        self.pos = pos
        self.code = code
        self.err = err

    def __str__(self):
        return f"<pos:{self.pos},code:{self.code},err:{self.err}>"

    def rel(self, scope):
        return f"<pos:{self.pos.rel(scope)},code:{self.code},err:{self.err}>"

    @staticmethod
    def unwrap(e):
        if not is_wrap_error(e):
            return e, False
        c = test_string.split(magic, 3)
        if len(c) < 4:
            return e, False
        pos, code, err = \
            StackPos.unwrap(c[1][4:-1]), \
            int(c[2][5:-1]), c[3][4:-1]
        return Frame(pos, code, err), True


import os.path


class Scope(object):
    def __init__(self, pn='', wd=''):
        self.pn = pn
        self.wd = wd

    def rel_pac(self, pac):
        if pac.startswith(self.pn):
            return pac[len(self.pn):].lstrip('.')
        return pac

    def rel_path(self, path):
        print(path)
        return os.path.relpath(path, self.wd)


if __name__ == '__main__':
    test_string = '<84f4446f>pos:<<github.com/Myriad-Dreamin/go-ves/lib/net/ves-client.(*VesClient).SendOpIntents,' \
                  '/home/kamiyoru/work/gosrc/src/github.com/Myriad-Dreamin/go-ves/lib/net/ves-client/client-send' \
                  '-opintents.go:19>,/home/kamiyoru/work/gosrc/src/github.com/Myriad-Dreamin/go-ves/lib/net/ves' \
                  '-client/client-send-opintents.go:37>,<84f4446f>code:17002,<84f4446f>err:rpc error: code = ' \
                  'Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection closed '
    f, ok = Frame.unwrap(test_string)
    if ok:
        print(f.rel(Scope('github.com/Myriad-Dreamin/go-ves',
                          '/home/kamiyoru/work/gosrc/src/github.com/Myriad-Dreamin/go-ves/lib'
                          '/net/ves-client')))
