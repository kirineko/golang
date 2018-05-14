package main  
  
import (  
    "encoding/json"  
    "fmt"  
    "io/ioutil"  
    "net/http"  
    "time"  
)  
  
type Item struct {  
    Seq    int  
    Result map[string]int  
}  
  
type Message struct {  
    Dept    string  
    Subject string  
    Time    int64  
    Detail  []Item  
}  
  
func main() {  
    url := "http://localhost:8085"  
    ret, err := http.Get(url)  
  
    if err != nil {  
        panic(err)  
    }  
    defer ret.Body.Close()  
  
    body, err := ioutil.ReadAll(ret.Body)  
    if err != nil {  
        panic(err)  
    }  
  
    var msg Message  
    err = json.Unmarshal(body, &msg)  
    if err != nil {  
        panic(err)  
    }  
  
    strTime := time.Unix(msg.Time, 0).Format("2006-01-02 15:04:05")  
    fmt.Println("Dept:", msg.Dept)  
    fmt.Println("Subject:", msg.Subject)  
    fmt.Println("Time:", strTime, "\n", msg.Detail)  
} 