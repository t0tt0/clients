package mock

//go:generate mkdir -p mock-control
//go:generate mockgen -source ../control/model-interface.go -mock_names "SessionKV=SessionKV,StorageHandler=StorageHandler" -package mock -destination model-interface.go
//go:generate mockgen -source ../control/gen-model-interface.go -mock_names "SessionDBI=SessionDB,SessionFSetI=SessionFSet,SessionAccountDBI=SessionAccountDB" -package mock -destination model-interface.go
//go:generate mockgen -source ../control/external-interface.go -mock_names "NSBClient=NSBClient,ChainDNS=ChainDNS,OpIntentInitializerI=OpIntentInitializer,CentralVESClient=CentralVESClient" -package mock -destination external-interface.go
