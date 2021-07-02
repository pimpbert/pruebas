package main

import (
	"fmt"
)

func main() {
	noticias := make(chan string, 3)
	noticias <- "México Sub-23 perfila dos partidos de preparación previo a su debut olímpico contra Francia"
	noticias <- "FIFA confirma convocatorias de 22 jugadores para JJOO, pero sólo 18 serán registrados cada partido"
	noticias <- "Pumas suma cinco positivos en pretemporada"
	close(noticias)
	noticias <- "El once de jugadores libres, liderado por Messi y Sergio Ramos"
	for noticia := range noticias {
		fmt.Println(noticia)
	}
}
