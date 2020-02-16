package mock

//go:generate mkdir -p mock-control
//go:generate mockgen -source ../control/model-interface.go -mock_names "SessionDBI=SessionDB,SessionFSetI=SessionFSet,SessionAccountDBI=SessionAccountDB" -package mock -destination model-interface.go
//go:generate mockgen -source ../control/external-interface.go -mock_names "NSBClient=NSBClient,OpIntentInitializerI=OpIntentInitializer" -package mock -destination external-interface.go


