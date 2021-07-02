package main

import (
	"fmt"
	"time"
)

func leerArchivo() string {
	time.Sleep(5 * time.Second)
	return "Datos del archivo"
}

func main() {
	datos := leerArchivo()
	fmt.Println(datos)
}
