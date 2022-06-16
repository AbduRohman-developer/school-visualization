package main

import (
	"fmt"
)

// version 0.2

func teacherLog(ts *[]teacher, ps *[]pupil, cls *[]class) {
	defer cleanUp()
	t := teacher{}
	name := ""
	fmt.Print("What is your name -->")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("can't read input")
	}
	check := 0
	for i := range *ts {
		if (*ts)[i].Name == name {
			check++
		}
	}
	if check == 0 {
		fmt.Println("There is not teacher with this name :(")
		return
	}
	for {
		fmt.Println("What do you want to do --> ")
		fmt.Print(`
			1 - list of pupils
			2 - add pupil
			3 - remove pupil
			4 - mark pupils
			5 - nothing to do(exit)
	`)
		choice := 0
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Println("can't read input")
			return
		}
		switch choice {
		case 1:
			t.viewStat(name, ps, cls)
		case 2:
			t.addPupil(ps, name, cls)
		case 3:
			t.removePupil(ps, name, cls)
		case 4:
			t.mark(ps, name, cls)
		case 5:
			return
		default:
			fmt.Println("can't find command")
			return
		}
	}
}
func (t teacher) viewStat(name string, ps *[]pupil, cls *[]class) {
	defer clean()
	clname, err := check(cls, name)
	if err != nil {
		fmt.Println(err)
		panic("failed to find class name")
	}
	if clname == "" {
		fmt.Println("can't find class")
		return
	}
	fmt.Println("Class name:", clname, "  Teacher name:", name, "\n")
	for i := range *ps {
		if clname == (*ps)[i].ClassName {
			fmt.Println("Name:", (*ps)[i].Name, "  Mark:", (*ps)[i].Mark)
		}
	}
}

func (t teacher) addPupil(ps *[]pupil, name string, cls *[]class) {
	defer clean()
	clname, err := check(cls, name)
	if err != nil {
		fmt.Println(err)
		panic("failed to find class name")
	}
	if clname == "" {
		fmt.Println("can't find class")
		return
	}
	p := pupil{}
	fmt.Print("What is his/her name --> ")
	_, err = fmt.Scan(&p.Name)
	if err != nil {
		fmt.Println("failed to read input")
		return
	}
	p.ClassName = clname
	fmt.Print("How old is he/she --> ")
	_, err = fmt.Scan(&p.Age)
	if err != nil {
		fmt.Println("failed to read input")
		return
	}
	fmt.Print("What is his/her parents name --> ")
	_, err = fmt.Scan(&p.ParentsName)
	if err != nil {
		fmt.Println("failed to read input")
		return
	}
	*ps = append(*ps, p)
}

func (t teacher) removePupil(ps *[]pupil, name string, cls *[]class) {
	defer clean()
	clname, err := check(cls, name)
	if err != nil {
		fmt.Println(err)
		panic("failed to find class name")
	}
	if clname == "" {
		fmt.Println("can't find class")
		return
	}
	pName := ""
	fmt.Print("Enter pupil name to remove --> ")
	_, err = fmt.Scan(&pName)
	if err != nil {
		fmt.Println("can't read input")
		return
	}
	for i := range *ps {
		if pName == (*ps)[i].Name && clname == (*ps)[i].ClassName {
			*ps = append((*ps)[:i], (*ps)[i+1:]...)
			return
		}
	}
	panic("can't find pupil in this group")
}

func (t teacher) mark(ps *[]pupil, name string, cls *[]class) {
	defer clean()
	clname, err := check(cls, name)
	pName := ""
	if err != nil {
		fmt.Println(err)
		panic("failed to find class name")
	}
	if clname == "" {
		fmt.Println("can't find class")
		return
	}
	fmt.Print("Enter name of pupil which you want to mark --> ")
	_, err = fmt.Scan(&pName)
	if err != nil {
		fmt.Println("can't read the input")
	}
	v := 0
	p := 0
	for i := range *ps {
		if (*ps)[i].Name == pName && (*ps)[i].ClassName == clname {
			p = i
			v++
		}
	}
	if v == 0 {
		fmt.Println("can't find pupil")
		return
	}
	mark := 0
	fmt.Print("Mark the pupil(0 - 100) --> ")
	_, err = fmt.Scan(&mark)
	if err != nil {
		panic("can't read the input")
	}
	if mark <= 100 && mark >= 0 {
		(*ps)[p].Mark = mark
	} else {
		fmt.Println("can't mark with this mark")
	}
}

func check(cls *[]class, name string) (string, error) {
	defer clean()
	check := 0
	var err error
	for i := range *cls {
		if name == (*cls)[i].TeacherName {
			check++
		}
	}
	clname := ""
	if check > 1 {
		check = 0
		for i := range *cls {
			if name == (*cls)[i].TeacherName {
				check++
				fmt.Println((*cls)[i].Name)
			}
		}
		fmt.Println("Type class name --> ")
		_, err = fmt.Scan(&clname)
		if err != nil {
			fmt.Println("can't find class")
			return clname, err
		}
	} else {
		for i := range *cls {
			if name == (*cls)[i].TeacherName {
				clname = (*cls)[i].Name
				return clname, err
			}
		}
	}
	return clname, err
}
