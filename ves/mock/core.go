package mock

//go:generate mockgen -source ../control/model-interface.go -mock_names "SessionKV=SessionKV,StorageHandler=StorageHandler" -package mock -destination model-interface.go
//go:generate mockgen -source ../model/internal/abstraction/session.go -mock_names "SessionDB=SessionDB" -package mock -destination model-session.go
//go:generate mockgen -source ../model/fset.go -mock_names "SessionFSetI=SessionFSet" -package mock -destination model-fset.go
//go:generate mockgen -source ../model/internal/abstraction/session-account.go -mock_names "SessionAccountDB=SessionAccountDB" -package mock -destination model-session-account.go
//go:generate mockgen -source ../control/external-interface.go -mock_names "NSBClient=NSBClient,ChainDNS=ChainDNS,CentralVESClient=CentralVESClient" -package mock -destination external-interface.go
//go:generate mockgen -source ../control/gen-external-interface.go -mock_names "BlockChainInterfaceI=BlockChainInterface,InitializerI=OpIntentInitializer" -package mock -destination gen-external-interface.go
