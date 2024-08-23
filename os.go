package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Alumno struct {
	nombre   string
	status   string
	notas    []float64
	promedio float64
}

func leerCSV(nombre_archivo string) ([]Alumno, error) {
	archivo, err := os.Open(nombre_archivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	lectura := csv.NewReader(archivo)
	lectura.Comma = ',' // Asegúrate de que coincide con el delimitador de tu CSV

	registros, err := lectura.ReadAll()
	if err != nil {
		return nil, err
	}

	var alumnos []Alumno

	for i, registro := range registros {
		if len(registro) < 3 {
			fmt.Printf("Advertencia: la fila %d no tiene suficientes columnas.\n", i+1)
			continue // Ignorar esta línea si no tiene al menos 3 columnas
		}

		nombre := strings.TrimSpace(registro[0])
		status := strings.TrimSpace(registro[1])

		// Convertir las notas de string a float64
		var notas []float64
		for _, notaStr := range registro[2:] {
			nota, err := strconv.ParseFloat(strings.TrimSpace(notaStr), 64)
			if err != nil {
				return nil, fmt.Errorf("error al convertir la nota '%s' en la fila %d: %v", notaStr, i+1, err)
			}
			notas = append(notas, nota)
		}

		// Crear un objeto Alumno y calcular el promedio
		alumno := Alumno{
			nombre:   nombre,
			status:   status,
			notas:    notas,
			promedio: calcularPromedio(notas),
		}
		alumnos = append(alumnos, alumno)
	}

	return alumnos, nil
}

func calcularPromedio(notas []float64) float64 {
	var suma float64

	for _, nota := range notas {
		suma += nota
	}

	return suma / float64(len(notas))
}

func crearArchivo(nombre_archivo string, alumnos []Alumno) error {
	archivo, err := os.Create(nombre_archivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escritura := csv.NewWriter(archivo)
	defer escritura.Flush()

	for _, alumno := range alumnos {
		// Crear un slice de strings para almacenar los valores a escribir en el archivo
		var registro []string
		registro = append(registro, alumno.nombre, alumno.status)
		for _, nota := range alumno.notas {
			registro = append(registro, fmt.Sprintf("%.2f", nota))
		}
		registro = append(registro, fmt.Sprintf("%.2f", alumno.promedio))

		// Escribir el registro en el archivo
		if err := escritura.Write(registro); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	nombre_archivo := "notas.csv"
	alumnos, err := leerCSV(nombre_archivo)
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	for _, alumno := range alumnos {
		fmt.Printf("Nombre: %s, Status: %s, Notas: %v, Promedio: %.2f\n", alumno.nombre, alumno.status, alumno.notas, alumno.promedio)
	}

	// Crear el nuevo archivo con los datos y promedios
	nuevo_archivo := "notas_con_promedio.csv"
	if err := crearArchivo(nuevo_archivo, alumnos); err != nil {
		fmt.Println("Error al crear el archivo CSV:", err)
		return
	}

	fmt.Println("Archivo creado exitosamente:", nuevo_archivo)
}
