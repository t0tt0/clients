import enum


class Code(enum.Enum):
    OK = 0
    BindError = 1
    UnserializeDataError = 2
    InvalidParameters = 3
    GetRawDataError = 4
    ToDo = 5
    InsertError = 100
    SelectError = 101
    NotFound = 102
    DeleteNoEffect = 103
    DuplicatePrimaryKey = 104
    UpdateError = 105
    DeleteError = 106
    BeginTransactionError = 107
    CommitTransactionError = 108
    DatabaseIncorrectStringValue = 109
    UpdateNoEffect = 110
    AuthGenerateTokenError = 1000
    AuthenticatePasswordError = 1001
    AuthenticatePolicyError = 1002
    ChangeOwnerError = 1003
    GroupCreateError = 1004
    AddReadPrivilegeError = 1005
    AddWritePrivilegeError = 1006
    GrantNoEffect = 1007
    GrantError = 1008
    UserIDMissing = 10000
    UserWrongPassword = 10001
    WeakPassword = 10002
    InvalidCityCode = 10003
    BadPhone = 10004
    SubmissionUploaded = 11000
    FSExecError = 11001
    UploadFileError = 11002
    ConfigModifyError = 11003
    StatError = 11004
    SessionInitError = 12000
    SessionRequestNSBError = 12001
    SessionInitGUIDError = 12002
    SessionInitOpIntentsError = 12003
    SessionRedisGetAckCountError = 12004
    SessionInsertAccountError = 12005
    SessionFindError = 12006
    SessionNotFind = 12007
    SessionAcknowledgeError = 12008
    SessionAccountFindError = 12009
    SessionAccountNotFound = 12010
    SessionAccountGetTotolError = 12011
    SessionAccountGetAcknowledgedError = 12012
    SessionSignTxsError = 12013
    SessionFreezeInfoError = 12014
    SessionInitInternalRequestError = 12015
    TransactionFindError = 13000
    DeserializeTransactionError = 13001
    AttestationSendError = 13002
    NotEnoughParamInformation = 13003
    EnsureTransactionValueError = 13004
    ParsePaymentOptionInconsistentValueError = 13005
    TransactionPrepareTranslateError = 13006
    TransactionTranslateError = 13007
    TransactionRawSerializeError = 13008
    ChainIDNotFound = 14000
    ChainTypeNotFound = 14001
    TransactionTypeNotFound = 14002
    ValueTypeNotFound = 14003
    GetBlockChainInterfaceError = 14004
    GetTransactionIntentError = 14005
    GetStorageError = 14006
    GetStorageTypeError = 14007
    SetStorageError = 14008
    DestinationRespUnknown = 14009
    ConvertSignerError = 15000
    DecodeAdditionError = 15001
    DecodeAddressError = 15002
    BadContractField = 15003
    BadPosField = 15004
    ReadMessageError = 16000
    ReadMessageIDError = 16001
    NotConnected = 17000
    GetVESHostError = 17001
    ExecuteError = 17002


code_desc = {
    0: 'CodeOK',
    1: 'CodeBindError',
    2: 'CodeUnserializeDataError',
    3: 'CodeInvalidParameters',
    4: 'CodeGetRawDataError',
    5: 'CodeToDo',
    100: 'CodeInsertError',
    101: 'CodeSelectError',
    102: 'CodeNotFound',
    103: 'CodeDeleteNoEffect',
    104: 'CodeDuplicatePrimaryKey',
    105: 'CodeUpdateError',
    106: 'CodeDeleteError',
    107: 'CodeBeginTransactionError',
    108: 'CodeCommitTransactionError',
    109: 'CodeDatabaseIncorrectStringValue',
    110: 'CodeUpdateNoEffect',
    1000: 'CodeAuthGenerateTokenError',
    1001: 'CodeAuthenticatePasswordError',
    1002: 'CodeAuthenticatePolicyError',
    1003: 'CodeChangeOwnerError',
    1004: 'CodeGroupCreateError',
    1005: 'CodeAddReadPrivilegeError',
    1006: 'CodeAddWritePrivilegeError',
    1007: 'CodeGrantNoEffect',
    1008: 'CodeGrantError',
    10000: 'CodeUserIDMissing',
    10001: 'CodeUserWrongPassword',
    10002: 'CodeWeakPassword',
    10003: 'CodeInvalidCityCode',
    10004: 'CodeBadPhone',
    11000: 'CodeSubmissionUploaded',
    11001: 'CodeFSExecError',
    11002: 'CodeUploadFileError',
    11003: 'CodeConfigModifyError',
    11004: 'CodeStatError',
    12000: 'CodeSessionInitError',
    12001: 'CodeSessionRequestNSBError',
    12002: 'CodeSessionInitGUIDError',
    12003: 'CodeSessionInitOpIntentsError',
    12004: 'CodeSessionRedisGetAckCountError',
    12005: 'CodeSessionInsertAccountError',
    12006: 'CodeSessionFindError',
    12007: 'CodeSessionNotFind',
    12008: 'CodeSessionAcknowledgeError',
    12009: 'CodeSessionAccountFindError',
    12010: 'CodeSessionAccountNotFound',
    12011: 'CodeSessionAccountGetTotolError',
    12012: 'CodeSessionAccountGetAcknowledgedError',
    12013: 'CodeSessionSignTxsError',
    12014: 'CodeSessionFreezeInfoError',
    12015: 'CodeSessionInitInternalRequestError',
    13000: 'CodeTransactionFindError',
    13001: 'CodeDeserializeTransactionError',
    13002: 'CodeAttestationSendError',
    13003: 'CodeNotEnoughParamInformation',
    13004: 'CodeEnsureTransactionValueError',
    13005: 'CodeParsePaymentOptionInconsistentValueError',
    13006: 'CodeTransactionPrepareTranslateError',
    13007: 'CodeTransactionTranslateError',
    13008: 'CodeTransactionRawSerializeError',
    14000: 'CodeChainIDNotFound',
    14001: 'CodeChainTypeNotFound',
    14002: 'CodeTransactionTypeNotFound',
    14003: 'CodeValueTypeNotFound',
    14004: 'CodeGetBlockChainInterfaceError',
    14005: 'CodeGetTransactionIntentError',
    14006: 'CodeGetStorageError',
    14007: 'CodeGetStorageTypeError',
    14008: 'CodeSetStorageError',
    14009: 'CodeDestinationRespUnknown',
    15000: 'CodeConvertSignerError',
    15001: 'CodeDecodeAdditionError',
    15002: 'CodeDecodeAddressError',
    15003: 'CodeBadContractField',
    15004: 'CodeBadPosField',
    16000: 'CodeReadMessageError',
    16001: 'CodeReadMessageIDError',
    17000: 'CodeNotConnected',
    17001: 'CodeGetVESHostError',
    17002: 'CodeExecuteError',
}

