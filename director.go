package main

import "fmt"

func (dr director) directorLog(ts *[]teacher, cls *classes) {
	// this function is for redirecting you on the path
	for {
		fmt.Print(`What do you want to do -->\n
                   1 - add teacher
                   2 - add class
                   3 - see teacher list
                   4 - nothing
                  `)
		option := ""
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("You should pick one not create")
		} else {
			switch option {
			case "1":
				dr.addTeacher(ts, cls)
			case "2":
				dr.addClass(ts)
			case "3":

			default:
				return
			}
		}
	}
}

func (dr director) addTeacher(ts *[]teacher, cls *classes) {
	// the body of function is related recursively and to continue you should enter inputs correctly
	// this will correct up all errors as possible
	for {
		t := teacher{}
		fmt.Print("What is his/her name -->")
		_, err := fmt.Scan(&t.name)
		if err != nil {
			fmt.Println("Please check out the name and reenter!")
		} else {
			for {
				fmt.Print("How old is he/or -->")
				_, err := fmt.Scan(&t.age)
				if err != nil {
					fmt.Println("Please check out the name and reenter!")
				} else {
					fmt.Println("If you want to create a new group fill it, else type 'not' to not create a new group!")
					_, err := fmt.Scan(&t.className)
					if err != nil {
						fmt.Println("Please check the input and type again!")
					} else if t.className == "not" {
						return
					} else {
						cl := class{}
						for {
							fmt.Print("What is the group name -->")
							_, err := fmt.Scan(&cl.name)
							if err != nil {
								fmt.Println("Please check out the name and enter again!")
							} else {
								cl.teacherName = t.name
								*cls = append(*cls, cl)
								t.className = cl.name
								*ts = append(*ts, t)
								return
							}
						}
					}
				}
			}
		}
	}
}

func (dr director) addClass(ts *[]teacher) {
	cl := class{}
	for {
		fmt.Print("What is name of class -->")
		_, err := fmt.Scan(&cl.name)
		if err != nil {
			fmt.Println("Please check your input, and try again later! ")
		} else {
			for {
				fmt.Print("If you need a new teacher type name, else type 'not' to continue --> ")
				_, err := fmt.Scan(&cl.teacherName)
				if err != nil {
					fmt.Println("Please check your input, and try again later! ")
				} else if cl.teacherName == "not" {
					for i, v := range *ts {
						fmt.Println(i+1, "-", v.name, "Class name:", v.className)
					}
					fmt.Print("Pick one of them -->")
					pick := 0
					_, err := fmt.Scan(&pick)
					if err != nil {
						fmt.Println("You could not use this!")
						cl.teacherName = ""
						return
					}
					if pick <= len(*ts) || pick >= 0 {
						cl.teacherName = (*ts)[pick-1].name
						return
					} else {
						fmt.Println("Your request is not proper!")
					}
				} else {
					t := teacher{
						name: cl.teacherName,
					}
					*ts = append(*ts, t)
					return
				}
			}
		}
	}
}
