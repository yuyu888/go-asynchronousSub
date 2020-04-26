package main


import (
    "fmt"
    "flag"
    //"log"
    "gas/httpServer"
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
    httpServer.Init();
}