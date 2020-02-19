package dblayer

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"go.uber.org/zap/zapcore"
	"log"
	"testing"
)

var dep = make(module.Module)

func TestMain(m *testing.M) {
	l, err := logger.NewZapLogger(logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		log.Fatal(err)
		return
	}
	dep.Provide(config.ModulePath.Minimum.Global.Logger, l)
	dep.Provide(config.ModulePath.Minimum.Global.Configuration, &config.ServerConfig{
		DatabaseConfig: config.DatabaseConfig{
			Escaper: `"`,
		},
	})
	InstallMock(dep, mockExpectation)
	m.Run()
}

func mockExpectation(dep module.Module, s sqlmock.Sqlmock) error {
	s.ExpectExec(`CREATE TABLE "transaction"`).WillReturnResult(
		sqlmock.NewResult(0, 1))
	s.ExpectExec(`CREATE TABLE "session_account"`).WillReturnResult(
		sqlmock.NewResult(0, 1))
	s.ExpectExec(`CREATE UNIQUE INDEX sa_sca`).WillReturnResult(
		sqlmock.NewResult(0, 1))
	s.ExpectExec(`CREATE TABLE "session"`).WillReturnResult(
		sqlmock.NewResult(0, 1))
	return nil
}
