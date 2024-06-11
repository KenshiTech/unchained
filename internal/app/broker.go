package app

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/scheduler"
	"github.com/TimeleapLabs/unchained/internal/service/frost"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	"github.com/TimeleapLabs/unchained/internal/transport/server"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket"
	"github.com/TimeleapLabs/unchained/internal/transport/server/websocket/store"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Broker starts the Unchained broker and contains its DI.
func Broker() {
	utils.Logger.
		With("Mode", "Broker").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithBlsIdentity(),
		crypto.WithEvmSigner(),
	)

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)

	//redisIns := redis.New()

	nativeSignerRepo := store.New()
	//redisSignerRepo := store.NewRedisStore(redisIns, nativeSignerRepo)

	frostService := frost.New(pos)

	scheduler := scheduler.New(
		scheduler.WithFrostEvents(frostService),
	)

	scheduler.Start(false)

	server.New(
		websocket.WithWebsocket(nativeSignerRepo),
	)
}
