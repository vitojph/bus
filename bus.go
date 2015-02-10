package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "time"
    "strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// carga la lista de servicios a partir de los ficheros *.data
func cargaServicios(filenames ...string) Servicios {
    var servicios[]Servicio

    // cargo los horarios
    for _, filename := range filenames {
        //fmt.Println(filename)
        lines, err := readLines(filename)
        check(err)
        for _, line := range lines {
            if !strings.HasPrefix(line, "//") {
                fields := strings.Split(line, ",")
                h_salida, err := strconv.Atoi(fields[3])
                check(err)
                m_salida, err := strconv.Atoi(fields[4])
                check(err)

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
    var horario string
    if !esFinde() {
        horario = "diario"
        if esVerano() {
            horario = "verano"
        }
    } else {
        horario = "festivo"
    }

    // recorremos la lista de servicios e imprimimos los próximos
    for _, servicio := range servicios {
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


func main() {
    // carga lista de servicios
    servicios := cargaServicios("data/528-539.data", "data/529.data")

    p := fmt.Printf
    now := time.Now()
    p("%s %d %s %d, a las %d:%d\n", now.Weekday(), now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute())

    busesAMadrid := filtraServicios(servicios, "madrid")
    p("a Madrid\n")
    for _, item := range busesAMadrid{
        p("  - %s -> %d:%d\n", item.linea, item.h_salida, item.m_salida)
    }

    busesANaval := filtraServicios(servicios, "navalcarnero")
    p("a Navalcarnero\n")
    for _, item := range busesANaval{
        p("  - %s -> %d:%d\n", item.linea, item.h_salida, item.m_salida)
   }

}
