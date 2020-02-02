package dbinstance

import (
	"fmt"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/go-ves/config"
	"github.com/Myriad-Dreamin/go-ves/lib/database/index"
	multi_index "github.com/Myriad-Dreamin/go-ves/lib/database/multi_index"
	"github.com/Myriad-Dreamin/go-ves/types"
	chain_dns "github.com/Myriad-Dreamin/go-ves/types/chain-dns"
	vesdb "github.com/Myriad-Dreamin/go-ves/types/database"
	"github.com/Myriad-Dreamin/go-ves/types/kvdb"
	"github.com/Myriad-Dreamin/go-ves/types/session"
	"github.com/Myriad-Dreamin/go-ves/types/storage-handler"
	//"github.com/Myriad-Dreamin/go-ves/types/user"
)

func XORMMigrate(muldb types.MultiIndex) (err error) {
	var xorm_muldb = muldb.(*multi_index.XORMMultiIndexImpl)
	//err = xorm_muldb.Register(&user.XORMUserAdapter{})
	//if err != nil {
	//	return
	//}
	err = xorm_muldb.Register(&session.SerialSession{})
	if err != nil {
		return
	}
	return nil
}

func MakeDB() model.VESDB {

	var db = new(vesdb.Database)
	var err error

	//TODO: SetEnv
	var muldb *multi_index.XORMMultiIndexImpl
	muldb, err = multi_index.GetXORMMultiIndex("mysql", "ves-admin:12345678@tcp(127.0.0.1:3306)/ves?charset=utf8")
	if err != nil {
		panic(fmt.Errorf("failed to get muldb: %v", err))
	}
	err = XORMMigrate(muldb)
	if err != nil {
		panic(fmt.Errorf("failed to migrate: %v", err))
	}

	var sindb *index.LevelDBIndex
	sindb, err = index.GetIndex("./index_data")
	if err != nil {
		panic(fmt.Errorf("failed to get sindb: %v", err))
	}

	db.SetIndex(sindb)
	db.SetMultiIndex(muldb)

	db.SetSessionBase(new(session.SerialSessionBase))
	db.SetSessionKVBase(new(kvdb.Database))
	db.SetStorageHandler(new(storage_handler.Database))
	db.SetChainDNS(chain_dns.NewDatabase(config.GetHostMap()))
	return db
}
