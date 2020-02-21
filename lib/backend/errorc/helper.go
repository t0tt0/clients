package errorc

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/go-sql-driver/mysql"
	"reflect"
)

type Code = types2.CodeRawType

func MaybeSelectError(anyObj interface{}, err error) serial.ErrorSerializer {
	if err != nil {
		return serial.ErrorSerializer{Code: types2.CodeSelectError, Err: err.Error()}
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return serial.ErrorSerializer{Code: types2.CodeNotFound, Err: "not found"}
	}
	return serial.ErrorSerializer{Code: types2.CodeOK}
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(err error) serial.ErrorSerializer {
	if err != nil {
		return serial.ErrorSerializer{Code: types2.CodeUpdateError, Err: err.Error()}
	}
	return serial.ErrorSerializer{Code: types2.CodeOK}
}

func CheckInsertError(err error) serial.ErrorSerializer {
	return serial.ErrorSerializer{Code: types2.CodeOK}
}

func CreateObj(aff int64,  err error) serial.ErrorSerializer {
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			switch mysqlError.Number {
			case 1062:
				return serial.ErrorSerializer{Code: types2.CodeDuplicatePrimaryKey, Err: err.Error()}
			case 1366:
				return serial.ErrorSerializer{Code: types2.CodeDatabaseIncorrectStringValue, Err: err.Error()}
			}
		}
		return serial.ErrorSerializer{Code: types2.CodeInsertError, Err: err.Error()}
	} else if aff == 0 {
		return serial.ErrorSerializer{Code: types2.CodeInsertError, Err: "affect nothing"}
	}
	return serial.ErrorSerializer{Code: types2.CodeOK}
}
