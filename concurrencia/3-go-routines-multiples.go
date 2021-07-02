package main

import (
	"fmt"
	"net/http"
	"time"
)

type RSS struct {
	Nombre string
	Url    string
}

func (rss *RSS) leer() {
	t1 := time.Now()
	resp, err := http.Get(rss.Url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	t2 := time.Since(t1).Seconds()
	fmt.Println(rss.Nombre, "> lectura completada en", t2, "segundos")
}

func main() {
	listaRSS := []*RSS{
		&RSS{Nombre: "ESPN", Url: "http://www.espn.com.mx/espn/rss/news"},
		&RSS{Nombre: "El Reforma", Url: "https://www.reforma.com/rss/portada.xml"},
		&RSS{Nombre: "El Financiero", Url: "https://www.elfinanciero.com.mx/arc/outboundfeeds/rss/?outputType=xml"},
	}
	for _, rss := range listaRSS {
		go rss.leer()
	}
	time.Sleep(time.Second * 5)
}
