package main

import "fmt"

// version 0.2

func pupilLog(ps *[]pupil, cls []class) {
	defer cleanUp()
	name := ""
	fmt.Print("What is your name --> ")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic("can't read the input")
	}
	check := 0
	for i := range *ps {
		if (*ps)[i].Name == name {
			check++
		}
	}
	if check == 0 {
		fmt.Println("can't find the name")
		return
	}
	fmt.Print(`
		1 - see my mark
		2 - see my teacher and class
		3 - change class
		4 - nothing
	`)
	cmd := 0
	_, err = fmt.Scan(&cmd)
	if err != nil {
		panic("can't")
	}
	p := pupil{}
	switch cmd {
	case 1:
		p.viewMark(name, ps)
	case 2:
		p.viewTeacher(name, cls, ps)
	case 3:
		p.changeClass(cls, name, ps)
	case 4:
		return
	default:
		panic("cheater!")
	}
}

func (p pupil) viewMark(name string, ps *[]pupil) {
	for i := range *ps {
		if (*ps)[i].Name == name {
			fmt.Println("Your mark is", (*ps)[i].Mark)
		}
	}
}

func (p pupil) viewTeacher(name string, cls []class, ps *[]pupil) {
	for i := range *ps {
		if name == (*ps)[i].Name {
			for j := range cls {
				if (*ps)[i].ClassName == cls[j].Name {
					fmt.Println("Class name:", (*ps)[i].ClassName, "	Teacher name", cls[j].TeacherName)
					return
				}
			}
			fmt.Println("Class name:", (*ps)[i].ClassName, "	Teacher name: nil")
			return
		}
	}
	fmt.Println("Class name:nil		Teacher name: nil")
	return
}

func (p pupil) changeClass(cls []class, name string, ps *[]pupil) {
	defer clean()
	fmt.Println("We have --> ")
	for i := range cls {
		fmt.Println(cls[i].Name, "course")
	}
	choice := ""
	fmt.Print("Enter exists class name to change up --> ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		panic("cheater!")
	}
	check := 0
	for i := range cls {
		if cls[i].Name == choice {
			check++
		}
	}
	if check != 0 {
		for i := range *ps {
			if name == (*ps)[i].Name {
				(*ps)[i].ClassName = choice
			}
		}
	} else {
		fmt.Println("can't find")
	}
}
