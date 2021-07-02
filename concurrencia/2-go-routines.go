package main

import (
	"fmt"
	"time"
)

func leerArchivo() string {
	time.Sleep(time.Second * 5)
	return "Datos del archivo"
}

func main() {
	go func() {
		datos := leerArchivo()
		fmt.Println(datos)
	}()
	time.Sleep(time.Second * 10)
	fmt.Println("Continuar con la ejecución")
}
