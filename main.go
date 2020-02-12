package main

import (
	"fmt"
	"reflect"
)

type Worker interface {
	printPosition()
	printNameOfWork()
	printNameOfWorker()
}

//person

type person struct {
	age         int
	gender      string
	sitizenShip string
	lastName    string
	firstName   string
}

func newPerson(age int, gender, sitizenShip, lastName, firstName string) person {
	return person{age: age, gender: gender, sitizenShip: sitizenShip, lastName: lastName, firstName: firstName}
}

//person

//programming

type programmer struct {
	person
	position   string
	nameOfWork string
}

func (p programmer) printPosition() {
	fmt.Println("Worker position: ", p.position)
}
func (p programmer) printNameOfWork() {
	fmt.Println("Name of work: ", p.nameOfWork)
}
func (p programmer) printNameOfWorker() {
	fmt.Println("Worker name: ", p.firstName, " ", p.lastName)
}
func newProgrammer(age int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) programmer {
	return programmer{
		person:     newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork: nameOfWork,
		position:   position,
	}
}

//programming

//bank

type banker struct {
	person
	position   string
	nameOfWork string
}

func (b banker) printPosition() {
	fmt.Println("Worker position: ", b.position)
}
func (b banker) printNameOfWork() {
	fmt.Println("Name of work: ", b.nameOfWork)
}
func (b banker) printNameOfWorker() {
	fmt.Println("Worker name: ", b.firstName, " ", b.lastName)
}
func newBanker(age int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) banker {
	return banker{
		person:     newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork: nameOfWork,
		position:   position,
	}
}

//bank

//taxi

type taxiDriver struct {
	person
	position   string
	nameOfWork string
}

func (t taxiDriver) printPosition() {
	fmt.Println("Worker position: ", t.position)
}
func (t taxiDriver) printNameOfWork() {
	fmt.Println("Name of work: ", t.nameOfWork)
}
func (t taxiDriver) printNameOfWorker() {
	fmt.Println("Worker name: ", t.firstName, " ", t.lastName)
}
func newTaxiDriver(age int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) taxiDriver {
	return taxiDriver{
		person:     newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork: nameOfWork,
		position:   position,
	}
}

//taxi

//school
type teacher struct {
	person
	position   string
	nameOfWork string
}

func (t teacher) printPosition() {
	fmt.Println("Worker position: ", t.position)
}
func (t teacher) printNameOfWork() {
	fmt.Println("Name of work: ", t.nameOfWork)
}
func (t teacher) printNameOfWorker() {
	fmt.Println("Worker name: ", t.firstName, " ", t.lastName)
}
func newTeacher(age int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) teacher {
	return teacher{
		person:     newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork: nameOfWork,
		position:   position,
	}
}

//school

//factory
type workerOfFactory struct {
	person
	position   string
	nameOfWork string
}

func (w workerOfFactory) printPosition() {
	fmt.Println("Worker position: ", w.position)
}
func (w workerOfFactory) printNameOfWork() {
	fmt.Println("Name of work: ", w.nameOfWork)
}
func (w workerOfFactory) printNameOfWorker() {
	fmt.Println("Worker name: ", w.firstName, " ", w.lastName)
}
func newWorkerOfFactory(age int, gender, sitizenShip, lastName, firstName string, nameOfWork, position string) workerOfFactory {
	return workerOfFactory{
		person:     newPerson(age, gender, sitizenShip, lastName, firstName),
		nameOfWork: nameOfWork,
		position:   position,
	}
}

//factory

//other

func getType(workers map[string]Worker) (types map[Worker]reflect.Type) {
	types = make(map[Worker]reflect.Type, len(workers))
	for _, value := range workers {
		types[value] = reflect.TypeOf(value)
	}
	return
}

//other

func main() {
	workers := make(map[string]Worker, 5)
	teacher := newTeacher(
		30,
		"Male",
		"Ukrainian",
		"Sydorenko",
		"Denis",
		"School",
		"Teacher",
	)
	workers[teacher.lastName+" "+teacher.firstName] = teacher
	taxiDriver := newTaxiDriver(
		45,
		"Male",
		"American",
		"Elton",
		"John",
		"Taxi",
		"Driver",
	)
	workers[taxiDriver.lastName+" "+taxiDriver.firstName] = taxiDriver
	programmer := newProgrammer(
		25,
		"Female",
		"Indian",
		"Ghas",
		"Matum",
		"Programming",
		"Programmer",
	)
	workers[programmer.lastName+" "+programmer.firstName] = programmer
	banker := newBanker(
		32,
		"Female",
		"Canadian",
		"Pokidko",
		"Danya",
		"Bank",
		"Banker",
	)
	workers[banker.lastName+" "+banker.firstName] = banker
	factoryWorker := newWorkerOfFactory(
		15,
		"Male",
		"Chinese",
		"Kolbaca",
		"Vasilek",
		"Factory",
		"Worker",
	)
	workers[factoryWorker.lastName+" "+factoryWorker.firstName] = factoryWorker
	for _, value := range workers {
		value.printNameOfWorker()
		value.printNameOfWork()
		value.printPosition()
		fmt.Println()
	}
	types := getType(workers)
	for _, value := range types {
		fmt.Println(value)
	}
}
