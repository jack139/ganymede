package client

import (
	"crypto/elliptic"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tjfoc/gmsm/sm2"
)

// AminoCdc amino编码类
var AminoCdc = codec.NewLegacyAmino()

func init() {
	AminoCdc.RegisterConcrete(sm2.PublicKey{}, "sm2/pubkey", nil)
	AminoCdc.RegisterConcrete(sm2.PrivateKey{}, "sm2/privkey", nil)
	AminoCdc.RegisterInterface((*elliptic.Curve)(nil), nil)
}
