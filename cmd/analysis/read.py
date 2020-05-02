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
            self.session_id = info['session id']
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
            self.session_id = info['session id']
            self.response_address = info['address']
        elif 'routing' in _msg:
            self.type = LogType.RouteSuccess
            self.receipt = info['receipt']
        elif 'route result' in _msg:
            self.type = LogType.RouteGotReceipt
            self.block_id = info['block id']
        elif 'adding merkle proof' in _msg:
            self.type = LogType.AddMerkleProof
            self.tx_hash = info['result']['hash']
        elif 'adding block check' in _msg:
            self.type = LogType.AddBlockCheck
            self.tx_hash = info['result']['hash']
        elif 'session closed' in _msg:
            self.type = LogType.SessionClose
        else:
            self.type = LogType.Unknown
        # if self.type != LogType.Unknown:
        #     print(self.__dict__)

fp = './res'

def parseLog(name):
    lines = []
    with open(f'{fp}/client.{name}.out') as f:
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
        pass
        # print(event.__dict__)
    return state, False


states = []

if __name__ == '__main__':
    events = []
    events += parseLog('a1') + parseLog('a2') + parseLog('a3') + parseLog('ves')
    good_events, bad_events = [], []
    for r in events:
        if r.type == LogType.Unknown:
            bad_events.append(r)
        else:
            good_events.append(r)
    good_events.sort(key=lambda r: r.timestamp)

    # for event in bad_events:
    #     print(event.type, event.info)

    state = StateBegin()
    for event in good_events:
        state, trans = matchState(state, event)
        if trans:
            states.append(state)
            print(event.timestamp, state.type)
        print(event.timestamp, event.source + '\t', str(event.type)+'\t', end='')
        if event.type == LogType.NewSession:
            print('\tses_id:', event.session_id[:8] + '\t', 'resp_addr:', event.response_address[:8])
            # event.response_address = info['responsible address']
        elif event.type == LogType.AckSession:
            print()
        elif event.type == LogType.GenerateAttestation:
            print('tid:', str(event.tid)+'\t\t\t', 'aid:', event.aid)
            # event.tid = info['tid']
            # event.aid = info['aid']
        elif event.type == LogType.InsuranceClaim:
            print('\ttid:', str(event.tid)+'\t\t\t', 'aid:', event.aid)
            # event.tid = info.get('tid')
            # event.aid = info.get('aid')
            # if event.tid is None:
            #     event.type = LogType.Unknown
        elif event.type == LogType.CloseTransaction:
            print('tid:', str(event.tid))
            # event.tid = info.get('tid')
        elif event.type == LogType.NewTransaction:
            print('\tses_id:', event.session_id[:8] + '\t', 'resp_addr:', event.response_address[:8])
        elif event.type == LogType.RouteSuccess:
            print('\treceipt:', event.receipt[:8])
        elif event.type == LogType.RouteGotReceipt:
            print('block_id:', event.block_id[:12])
        elif event.type == LogType.AddMerkleProof:
            print('\tnsb_tx_hash:', event.tx_hash[:8])
        elif event.type == LogType.AddBlockCheck:
            print('\tnsb_tx_hash:', event.tx_hash[:8])
        elif event.type == LogType.SessionClose:
            print()
        elif event.type == LogType.Unknown:
            print()

        # if self.type != LogType.Unknown:
        #     print(self.__dict__)
    


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

    