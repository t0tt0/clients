package vesclient

func (vc *VesClient) Boot() (err error) {
	//if err = vc.load(dataPrefix + "/" + string(vc.name)); err != nil {
	//	vc.logger.Error("load config from filepath error", "path", dataPrefix+"/"+string(vc.name), "error", err)
	//	return
	//}
	//phandler.register(vc.save)

	go vc.read()
	if err = vc.SayClientHello(vc.name); err != nil {
		return
	}
	return
}
