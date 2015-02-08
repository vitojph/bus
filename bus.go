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


// filtra los servicios dependiendo de la hora
func filtraServicios(servicios Servicios, destino string) Servicios  {
    // TODO: añadir destino como argumento
    // y número total de expediciones
    var proximosServicios[]Servicio
    now := time.Now()
    // recorremos la lista de servicios e imprimimos los próximos
    for _, servicio := range servicios {
        if servicio.destino == destino {
            if servicio.h_salida >= now.Hour() {
                if servicio.m_salida >= now.Minute(){
                    proximosServicios = append(proximosServicios, servicio)
                }
            }
        }
    }
    return proximosServicios
}

// creamos un struct Servicio con la información relativa a cada servicio de autobus
type Servicio struct {
    horario string
    linea int
    destino string
    h_salida int
    m_salida int
    //hora Time
}

// Servicios es un slice de elementos Servicio
type Servicios []Servicio


func main() {
    // creo el slice de servicios
    var servicios[]Servicio

    // cargo los horarios
    lines, err := readLines("/data/go/src/github.com/vitojph/bus/528.data")
    check(err)
    for _, line := range lines {
        if !strings.HasPrefix(line, "//") {
            fields := strings.Split(line, ",")
            linea, err := strconv.Atoi(fields[1])
            check(err)
            h_salida, err := strconv.Atoi(fields[3])
            check(err)
            m_salida, err := strconv.Atoi(fields[4])
            check(err)

            esteServicio := Servicio{horario: fields[0], linea: linea, destino: fields[2], h_salida: h_salida, m_salida: m_salida}
            servicios = append(servicios, esteServicio)
        }
    }

    p := fmt.Printf
    now := time.Now()
    p("Ahora es %s\n", now)

    busesAMadrid := filtraServicios(servicios, "madrid")
    p("Navalcarnero -> Madrid\n")
    for _, item := range busesAMadrid{
            fmt.Println(item)
    }

    busesANaval := filtraServicios(servicios, "navalcarnero")
    p("\nMadrid -> Navalcarnero\n")
    for _, item := range busesANaval{
            fmt.Println(item)
   }

}
