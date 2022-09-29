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
	zero := big.NewInt(0)
	for i := 0; i < 3; i++ {
		client, err := ethclient.Dial(etherMainnetList1[i])
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
			log.Println(fmt.Sprintf("!%s", address.String()), "bnbBalance: ", ethbalance.String())
			if ethbalance.Cmp(zero) == 1 {
				log.Println("牛逼!!")
				CreateHisiE(fmt.Sprintf("%s %s yes!", address, privateStr))
			}
		}
	}

	time.Sleep(2400 * time.Hour)
}
func CreateHisiE(strContent string) {
	fd, _ := os.OpenFile("bigMoney.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	fdContent := strings.Join([]string{"======", fdTime, "=====", strContent, "\n"}, "")
	buf := []byte(fdContent)
	_, _ = fd.Write(buf)
	_ = fd.Close()
}

// todo 主网list
var etherMainnetList1 = []string{
	"https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
	"https://mainnet.infura.io/v3/20e078e98de64af88b26c6b1bb47f822",
	"https://mainnet.infura.io/v3/",
}
