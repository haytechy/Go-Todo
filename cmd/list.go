package cmd

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List notes",
	Run: listNote,
}

func listNote(cmd *cobra.Command, args []string) {
	all, err := cmd.Flags().GetBool("all")
	notes, err := getNotes("notes.csv")
	if err != nil {
		panic(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	if all {
		fmt.Fprintln(w, "ID\t Task\t Created\t Done" )
	} else {
		fmt.Fprintln(w, "ID\t Task\t Created" )
	}
	for _, i := range slices.Sorted(maps.Keys(notes)) {
		note := notes[i]
		t, err := time.Parse("2006-01-02T15:04:05-07:00", note.CreatedAt)
		if err != nil {
			panic(err)
		}
		if all {
			fmt.Fprintln(w, strconv.Itoa(i), "\t", note.Description, "\t", t.Format("15:04:05 02-01-2006"), "\t", note.IsComplete)
		} else {
			fmt.Fprintln(w, strconv.Itoa(i), "\t", note.Description, "\t", t.Format("15:04:05 02-01-2006"))
		}
	}
}

func init() {
	listCmd.Flags().BoolP("all", "a",  false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}
