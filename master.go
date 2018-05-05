package main

/* Al useful imports */
import (
    "log"
    "net"
    "google.golang.org/grpc"
    "time"
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


    for{
        if checkAvailability(address[0]) {
            actAddress = address[0]
        }else if checkAvailability(address[1]){
            actAddress = address[1]
        }else if checkAvailability(address[2]){
            actAddress = address[2]
        }else{
            actAddress = ""
        }

        getAddress()
    }
    
}

func checkAvailability(address string) bool{
    timeout := time.Duration(10 * time.Second)
    _, err := net.DialTimeout("tcp",address, timeout)
    if err != nil {
        log.Println("Site unreachable, error: ", err)
        return false
    }else{
        log.Println("Reach Site: ", address)
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



