package httpServer

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
)


func httpHandlerIndex(w http.ResponseWriter, r *http.Request) {
    resp := "httpserver is runing"
    fmt.Fprintln(w, resp)
}

func httpHandlerTest(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    msgData := query.Get("msgData")
    data := "this is test return, data:"+msgData
    resp := `{"status":200, "message":"success","data":"`+data+`"}`
    fmt.Fprintln(w, resp)

}

func httpHandlerTest1(w http.ResponseWriter, r *http.Request) {
    time.Sleep(3 * time.Second)
    query := r.URL.Query()
    msgData := query.Get("msgData")
    data := "this is test1 return, data:"+msgData
    resp := `{"status":200, "message":"success","data":"`+data+`"}`
    fmt.Fprintln(w, resp)
}

func httpHandlerCall(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    msgData := query.Get("msgData")
    action := query.Get("action")
    if len(action)==0{
        display(w, 4001, "action is null", "")
        return
    }
    if url, ok := UrlList[action]; !ok {
        display(w, 4002, "Illegal request, url is not allowed", "")
    }else{
        requestData := &RequestData{
            url : url,
            msgData : msgData,
        }
        Task.pub(requestData)
        display(w, 200, "request is accepted", "")
    }

}

func display(w http.ResponseWriter, status int, msg string, data string){
    resp := `{"status":`+strconv.Itoa(status)+`, "message":"`+msg+`","data":"`+data+`"}`
    fmt.Fprintln(w, resp)
}


func Init() {
    fmt.Println("httpServer is run")
    go Task.ProcLoop()
    http.HandleFunc("/", httpHandlerIndex)
    http.HandleFunc("/test", httpHandlerTest)
    http.HandleFunc("/test1", httpHandlerTest1)
    http.HandleFunc("/call", httpHandlerCall)
    http.ListenAndServe("0.0.0.0:80", nil)
}