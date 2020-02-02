package errorc

import (
	"github.com/Myriad-Dreamin/go-ves/central-ves/lib/serial"
	"github.com/Myriad-Dreamin/go-ves/central-ves/types"
	"github.com/go-sql-driver/mysql"
	"reflect"
)

type Code = types.CodeRawType

func MaybeSelectError(anyObj interface{}, err error) serial.ErrorSerializer {
	if err != nil {
		return serial.ErrorSerializer{Code: types.CodeSelectError, Err: err.Error()}
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return serial.ErrorSerializer{Code: types.CodeNotFound, Err: "not found"}
	}
	return serial.ErrorSerializer{Code: types.CodeOK}
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(obj UpdateFieldsable, fields []string) serial.ErrorSerializer {
	_, err := obj.UpdateFields(fields)
	if err != nil {
		return serial.ErrorSerializer{Code: types.CodeUpdateError, Err: err.Error()}
	}
	return serial.ErrorSerializer{Code: types.CodeOK}
}

type Creatable interface {
	Create() (int64, error)
}

func CheckInsertError(err error) serial.ErrorSerializer {
	return serial.ErrorSerializer{Code: types.CodeOK}
}

func CreateObj(createObj Creatable) serial.ErrorSerializer {
	affected, err := createObj.Create()
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			switch mysqlError.Number {
			case 1062:
				return serial.ErrorSerializer{Code: types.CodeDuplicatePrimaryKey, Err: err.Error()}
			case 1366:
				return serial.ErrorSerializer{Code: types.CodeDatabaseIncorrectStringValue, Err: err.Error()}
			}
		}
		return serial.ErrorSerializer{Code: types.CodeInsertError, Err: err.Error()}
	} else if affected == 0 {
		return serial.ErrorSerializer{Code: types.CodeInsertError, Err: "affect nothing"}
	}
	return serial.ErrorSerializer{Code: types.CodeOK}
}
