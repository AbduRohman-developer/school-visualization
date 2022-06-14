package main

import (
	"fmt"
)

// version 0.1

func main() {
	// teacher initial values as string, we are going to append values later
	ts := make([]teacher, 10)
	t1 := teacher{
		name:      "Shokir",
		age:       "27",
		className: "Node.js",
	}
	t2 := teacher{
		name:      "Tohir",
		age:       "27",
		className: "Go",
	}
	ts = []teacher{t1, t2}
	// school director init
	dr := director{
		name: "Akbarshoh",
		age:  "19",
	}
	// classes initial values as class, we are going to append values later
	cls := classes{
		class{
			name:        "Go",
			teacherName: "Tohir",
		},
		{
			name:        "Node.js",
			teacherName: "Shokir",
		},
	}
	// the school, we can add our own school later
	sch := school{
		name:     "Restart",
		director: dr,
		classes:  cls,
	}

	// initial pupil values, again we can add our pupils
	p1 := pupil{
		name:        "Akbarshoh",
		age:         "19",
		mark:        "91",
		parentsName: "Isoqjon",
	}
	p2 := pupil{
		name:        "Asror",
		age:         "21",
		mark:        "88",
		parentsName: "Dilshod",
	}

	p3 := pupil{
		name:        "Abror",
		age:         "20",
		mark:        "91",
		parentsName: "Sohibjon",
	}
	p4 := pupil{
		name:        "Timur",
		age:         "18",
		parentsName: "Qamarjon",
		mark:        "95",
	}
	p5 := pupil{
		name:        "Yodgor",
		age:         "19",
		parentsName: "Rovshan",
		mark:        "87",
	}
	p6 := pupil{
		name:        "Muhammad",
		age:         "20",
		parentsName: "Tohirjon",
		mark:        "83",
	}
	p := pupil{}
	t := teacher{}
	ps := []pupil{p1, p2, p3, p4, p5, p6}
	ps = append(ps, p1)
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
			fmt.Println("You should pick one not create")
		} else {
			switch person {
			case "1":
				dr.directorLog(&ts, &cls)
			case "2":
				t.teacherLog()
			case "3":
				p.pupilLog()
			default:
				return
			}
		}
	}
}
