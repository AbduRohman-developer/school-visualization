package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func directorLog(classes classes, teachers teachers, sch school, NPup int) {
	fmt.Print("1 - classes list\n2 - teachers list\n3 - number of pupil\n")
	var cmd int
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		w := tabwriter.NewWriter(os.Stdout, 5, 2, 2, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "\tClass name\tTeacher name\t")
		for i := range classes {
			fmt.Fprintf(w, "\t%s\t%s\t\n", classes[i].className, classes[i].teacherName)
		}
		w.Flush()
	case 2:
		w := tabwriter.NewWriter(os.Stdout, 5, 2, 2, ' ', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, "\tTeacher name\tClass name\t")
		for i := range teachers {
			fmt.Fprintf(w, "\t%s\t%s\t\n", teachers[i].name, teachers[i].className)
		}
		w.Flush()
	case 3:
		fmt.Printf("Number of pupil in %s is %d", sch.name, NPup)
	}
}
