package main

import (
	"os"
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"container/list"
	"io/ioutil"
	"strings"
)

func funcion_sha256(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}
func obtener_lista_path(path string ) (lista list.List) {
	archivos, _ := ioutil.ReadDir(path)
    for _, f := range archivos {
            lista.PushBack(f.Name())
    }
	return 
}
func obtener_lista_password() (lista list.List) {
	archivo, _ := os.Open("test_sha256/sha256.txt")
	scanner := bufio.NewScanner(archivo)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		lista.PushBack(scanner.Text())
	}
	archivo.Close()
	return
}
func comprobar_existencia(palabra string)  {
	palabra_hash := funcion_sha256("IEEE"+palabra+"Xtreme")
	lista_password :=  obtener_lista_password()

	for e := lista_password.Front(); e != nil; e=e.Next() {
		elemento_sha256 :=e.Value.(string)

		if strings.Compare(palabra_hash,elemento_sha256) == 0 {
			escribir_palabra_encontrada("palabras_encontradas.txt",palabra)
		}



    }
	for e := lista_password.Front(); e != nil; e=e.Next() {

    }
}
func decifrar_password(path string)  {
	archivo, _ := os.Open(path)
	scanner := bufio.NewScanner(archivo)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		comprobar_existencia(scanner.Text())

	}
}

func escribir_palabra_encontrada(path ,linea string) {
	lista := list.List{}
	validar := 0

	archivo_lectura, _ := os.Open(path)
	scanner := bufio.NewScanner(archivo_lectura)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == linea{
			validar = 1
		}
		lista.PushBack(scanner.Text())
	}
	if validar == 0 {
		lista.PushBack(linea)
	}
	archivo_lectura.Close()
	archivo, _ := os.Create(path)
	escritura := bufio.NewWriter(archivo)

	for e := lista.Front(); e != nil; e=e.Next() {
		escritura.WriteString(e.Value.(string)+"\n")
    }
	escritura.Flush()
	archivo.Close()
}

func main() {
	lista_rutas := obtener_lista_path("/root/PycharmProjects/hash_diccionario/diccionarios")
	for e := lista_rutas.Front(); e != nil; e=e.Next() {
			decifrar_password("diccionarios/"+e.Value.(string))
    }
}
