package cmd

import (
	"fmt"
	"os"
	"syscall"
	"strconv"
	"encoding/csv"
	"slices"
	"maps"
)

func loadFile(filepath string, mode string) (*os.File, error) {
	option := 0
	if mode == "read" {
		option = os.O_RDWR|os.O_CREATE
	} else if mode == "write" {
		option = os.O_RDWR|os.O_CREATE|os.O_TRUNC
	} else {
		return nil, fmt.Errorf("No Mode")
	}
	f, err := os.OpenFile(filepath, option, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func getNotes(filepath string) (map[int]Note, error) {
	file, err := loadFile(filepath, "read")
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	closeFile(file)
	if err != nil {
		panic(err)
	}
	if len(data) > 0 {
		data = data[1:]
	}
	notes := make(map[int]Note)
	for _, row := range data {
		bool, err := strconv.ParseBool(row[3])
		if err != nil {
			panic(err)
		}
		ID, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}
		notes[ID] = Note{row[1], row[2], bool}
	}
	return notes, nil
}

func writeNotes(notes map[int]Note) {
	file, err := loadFile("notes.csv", "write")
	if err != nil {
		panic(err)
	}
	defer closeFile(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"ID","Description","CreatedAt","IsComplete"})
	for _, i := range slices.Sorted(maps.Keys(notes)) {
		note := notes[i]
		writer.Write([]string{strconv.Itoa(i), note.Description, note.CreatedAt, strconv.FormatBool(note.IsComplete)})
	}
}
