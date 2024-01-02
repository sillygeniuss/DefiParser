package balance

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"walletForCurve/node"
)

func GetAddressBalance(ctx context.Context, client node.LocalClient, address string) (string, error) {
	hexAddr := common.HexToAddress(address)
	balance, err := client.BalanceAt(ctx, hexAddr, nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}
