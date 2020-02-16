//go:generate stringer -type=CodeType
package types

type CodeRawType = int
type CodeType CodeRawType
type Code = CodeRawType

const (
	// Generic Code

	CodeOK Code = iota
	// CodeBindError indicates a parameter missing error
	CodeBindError
	// CodeUnserializeDataError indicates a parsing data error
	CodeUnserializeDataError
	// CodeInvalidParameters tells some wrong data was in the request
	CodeInvalidParameters
	// GetRawDataError tells some wrong data was in the request
	CodeGetRawDataError
	CodeToDo

	CodeGenericErrorR
	CodeGenericErrorL = CodeOK
)

const (
	// Generic Code -- Database
	// CodeInsertError occurs when insert object into database
	CodeInsertError Code = iota + 100
	// CodeSelectError occurs when select object from database
	CodeSelectError
	// CodeNotFound occurs when object with specific condition is not in the
	// database
	CodeNotFound
	// CodeDeleteNoEffect occurs when deleting object has no effect
	CodeDeleteNoEffect
	// CodeDuplicatePrimaryKey occurs when the object's primary key conflicts
	// with something that was already in the database
	CodeDuplicatePrimaryKey
	// CodeUpdateError occurs when update object to database
	CodeUpdateError // 105
	// CodeDeleteError occurs when delete object from database
	CodeDeleteError

	// CodeDeleteError occurs when begin a transaction
	CodeBeginTransactionError

	// CodeDeleteError occurs when commit a transaction
	CodeCommitTransactionError

	//
	CodeDatabaseIncorrectStringValue

	CodeUpdateNoEffect

	CodeDatabaseErrorR
	CodeDatabaseErrorL = CodeInsertError
)

const (
	// Generic Code -- Authentication
	// CodeAuthGenerateTokenError occurs when insert object into database
	CodeAuthGenerateTokenError Code = iota + 1000
	CodeAuthenticatePasswordError
	CodeAuthenticatePolicyError

	CodeChangeOwnerError
	CodeGroupCreateError
	CodeAddReadPrivilegeError
	CodeAddWritePrivilegeError

	CodeGrantNoEffect
	CodeGrantError

	CodeAuthenticationErrorR
	CodeAuthenticationErrorL = CodeAuthGenerateTokenError
)

const (
	CodeUserIDMissing Code = iota + 10000
	CodeUserWrongPassword
	CodeWeakPassword
	CodeInvalidCityCode
	CodeBadPhone

	CodeUserServiceErrorR
	CodeUserServiceErrorL = CodeUserIDMissing
)

const (
	CodeSubmissionUploaded Code = iota + 11000
	CodeFSExecError
	CodeUploadFileError
	CodeConfigModifyError
	CodeStatError

	CodeFileSystemErrorR
	CodeFileSystemErrorL = CodeSubmissionUploaded
)

const (
	CodeSessionInitError Code = iota + 12000
	CodeSessionRequestNSBError
	CodeSessionInitGUIDError
	CodeSessionInitOpIntentsError
	CodeSessionRedisGetAckCountError
	CodeSessionInsertAccountError
	CodeSessionFindError
	CodeSessionNotFind
	CodeSessionAcknowledgeError
	CodeSessionAccountFindError
	CodeSessionAccountNotFound
	CodeSessionAccountGetTotolError
	CodeSessionAccountGetAcknowledgedError
	CodeSessionSignTxsError
	CodeSessionFreezeInfoError
	CodeSessionInitInternalRequestError

	CodeSessionServiceErrorR
	CodeSessionServiceErrorL = CodeSessionInitError
)

const (
	CodeTransactionFindError Code = iota + 13000
	CodeDeserializeTransactionError
	CodeAttestationSendError
	CodeNotEnoughParamInformation
	CodeEnsureTransactionValueError
	CodeParsePaymentOptionInconsistentValueError
	CodeTransactionPrepareTranslateError
	CodeTransactionTranslateError
	CodeTransactionRawSerializeError

	CodeTransactionServiceErrorR
	CodeTransactionServiceErrorL = CodeTransactionFindError
)

const (
	CodeChainIDNotFound Code = iota + 14000
	CodeChainTypeNotFound
	CodeTransactionTypeNotFound
	CodeValueTypeNotFound
	CodeGetBlockChainInterfaceError
	CodeGetTransactionIntentError
	CodeGetStorageError
	CodeGetStorageTypeError
	CodeSetStorageError

	CodeBlockChainErrorR
	CodeBlockChainErrorL = CodeChainIDNotFound
)

const (
	CodeConvertSignerError Code = iota + 15000
	CodeDecodeAdditionError
	CodeDecodeAddressError
	CodeBadContractField
	CodeBadPosField

	CodeConvertErrorR
	CodeConvertErrorL = CodeConvertSignerError
)

const (
	CodeReadMessageError Code = iota + 16000
	CodeReadMessageIDError

	CodeWebSocketErrorR
	CodeWebSocketErrorL = CodeReadMessageError
)

const (
	CodeNotConnected Code = iota + 17000
	CodeGetVESHostError
	CodeExecuteError

	CodeGRPCErrorR
	CodeGRPCErrorL = CodeNotConnected
)

//

var CodeDesc map[Code]string

func init() {
	CodeDesc = make(map[Code]string)
	for _, groupCode := range []struct {
		L Code
		R Code
	}{
		{CodeGenericErrorL, CodeGenericErrorR},
		{CodeDatabaseErrorL, CodeDatabaseErrorR},
		{CodeAuthenticationErrorL, CodeAuthenticationErrorR},
		{CodeFileSystemErrorL, CodeFileSystemErrorR},
		{CodeUserServiceErrorL, CodeUserServiceErrorR},
		{CodeSessionServiceErrorL, CodeSessionServiceErrorR},
		{CodeTransactionServiceErrorL, CodeTransactionServiceErrorR},
		{CodeBlockChainErrorL, CodeBlockChainErrorR},
		{CodeConvertErrorL, CodeConvertErrorR},
		{CodeWebSocketErrorL, CodeWebSocketErrorR},
		{CodeGRPCErrorL, CodeGRPCErrorR},
	} {
		for i := groupCode.L; i < groupCode.R; i++ {
			CodeDesc[i] = CodeType(i).String()
		}
	}
}
