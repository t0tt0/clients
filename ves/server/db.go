package server

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/HyperService-Consortium/go-ves/ves/model/index"
	"github.com/Myriad-Dreamin/functional-go"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/minimum-lib/rbac"
	"reflect"
)

type dbResult struct {
	dbName string
	functional.DecayResult
}

func (srv *Server) registerDatabaseService() bool {

	for _, dbResult := range []dbResult{
		{"transactionDB", functional.Decay(model.NewTransactionDB(srv.Module))},
		{"sessionAccountDB", functional.Decay(model.NewSessionAccountDB(srv.Module))},
		{"sessionDB", functional.Decay(model.NewSessionDB(srv.Module))},
		{"objectDB", functional.Decay(model.NewObjectDB(srv.Module))},
	} {
		if dbResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("init %T DB error", dbResult.First), "error", dbResult.Err)
			return false
		}
		srv.ModelProvider.Register(dbResult.dbName, dbResult.First)
	}
	return true
}

func (srv *Server) PrepareDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	if !model.Install(srv.Module) {
		return false
	}

	//srv.RedisPool, err = model.OpenRedis(cfg)
	//if err != nil {
	//	srv.Logger.Debug("create redis pool error", "error", err)
	//	return false
	//}
	//
	//srv.Logger.Info("connected to redis",
	//	"connection-type", cfg.RedisConfig.ConnectionType,
	//	"host", cfg.RedisConfig.Host,
	//	"connection-timeout", cfg.RedisConfig.ConnectionTimeout,
	//	"database", cfg.RedisConfig.Database,
	//	"read-timeout", cfg.RedisConfig.ReadTimeout,
	//	"write-timeout", cfg.RedisConfig.WriteTimeout,
	//	"idle-timeout", cfg.RedisConfig.IdleTimeout,
	//	"wait", cfg.RedisConfig.Wait,
	//	"max-active", cfg.RedisConfig.MaxActive,
	//	"max-idle", cfg.RedisConfig.MaxIdle,
	//)
	//err = model.RegisterRedis(srv.RedisPool, srv.Logger)
	//if err != nil {
	//	srv.Logger.Debug("register redis error", "error", err)
	//	return false
	//}
	err := rbac.InitGorm(
		srv.Module.Require(config.ModulePath.Minimum.DBInstance.GormDB).(*model.GormDB),
	)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())

	return srv.initIndex() && srv.registerDatabaseService()
}

func (srv *Server) MockDatabase() bool {
	srv.Cfg.DatabaseConfig.Debug(srv.Logger)

	if !model.InstallMock(srv.Module, func(dep module.Module, s sqlmock.Sqlmock) error {
		s.ExpectExec(`CREATE TABLE "transaction"`).WillReturnResult(
			sqlmock.NewResult(0, 1))
		s.ExpectExec(`CREATE TABLE "session_account"`).WillReturnResult(
			sqlmock.NewResult(0, 1))
		s.ExpectExec(`CREATE UNIQUE INDEX sa_sca`).WillReturnResult(
			sqlmock.NewResult(0, 1))
		s.ExpectExec(`CREATE TABLE "session"`).WillReturnResult(
			sqlmock.NewResult(0, 1))
		return nil
	}) {
		return false
	}

	err := rbac.InitGorm(
		srv.Module.Require(config.ModulePath.Minimum.DBInstance.GormDB).(*model.GormDB),
	)
	if err != nil {
		srv.Logger.Debug("rbac to database error", "error", err)
		return false
	}
	srv.ModelProvider.Register("enforcer", rbac.GetEnforcer())
	return srv.initIndex() && srv.registerDatabaseService()
}

func (srv *Server) initIndex() bool {
	var err error

	srv.levelDB, err = model.NewLevelDB(srv.Cfg.LevelDBConfig.LocalPath, nil)
	if err != nil {
		srv.Logger.Debug("create oss engine error", "error", err)
		return false
	}
	err = model.RegisterIndex(srv.levelDB)
	if err != nil {
		srv.Logger.Debug("register oss engine error", "error", err)
		return false
	}

	if i := index.ToIndex(srv.levelDB); i == nil {
		srv.Logger.Debug("oss engine cannot convert to index", "type", reflect.TypeOf(srv.levelDB))
	} else {
		srv.Module.Provide(config.ModulePath.DBInstance.Index, i)
		srv.Module.Provide(config.ModulePath.Global.Storage, index.NewSessionKV(i))
		srv.Module.Provide(config.ModulePath.Global.StorageHandler, index.NewStorageHandler(i))
	}

	return true
}
