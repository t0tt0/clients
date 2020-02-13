// Code generated by "stringer -type=CodeType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeOK-0]
	_ = x[CodeBindError-1]
	_ = x[CodeUnserializeDataError-2]
	_ = x[CodeInvalidParameters-3]
	_ = x[CodeGetRawDataError-4]
	_ = x[CodeToDo-5]
	_ = x[CodeGenericErrorR-6]
	_ = x[CodeInsertError-100]
	_ = x[CodeSelectError-101]
	_ = x[CodeNotFound-102]
	_ = x[CodeDeleteNoEffect-103]
	_ = x[CodeDuplicatePrimaryKey-104]
	_ = x[CodeUpdateError-105]
	_ = x[CodeDeleteError-106]
	_ = x[CodeBeginTransactionError-107]
	_ = x[CodeCommitTransactionError-108]
	_ = x[CodeDatabaseIncorrectStringValue-109]
	_ = x[CodeUpdateNoEffect-110]
	_ = x[CodeDatabaseErrorR-111]
	_ = x[CodeAuthGenerateTokenError-1000]
	_ = x[CodeAuthenticatePasswordError-1001]
	_ = x[CodeAuthenticatePolicyError-1002]
	_ = x[CodeChangeOwnerError-1003]
	_ = x[CodeGroupCreateError-1004]
	_ = x[CodeAddReadPrivilegeError-1005]
	_ = x[CodeAddWritePrivilegeError-1006]
	_ = x[CodeGrantNoEffect-1007]
	_ = x[CodeGrantError-1008]
	_ = x[CodeAuthenticationErrorR-1009]
	_ = x[CodeUserIDMissing-10000]
	_ = x[CodeUserWrongPassword-10001]
	_ = x[CodeWeakPassword-10002]
	_ = x[CodeInvalidCityCode-10003]
	_ = x[CodeBadPhone-10004]
	_ = x[CodeUserServiceErrorR-10005]
	_ = x[CodeSubmissionUploaded-11000]
	_ = x[CodeFSExecError-11001]
	_ = x[CodeUploadFileError-11002]
	_ = x[CodeConfigModifyError-11003]
	_ = x[CodeStatError-11004]
	_ = x[CodeFileSystemErrorR-11005]
	_ = x[CodeSessionInitError-12000]
	_ = x[CodeSessionRequestNSBError-12001]
	_ = x[CodeSessionInitGUIDError-12002]
	_ = x[CodeSessionInitOpIntentsError-12003]
	_ = x[CodeSessionRedisGetAckCountError-12004]
	_ = x[CodeSessionInsertAccountError-12005]
	_ = x[CodeSessionFindError-12006]
	_ = x[CodeSessionNotFindError-12007]
	_ = x[CodeSessionAcknowledgeError-12008]
	_ = x[CodeSessionAccountFindError-12009]
	_ = x[CodeSessionAccountNotFound-12010]
	_ = x[CodeSessionAccountGetTotolError-12011]
	_ = x[CodeSessionAccountGetAcknowledgedError-12012]
	_ = x[CodeSessionSignTxsError-12013]
	_ = x[CodeSessionFreezeInfoError-12014]
	_ = x[CodeSessionInitInternalRequestError-12015]
	_ = x[CodeSessionServiceErrorR-12016]
	_ = x[CodeTransactionFindError-13000]
	_ = x[CodeDeserializeTransactionError-13001]
	_ = x[CodeAttestationSendError-13002]
	_ = x[CodeTransactionServiceErrorR-13003]
	_ = x[CodeChainIDNotFound-14000]
	_ = x[CodeChainTypeNotFound-14001]
	_ = x[CodeChainDNSErrorR-14002]
	_ = x[CodeConvertSignerError-15000]
	_ = x[CodeDecodeAdditionError-15001]
	_ = x[CodeDecodeAddressError-15002]
	_ = x[CodeConvertErrorR-15003]
}

