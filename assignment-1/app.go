package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id      string
	Name    string
	Address string
	Job     string
	Reason  string
}

var students []Student = []Student{
	{
		Id:      "S1",
		Name:    "Giuseppe",
		Address: "Surabaya",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding",
	},
	{
		Id:      "S2",
		Name:    "Mefaza",
		Address: "Surabaya",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding",
	},
	{
		Id:      "S3",
		Name:    "Fauziann",
		Address: "Jombang",
		Job:     "Student",
		Reason:  "Karena saya suka ngoding",
	},
}

func main() {
	var inputs = os.Args

	result, err := FindStudent(inputs[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("\nID :" + result.Id)
	fmt.Printf("\nNama :" + result.Name)
	fmt.Printf("\nAddress :" + result.Address)
	fmt.Printf("\nJob :" + result.Job)
	fmt.Printf("\nReason :" + result.Reason)
}

func FindStudent(studentId string) (Student, error) {
	for _, value := range students {
		if value.Id == studentId {
			return value, nil
		}
	}

	return Student{}, errors.New("Student not found")
}
