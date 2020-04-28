import json
import datetime
import enum
class LogType(enum.Enum):
    Unknown = 0
    NewSession = 1
    AckSession = 2
    GenerateAttestation = 3
    InsuranceClaim = 4
    CloseTransaction = 5
    RouteSuccess = 6
    RouteGotReceipt = 7
    AddMerkleProof = 8
    AddBlockCheck = 9
    SessionClose = 10
    NewTransaction = 11

class LoggerInfo:
    def __init__(self, source, timestamp, info):
        self.timestamp = timestamp
        self.source = source
        self.info = info
        _msg = info['_msg']
        if 'new session' in _msg:
            self.type = LogType.NewSession
            self.response_address = info['responsible address']
        elif 'user ack to server' in _msg:
            self.type = LogType.AckSession
        elif 'generate' in _msg:
            self.type = LogType.GenerateAttestation
            self.tid = info['tid']
            self.aid = info['aid']
        elif 'insurance' in _msg:
            self.type = LogType.InsuranceClaim
            self.tid = info.get('tid')
            self.aid = info.get('aid')
            if self.tid is None:
                self.type = LogType.Unknown
        elif 'status closed' in _msg:
            self.type = LogType.CloseTransaction
            self.tid = info.get('tid')
        elif 'new transaction' in _msg:
            self.type = LogType.NewTransaction
        elif 'routing' in _msg:
            self.type = LogType.RouteSuccess
        elif 'route result' in _msg:
            self.type = LogType.RouteGotReceipt
        elif 'adding merkle proof' in _msg:
            self.type = LogType.AddMerkleProof
        elif 'adding block check' in _msg:
            self.type = LogType.AddBlockCheck
        elif 'session closed' in _msg:
            self.type = LogType.SessionClose
        else:
            self.type = LogType.Unknown
        # if self.type != LogType.Unknown:
        #     print(self.__dict__)

def parseLog(name):
    lines = []
    with open(f'./res/client.{name}.out') as f:
        lines = f.readlines()
    logger_raw_events, others = [], []
    for line in lines:
        tok = line.split('\t')
        if len(tok) == 5:
            logger_raw_events.append(tok)
        else:
            others.append(tok)
            
    # print(others)

    logger_events = []
    for tok in logger_raw_events:
        logger_events.append(LoggerInfo(name, datetime.datetime.strptime(tok[0], '%Y-%m-%dT%H:%M:%S.%f+0800'), json.loads(tok[4])))

    return logger_events


class StateType(enum.Enum):
    Begin = 0
    CreateSession = 1
    Transaction = 2
    ClosedSession = 3

class StateBegin:
    def __init__(self):
        self.type = StateType.Begin

class StateCreateSession:
    def __init__(self, timestamp):
        self.type = StateType.CreateSession
        self.timestamp = timestamp
        self.duration = None

class StateTransaction:
    def __init__(self, timestamp):
        self.type = StateType.Transaction
        self.timestamp = timestamp
        self.duration = None
        self.tid = 0

class StateClosedSession:
    def __init__(self, timestamp):
        self.type = StateType.ClosedSession
        self.timestamp = timestamp

def matchState(state, event):
    if state.type == StateType.Begin:
        if event.type == LogType.NewSession:
            return StateCreateSession(event.timestamp), True
    elif state.type == StateType.CreateSession:
        if event.type == LogType.NewTransaction:
            state.duration = event.timestamp - state.timestamp
            return StateTransaction(event.timestamp), True
    elif state.type == StateType.Transaction:
        if (event.type == LogType.InsuranceClaim and event.aid == 'Closed') or \
            event.type == LogType.CloseTransaction:
            state.tid = event.tid
            state.duration = event.timestamp - state.timestamp
        elif event.type == LogType.NewTransaction:
            state.duration = event.timestamp - state.timestamp
            return StateTransaction(event.timestamp), True
        elif event.type == LogType.SessionClose:
            return StateClosedSession(event.timestamp), True
            # return StateTransaction(event.timestamp), True
    else:
        print(event.__dict__)
    return state, False


states = []

if __name__ == '__main__':
    events = []
    events += parseLog('a1') + parseLog('a2') + parseLog('ves')
    good_events, bad_events = [], []
    for r in events:
        if r.type == LogType.Unknown:
            bad_events.append(r)
        else:
            good_events.append(r)
    good_events.sort(key=lambda r: r.timestamp)
    state = StateBegin()
    for event in good_events:
        state, trans = matchState(state, event)
        if trans:
            states.append(state)
            print(event.timestamp, state.type)
        print(event.timestamp, event.source + '\t', event.type)

    s, t = None, None
    for state in states:
        if state.type == StateType.CreateSession:
            print("Create Session", "Duration:", state.duration, "timestamp:", state.timestamp)
            s = state
        elif state.type == StateType.Transaction:
            print("Execute Transaction", "Index:", state.tid, "Duration:", state.duration)
        # elif state.type == StateType.ClosedSession:
        #     print("CreateSession", "Duration:", state.duration)
        elif state.type == StateType.ClosedSession:
            print("Close Session", "timestamp:", state.timestamp)
            t = state
    
    if s and t:
        print('total:', t.timestamp - s.timestamp)

    