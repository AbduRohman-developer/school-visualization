package main

// varsion 0.2

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func configPupil(filename string, pupils []pupil) ([]pupil, error) {
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Can't close file!")
			return
		}
	}(file)
	if err != nil {
		return pupils, err
	}
	touch, err2 := io.ReadAll(file)
	if err2 != nil {
		return pupils, err2
	}
	err3 := json.Unmarshal(touch, &pupils)
	if err3 != nil {
		return pupils, err3
	}
	return pupils, err3
}

func configTeacher(filename string, teachers []teacher) ([]teacher, error) {
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Can't close file!")
			return
		}
	}(file)
	if err != nil {
		return teachers, err
	}
	touch, err2 := io.ReadAll(file)
	if err2 != nil {
		return teachers, err2
	}
	err3 := json.Unmarshal(touch, &teachers)
	if err3 != nil {
		return teachers, err3
	}
	return teachers, err3
}

func configClass(filename string, cls []class) ([]class, error) {
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Can't close file!")
			return
		}
	}(file)
	if err != nil {
		return cls, err
	}
	touch, err2 := io.ReadAll(file)
	if err2 != nil {
		return cls, err2
	}
	err3 := json.Unmarshal(touch, &cls)
	if err3 != nil {
		return cls, err3
	}
	return cls, err3
}
