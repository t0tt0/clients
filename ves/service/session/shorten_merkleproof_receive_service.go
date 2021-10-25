package sessionservice

//type ShortenMerkleProofReceiveService struct {
//	*vs.VServer
//	context.Context
//	*uiprpc.ShortenMerkleProofReceiveRequest
//}
//
//func NewShortenMerkleProofReceiveService(server *vs.VServer, context context.Context, shortenMerkleProofReceiveRequest *uiprpc.ShortenMerkleProofReceiveRequest) ShortenMerkleProofReceiveService {
//	return ShortenMerkleProofReceiveService{VServer: server, Context: context, ShortenMerkleProofReceiveRequest: shortenMerkleProofReceiveRequest}
//}
//
//func (s ShortenMerkleProofReceiveService) Serve() (*uiprpc.ShortenMerkleProofReceiveReply, error) {
//	s.DB.ActivateSession(s.GetSessionId())
//	ses, err := s.DB.FindSessionInfo(s.GetSessionId())
//	if err != nil {
//		s.DB.InactivateSession(s.GetSessionId())
//		s.Logger.Error("find session info", "sid", hex.EncodeToString(s.GetSessionId()), "err", err)
//		return nil, err
//	}
//
//	defer func() {
//		err = s.DB.UpdateSessionInfo(ses)
//		if err != nil {
//			s.Logger.Error("update failed", "sid", hex.EncodeToString(ses.GetGUID()), "err", err)
//		}
//		s.DB.InactivateSession(s.GetSessionId())
//	}()
//
//	ses.SetSigner(s.Signer)
//
//	var merkle = s.GetMerkleproof()
//
//	// to do: verify merkle proof
//
//	err = s.DB.SetKV(
//		ses.GetGUID(),
//		merkle.GetKey(),
//		merkle.GetValue(),
//	)
//
//	if err != nil {
//		s.Logger.Error("error", "error", err)
//		return nil, err
//	}
//
//	return &uiprpc.ShortenMerkleProofReceiveReply{
//		Ok: true,
//	}, nil
//
//}
