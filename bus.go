package main

import (
    "fmt"
    "time"
    )


// creamos un struct Servicio con la información relativa a cada servicio de autobus
type Servicio struct {
    linea int
    destino string
    salida int
    llegada string
    recorrido string
}

// Servicios es un slice de elementos Servicio
type Servicios []Servicio

type myTime struct {
    time.Time // anonymous field
}

// Carga un listado de horarios de una línea
func cargaLinea() Servicios {

    var servicios[]Servicio

    for i := 0; i < 24; i++ {
        servicio := Servicio{linea:528, destino:"Madrid", salida:i, recorrido:"Cuesta del Águila"}
        servicios = append(servicios, servicio)
    }
    return servicios
}

// filtra los servicios dependiendo de la hora
func filtraServicios() Servicios  {
    // TODO: añadir destino como argumento
    // carga los horarios
    servicios := cargaLinea()
    var proximosServicios[]Servicio
    now := time.Now()
    // recorremos la lista de servicios e imprimimos los próximos
    for _, servicio := range servicios {
        if servicio.salida >= now.Hour() {
            proximosServicios = append(proximosServicios, servicio)
        }
    }
    return proximosServicios
}

func main() {

    p := fmt.Printf
    now := time.Now()
    p("Ahora es %s\n", now)

    //t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
    //p("Go launched at %s\n", t.Local())

    //salida := time.Date(2015, 1, 1, 9, 0, 0, 0, time.UTC)
    //p("La hora que he creado es %s\n", salida)

    proximosBuses := filtraServicios()
    fmt.Println(proximosBuses)

}
