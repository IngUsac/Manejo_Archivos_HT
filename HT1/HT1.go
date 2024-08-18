package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type Profesor struct {
	IdProfesor int32
	Cui        int32
	Nombre     [50]byte
	Curso      [50]byte
}

type Estudiante struct {
	IdEstudiante int32
	Cui          int32
	Nombre       [50]byte
	Carne        [50]byte
}

const fileName = "registro.bin"

func main() {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error abriendo el archivo:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Registro de Profesor")
		fmt.Println("2. Registro de Estudiante")
		fmt.Println("3. Ver Registros")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			registrarProfesor(file, reader)
		case "2":
			registrarEstudiante(file, reader)
		case "3":
			verRegistros(file)
		case "4":
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func registrarProfesor(file *os.File, reader *bufio.Reader) {
	var profesor Profesor
	var marker byte = 'P'
	fmt.Print("Ingrese ID Profesor: ")
	fmt.Fscan(reader, &profesor.IdProfesor)

	if !idUnicoProfesor(file, profesor.IdProfesor) {
		fmt.Println("ID Profesor ya existe. Intente con otro ID.")
		return
	}

	fmt.Print("Ingrese CUI: ")
	fmt.Fscan(reader, &profesor.Cui)
	reader.ReadString('\n') // Consume el salto de línea

	fmt.Print("Ingrese Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
	copy(profesor.Nombre[:], nombre)

	fmt.Print("Ingrese Curso: ")
	curso, _ := reader.ReadString('\n')
	curso = strings.TrimSpace(curso)
	copy(profesor.Curso[:], curso)

	file.Seek(0, os.SEEK_END) // Asegurar que escribimos al final del archivo
	file.Write([]byte{marker})
	if err := binary.Write(file, binary.LittleEndian, &profesor); err != nil {
		fmt.Println("Error escribiendo registro de profesor:", err)
	}
}

func registrarEstudiante(file *os.File, reader *bufio.Reader) {
	var estudiante Estudiante
	var marker byte = 'E'
	fmt.Print("Ingrese ID Estudiante: ")
	fmt.Fscan(reader, &estudiante.IdEstudiante)

	if !idUnicoEstudiante(file, estudiante.IdEstudiante) {
		fmt.Println("ID Estudiante ya existe. Intente con otro ID.")
		return
	}

	fmt.Print("Ingrese CUI: ")
	fmt.Fscan(reader, &estudiante.Cui)
	reader.ReadString('\n') // Consume el salto de línea

	fmt.Print("Ingrese Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
	copy(estudiante.Nombre[:], nombre)

	fmt.Print("Ingrese Carne: ")
	carne, _ := reader.ReadString('\n')
	carne = strings.TrimSpace(carne)
	copy(estudiante.Carne[:], carne)

	file.Seek(0, os.SEEK_END) // Asegurar que escribimos al final del archivo
	file.Write([]byte{marker})
	if err := binary.Write(file, binary.LittleEndian, &estudiante); err != nil {
		fmt.Println("Error escribiendo registro de estudiante:", err)
	}
}

func verRegistros(file *os.File) {
	file.Seek(0, 0)
	var profesor Profesor
	var estudiante Estudiante
	var profesores []Profesor
	var estudiantes []Estudiante

	for {
		var marker byte
		if err := binary.Read(file, binary.LittleEndian, &marker); err != nil {
			break
		}
		if marker == 'P' {
			if err := binary.Read(file, binary.LittleEndian, &profesor); err != nil {
				break
			}
			profesores = append(profesores, profesor)
		} else if marker == 'E' {
			if err := binary.Read(file, binary.LittleEndian, &estudiante); err != nil {
				break
			}
			estudiantes = append(estudiantes, estudiante)
		}
	}

	fmt.Println("Registros de Profesores:")
	for _, profesor := range profesores {
		fmt.Printf("Profesor - ID: %d, CUI: %d, Nombre: %s, Curso: %s\n", profesor.IdProfesor, profesor.Cui, profesor.Nombre[:], profesor.Curso[:])
	}

	fmt.Println("Registros de Estudiantes:")
	for _, estudiante := range estudiantes {
		fmt.Printf("Estudiante - ID: %d, CUI: %d, Nombre: %s, Carne: %s\n", estudiante.IdEstudiante, estudiante.Cui, estudiante.Nombre[:], estudiante.Carne[:])
	}
}

func idUnicoProfesor(file *os.File, id int32) bool {
	file.Seek(0, 0)
	var profesor Profesor
	var estudiante Estudiante

	for {
		var marker byte
		if err := binary.Read(file, binary.LittleEndian, &marker); err != nil {
			break
		}
		if marker == 'P' {
			if err := binary.Read(file, binary.LittleEndian, &profesor); err != nil {
				break
			}
			if profesor.IdProfesor == id {
				return false
			}
		} else if marker == 'E' {
			binary.Read(file, binary.LittleEndian, &estudiante)
		}
	}
	return true
}

func idUnicoEstudiante(file *os.File, id int32) bool {
	file.Seek(0, 0)
	var profesor Profesor
	var estudiante Estudiante

	for {
		var marker byte
		if err := binary.Read(file, binary.LittleEndian, &marker); err != nil {
			break
		}
		if marker == 'E' {
			if err := binary.Read(file, binary.LittleEndian, &estudiante); err != nil {
				break
			}
			if estudiante.IdEstudiante == id {
				return false
			}
		} else if marker == 'P' {
			binary.Read(file, binary.LittleEndian, &profesor)
		}
	}
	return true
}
