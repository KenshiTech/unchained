package net

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/KenshiTech/unchained/config"
	"github.com/KenshiTech/unchained/constants"
	"github.com/KenshiTech/unchained/constants/opcodes"
	"github.com/KenshiTech/unchained/crypto/bls"
	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/kosk"
	"github.com/KenshiTech/unchained/log"
	"github.com/KenshiTech/unchained/net/repository"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/vmihailenco/msgpack/v5"
)

var challenges *xsync.MapOf[*websocket.Conn, kosk.Challenge]
var signers *xsync.MapOf[*websocket.Conn, bls.Signer]
var upgrader = websocket.Upgrader{} // use default options

func processKosk(conn *websocket.Conn, payload []byte) error {
	var challenge kosk.Challenge
	err := msgpack.Unmarshal(payload, &challenge)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return constants.ErrInvalidPacket
	}

	signer, ok := signers.Load(conn)
	if !ok {
		return constants.ErrMissingHello
	}

	challenge.Passed, err = kosk.VerifyChallenge(challenge.Random, signer.PublicKey, challenge.Signature)
	if err != nil {
		return constants.ErrInvalidKosk
	}
	if !challenge.Passed {
		log.Logger.Error("challenge is Passed")
		return constants.ErrInvalidKosk
	}

	challenges.Store(conn, challenge)

	return nil
}

