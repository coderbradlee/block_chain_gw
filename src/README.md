1、调用bitcoin的http rpc接口
2、调用ethereum的http rpc接口
3、实现两者互通转账
4、实现智能合约存储比特币的信息


用bitcoin-testnet-box
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 help
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 sendtoaddress "n1bnrtuqShki6wUSUVHRJ6dvc3EH46H38a" 12

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getinfo

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getaccountaddress ""
output::mkxoF3Vqa3D2EtBPEYbhsZgXfpKuoHyqP2
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19011 -rpcuser=admin2 -rpcpassword=123 getaccountaddress ""
output::n1Dgewo57eFPcRRPanSwzcAdYbLQ8saH8S

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 sendtoaddress "n1bnrtuqShki6wUSUVHRJ6dvc3EH46H38a" 12
output::3e3b7baf784266230691a0939fcb8409766e82a0e79bca20b1dc29f38c3840d7

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getaccount mkxoF3Vqa3D2EtBPEYbhsZgXfpKuoHyqP2 
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress n1bnrtuqShki6wUSUVHRJ6dvc3EH46H38a

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 gettransaction 3e3b7baf784266230691a0939fcb8409766e82a0e79bca20b1dc29f38c3840d7

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getbalance n1bnrtuqShki6wUSUVHRJ6dvc3EH46H38a
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 generate 100


bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 sendtoaddress "n1bnrtuqShki6wUSUVHRJ6dvc3EH46H38a" 1000
output::fdfaf5f1f925c75d40e9094cfe7f4ed011a3fb23ed43117fb207eed691012d7e

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 sendtoaddress "mkxoF3Vqa3D2EtBPEYbhsZgXfpKuoHyqP2" 1000
output::39a4eb3135f894ded592b45e3e89a95a9ead03e4b370bd8664bcaeda31747031

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getbalance mkxoF3Vqa3D2EtBPEYbhsZgXfpKuoHyqP2

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getmininginfo

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 listaddressgroupings
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19011 -rpcuser=admin2 -rpcpassword=123 listaddressgroupings

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 sendtoaddress "mzbnA97gw6L32bShQeNemS52ZhhCukTupS" 120 
output::f32f0d6e1fd8b30cbdee2e6e15a6d1aa2de20a48eb677439b14d27c4bfd6ed41


bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getbalance mzbnA97gw6L32bShQeNemS52ZhhCukTupS

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 listtransactions


getaddressesbyaccount "" 获取钱包里面的地址
getreceivedbyaddress 获取可以支配的余额

/////至少6个确认块
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress mkxoF3Vqa3D2EtBPEYbhsZgXfpKuoHyqP2 6
1000
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress myQoCs8vRqnBFAG36v4BtUGNQAzbLpJ6dF
0
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress mu9ESBCd53C3z9H81YVCMMQcYmQUYHgLQE
0
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress mzbnA97gw6L32bShQeNemS52ZhhCukTupS
0
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19001 -rpcuser=admin1 -rpcpassword=123 getreceivedbyaddress n1Dgewo57eFPcRRPanSwzcAdYbLQ8saH8S
0

bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19011 -rpcuser=admin2 -rpcpassword=123 getreceivedbyaddress mzbnA97gw6L32bShQeNemS52ZhhCukTupS
bitcoin-cli -rpcconnect=127.0.0.1 -rpcport=19011 -rpcuser=admin2 -rpcpassword=123 getreceivedbyaddress n1Dgewo57eFPcRRPanSwzcAdYbLQ8saH8S

以上两个地址是19011这个钱包里的，用19001来调用得到的是0