package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type MBR struct {
	Size          int32
	CreationDate  int64
	DiskSignature int32
}

func executeCommand(args []string) { // comando de consola "execute" para el script  ej.  'execute archivo_script.txt'
	if len(args) == 0 {
		fmt.Println("Se requiere un archivo de script.txt ")
		return
	}

	scriptFile := args[0]
	file, err := os.Open(scriptFile)
	if err != nil {
		fmt.Println("Error al abrir el archivo:  ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Ejecutando commando: ", line)
		if strings.HasPrefix(line, "# ") {
			fmt.Println(line) // Mostrar comentarios
		} else {
			processCommand(line) // Procesar comando
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo: ", err)
	}
}

func mkdiskCommand(args []string) { // comando de consola "mkdisk" dentro del script  ej.  'mkdisk -size 2 -name disco1'
	if len(args) < 4 {
		fmt.Println("Uso: mkdisk -size [tamaño_bytes] -name [nombre_disco]")
		return
	}

	var size int
	var name string

	for i := 0; i < len(args); i += 2 {
		if args[i] == "-size" {
			size, _ = strconv.Atoi(args[i+1])
		} else if args[i] == "-name" {
			name = args[i+1]
		}
	}

	if size == 0 || name == "" {
		fmt.Println("Parámetros insuficientes. ")
		return
	}

	createDisk(int32(size*1024*1024), name) // llamar metodo para crear el disco simulado en archivo binario.adjs
}

func createDisk(size int32, name string) {
	file, err := os.Create(name + ".adsj")
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err) // no se pudo crear el archivo binario.adjs
		return
	}
	defer file.Close()

	emptyChar := byte(0)
	for i := int32(0); i < size; i++ {
		file.Write([]byte{emptyChar}) // Inicializar el archivo con el carácter vacío
	}

	mbr := MBR{
		Size:          size,
		CreationDate:  time.Now().Unix(),
		DiskSignature: int32(rand.Intn(100000)),
	}

	file.Seek(0, 0)
	err = binary.Write(file, binary.LittleEndian, &mbr)
	if err != nil {
		fmt.Println("Error al escribir el MBR: ", err) // no se pudo escribir en el archivo binario.adjs
		return
	}

	fmt.Println("Disco creado exitosamente. ") //  MBR creado
}

func repCommand(args []string) { // comando de consola "rep" dentro del script  ej.  'rep -name disco1'
	if len(args) == 0 {
		fmt.Println("Uso: rep -name [nombre_disco]")
		return
	}

	var name string
	for i := 0; i < len(args); i += 2 {
		if args[i] == "-name" {
			name = args[i+1]
		}
	}

	if name == "" {
		fmt.Println("Nombre de disco no proporcionado. ")
		return
	}

	file, err := os.Open(name + ".adsj")
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}
	defer file.Close()

	var mbr MBR
	err = binary.Read(file, binary.LittleEndian, &mbr)
	if err != nil {
		fmt.Println("Error al leer el MBR: ", err)
		return
	}
	//para desplegar informacion del disco
	fmt.Println("Reporte del MBR: ")
	fmt.Println("Tamaño del disco: ", mbr.Size, " bytes")
	fmt.Println("Fecha de creación: ", time.Unix(mbr.CreationDate, 0))
	fmt.Println("Firma del disco: ", mbr.DiskSignature)
}
