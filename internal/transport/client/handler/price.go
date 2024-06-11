package handler

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/model"
)

func (h *consumer) PriceReport(ctx context.Context, message []byte) {
	packet := new(model.BroadcastPricePacket).FromBytes(message)

	priceInfoHash, err := packet.Info.Bls()
	if err != nil {
		return
	}

	err = h.uniswap.RecordSignature(
		ctx,
		packet.Signature[:],
		packet.Signer,
		priceInfoHash,
		packet.Info,
		true,
		false,
	)
	if err != nil {
		return
	}
}

func (w *worker) PriceReport(_ context.Context, _ []byte) {}
