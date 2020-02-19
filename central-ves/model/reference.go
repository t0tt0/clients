package model

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
	"github.com/jinzhu/gorm"
)

type Enforcer = database.Enforcer
type ORMObject = dorm.ORMObject
type GormDB = gorm.DB
