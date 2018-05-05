package main

/* Al useful imports */
import (
    "log"
    "net"
    "google.golang.org/grpc"
    "time"
    "net/http"
    "encoding/json"
    "os"
)

var actAddress string

const (
    port = ":8000" //master running port
)
//server address

    

func main(){
    address := [3]string{"localhost:8001", "localhost:8002", "localhost:8003"} //master running port

    //setup 
    _, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpc.NewServer()
    log.Printf("master coordinator created")

    //make a channel for actual server address
    for {
        if checkAvailability(address[0]) {
            actAddress = address[0]

            //get previous log
        }else if checkAvailability(address[1]){
            actAddress = address[1]

            //get previous log
        }else if checkAvailability(address[2]){
            actAddress = address[2]

             //get previous log
        }else{
            actAddress = ""
        }
        
        time.Sleep(10 * time.Second)
        getAddress()
    }   
    http.HandleFunc("/actAddress", showAddress)
    
}

func showAddress(w http.ResponseWriter, req *http.Request) {
    //get json api

    json.NewEncoder(w).Encode(actAddress)


}

func checkAvailability(address string) bool{
    timeout := time.Duration(10 * time.Second)
    _, err := net.DialTimeout("tcp",address, timeout)
    if err != nil {
        log.Println("Site unreachable, error: ", err)
        return false
    }else{
        log.Println("Reach Site: ", address)

        //write address to log//////////////////////////////////////////////////////////////////////////////
        //convert from []string to []byte
        f, err := os.OpenFile("global.log", os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil {
            panic(err)
        }

        defer f.Close()

        if _, err = f.WriteString(address); err != nil {
            panic(err)
        }

            return true
        }

}

func getAddress() {
    if actAddress != ""{
        log.Println("Using server:", actAddress)
    }else{
        log.Println("No server is available, please wait a moment.")
    }
    
}

//copy previous servers' log
func copyLog() {
    
}



