package main

import (
    "fmt"
    "time"
    )


// creamos un struc Servicio con la información relativa a cada servicio de autobus
type Servicio struct {
    linea int
    destino string
    salida int
    llegada string
    recorrido string
}

// Servicios es un slice de elementos Servicio
type Servicios []Servicio

// Carga un listado de horarios de una línea
func loadLine() Servicios {

    var servicios[]Servicio
    for i := 1; i < 11; i++ {
        servicio := Servicio{linea:528, destino:"Madrid", salida:i, recorrido:"Cuesta del Águila"}
        servicios = append(servicios, servicio)
    }
    return servicios
}



func main() {

    p := fmt.Printf
    now := time.Now()
    p("Ahora es %s\n", now)

    t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    p("Go launched at %s\n", t.Local())

    salida := time.Date(2015, 1, 1, 9, 0, 0, 0, time.UTC)
    p("La hora que he creado es %s\n", salida)


    servicios := loadLine()
    p("\n\n")
    fmt.Println(servicios)
}
