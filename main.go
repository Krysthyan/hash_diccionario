package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"container/list"
)

func diccionario() (lista list.List) {

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