code_revert_desc = {
    'CodeOK': 0,
    'CodeBindError': 1,
    'CodeUnserializeDataError': 2,
    'CodeInvalidParameters': 3,
    'CodeGetRawDataError': 4,
    'CodeToDo': 5,
    'CodeInsertError': 100,
    'CodeSelectError': 101,
    'CodeNotFound': 102,
    'CodeDeleteNoEffect': 103,
    'CodeDuplicatePrimaryKey': 104,
    'CodeUpdateError': 105,
    'CodeDeleteError': 106,
    'CodeBeginTransactionError': 107,
    'CodeCommitTransactionError': 108,
    'CodeDatabaseIncorrectStringValue': 109,
    'CodeUpdateNoEffect': 110,
    'CodeAuthGenerateTokenError': 1000,
    'CodeAuthenticatePasswordError': 1001,
    'CodeAuthenticatePolicyError': 1002,
    'CodeChangeOwnerError': 1003,
    'CodeGroupCreateError': 1004,
    'CodeAddReadPrivilegeError': 1005,
    'CodeAddWritePrivilegeError': 1006,
    'CodeGrantNoEffect': 1007,
    'CodeGrantError': 1008,
    'CodeUserIDMissing': 10000,
    'CodeUserWrongPassword': 10001,
    'CodeWeakPassword': 10002,
    'CodeInvalidCityCode': 10003,
    'CodeBadPhone': 10004,
    'CodeSubmissionUploaded': 11000,
    'CodeFSExecError': 11001,
    'CodeUploadFileError': 11002,
    'CodeConfigModifyError': 11003,
    'CodeStatError': 11004,
    'CodeSessionInitError': 12000,
    'CodeSessionRequestNSBError': 12001,
    'CodeSessionInitGUIDError': 12002,
    'CodeSessionInitOpIntentsError': 12003,
    'CodeSessionRedisGetAckCountError': 12004,
    'CodeSessionInsertAccountError': 12005,
    'CodeSessionFindError': 12006,
    'CodeSessionNotFind': 12007,
    'CodeSessionAcknowledgeError': 12008,
    'CodeSessionAccountFindError': 12009,
    'CodeSessionAccountNotFound': 12010,
    'CodeSessionAccountGetTotolError': 12011,
    'CodeSessionAccountGetAcknowledgedError': 12012,
    'CodeSessionSignTxsError': 12013,
    'CodeSessionFreezeInfoError': 12014,
    'CodeSessionInitInternalRequestError': 12015,
    'CodeTransactionFindError': 13000,
    'CodeDeserializeTransactionError': 13001,
    'CodeAttestationSendError': 13002,
    'CodeNotEnoughParamInformation': 13003,
    'CodeEnsureTransactionValueError': 13004,
    'CodeParsePaymentOptionInconsistentValueError': 13005,
    'CodeTransactionPrepareTranslateError': 13006,
    'CodeTransactionTranslateError': 13007,
    'CodeTransactionRawSerializeError': 13008,
    'CodeChainIDNotFound': 14000,
    'CodeChainTypeNotFound': 14001,
    'CodeTransactionTypeNotFound': 14002,
    'CodeValueTypeNotFound': 14003,
    'CodeGetBlockChainInterfaceError': 14004,
    'CodeGetTransactionIntentError': 14005,
    'CodeGetStorageError': 14006,
    'CodeGetStorageTypeError': 14007,
    'CodeSetStorageError': 14008,
    'CodeDestinationRespUnknown': 14009,
    'CodeConvertSignerError': 15000,
    'CodeDecodeAdditionError': 15001,
    'CodeDecodeAddressError': 15002,
    'CodeBadContractField': 15003,
    'CodeBadPosField': 15004,
    'CodeReadMessageError': 16000,
    'CodeReadMessageIDError': 16001,
    'CodeNotConnected': 17000,
    'CodeGetVESHostError': 17001,
    'CodeExecuteError': 17002,
}


