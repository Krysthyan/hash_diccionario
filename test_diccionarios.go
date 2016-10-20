package main

import (
	"os"
	"bufio"
	"fmt"
	"crypto/sha256"
	"encoding/base64"
	"container/list"
	"io/ioutil"
)

func encriptar_palabra(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}
func descubir_pababra()  {
	
}

func obtener_lista_path(path string ) (lista list.List) {
	archivos, _ := ioutil.ReadDir(path)
    for _, f := range archivos {
            lista.PushBack(f.Name())
    }
	return 
}
func obtener_sha256() (lista list.List) {
	archivo, _ := os.Open("test_sha256/sha256.txt")
	scanner := bufio.NewScanner(archivo)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		lista.PushBack(scanner.Text())
	}
	return
}

func leer_diccionario(path string)  {
	archivo, _ := os.Open(path)
	scanner := bufio.NewScanner(archivo)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
}

func escribir_password(path ,linea string) {
	archivo, _ := os.Create(path)
	scanner := bufio.NewWriter(archivo)
	scanner.WriteString(linea)
	scanner.Flush()
}

func main() {

}
