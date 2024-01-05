package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"walletForCurve/curve/registry/meta"
	"walletForCurve/node"
)

// registry address lit
var addressProviderRegistry = "0x0000000022D53366457F9d5E68Ec105046FC4383"

var metaRegistryAddress = "0xF98B45FA17DE75FB1aD0e7aFD971b0ca00e379fC"

// pool:stetch address list
var stETHPoolAddress = "0xdc24316b9ae028f1497c275eb9192a3ea0f67022"
var stETHTokenAddress = "0x06325440d014e39736583c165c2963ba99faf14e"
var stETHGaugeAddress = "0x182b723a58739a9c974cfdb385ceadb237453c28"

var (
	flagConfig string
)

func initArgs() {
	flag.StringVar(&flagConfig, "config", "./conf/config.yaml", "config file path")
	flag.Parse()
}

func init() {
	initArgs()
	LoadConfig(flagConfig)
}

func main() {
	// Create a new client
	//client, err := node.NewLocalClient()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(Config.Wallet)
	return
	// Get the latest block
	//block, err := client.GetCurrentBlock()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("current block number: %d\n", block.Number())

	//ctx := context.Background()
	//asset, e := balance.GetAddressBalance(ctx, client, curveWalletAddr)
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Printf("curve wallet balance: %s\n", asset)

	// Output:获取功能registry列表
	//addressProvider, e := address.NewAddressProvider(common.HexToAddress(addressProviderRegistry), client)
	//if e != nil {
	//	panic(e)
	//}
	//
	//// Output: 创建pool registry实例
	//poolRegistryAddr, _ := addressProvider.GetAddress(nil, big.NewInt(0))
	//fmt.Printf("main registry address: %s\n", poolRegistryAddr.Hex())
	//
	//poolIns, _ := pool.NewPool(poolRegistryAddr, client)
	//poolCount, _ := poolIns.PoolCount(nil)
	//// todo 这里poolCount拿到的值跟官方浏览器对不上，需要进一步确认
	//fmt.Printf("pool account: %d\n", poolCount.Int64())
	//
	///** ====================================================================================================================================== */
	//// 使用Meta Registry进行Curve交互
	//metaRegIns, e := InitMetaIns(client)
	//if e != nil {
	//	panic("init meta registry error: " + e.Error())
	//}
	//
	//// 获取pool名称
	//addr := "0xDC24316b9AE028F1497c275EB9192a3Ea0f67022"
	//poolName, e := metaRegIns.GetPoolName(nil, common.HexToAddress(addr))
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Printf("pool name: %s\n", poolName)
	//
	//// output wallet steCRV balance
	//tokenIns, _ := steth.NewToken(common.HexToAddress(stETHTokenAddress), client)
	////balance, _ := tokenIns.BalanceOf(nil, common.HexToAddress(curveWalletAddr))
	//
	//decimal, _ := tokenIns.Decimals(nil)
	//
	//fBalance := new(big.Float).SetInt(balance)
	//realBalance := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(int(decimal.Int64()))))
	//fmt.Printf("stETH real steCRV balance: %f\n", realBalance)
	//
	//// output lp token total supply
	//totalSupply, _ := tokenIns.TotalSupply(nil)
	//
	//// output: pool coin address list
	//poolCoinAddrList, _ := poolIns.GetCoins(nil, common.HexToAddress(addr))
	//fmt.Printf("pool coin address list: %s\n", poolCoinAddrList)
	//
	//// output pool coin info
	//poolCoinBalanceList, _ := poolIns.GetBalances(nil, common.HexToAddress(addr))
	//fmt.Printf("pool coin balance list: %s\n", poolCoinBalanceList)
	//
	////// 分别计算池中两种Token的单位比例
	//coin1PerRatio := new(big.Float).Quo(new(big.Float).SetInt(poolCoinBalanceList[0]), new(big.Float).SetInt(totalSupply))
	//fmt.Printf("coin1 per ratio: %f\n", coin1PerRatio)
	//
	//coin2PerRatio := new(big.Float).Quo(new(big.Float).SetInt(poolCoinBalanceList[1]), new(big.Float).SetInt(totalSupply))
	//fmt.Printf("coin2 per ratio: %f\n", coin2PerRatio)
	//
	//// 计算当前钱包中的lp Token总量下对应的各币种数量
	//coin1Balance := new(big.Float).Mul(coin1PerRatio, fBalance)
	//coin1RealBalance := new(big.Float).Quo(coin1Balance, big.NewFloat(math.Pow10(int(decimal.Int64()))))
	//fmt.Printf("coin1 real balance: %f\n", coin1RealBalance)
	//
	//coin2Balance := new(big.Float).Mul(coin2PerRatio, fBalance)
	//coin2RealBalance := new(big.Float).Quo(coin2Balance, big.NewFloat(math.Pow10(int(decimal.Int64()))))
	//fmt.Printf("coin1 real balance: %f\n", coin2RealBalance)

	//poolRatio := new(big.Float).Quo(new(big.Float).SetInt(totalSupply), new(big.Float).SetInt(new(big.Int).Add(coinBalanceList[0], coinBalanceList[1])))
	//fmt.Printf("pool ratio: %s\n", poolRatio.String())

	//coinTotalCount := new(big.Int).Add(coinBalanceList[0], coinBalanceList[1])
	//fmt.Printf("total supply: %s\n", totalSupply.String())
	//fmt.Printf("coin total: %s\n", new(big.Int).Add(coinBalanceList[0], coinBalanceList[1]).String())

	//fmt.Printf("coin2 total: %d\n", new(big.Int).Sub(coinTotalCount, coinBalanceList[0]))
	//current := new(big.Float).Mul(poolRatio, new(big.Float).SetInt(balance))
	//
	//coin1 := new(big.Float).Mul(new(big.Float).SetInt(coinBalanceList[0]), current)
	//coin2 := new(big.Float).Mul(new(big.Float).SetInt(coinBalanceList[1]), current)
	//
	//fmt.Printf("coin1: %s, coin2: %s\n", tool.DecConverter2(coin1, decimal).String(), tool.DecConverter2(coin2, decimal).String())
	//coinAddrList, _ := poolIns.GetCoins(nil, common.HexToAddress(stETHPoolAddress))
	//fmt.Printf("pool coin address list: %s\n", coinAddrList)

	/** <==============================================================================================================> */
	//tokenName, _ := tokenIns.Name(nil)

	//tokenTicker, _ := tokenIns.Symbol(nil)

	//gaugeIns, _ := steth.NewGauge(common.HexToAddress(stETHGaugeAddress), client)
	//gaugeBalance, _ := gaugeIns.BalanceOf(nil, common.HexToAddress(curveWalletAddr))
	//fmt.Printf("stETH gauge balance: %s\n", gaugeBalance.String())

	//poolCount, _ = metaRegIns.PoolCount(nil)
	//fmt.Printf("get pool count: %d from meta registry\n", poolCount)

	//// output pool index and address
	//var counter int
	//for i := int64(0); i < poolCount.Int64(); i++ {
	//	if counter == 10 {
	//		break
	//	}
	//
	//	poolAddr, _ := metaRegIns.PoolList(nil, big.NewInt(i))
	//	if ok, _ := metaRegIns.IsRegistered(nil, poolAddr); !ok {
	//		continue
	//	}
	//
	//	poolName, _ := metaRegIns.GetPoolName(nil, poolAddr)
	//	poolLPToken, _ := metaRegIns.GetLpToken(nil, poolAddr)
	//
	//	fmt.Printf("pool info detail index: %d, address: %s, name: %s, lpToken: %s\n", i, poolAddr.Hex(), poolName, poolLPToken.Hex())
	//	counter++
	//}
}

func InitMetaIns(baseClient node.LocalClient) (*meta.Meta, error) {
	metaRegistry, e := meta.NewMeta(common.HexToAddress(metaRegistryAddress), baseClient)
	if e != nil {
		return nil, e
	}
	return metaRegistry, nil
}

func GetPoolName(ins *meta.Meta, poolAddr common.Address) (string, error) {
	poolName, e := ins.GetPoolName(nil, poolAddr)
	if e != nil {
		return "", fmt.Errorf("fetch pool name false: %s", e.Error())
	}

	return poolName, nil
}
