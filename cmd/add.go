package cmd

import (
	"maps"
	"slices"
	"github.com/spf13/cobra"
	"time"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a note",
	Args: cobra.ExactArgs(1),
	Run: addNote,
}

func addNote(cmd *cobra.Command, args []string) {
	notes, err := getNotes("notes.csv")
	if err != nil {
		panic(err)
	}
	newID := 1
	for _, i := range slices.Sorted(maps.Keys(notes)) {
		if newID != i {
			break
		}
		newID += 1
	}
	notes[newID] = Note{args[0], time.Now().Format("2006-01-02T15:04:05-07:00"), false}
	writeNotes(notes)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