const (
	_CodeType_name_0 = "CodeOKCodeBindErrorCodeUnserializeDataErrorCodeInvalidParametersCodeGetRawDataErrorCodeToDoCodeGenericErrorR"
	_CodeType_name_1 = "CodeInsertErrorCodeSelectErrorCodeNotFoundCodeDeleteNoEffectCodeDuplicatePrimaryKeyCodeUpdateErrorCodeDeleteErrorCodeBeginTransactionErrorCodeCommitTransactionErrorCodeDatabaseIncorrectStringValueCodeUpdateNoEffectCodeDatabaseErrorR"
	_CodeType_name_2 = "CodeAuthGenerateTokenErrorCodeAuthenticatePasswordErrorCodeAuthenticatePolicyErrorCodeChangeOwnerErrorCodeGroupCreateErrorCodeAddReadPrivilegeErrorCodeAddWritePrivilegeErrorCodeGrantNoEffectCodeGrantErrorCodeAuthenticationErrorR"
	_CodeType_name_3 = "CodeUserIDMissingCodeUserWrongPasswordCodeWeakPasswordCodeInvalidCityCodeCodeBadPhoneCodeUserServiceErrorR"
	_CodeType_name_4 = "CodeSubmissionUploadedCodeFSExecErrorCodeUploadFileErrorCodeConfigModifyErrorCodeStatErrorCodeFileSystemErrorR"
	_CodeType_name_5 = "CodeSessionInitErrorCodeSessionRequestNSBErrorCodeSessionInitGUIDErrorCodeSessionInitOpIntentsErrorCodeSessionRedisGetAckCountErrorCodeSessionInsertAccountErrorCodeSessionFindErrorCodeSessionNotFindErrorCodeSessionAcknowledgeErrorCodeSessionAccountFindErrorCodeSessionAccountNotFoundCodeSessionAccountGetTotolErrorCodeSessionAccountGetAcknowledgedErrorCodeSessionSignTxsErrorCodeSessionFreezeInfoErrorCodeSessionInitInternalRequestErrorCodeSessionServiceErrorR"
	_CodeType_name_6 = "CodeTransactionFindErrorCodeDeserializeTransactionErrorCodeAttestationSendErrorCodeTransactionServiceErrorR"
	_CodeType_name_7 = "CodeChainIDNotFoundCodeChainTypeNotFoundCodeChainDNSErrorR"
	_CodeType_name_8 = "CodeConvertSignerErrorCodeDecodeAdditionErrorCodeDecodeAddressErrorCodeConvertErrorR"
)

var (
	_CodeType_index_0 = [...]uint8{0, 6, 19, 43, 64, 83, 91, 108}
	_CodeType_index_1 = [...]uint8{0, 15, 30, 42, 60, 83, 98, 113, 138, 164, 196, 214, 232}
	_CodeType_index_2 = [...]uint8{0, 26, 55, 82, 102, 122, 147, 173, 190, 204, 228}
	_CodeType_index_3 = [...]uint8{0, 17, 38, 54, 73, 85, 106}
	_CodeType_index_4 = [...]uint8{0, 22, 37, 56, 77, 90, 110}
	_CodeType_index_5 = [...]uint16{0, 20, 46, 70, 99, 131, 160, 180, 203, 230, 257, 283, 314, 352, 375, 401, 436, 460}
	_CodeType_index_6 = [...]uint8{0, 24, 55, 79, 107}
	_CodeType_index_7 = [...]uint8{0, 19, 40, 58}
	_CodeType_index_8 = [...]uint8{0, 22, 45, 67, 84}
)

func (i CodeType) String() string {
	switch {
	case 0 <= i && i <= 6:
		return _CodeType_name_0[_CodeType_index_0[i]:_CodeType_index_0[i+1]]
	case 100 <= i && i <= 111:
		i -= 100
		return _CodeType_name_1[_CodeType_index_1[i]:_CodeType_index_1[i+1]]
	case 1000 <= i && i <= 1009:
		i -= 1000
		return _CodeType_name_2[_CodeType_index_2[i]:_CodeType_index_2[i+1]]
	case 10000 <= i && i <= 10005:
		i -= 10000
		return _CodeType_name_3[_CodeType_index_3[i]:_CodeType_index_3[i+1]]
	case 11000 <= i && i <= 11005:
		i -= 11000
		return _CodeType_name_4[_CodeType_index_4[i]:_CodeType_index_4[i+1]]
	case 12000 <= i && i <= 12016:
		i -= 12000
		return _CodeType_name_5[_CodeType_index_5[i]:_CodeType_index_5[i+1]]
	case 13000 <= i && i <= 13003:
		i -= 13000
		return _CodeType_name_6[_CodeType_index_6[i]:_CodeType_index_6[i+1]]
	case 14000 <= i && i <= 14002:
		i -= 14000
		return _CodeType_name_7[_CodeType_index_7[i]:_CodeType_index_7[i+1]]
	case 15000 <= i && i <= 15003:
		i -= 15000
		return _CodeType_name_8[_CodeType_index_8[i]:_CodeType_index_8[i+1]]
	default:
		return "CodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}