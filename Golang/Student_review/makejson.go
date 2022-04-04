package main

import (
    "fmt"
    "encoding/json"
    "os"
    "bufio"
)

func main() {
    var a map[string]string
    a = make(map[string]string)
    fmt.Println("Name:")
    br := bufio.NewReader(os.Stdin)
    bname, _, _ := br.ReadLine()
    name:=string(bname)
    a["name"]=name
    fmt.Println("Adress: ")
    br1 := bufio.NewReader(os.Stdin)
    badress, _, _ := br1.ReadLine()
    adress:=string(badress)
    a["adress"]=adress
    x, err := json.Marshal(a)
    if err != nil {
        fmt.Println("Wrong")
    } else {
        fmt.Println("Encoded data : ")
        fmt.Println(x)
        fmt.Println("Decoded data : ")
        fmt.Println(string(x))
    }
}
