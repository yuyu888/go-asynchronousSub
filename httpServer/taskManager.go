package httpServer

import (
    "fmt"
    "gas/libs/curl"
    "gas/libs/log"
    "go.uber.org/zap"
)

//任务管理
type TaskManager struct {
    inchan chan *RequestData
    maxGchNum int
}

type RequestData struct {
    url string
    msgData string

}


//创建任务管理者
var Task = TaskManager{
    inchan:   make(chan *RequestData, 10000),
}

func SetMaxTaskNum(maxNum  int){
    Task.maxGchNum = maxNum
}

func (Task *TaskManager) ProcLoop() {
    fmt.Println(Task.maxGchNum)

    gch := make(chan bool, Task.maxGchNum)
    for {
        requestInfo := Task.sub()
        gch <- true
        go Task.doRequest(requestInfo, gch)
    }
}

func (Task *TaskManager) doRequest(requestInfo *RequestData, gch chan bool){
    RequestUrl := requestInfo.url+"?msgData="+requestInfo.msgData
    fmt.Println(RequestUrl)

    req := curl.NewRequest()
    resp, err := req.SetUrl(RequestUrl).
        SetMethod("GET").
        Send()
    if err != nil {
        log.InfoLogger.Info("request fail, service is bad", zap.String("errorCode", "E500"), zap.String("RequestUrl", RequestUrl))
    } else {
        if resp.IsOk() {
            log.InfoLogger.Info("request is Ok",
                zap.String("errorCode", "E200"),
                zap.String("RequestUrl", RequestUrl),
                zap.String("response", resp.Body),
                zap.Int("httpStatus", resp.Raw.StatusCode))
        } else {
            log.InfoLogger.Info("request is fail",
                zap.String("errorCode", "E400"),
                zap.String("RequestUrl", RequestUrl),
                zap.Int("httpStatus", resp.Raw.StatusCode),)
        }
    }
    <-gch
}

func (Task *TaskManager) sub() (*RequestData){
    select {
        case requestInfo := <-Task.inchan:
            return requestInfo
    }
    return nil
}

func (Task *TaskManager) pub(requestInfo *RequestData) error{
    select {
        case   Task.inchan<-requestInfo:
    }
    return nil
}