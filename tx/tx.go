/*
reference
https://blog.csdn.net/u010662978/article/details/79195284
*/
package main

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"fmt"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"encoding/hex"
	"bytes"
)

func main(){
	address:="mtNELkqMFJUHyKYxdcGJN6mszNwYH1JHyW"
	var balance int64 = 75845392//utxo
	var fee int64 = 0.001*1e8//交易费
	var leftToMe = balance - fee
	//1.构造输出
	outputs:=[]*wire.TxOut{}
	//1.1输出1，给自己转付完交易费之后的钱
	addr,_:=btcutil.DecodeAddress(address,&chaincfg.SimNetParams)
	pkScript,_:=txscript.PayToAddrScript(addr)
	outputs = append(outputs,wire.NewTxOut(leftToMe,pkScript))
	//1.2 添加文字
	comment:="币加锁，到此一游"
	pkScript,_=txscript.NullDataScript([]byte(comment))
	outputs = append(outputs,wire.NewTxOut(int64(0),pkScript))
	//2.构造输入
	prevTxHash :="23f427cfd4b4b6214d04f1642f4e08421fb693f443914e247b9ec3fc3bc71895"
	prevPkScriptHex:="76a9148cf3495bc3fde7d99a3039772a2924c5dbf4052188ac"
	prevTxOutputN:=uint32(0)

	hash,_:=chainhash.NewHashFromStr(prevTxHash)
	outPoint:=wire.NewOutPoint(hash,prevTxOutputN)
	txIn :=wire.NewTxIn(outPoint,nil,nil)
	inputs:=[]*wire.TxIn{txIn}

	prevPkScript,_:=hex.DecodeString(prevPkScriptHex)
	prevPkScripts:=make([][]byte,1)
	prevPkScripts[0] = prevPkScript

	tx:=&wire.MsgTx{
		Version:wire.TxVersion,
		TxIn:inputs,
		TxOut:outputs,
		LockTime:0,
	}
	privKey:="92TNcm4wXSb6zyov6x8TFg1NuQRZ6j7DMPNXkW2FeAYDrxLhSE7"
	sign(tx,privKey,prevPkScripts)
	buf :=bytes.NewBuffer(make([]byte,0,tx.SerializeSize()))
	if err:=tx.Serialize(buf);err != nil{

	}
	txHex := hex.EncodeToString(buf.Bytes())
	fmt.Println("hex",txHex)
	//01000000019518c73bfcc39e7b244e9143f493b61f42084e2f64f1044d21b6b4d4cf27f423000000008a4730440220735082cd5f2d0b55f7db43475eba244bce0ff7b57a379ee2d677fbf5e9e57bef0220373a9c45cc07d6e615ef577aa80ab9a9e1745b10d9e305386793d24311c2f8be014104159b9e8fdf9c91eafe4f2480ff2fcb5b35f58bb0a3fffa13e72817ae7bb3f397f6bb57a8353066417ffdf56511d9c9ed96f833694e08f70a74b811bd7ef03b3fffffffff0270c88304000000000000000000000000001a6a18e5b881e58aa0e99481efbc8ce588b0e6ada4e4b880e6b8b800000000
	//broadcast website https://tbtc.blockdozer.com/tx/send
	//txid  9851697828814f9e87797e581211c4a36f50fcedd55da1ff963994a69449f4b5
	//https://www.blocktrail.com/tBTC/tx/9851697828814f9e87797e581211c4a36f50fcedd55da1ff963994a69449f4b5
}

//签名
func sign(tx *wire.MsgTx,privKeyStr string,prevPKScripts [][]byte){
	inputs := tx.TxIn
	wif,err:=btcutil.DecodeWIF(privKeyStr)
	if err!=nil{
		fmt.Println("wif err",err)
	}
	privKey :=wif.PrivKey
	for i:=range inputs{
		pkScript:=prevPKScripts[i]
		var script []byte
		script,err = txscript.SignatureScript(tx,i,pkScript,txscript.SigHashAll,privKey,false)
		inputs[i].SignatureScript = script
	}
}
