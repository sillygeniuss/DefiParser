package tool

import (
	"math"
	"math/big"
)

func DecConverter(aim, decimal *big.Int) *big.Float {
	fAim := new(big.Float).SetInt(aim)
	return new(big.Float).Quo(fAim, big.NewFloat(math.Pow10(int(decimal.Int64()))))
}

func DecConverter2(aim *big.Float, decimal *big.Int) *big.Float {
	return new(big.Float).Quo(aim, big.NewFloat(math.Pow10(int(decimal.Int64()))))
}
