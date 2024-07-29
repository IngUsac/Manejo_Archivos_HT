package main

import "fmt"

// Estructura Maestro
type Maestro struct {
	ID      int
	Nombre  string
	Materia string
}

// Estructura Estudiante
type Estudiante struct {
	ID     int
	Nombre string
	Grado  string
}

// Mapa para gestionar la relación muchos a muchos
var maestrosEstudiantes = make(map[int][]int)
var estudiantesMaestros = make(map[int][]int)

// Función para asignar un estudiante a un maestro
func asignarEstudianteAMaestro(maestroID, estudianteID int) {
	maestrosEstudiantes[maestroID] = append(maestrosEstudiantes[maestroID], estudianteID)
	estudiantesMaestros[estudianteID] = append(estudiantesMaestros[estudianteID], maestroID)
}

// Función para imprimir la lista de estudiantes de un maestro
func listarEstudiantesDeMaestro(maestroID int, estudiantes []Estudiante) {
	fmt.Printf("Estudiantes del Maestro %d:\n", maestroID)
	for _, estudianteID := range maestrosEstudiantes[maestroID] {
		for _, estudiante := range estudiantes {
			if estudiante.ID == estudianteID {
				fmt.Printf("- %s\n", estudiante.Nombre)
			}
		}
	}
}

// Función para imprimir la lista de maestros de un estudiante
func listarMaestrosDeEstudiante(estudianteID int, maestros []Maestro) {
	fmt.Printf("Maestros del Estudiante %d:\n", estudianteID)
	for _, maestroID := range estudiantesMaestros[estudianteID] {
		for _, maestro := range maestros {
			if maestro.ID == maestroID {
				fmt.Printf("- %s\n", maestro.Nombre)
			}
		}
	}
}

func main() {
	// Crear ejemplos de maestros
	maestro1 := Maestro{ID: 1, Nombre: "Juan Pérez", Materia: "Matemáticas"}
	maestro2 := Maestro{ID: 2, Nombre: "Ana Gómez", Materia: "Historia"}

	// Crear ejemplos de estudiantes
	estudiante1 := Estudiante{ID: 1, Nombre: "Carlos Díaz", Grado: "5to"}
	estudiante2 := Estudiante{ID: 2, Nombre: "María López", Grado: "5to"}

	// Asignar estudiantes a maestros
	asignarEstudianteAMaestro(maestro1.ID, estudiante1.ID)
	asignarEstudianteAMaestro(maestro1.ID, estudiante2.ID)
	asignarEstudianteAMaestro(maestro2.ID, estudiante1.ID)

	// Listar estudiantes de un maestro
	listarEstudiantesDeMaestro(maestro1.ID, []Estudiante{estudiante1, estudiante2})

	// Listar maestros de un estudiante
	listarMaestrosDeEstudiante(estudiante1.ID, []Maestro{maestro1, maestro2})
}
