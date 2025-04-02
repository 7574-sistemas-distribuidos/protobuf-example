package main

import (
	"fmt"
	"log"
	"protobuf/personpb"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := &personpb.Person{
		Name:  proto.String("Juan Pérez"),
		Id:    proto.Int32(123),
		Email: proto.String("juan.perez@example.com"),
	}

	// La serializamos en binario usando proto
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("Error al serializar: ", err)
	}

	// Imprimir el tamaño del dato serializado
	fmt.Printf("Tamaño de los datos serializados con protobuf: %d bytes\n", len(data))

	// Deserializar desde binario
	newPerson := &personpb.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Error al deserializar: ", err)
	}

	// Mostrar el objeto deserializado
	fmt.Printf("Objeto deserializado: %+v\n", newPerson)

	/*
		¿Qué pasa si quisieramos represtar esto mismo en un protocolo de texto tipo TLV?

		Tipo (T): 1 byte ("1" para el campo name, "2" para id, y "3" para email).
		Longitud (L): Un entero de 4 bytes
		Valor (V): Cadena de texto, tendrá tantos bytes como Longitud indique.

		En el ejemplo:
		Name = 	1 byte (tipo) + 4 byte (longitud) + 10 byte (nombre) 	= 15 bytes
		Id = 	1 byte (tipo) + 4 byte (longitud) + 3 bytes (id) 		= 8 bytes
		Email = 1 byte (tipo) + 4 byte (longitud) + 24 bytes (mail) 	= 29 bytes

		TOTAL = 52 bytes
	*/
}
