package main

import (
	"fmt"
)

// version 0.2

func (dr director) directorLog(ts *[]teacher, ps *[]pupil, cls *[]class) {
	defer cleanUp()
	for {
		fmt.Print(`
		1 - view what we have 
		2 - add teacher
		3 - add class
		4 - remove teacher
		5 - remove class
		6 - nothing
	`)
		fmt.Println("What do you want to do -->")
		cmd := 0
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch cmd {
		case 1:
			dr.viewStat(cls, ps)
		case 2:
			dr.addTeacher(ts, cls)
		case 3:
			dr.addClass(ts, cls)
		case 4:
			dr.removeTeacher(ts, cls)
		case 5:
			dr.removeClass(ts, cls, ps)
		case 6:
			return
		default:
			fmt.Println("command not found")
			return
		}
	}
}

func (dr director) removeClass(ts *[]teacher, cls *[]class, ps *[]pupil) {
	defer clean()
	name := ""
	fmt.Println("Enter name of class to remove --> ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("incorrect typing!")
		return
	}
	check := 0
	for i := range *cls {
		if name == (*cls)[i].Name {
			if len(*cls)-1 == 0 {
				*cls = []class{}
				check++
				break
			}
			if i == len(*cls)-1 {
				*cls = (*cls)[:i]
				check++
				break
			}
			*cls = append((*cls)[:i], (*cls)[i+1:]...)
			check++
			break
		}
	}
	if check == 0 {
		fmt.Println("can't find!")
		return
	}
	for i := range *ts {
		if (*ts)[i].ClassName == name {
			(*ts)[i].ClassName = ""
		}
	}
	for i := range *ps {
		if (*ps)[i].ClassName == name {
			if len(*ps)-1 == 0 {
				*ps = []pupil{}
				return
			}
			if len(*ps)-1 == i {
				*ps = (*ps)[:i-1]
				print("hey")
				break
			}
			*ps = append((*ps)[:i], (*ps)[i+1:]...)
		}
	}
}

func (dr director) removeTeacher(ts *[]teacher, cls *[]class) {
	defer clean()
	name := ""
	fmt.Println("Type teacher name to remove --> ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("failed to read")
	}
	check := 0
	for i := range *ts {
		if name == (*ts)[i].Name {
			if len(*ts) == 1 {
				*ts = []teacher{}
				check++
				break
			}
			if i == len(*ts)-1 {
				*ts = (*ts)[:i-1]
				check++
				break
			}
			*ts = append((*ts)[:i], (*ts)[i+1:]...)
			check++
		}
	}
	if check == 0 {
		fmt.Println("can't find")
	}
	for i := range *cls {
		if name == (*cls)[i].TeacherName {
			(*cls)[i].TeacherName = ""
		}
	}
}

func (dr director) addClass(ts *[]teacher, cls *[]class) {
	defer clean()
	c := class{}
	fmt.Print("Type class name to add --> ")
	_, err := fmt.Scan(&c.Name)
	if err != nil {
		fmt.Println("can't read")
		return
	}
	fmt.Print("Type teacher name that will given to the class \n(if you enter an existing one, its teacher will replace with that one!!!)--> ")
	_, err = fmt.Scan(&c.TeacherName)
	if err != nil {
		fmt.Println("can't read")
		return
	}
	check := 0
	for i := range *ts {
		if c.TeacherName == (*ts)[i].Name {
			(*ts)[i].ClassName = c.Name
			check++
		}
	}
	if check != 0 {
		for i := range *cls {
			if (*cls)[i].TeacherName == c.TeacherName {
				(*cls)[i].Name = c.Name
				return
			}
		}
	}
	*cls = append(*cls, c)
}

func (dr director) viewStat(cls *[]class, ps *[]pupil) {
	defer clean()
	counter := 0
	//w := tabwriter.NewWriter(os.Stdout,5,1,1,' ',tabwriter.AlignRight|tabwriter.Debug)
	for i := range *cls {
		counter = 0
		fmt.Println("Class name:", (*cls)[i].Name, "Teacher name:", (*cls)[i].TeacherName)
		for j := range *ps {
			if (*ps)[j].ClassName == (*cls)[i].Name {
				counter++
				fmt.Println((*ps)[j].Name)
			}
		}
		//if counter == 0 {
		//	fmt.Println("nil")
		//}
	}
}

func (dr director) addTeacher(ts *[]teacher, cls *[]class) {
	defer clean()
	t := teacher{}
	fmt.Print("What his/her name -->")
	_, err := fmt.Scan(&t.Name)
	if err != nil {
		fmt.Println("Type properly! error = ", err)
		return
	}
	fmt.Print("How old is he/she -->")
	_, err = fmt.Scan(&t.Age)
	if err != nil {
		panic(err)
	}
	fmt.Print("Which class is for him (if you type exists class name, the teacher name will change!!!) \nOr else type 'not' to create empty class name --> ")
	_, err = fmt.Scan(&t.ClassName)
	if err != nil {
		fmt.Println("Type properly! error = ", err)
		print("err")
		return
	} else if t.ClassName == "not" {
		print("not")
		t.ClassName = ""
		*ts = append(*ts, t)
		return
	}
	for i := range *cls {
		if (*cls)[i].Name == t.ClassName {
			(*cls)[i].TeacherName = t.Name
			*ts = append(*ts, t)
			return
		}
	}
	cl := class{
		Name:        t.ClassName,
		TeacherName: t.Name,
	}
	*cls = append(*cls, cl)
	// append data to teacher array
	*ts = append(*ts, t)
}

func clean() {
	if err := recover(); err != nil {
		fmt.Println("we can't do this(but commands until there is saved)")
	}
}

func cleanUp() bool {
	if err := recover(); err != nil {
		fmt.Println("we can't do this(but commands until there is saved)")
		return true
	} else {
		return false
	}
}
