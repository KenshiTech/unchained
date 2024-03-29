package datasets

import (
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type Signer struct {
	Name           string
	EvmWallet      string
	PublicKey      [96]byte
	ShortPublicKey [48]byte
}

type Signature struct {
	Signature bls12381.G1Affine
	Signer    Signer
	Processed bool
}

func (s *Signature) Sia() *sia.Sia {
	return new(sia.Sia).
		AddByteArray8(s.Signature.Marshal()).
		EmbedSia(s.Signer.Sia()).
		AddBool(s.Processed)
}

func (s *Signature) DeSia(sia *sia.Sia) *Signature {
	err := s.Signature.Unmarshal(sia.ReadByteArray8())

	if err != nil {
		s.Signature = bls12381.G1Affine{}
	}

	s.Signer.DeSia(sia)
	s.Processed = sia.ReadBool()

	return s
}

func (s *Signer) Sia() *sia.Sia {
	return new(sia.Sia).
		AddString8(s.Name).
		AddString8(s.EvmWallet).
		AddByteArray8(s.PublicKey[:]).
		AddByteArray8(s.ShortPublicKey[:])
}

func (s *Signer) DeSia(sia *sia.Sia) *Signer {
	s.Name = sia.ReadString8()
	s.EvmWallet = sia.ReadString8()
	copy(s.PublicKey[:], sia.ReadByteArray8())
	copy(s.ShortPublicKey[:], sia.ReadByteArray8())

	return s
}
