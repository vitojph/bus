package main

import (
    "fmt"
    "time"
    )


// creamos un struc servicio con la información relativa a cada servicio de autobus
type servicio struct {
    linea int
    origen string
    destino string
    //salida time.Clock
    //llegada time.Clock
    recorrido string
}

func main() {
    now := time.Now()
    fmt.Printf("Ahora es %s\n", now)

    t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    fmt.Printf("Go launched at %s\n", t.Local())

    //salida := time.Clock(9, 0, 0)
    //fmt.Println(salida)

}
