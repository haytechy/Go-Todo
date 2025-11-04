package cmd

import (
	"fmt"
	"strconv"
	"log"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
)

var completeCmd = &cobra.Command{
	Use: "complete",
	Short: "Mark a note as complete",
	Run: completeNote,
}

func completeNote(cmd *cobra.Command, args []string) {
	ID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Value not integer: %s", err)
	}
	notes, err := getNotes("notes.csv")
	if err != nil {
		panic(err)
	}
	if _, exists := notes[ID]; exists {
		note := notes[ID]
		note.IsComplete = true
		notes[ID] = note
		writeNotes(notes)
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "ID\t Task\t Created\t Done" )
		t, err := time.Parse("2006-01-02T15:04:05-07:00", note.CreatedAt)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, strconv.Itoa(ID), "\t", note.Description, "\t", t.Format("15:04:05 02-01-2006"), "\t", note.IsComplete)

	} else  {
		fmt.Println("No ID")
	}

}

func init() {
	rootCmd.AddCommand(completeCmd)
}
