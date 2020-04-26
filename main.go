package main


import (
    "fmt"
    "flag"
    //"log"
    "gas/httpServer"
    "gas/libs/log"
    "strconv"
)

var (
    Cnum = flag.String("cnum", "10", "max concurrent num")
)

func main(){
    flag.Parse()
    fmt.Println(*Cnum)
    maxNum, _ := strconv.Atoi(*Cnum)
    httpServer.SetMaxTaskNum(maxNum)
    log.Init()
    httpServer.Init();
}