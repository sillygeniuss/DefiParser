package node

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcAddr = "https://eth.public-rpc.com"
	ctx     = context.Background()
)

type LocalClient struct {
	*ethclient.Client
}

type Block struct {
	*types.Block
}

func NewLocalClient() (LocalClient, error) {
	client, err := ethclient.Dial(rpcAddr)
	if err != nil {
		return LocalClient{}, err
	}
	return LocalClient{client}, nil
}

func (c LocalClient) GetCurrentBlock() (*Block, error) {
	block, e := c.BlockByNumber(ctx, nil)
	if e != nil {
		return nil, e
	}
	return &Block{block}, nil
}