class VESError(Exception):
    pass


class OK(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.OK
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class BindError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.BindError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UnserializeDataError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UnserializeDataError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class InvalidParameters(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.InvalidParameters
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetRawDataError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetRawDataError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ToDo(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ToDo
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class InsertError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.InsertError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SelectError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SelectError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class NotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.NotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DeleteNoEffect(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DeleteNoEffect
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DuplicatePrimaryKey(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DuplicatePrimaryKey
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UpdateError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UpdateError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DeleteError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DeleteError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class BeginTransactionError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.BeginTransactionError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class CommitTransactionError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.CommitTransactionError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DatabaseIncorrectStringValue(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DatabaseIncorrectStringValue
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UpdateNoEffect(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UpdateNoEffect
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AuthGenerateTokenError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AuthGenerateTokenError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AuthenticatePasswordError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AuthenticatePasswordError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AuthenticatePolicyError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AuthenticatePolicyError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ChangeOwnerError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ChangeOwnerError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GroupCreateError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GroupCreateError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AddReadPrivilegeError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AddReadPrivilegeError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AddWritePrivilegeError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AddWritePrivilegeError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GrantNoEffect(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GrantNoEffect
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GrantError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GrantError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UserIDMissing(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UserIDMissing
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UserWrongPassword(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UserWrongPassword
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class WeakPassword(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.WeakPassword
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class InvalidCityCode(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.InvalidCityCode
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class BadPhone(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.BadPhone
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SubmissionUploaded(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SubmissionUploaded
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class FSExecError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.FSExecError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class UploadFileError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.UploadFileError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ConfigModifyError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ConfigModifyError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class StatError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.StatError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionInitError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionInitError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionRequestNSBError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionRequestNSBError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionInitGUIDError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionInitGUIDError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionInitOpIntentsError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionInitOpIntentsError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionRedisGetAckCountError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionRedisGetAckCountError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionInsertAccountError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionInsertAccountError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionFindError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionFindError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionNotFind(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionNotFind
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionAcknowledgeError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionAcknowledgeError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionAccountFindError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionAccountFindError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionAccountNotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionAccountNotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionAccountGetTotolError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionAccountGetTotolError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionAccountGetAcknowledgedError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionAccountGetAcknowledgedError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionSignTxsError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionSignTxsError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionFreezeInfoError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionFreezeInfoError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SessionInitInternalRequestError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SessionInitInternalRequestError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class TransactionFindError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.TransactionFindError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DeserializeTransactionError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DeserializeTransactionError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class AttestationSendError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.AttestationSendError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class NotEnoughParamInformation(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.NotEnoughParamInformation
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class EnsureTransactionValueError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.EnsureTransactionValueError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ParsePaymentOptionInconsistentValueError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ParsePaymentOptionInconsistentValueError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class TransactionPrepareTranslateError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.TransactionPrepareTranslateError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class TransactionTranslateError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.TransactionTranslateError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class TransactionRawSerializeError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.TransactionRawSerializeError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ChainIDNotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ChainIDNotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ChainTypeNotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ChainTypeNotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class TransactionTypeNotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.TransactionTypeNotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ValueTypeNotFound(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ValueTypeNotFound
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetBlockChainInterfaceError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetBlockChainInterfaceError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetTransactionIntentError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetTransactionIntentError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetStorageError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetStorageError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetStorageTypeError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetStorageTypeError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class SetStorageError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.SetStorageError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DestinationRespUnknown(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DestinationRespUnknown
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ConvertSignerError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ConvertSignerError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DecodeAdditionError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DecodeAdditionError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class DecodeAddressError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.DecodeAddressError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class BadContractField(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.BadContractField
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class BadPosField(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.BadPosField
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ReadMessageError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ReadMessageError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ReadMessageIDError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ReadMessageIDError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class NotConnected(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.NotConnected
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class GetVESHostError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.GetVESHostError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


class ExecuteError(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.ExecuteError
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


def response_to_error(r):
    return eval(code_desc[r.code][4:])(r.get_error())

