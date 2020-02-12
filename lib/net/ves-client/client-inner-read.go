package vesclient

func (vc *VesClient) read() {
	for {
		message, messageID, err := vc.ReadMessage()
		if err != nil {
			vc.logger.Error("read message error", err)
			continue
		}

		vc.processMessage(message, messageID)
	}
}
