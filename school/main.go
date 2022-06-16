package main

import (
	"fmt"
)

// version 0.2

var (
	cls []class
	ts  []teacher
	ps  []pupil
	dr  = director{
		name: "Akbarshoh",
		age:  "19",
	}
	sch = school{
		name:     "Restart",
		director: dr,
		cls:      cls,
	}
)

func main() {
	defer clean()
	// reads items from json file
	pupils, err := configPupil("pupiljs.json", ps)
	if err != nil {
		fmt.Println(err)
	}
	classes, err := configClass("classjs.json", cls)
	if err != nil {
		fmt.Println(err)
	}
	teachers, err := configTeacher("teacherjs.json", ts)
	if err != nil {
		fmt.Println(err)
	}
	// give you a choice
	fmt.Printf("Welcome to %s\n", sch.name)
	for {
		fmt.Println("Who are you -->")
		fmt.Print(`
                          1 - director
                          2 - teacher
                          3 - pupil
                          4 - no one (exit)
                          `)
		person := ""
		_, err := fmt.Scan(&person)
		if err != nil {
			fmt.Printf("failed to perform operation: %s", err)
			return
		} else {
			switch person {
			case "1":
				dr.directorLog(&teachers, &pupils, &classes)
			case "2":
				teacherLog(&teachers, &pupils, &classes)
			case "3":
				pupilLog(&pupils, classes)
			default:
				return
			}
		}
		writeTeacher("teacherjs.json", teachers)
		writeClass("classjs.json", classes)
		writePupil("pupiljs.json", pupils)
		if cleanUp() {
			return
		}
	}
}
