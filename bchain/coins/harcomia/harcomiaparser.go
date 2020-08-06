package harcomia

import (
	"github.com/scryptachain/blockbook-scrypta/bchain"
	"github.com/scryptachain/blockbook-scrypta/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0x0c01acc1
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{100}
	MainNetParams.ScriptHashAddrID = []byte{87}
}

type HarcomiaParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewHarcomiaParser(params *chaincfg.Params, c *btc.Configuration) *HarcomiaParser {
	return &HarcomiaParser{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

func (p *HarcomiaParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

func (p *HarcomiaParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