func processHello(conn *websocket.Conn, payload []byte) ([]byte, error) {
	var signer bls.Signer
	err := msgpack.Unmarshal(payload, &signer)
	if err != nil {
		log.Logger.Error("Cant unmarshal packet: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	if signer.Name == "" {
		log.Logger.Error("Signer name is empty Or public key is invalid")
		return []byte{}, constants.ErrInvalidConfig
	}

	signers.Range(func(conn *websocket.Conn, signerInMap bls.Signer) bool {
		publicKeyInUse := signerInMap.PublicKey == signer.PublicKey
		if publicKeyInUse {
			Close(conn)
		}
		return !publicKeyInUse
	})

	signers.Store(conn, signer)

	// Start KOSK verification
	challenge := kosk.Challenge{Random: kosk.NewChallenge()}
	challenges.Store(conn, challenge)
	koskPayload, err := msgpack.Marshal(challenge)
	if err != nil {
		log.Logger.Error("Cant marshal challenge: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	return koskPayload, nil
}

func checkPublicKey(conn *websocket.Conn) (*bls.Signer, error) {
	challenge, ok := challenges.Load(conn)
	if !ok || !challenge.Passed {
		return nil, constants.ErrMissingKosk
	}

	signer, ok := signers.Load(conn)
	if !ok {
		return nil, constants.ErrMissingHello
	}

	return &signer, nil
}

// TODO: Can we use any part of this?
func processPriceReport(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	var report datasets.PriceReport
	err = msgpack.Unmarshal(payload, &report)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	toHash, err := msgpack.Marshal(&report.PriceInfo)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	hash, err := bls.Hash(toHash)
	if err != nil {
		log.Logger.Error("Cant hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Cant recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Cant recover bls pub-key: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.Error("Cant recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	priceInfo := datasets.BroadcastPricePacket{
		Info:      report.PriceInfo,
		Signature: report.Signature,
		Signer:    *signer,
	}

	priceInfoByte, err := msgpack.Marshal(&priceInfo)
	// TODO: Handle this error properly
	// TODO: Maybe notify the peer so they can resend
	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Cannot marshal the broadcast packet")
		return []byte{}, constants.ErrInternalError
	}

	return priceInfoByte, nil
}

func processEventLog(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	var report datasets.EventLogReport
	err = msgpack.Unmarshal(payload, &report)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	toHash, err := msgpack.Marshal(&report.EventLog)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	hash, err := bls.Hash(toHash)
	if err != nil {
		log.Logger.Error("Cant hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Cant recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Cant recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.Error("Cant recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	broadcastPacket := datasets.BroadcastEventPacket{
		Info:      report.EventLog,
		Signature: report.Signature,
		Signer:    *signer,
	}

	broadcastPayload, err := msgpack.Marshal(&broadcastPacket)
	// TODO: Handle this error properly
	// TODO: Maybe notify the peer so they can resend
	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Cannot marshal the broadcast packet")

		return []byte{}, constants.ErrInternalError
	}

	return broadcastPayload, nil
}

func processCorrectnessRecord(conn *websocket.Conn, payload []byte) ([]byte, error) {
	signer, err := checkPublicKey(conn)
	if err != nil {
		return []byte{}, err
	}

	var report datasets.CorrectnessReport
	err = msgpack.Unmarshal(payload, &report)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	toHash, err := msgpack.Marshal(&report.Correctness)
	if err != nil {
		log.Logger.Error("Cant unmarshal Msgpack: %v", err)
		return []byte{}, constants.ErrInvalidPacket
	}

	hash, err := bls.Hash(toHash)
	if err != nil {
		log.Logger.Error("Cant hash bls: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	signature, err := bls.RecoverSignature(report.Signature)
	if err != nil {
		log.Logger.Error("Cant recover bls signature: %v", err)
		return []byte{}, constants.ErrInternalError
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)
	if err != nil {
		log.Logger.Error("Cant recover bls pub-key: %v", err)
		return []byte{}, constants.ErrCantVerifyBls
	}

	ok, err := bls.Verify(signature, hash, pk)
	if err != nil {
		log.Logger.With("Error", err).Error("Cant verify bls")
		return []byte{}, constants.ErrCantVerifyBls
	}
	if !ok {
		return []byte{}, constants.ErrInvalidSignature
	}

	broadcastPacket := datasets.BroadcastCorrectnessPacket{
		Info:      report.Correctness,
		Signature: report.Signature,
		Signer:    *signer,
	}

	broadcastPayload, err := msgpack.Marshal(&broadcastPacket)

	// TODO: Handle this error properly
	// TODO: Maybe notify the peer so they can resend
	if err != nil {
		log.Logger.
			With("Error", err).
			Error("Cannot marshal the broadcast packet")

		return []byte{}, constants.ErrInternalError
	}

	return broadcastPayload, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Error("cant upgrade connection: %v", err)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Logger.Error("cant close connection: %v", err)
		}
	}(conn)

	defer signers.Delete(conn)
	defer challenges.Delete(conn)
	defer repository.Consumers.Delete(conn)
	defer repository.BroadcastMutex.Delete(conn)

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Logger.Error("cant read message: %v", err)
			break
		}

		switch opcodes.OpCode(payload[0]) {
		// TODO: Make a table of call codes
		case opcodes.Hello:
			result, err := processHello(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			SendMessage(conn, messageType, opcodes.Feedback, "conf.ok")
			Send(conn, messageType, opcodes.KoskChallenge, result)
		case opcodes.PriceReport:
			// TODO: Maybe this is unnecessary
			if payload[1] != 0 {
				SendError(conn, messageType, opcodes.Feedback, constants.ErrNotSupportedDataset)
				continue
			}

			result, err := processPriceReport(conn, payload[2:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.PriceReportBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")
		case opcodes.EventLog:
			// TODO: Maybe this is unnecessary
			if payload[1] != 0 {
				SendError(conn, messageType, opcodes.Feedback, constants.ErrNotSupportedDataset)
				continue
			}

			result, err := processEventLog(conn, payload[2:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.EventLogBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")
		case opcodes.CorrectnessReport:
			result, err := processCorrectnessRecord(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}

			BroadcastPayload(opcodes.CorrectnessReportBroadcast, result)
			SendMessage(conn, messageType, opcodes.Feedback, "signature.accepted")
		case opcodes.KoskResult:
			err := processKosk(conn, payload[1:])
			if err != nil {
				SendError(conn, messageType, opcodes.Error, err)
			}
			SendMessage(conn, messageType, opcodes.Feedback, "kosk.ok")
		case opcodes.RegisterConsumer:
			// TODO: Consumers must specify what they're subscribing to
			repository.Consumers.Store(conn, true)
			repository.BroadcastMutex.Store(conn, new(sync.Mutex))
		default:
			SendError(conn, messageType, opcodes.Error, constants.ErrNotSupportedInstruction)
		}
	}
}

func StartServer() {
	flag.Parse()
	versionedRoot := fmt.Sprintf("/%s", constants.ProtocolVersion)
	http.HandleFunc(versionedRoot, rootHandler)
	addr := config.Config.GetString("broker.bind")

	readHeaderTimeoutInSecond := 3
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutInSecond) * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func init() {
	signers = xsync.NewMapOf[*websocket.Conn, bls.Signer]()
	challenges = xsync.NewMapOf[*websocket.Conn, kosk.Challenge]()
}
