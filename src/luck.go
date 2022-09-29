package main

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	fmt.Println("幸运原理: ")
	fmt.Println("1: 生成随机密钥")
	fmt.Println("2: 密钥算出地址, 循环比对1.8万个地址")
	fmt.Println("3: 成功后,生成 bigMoney.txt 文件存储地址与私钥")
	fmt.Println("此举我称之为 \"大海捞针\" , 概率极低, 不过运气这种事, 谁说得好呢 ^_^ ")
	fmt.Println("")
	fmt.Println("注意 : 5 秒后执行..")
	time.Sleep(5 * time.Second)

	var addrMap = &addrMapList
	for i := 0; i < 4; i++ {

		go func() {
			for true {
				privateKey, _ := crypto.GenerateKey()
				privateKeyBytes := crypto.FromECDSA(privateKey)
				privateStr := hexutil.Encode(privateKeyBytes)[2:]
				publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
				address := crypto.PubkeyToAddress(*publicKeyECDSA)
				s := address.String()
				sf := s[2:10]
				sl := s[34:42]
				for _, v := range *addrMap {
					if sf == v {
						PrintLog("牛逼!!")
						CreateHis(fmt.Sprintf("%s %s yes!", address, privateStr))
					}
					if sl == v {
						PrintLog("牛逼!!")
						CreateHis(fmt.Sprintf("%s %s yes!", address, privateStr))
					}
				}
			}
		}()
	}

	time.Sleep(2400 * time.Hour)
}

func CreateHis(strContent string) {
	fd, _ := os.OpenFile("bigMoney.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fdTime := time.Now().Format("2006-01-02 15:04:05")
	fdContent := strings.Join([]string{"======", fdTime, "=====", strContent, "\n"}, "")
	buf := []byte(fdContent)
	_, _ = fd.Write(buf)
	_ = fd.Close()
}
func PrintLog(logStr string) {
	nowTime := time.Now()
	sStr := nowTime.Format("15:04:05")
	ms := fmt.Sprintf("%v", nowTime.UnixNano()/1e6)
	msStr := ms[10:]
	printStr := fmt.Sprintf("%s:%s: %s", sStr, msStr, logStr)
	fmt.Println(printStr)
}

// todo 有金额地址list  大于10万美元
var addrMapList = []string{
	"00000000",
	"11111111",
	"22222222",
	"33333333",
	"44444444",
	"55555555",
	"66666666",
	"77777777",
	"88888888",
	"99999999",
}
