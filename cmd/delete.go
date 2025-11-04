package cmd

import (
	"fmt"
	"strconv"
	"log"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete a note",
	Run: deleteNote,
}

func deleteNote(cmd *cobra.Command, args []string) {
	ID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Value not integer: %s", err)
	}
	notes, err := getNotes("notes.csv")
	if err != nil {
		panic(err)
	}
	if _, exists := notes[ID]; exists {
		delete(notes, ID)
		writeNotes(notes)

	} else  {
		fmt.Println("ID does not exist")
	}

}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
