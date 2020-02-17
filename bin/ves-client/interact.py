import copy
import rlcompleter
import readline
import code


def init():
    _ = rlcompleter
    # tab completion
    readline.parse_and_bind('tab: complete')


def create_local():
    _local = copy.copy(locals())
    _local.update(globals())
    return _local


def interact(*args, **kwargs):
    kwargs['local'] = kwargs.get('local', create_local())
    code.interact(*args, **kwargs)
