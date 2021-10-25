package sessionservice

//type MerkleProofReceiveService struct {
//	*vs.VServer
//	context.Context
//	*uiprpc.MerkleProofReceiveRequest
//}
//
//func NewMerkleProofReceiveService(server *vs.VServer, context context.Context, merkleProofReceiveRequest *uiprpc.MerkleProofReceiveRequest) MerkleProofReceiveService {
//	return MerkleProofReceiveService{VServer: server, Context: context, MerkleProofReceiveRequest: merkleProofReceiveRequest}
//}
//
//func (s MerkleProofReceiveService) Serve() (*uiprpc.MerkleProofReceiveReply, error) {
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
//		s.Logger.Error("set kv error", "sid", hex.EncodeToString(ses.GetGUID()), "err", err)
//		return nil, err
//	}
//
//	return &uiprpc.MerkleProofReceiveReply{
//		Ok: true,
//	}, nil
//}
