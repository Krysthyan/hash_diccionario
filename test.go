package main

import (
	"fmt"
	"container/list"
)

func main() {
	var x list.List
    x.PushBack("hola mundo")
    x.PushBack("como estas")
    x.PushBack("en donde")

    for e := x.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(string))
    }
}
