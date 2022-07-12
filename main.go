package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid1, _ := uuid.NewUUID()
	uuid2, _ := uuid.NewDCEGroup()
	uuid3 := uuid.NewMD5(uuid.UUID{}, []byte("test"))
	uuid4, _ := uuid.NewRandom()
	uuid5 := uuid.NewSHA1(uuid.UUID{}, []byte("test"))

	fmt.Println(uuid1.Version().String())
	fmt.Println(uuid1.String())

	fmt.Println()

	fmt.Println(uuid2.Version().String())
	fmt.Println(uuid2.String())

	fmt.Println()

	fmt.Println(uuid3.Version().String())
	fmt.Println(uuid3.String())

	fmt.Println()

	fmt.Println(uuid4.Version().String())
	fmt.Println(uuid4.String())

	fmt.Println()

	fmt.Println(uuid5.Version().String())
	fmt.Println(uuid5.String())
}
