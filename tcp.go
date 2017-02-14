package main
//A tcp server which listens on the port 3333 for incoming messages. 
//It can only handle ONE incoming message. How to handle multiple is yet to be solved.

import(
    "fmt"
    "net"
    "os"
)

func main(){

    listener, err := net.Listen("tcp", "localhost:3333")
    if(err != nil){
      fmt.Println("error while listening: "+err.Error())
      os.Exit(1)
    }
    fmt.Println("listening on: "+"localhost:3333")
    defer listener.Close()

    for {
      fmt.Println("trying to accept: "+"localhost:3333")
      conn, err := listener.Accept()
      if(err != nil){
        fmt.Println("error while accepting: "+err.Error())
        os.Exit(1)
      }
      go handleRequests(conn)
      fmt.Println("Code after handleRequests()");
    }
    fmt.Println("tcp acceptance of incoming messages stopped");
}


    func handleRequests(conn net.Conn){
      defer conn.Close()
      b := make([]byte, 1024)
      n, err := conn.Read(b)
      if(err != nil){
        fmt.Println("error while reading: "+err.Error())
        os.Exit(1)
      }
      fmt.Println("data read from tcp client: "+string(b[:n]))

      conn.Write([]byte("Got your message!"))
    }
