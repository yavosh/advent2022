package terminal

import (
	"fmt"
)

type Entry struct {
	EntryType string
	Path      string
	Name      string
	Size      int
}

func (e Entry) String() string {
	if e.EntryType == "D" {
		return fmt.Sprintf("[%s] %q", e.EntryType, e.Path)
	}

	return fmt.Sprintf("[%s] %q %d", e.EntryType, e.Path, e.Size)
}

func (e Entry) IsDir() bool {
	return e.EntryType == "D"
}

func (e Entry) IsFile() bool {
	return e.EntryType == "F"
}
