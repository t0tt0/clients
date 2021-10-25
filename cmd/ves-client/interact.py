import rlcompleter
import readline
import code


def init():
    _ = rlcompleter
    # tab completion
    readline.parse_and_bind('tab: complete')



def interact(*args, **kwargs):
    # kwargs['local'] = kwargs.get('local')
    code.interact(*args, **kwargs)
