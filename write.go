package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func writeTeacher(filename string, teachers []teacher) {
	data, err := json.Marshal(teachers)
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("can't write data to file")
	}
}

func writeClass(filename string, classes []class) {
	data, err := json.Marshal(classes)
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("can't write data to file")
	}
}

func writePupil(filename string, pupils []pupil) {
	data, err := json.Marshal(pupils)
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("can't write data to file")
	}
}
