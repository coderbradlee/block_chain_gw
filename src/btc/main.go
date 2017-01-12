package main
 
import (
    "github.com/btcsuite/btcd/chaincfg"
    "github.com/btcsuite/btcrpcclient"
    "github.com/btcsuite/btcutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "time"
    "strings"
    "io/ioutil"
    "flag"
    "runtime"
)
type Configuration struct {
    Ip string,
    Port string,
    User string,
    Pass string,
    Log_name string
}
var configuration Configuration
// var (
//     logFileName string
// )
func init() {
    
    file, _ := os.Open("conf.json")
    decoder := json.NewDecoder(file)
    configuration = Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    log_init()
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
    //fmt.Printf("%s\n",configuration.Exec_time)
}
func log_init() {
    log_name:=fmt.Sprintf("%s",configuration.Log_name)
    logFileName := flag.String("log", log_name, "Log file name")
    runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()

    //set logfile Stdout
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
 
    // prepare a sendMany transaction
    // receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
    // if err != nil {
    //     log.Fatalf("address receiver1 seems to be invalid: %v", err)
    // }
    // receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
    // if err != nil {
    //     log.Fatalf("address receiver2 seems to be invalid: %v", err)
    // }
    // receivers := map[btcutil.Address]btcutil.Amount{
    //     receiver1: 42,  // 42 satoshi
    //     receiver2: 100, // 100 satoshi
    // }
 
    // // create and send the sendMany tx
    // txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
    // if err != nil {
    //     log.Fatalf("error sendMany: %v", err)
    // }
    // log.Printf("sendMany completed! tx sha is: %s", txSha.String())
}


