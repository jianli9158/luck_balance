package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ackermanx/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	log.Println("幸运原理: ")
	log.Println("1: 生成随机密钥")
	log.Println("2: 密钥算出地址, 循环查询bnb余额")
	log.Println("3: 成功后,生成 bigMoney.txt 文件存储地址与私钥")
	log.Println("此举我称之为 \"大海捞针\" , 概率极低, 不过运气这种事, 谁说得好呢 ^_^ ")
	log.Println("")
	log.Println("注意 : 3 秒后执行..")
	time.Sleep(3 * time.Second)
	for i := 0; i < 128; i++ {
		//bnb
		go balance_bnb()
		//eth
		go balance_eth()
	}
	time.Sleep(2400 * time.Hour)
}
func balance_eth() {
	zero := big.NewInt(0)
	for i := 0; i < 3; i++ {
		client, err := ethclient.Dial(etherMainnetList[i])
		if err != nil {
			log.Println(err)
			if i == 2 {
				i = -1
			}
			continue
		}

		for true {
			privateKey, _ := crypto.GenerateKey()
			privateKeyBytes := crypto.FromECDSA(privateKey)
			privateStr := hexutil.Encode(privateKeyBytes)[2:]
			publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
			address := crypto.PubkeyToAddress(*publicKeyECDSA)
			account := common.HexToAddress(address.String())
			// get eth balance
			ethbalance, err := client.BalanceAt(context.Background(), account, nil)
			if err != nil {
				log.Println(err)
				if i == 2 {
					i = -1
				}
				break
			}
			log.Println(fmt.Sprintf("!%s", address.String()), "ethBalance: ", ethbalance.String())
			if ethbalance.Cmp(zero) == 1 {
				log.Println("牛逼!!")
				CreateHisi(fmt.Sprintf("%s %s ethBalance>0!!!!", address, privateStr))
			}
		}
	}

}

func balance_bnb() {
	zero := big.NewInt(0)
	for i := 0; i < 13; i++ {
		fmt.Println("bnbBalance: ", binanceMainnetList[i], i)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		c, err := ethclient.DialContext(ctx, binanceMainnetList[i])
		cancel()
		if err != nil {
			log.Println(err)
			if i == 12 {
				i = -1
			}
			continue
		}

		for true {
			privateKey, _ := crypto.GenerateKey()
			privateKeyBytes := crypto.FromECDSA(privateKey)
			privateStr := hexutil.Encode(privateKeyBytes)[2:]
			publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
			address := crypto.PubkeyToAddress(*publicKeyECDSA)
			// get bnb balance
			bnbBalance, err := c.BalanceAt(context.Background(), common.HexToAddress(address.String()), nil)
			if err != nil {
				log.Println(err)
				if i == 12 {
					i = -1
				}
				break
			}
			log.Println(fmt.Sprintf("!%s", address.String()), "bnbBalance: ", bnbBalance.String())
			if bnbBalance.Cmp(zero) == 1 {
				fmt.Println("牛逼!!")
				CreateHisi(fmt.Sprintf("%s %s ethBalance>0!!!!", address, privateStr))
			}
		}

	}
}
func CreateHisi(strContent string) {
	fd, _ := os.OpenFile("bigMoney.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	fdContent := strings.Join([]string{"======", fdTime, "=====", strContent, "\n"}, "")
	buf := []byte(fdContent)
	_, _ = fd.Write(buf)
	_ = fd.Close()
}

// todo 主网list
var binanceMainnetList = []string{
	"https://bsc-dataseed.binance.org/",
	"https://bsc-dataseed1.defibit.io/",
	"https://bsc-dataseed1.ninicoin.io/",
	"https://bsc-dataseed2.defibit.io/",
	"https://bsc-dataseed3.defibit.io/",
	"https://bsc-dataseed4.defibit.io/",
	"https://bsc-dataseed2.ninicoin.io/",
	"https://bsc-dataseed3.ninicoin.io/",
	"https://bsc-dataseed4.ninicoin.io/",
	"https://bsc-dataseed1.binance.org/",
	"https://bsc-dataseed2.binance.org/",
	"https://bsc-dataseed3.binance.org/",
	"https://bsc-dataseed4.binance.org/",
}

// todo 主网list
var etherMainnetList = []string{
	"https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
	"https://mainnet.infura.io/v3/20e078e98de64af88b26c6b1bb47f822",
	"https://mainnet.infura.io/v3/",
}
