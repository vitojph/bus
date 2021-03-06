package main

import (
    "fmt"
    "sort"
    "strings"
    "time"
    "strconv"
    "github.com/vitojph/myutils"
)

// carga la lista de servicios a partir de los ficheros *.data
func cargaServicios(filenames ...string) Servicios {
    var servicios[]Servicio

    // cargo los horarios
    for _, filename := range filenames {
        //fmt.Println(filename)
        lines, err := myutils.ReadLines(filename)
        myutils.Check(err)
        for _, line := range lines {
            if !strings.HasPrefix(line, "//") {
                fields := strings.Split(line, ",")
                h_salida, err := strconv.Atoi(fields[3])
                myutils.Check(err)
                m_salida, err := strconv.Atoi(fields[4])
                myutils.Check(err)

                esteServicio := Servicio{horario: fields[0], linea: fields[1], destino: fields[2], h_salida: h_salida, m_salida: m_salida}
                servicios = append(servicios, esteServicio)
            }
        }
    }
    return servicios
}

// comprueba si la fecha actual es durante el horario de verano
func esVerano() bool {
    now := time.Now()
    if now.Month() == 8 {
        return true
    } else {
        return false
    }
}

// comprueba si la fecha actual es fin de semana
func esFinde() bool {
    findes := map[string]int{"Saturday": 1, "Sunday": 1}
    now := time.Now()
    day := fmt.Sprintf("%s", now.Weekday())
    _, ok := findes[day]
    return ok
}

// filtra los servicios dependiendo de la hora
func filtraServicios(servicios Servicios, destino string) Servicios  {
    // TODO: añadir destino como argumento
    // y número total de expediciones
    var proximosServicios[]Servicio
    now := time.Now()

    // comprobamos el horario que necesitamos
    horario := "diario"
    if esFinde() {
        horario = "festivo"
        if esVerano() {
            horario = "verano"
        }
    }
    // recorremos la lista de servicios e imprimimos los próximos
    fmt.Println(horario, destino)
    for _, servicio := range servicios {
        //fmt.Println(servicio)
        if servicio.destino == destino {
            if servicio.horario == horario || servicio.horario == "*" {
                if servicio.h_salida >= now.Hour() {
                    if servicio.m_salida >= now.Hour() {
                        proximosServicios = append(proximosServicios, servicio)
                    } else {
                    }
                }
            }
        }
    }
    return proximosServicios
}

// creamos un struct Servicio con la información relativa a cada servicio de autobus
type Servicio struct {
    horario string
    linea string
    destino string
    h_salida int
    m_salida int
    //hora Time
}

// Servicios es un slice de elementos Servicio
type Servicios []Servicio

// ordenados Servicios por hora
type PorHora []Servicio

func (a PorHora) Len() int {
    return len(a)
}

func (a PorHora) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a PorHora) Less(i, j int) bool {
    return a[i].h_salida < a[j].h_salida || (a[i].h_salida == a[j].h_salida && a[i].m_salida < a[j].m_salida )
}

func main() {
    // carga lista de servicios
    servicios := cargaServicios("data/528.data", "data/539.data", "data/54N.data", "data/529.data")

    p := fmt.Printf
    now := time.Now()
    p("%s %d %s %d, a las %d:%d\n", now.Weekday(), now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute())

    busesAMadrid := filtraServicios(servicios, "madrid")
    p("a Madrid\n")
    sort.Sort(PorHora(busesAMadrid))
    for _, item := range busesAMadrid{
        p("  - %s -> %d:%d\n", item.linea, item.h_salida, item.m_salida)
    }

    busesAMostoles := filtraServicios(servicios, "móstoles")
    p("a Móstoles\n")
    sort.Sort(PorHora(busesAMadrid))
    for _, item := range busesAMostoles{
        p("  - %s -> %d:%d\n", item.linea, item.h_salida, item.m_salida)
    }

    busesANaval := filtraServicios(servicios, "navalcarnero")
    p("a Navalcarnero\n")

    sort.Sort(PorHora(busesANaval))
    for _, item := range busesANaval{
        p("  - %s -> %d:%d\n", item.linea, item.h_salida, item.m_salida)
   }

}
