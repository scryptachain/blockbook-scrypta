package scrypta

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/juju/errors"
	"github.com/scryptachain/blockbook-scrypta/bchain"
	"github.com/scryptachain/blockbook-scrypta/bchain/coins/btc"
)

type ScryptaRPC struct {
	*btc.BitcoinRPC
}

// sendrawtransaction

type CmdMasternodeList struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type ResMasternodeList struct {
	Error  *bchain.RPCError `json:"error"`
	Result string           `json:"result"`
}

const firstBlockWithSpecialTransactions = 454000

func NewScryptaRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}
	s := &ScryptaRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateSmartFee = false
	return s, nil
}

func (b *ScryptaRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain
	params := GetChainParams(chainName)
	b.Parser = NewScryptaParser(params, b.ChainConfig)
	b.Testnet = false
	b.Network = "livenet"
	glog.Info("rpc: block chain ", params.Name)
	return nil
}

func (b *ScryptaRPC) GetBlock(hash string, height uint32) (*bchain.Block, error) {
	if hash == "" && height < firstBlockWithSpecialTransactions {
		return b.BitcoinRPC.GetBlock(hash, height)
	}
	var err error
	if hash == "" && height > 0 {
		hash, err = b.GetBlockHash(height)
		if err != nil {
			return nil, err
		}
	}
	glog.V(1).Info("rpc: getblock (verbosity=1) ", hash)
	res := btc.ResGetBlockThin{}
	req := btc.CmdGetBlock{Method: "getblock"}
	req.Params.BlockHash = hash
	req.Params.Verbosity = 1
	err = b.Call(&req, &res)
	if err != nil {
		return nil, errors.Annotatef(err, "hash %v", hash)
	}
	if res.Error != nil {
		return nil, errors.Annotatef(res.Error, "hash %v", hash)
	}
	txs := make([]bchain.Tx, 0, len(res.Result.Txids))
	for _, txid := range res.Result.Txids {
		tx, err := b.GetTransaction(txid)
		if err != nil {
			if err == bchain.ErrTxNotFound {
				glog.Errorf("rpc: getblock: skipping transanction in block %s due error: %s", hash, err)
				continue
			}
			return nil, err
		}
		txs = append(txs, *tx)
	}
	block := &bchain.Block{
		BlockHeader: res.Result.BlockHeader,
		Txs:         txs,
	}
	return block, nil
}

func (b *ScryptaRPC) GetTransactionForMempool(txid string) (*bchain.Tx, error) {
	return b.GetTransaction(txid)
}

// MasternodeList returns masternodes list
func (b *ScryptaRPC) MasternodeList() (string, error) {
	glog.V(1).Info("rpc: masternodelist")

	res := ResMasternodeList{}
	req := CmdMasternodeList{Method: "masternodelist"}
	err := b.Call(&req, &res)

	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}

	return res.Result, nil
}
