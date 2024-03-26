package datasets

import (
	"github.com/KenshiTech/unchained/constants"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

func (m *Correctness) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}

func (m *BroadcastCorrectnessPacket) Protobuf() ([]byte, error) {
	protoModel, err := proto.Marshal(m)
	if err != nil {
		log.Err(err)
		return nil, constants.ErrInternalError
	}

	return protoModel, nil
}
