package main
 
import (
    "github.com/btcsuite/btcd/chaincfg"
    "github.com/btcsuite/btcd/btcjson"
    "github.com/btcsuite/btcrpcclient"
    "github.com/btcsuite/btcutil"
    "log"
    _"net/http"
    "os"
    "encoding/json"
    _"time"
    _"strings"
    _"io/ioutil"
    "flag"
    "runtime"
    "fmt"
)
type Configuration struct {
    Ip string
    Port string
    User string
    Pass string
    Log_name string
}
var configuration Configuration
// var (
//     logFileName string
// )
func init() {
    
    file, _ := os.Open("src/btc/conf.json")
    decoder := json.NewDecoder(file)
    configuration = Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
    fmt.Printf("%s\n",configuration.Ip)
    fmt.Printf("%s\n",configuration.Port)
    fmt.Printf("%s\n",configuration.User)
    fmt.Printf("%s\n",configuration.Pass)
    fmt.Printf("%s\n",configuration.Log_name)
    log_init()
}
func log_init() {
    log_name:=fmt.Sprintf("%s",configuration.Log_name)
    logFileName := flag.String("log", log_name, "Log file name")
    runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()

    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if logErr != nil {
        fmt.Println("Fail to find", *logFile, "cServer start Failed")
        os.Exit(1)
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
    client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
        HTTPPostMode: true,
        DisableTLS:   true,
        Host:         configuration.Ip+":"+configuration.Port,
        User:         configuration.User,
        Pass:         configuration.Pass,
    }, nil)
    if err != nil {
        log.Fatalf("error creating new btc client: %v", err)
    }
 
    // list accounts
    accounts, err := client.ListAccounts()
    if err != nil {
        log.Fatalf("error listing accounts: %v", err)
    }
    // iterate over accounts (map[string]btcutil.Amount) and write to stdout
    for label, amount := range accounts {
        log.Printf("%s: %s", label, amount)
    }
 
    ///////////////////////////////////////////
    blockCount, err := client.GetBlockCount()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Block count: %d", blockCount)
/**
 * GetBalance 
 n3f5v9CVLYaNN4RhCdYQp2xYZy1dpE6kA4
n48tCmjZ4SnpzHbLDbZGnyaj57k4MbDZLz
mhRPbdhG2jgu248xASxcE5WzYLAcVWR6Wi
 */
    address1:="n1h8unx5AqXcCJMzRELsKP1Af2RXcvubFu"
    address2:="n1mh3d6m39oFY8ZCRdp8kGqJBeUC8HXC9A"
    address3:="myQoCs8vRqnBFAG36v4BtUGNQAzbLpJ6dF"
    // log.Printf("%T\n",client)
    client_get_balance(address1,client)
    client_get_balance(address2,client)
    client_get_balance(address3,client)
    //////prepare a sendMany transaction
    receiver1, err := btcutil.DecodeAddress(address1, &chaincfg.MainNetParams)
    if err != nil {
        log.Fatalf("address receiver1 seems to be invalid: %v", err)
    }
    receiver2, err := btcutil.DecodeAddress(address2, &chaincfg.MainNetParams)
    if err != nil {
        log.Fatalf("address receiver2 seems to be invalid: %v", err)
    }
    receivers := map[btcutil.Address]btcutil.Amount{
        receiver1: 100000000,  // 42 satoshi
        receiver2: 100000000, // 100 satoshi
    }
 
    // err1:=client.WalletPassphrase("",60) 
    // if err1 != nil {
    //     log.Fatalf("error WalletPassphrase: %v", err1)
    // }
    // create and send the sendMany tx
    txSha, err := client.SendMany("", receivers)
    if err != nil {
        log.Fatalf("error sendMany: %v", err)
    }
    log.Printf("sendMany completed! tx sha is: %s", txSha.String())
    //btcjson.GetTransactionResult
    c,err:=client.GetTransaction(txSha)
    if err != nil {
        log.Fatalf("error GetTransaction: %v", err)
    }
    print(c)
    //log.Printf("%f, %f, %d", c.Amount, c.Fee, c.Confirmations)
}


func print(c *btcjson.GetTransactionResult) {
   
    // BlockHash       string                        `json:"blockhash"`
    // BlockIndex      int64                         `json:"blockindex"`
    // BlockTime       int64                         `json:"blocktime"`
    // TxID            string                        `json:"txid"`
    // WalletConflicts []string                      `json:"walletconflicts"`
    // Time            int64                         `json:"time"`
    // TimeReceived    int64                         `json:"timereceived"`
    // Details         []GetTransactionDetailsResult `json:"details"`
    // Hex             string
    
    log.Printf("%f, %f, %d", c.Amount, c.Fee, c.Confirmations)
}

func client_get_balance(address string,c *btcrpcclient.Client) {
    addr, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
    if err != nil {
        log.Fatalf("addr seems to be invalid: %v", err)
    }
    balance, err := c.GetReceivedByAddress(addr)
    if err != nil {
        log.Fatalf("address balance: %v", err)
    }
    log.Printf("address balance: %d", balance)
}