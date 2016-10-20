package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"container/list"
)

func diccionario() (lista list.List) {
	lista.PushBack("Password1")
	lista.PushBack("perro")
	lista.PushBack("stand5")
	lista.PushBack("demons8")
	lista.PushBack("binary2")
	lista.PushBack("maze3")
	lista.PushBack("real7")
	lista.PushBack("me4t4")
	lista.PushBack("h0rs30")
	lista.PushBack("4rt5")
	lista.PushBack("bl0ck5")
	lista.PushBack("sad6")
	lista.PushBack("cow3")
	lista.PushBack("sh3lls4")
	lista.PushBack("land6")
	lista.PushBack("password1")
	lista.PushBack("world4")
	lista.PushBack("crashing")
	lista.PushBack("game9")
	lista.PushBack("c4v35")
	lista.PushBack("url5")
	lista.PushBack("g4mm42")
	lista.PushBack("b4r4")
	lista.PushBack("shortening")
	lista.PushBack("dice9")
	lista.PushBack("esport6")
	lista.PushBack("bon8")
	lista.PushBack("fin1te4")
	lista.PushBack("b34ns1")
	lista.PushBack("roaming7")
	lista.PushBack("t0k3n7")
	lista.PushBack("square7")
	lista.PushBack("winn3r0")
	lista.PushBack("metroc0")
	lista.PushBack("character1")
	lista.PushBack("voyage8")
	lista.PushBack("minimum8")
	lista.PushBack("television")
	lista.PushBack("security")
	lista.PushBack("xtreme")
	lista.PushBack("salted")
	lista.PushBack("peppered")
	lista.PushBack("encoded")
	lista.PushBack("sha256")
	lista.PushBack("possible")
	lista.PushBack("base64")
	lista.PushBack("values")
	lista.PushBack("output")
	lista.PushBack("input")
	lista.PushBack("password")
	lista.PushBack("IEEEXtreme")


	return

}
func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	var text string
	fmt.Scanln(&text)
	dic := diccionario()
	for e := dic.Front(); e != nil; e=e.Next() {
		if encryptPassword("IEEE"+e.Value.(string)+"Xtreme") == text {
			fmt.Println(e.Value.(string))
			break

		}
    }

}
