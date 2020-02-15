package router


//func GetHTTPServeMux(server *Server) *gin.Engine {
//	r := gin.Default()
//
//	r.Use(cors.New(cors.Config{
//		AllowOriginFunc:  func(origin string) bool { return true },
//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
//		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Grpc-Web", "X-User-Agent"},
//		ExposeHeaders:    []string{"grpc-status", "grpc-message"},
//		AllowCredentials: true,
//		MaxAge:           12 * time.Hour,
//	}))
//
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	vGroup := r.Group("/uiprpc.VES")
//
//	tieMethod(
//		vGroup, "UserRegister",
//		func() proto.Message { return new(uiprpc.UserRegisterRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.UserRegister(context.Background(), in.(*uiprpc.UserRegisterRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "SessionStart",
//		func() proto.Message { return new(uiprpc.SessionStartRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.SessionStart(context.Background(), in.(*uiprpc.SessionStartRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "SessionAckForInit",
//		func() proto.Message { return new(uiprpc.SessionAckForInitRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.SessionAckForInit(context.Background(), in.(*uiprpc.SessionAckForInitRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "SessionRequireTransact",
//		func() proto.Message { return new(uiprpc.SessionRequireTransactRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.SessionRequireTransact(context.Background(), in.(*uiprpc.SessionRequireTransactRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "SessionRequireRawTransact",
//		func() proto.Message { return new(uiprpc.SessionRequireRawTransactRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.SessionRequireRawTransact(context.Background(), in.(*uiprpc.SessionRequireRawTransactRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "AttestationReceive",
//		func() proto.Message { return new(uiprpc.AttestationReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.AttestationReceive(context.Background(), in.(*uiprpc.AttestationReceiveRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "MerkleProofReceive",
//		func() proto.Message { return new(uiprpc.MerkleProofReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.MerkleProofReceive(context.Background(), in.(*uiprpc.MerkleProofReceiveRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "ShrotenMerkleProofReceive",
//		func() proto.Message { return new(uiprpc.ShortenMerkleProofReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.ShrotenMerkleProofReceive(context.Background(), in.(*uiprpc.ShortenMerkleProofReceiveRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "InformMerkleProof",
//		func() proto.Message { return new(uiprpc.MerkleProofReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.InformMerkleProof(context.Background(), in.(*uiprpc.MerkleProofReceiveRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "InformShortenMerkleProof",
//		func() proto.Message { return new(uiprpc.ShortenMerkleProofReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.InformShortenMerkleProof(context.Background(), in.(*uiprpc.ShortenMerkleProofReceiveRequest))
//		},
//	)
//
//	tieMethod(
//		vGroup, "InformAttestation",
//		func() proto.Message { return new(uiprpc.AttestationReceiveRequest) },
//		func(ctx context.Context, in proto.Message) (proto.Message, error) {
//			return server.InformAttestation(context.Background(), in.(*uiprpc.AttestationReceiveRequest))
//		},
//	)
//
//	return r
//}
//
//func tieMethod(
//	vGroup *gin.RouterGroup, method string,
//	objectFactory func() proto.Message,
//	serveFunc func(ctx context.Context, in proto.Message) (proto.Message, error),
//) {
//	vGroup.OPTIONS(method, func(c *gin.Context) {
//		c.Status(200)
//	})
//
//	vGroup.POST(method, func(c *gin.Context) {
//		p, err := c.GetRawData()
//		if err != nil && err != io.EOF {
//			c.AbortWithError(400, err)
//			return
//		}
//		var in = objectFactory()
//		err = decode(p, in)
//		if err != nil {
//			c.AbortWithError(400, err)
//			return
//		}
//
//		ret, err := serveFunc(context.Background(), in)
//		if err != nil {
//			c.AbortWithError(400, err)
//			return
//		}
//		b, err := encode(c, ret)
//		if err != nil {
//			c.AbortWithError(500, err)
//			return
//		}
//		c.Data(200, "application/grpc-web-text", b)
//		// fmt.Println(c.Request.Method)
//		// fmt.Println(c.Request.URL)
//		// fmt.Println(c.Request.Proto)
//		// fmt.Println(c.Request.ProtoMajor)
//		// fmt.Println(c.Request.Header)
//		// fmt.Println(c.Request.ProtoMajor)
//	})
//}
//
//// now only decode one message
//func decode(b []byte, in proto.Message) error {
//	l := base64.StdEncoding.DecodedLen(len(b))
//	if l < 5 {
//		return errors.New("too short")
//	}
//	g := make([]byte, l)
//	_, err := base64.StdEncoding.Decode(g, b)
//	if err != nil {
//		return err
//	}
//
//	if g[0] != 0 {
//		return errors.New("needing data frame")
//	}
//
//	var gg = binary.BigEndian.Uint32(g[1:5])
//	if len(g) != 5+int(gg) {
//		return errors.New("malformed")
//	}
//	return proto.Unmarshal(g[5:5+gg], in)
//}
//
//func encode(c *gin.Context, ret proto.Message) (t []byte, err error) {
//	s, err := proto.Marshal(ret)
//	if err != nil {
//		return nil, err
//	}
//	t = make([]byte, len(s)+5)
//	t[0] = 0
//	binary.BigEndian.PutUint32(t[1:5], uint32(len(s)))
//	copy(t[5:], s)
//	l := base64.StdEncoding.EncodedLen(len(t))
//	g := make([]byte, l)
//	base64.StdEncoding.Encode(g, t)
//	return g, nil
//}
