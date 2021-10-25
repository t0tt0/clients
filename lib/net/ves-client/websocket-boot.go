package vesclient

func (vc *VesClient) Boot() (err error) {
	go vc.conn.ReadRoutine()
	if err = vc.SayClientHello(vc.name); err != nil {
		return
	}
	return
}
