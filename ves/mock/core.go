package mock
//go:generate mockgen -source ../control/model-interface.go -mock_names "SessionKV=SessionKV,StorageHandler=StorageHandler" -package mock -destination model-interface.go
//go:generate mockgen -source ../control/gen-model-interface.go -mock_names "SessionDBI=SessionDB,SessionFSetI=SessionFSet,SessionAccountDBI=SessionAccountDB" -package mock -destination gen-model-interface.go
//go:generate mockgen -source ../control/external-interface.go -mock_names "NSBClient=NSBClient,ChainDNS=ChainDNS,CentralVESClient=CentralVESClient" -package mock -destination external-interface.go
//go:generate mockgen -source ../control/gen-external-interface.go -mock_names "BlockChainInterfaceI=BlockChainInterface,OpIntentInitializerI=OpIntentInitializer" -package mock -destination gen-external-interface.go
