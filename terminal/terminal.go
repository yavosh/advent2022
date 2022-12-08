package terminal

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

type Entry struct {
	EntryType string
	Entries   []Entry
	Path      string
	Name      string
	Size      int
}

func (e Entry) Print() {

	if e.EntryType == "F" {
		fmt.Printf("%d %q\n", e.Size, path.Join("/", e.Path, e.Name))
		return
	}

	fmt.Printf("dir %q\n", path.Join("/", e.Path, e.Name))

	for _, e := range e.Entries {
		e.Print()
	}
}

func (e Entry) String() string {
	if e.EntryType == "D" {
		return fmt.Sprintf("dir %q", e.Name)
	}

	return fmt.Sprintf("%d %q", e.Size, e.Name)
}

type Terminal struct {
	cwd     []string
	entries map[string][]Entry
	dirs    map[string]*Entry
	du      map[string]int
	cmd     string
}

func New() *Terminal {
	t := Terminal{
		cwd:     make([]string, 0),
		entries: make(map[string][]Entry, 0),
		dirs:    make(map[string]*Entry, 0),
		du:      make(map[string]int, 0),
	}

	t.dirs["/"] = &Entry{
		EntryType: "D",
		Path:      "/",
		Name:      "/",
		Entries:   make([]Entry, 0),
	}

	return &t
}

func (t *Terminal) Path() string {
	return "/" + strings.Join(t.cwd, "/")
}

func (t *Terminal) Read(l string) error {
	//fmt.Printf("[%q] %q \n", t.Path(), l)

	if IsCommand(l) {
		t.Command(l[2:])
		return nil
	}

	if t.cmd == "ls" {
		if _, ok := t.entries[t.Path()]; !ok {
			t.entries[t.Path()] = make([]Entry, 0)
		}

		// dir
		if strings.HasPrefix(l, "dir ") {
			entry := Entry{
				EntryType: "D",
				Path:      path.Join(t.Path(), l[4:]),
				Name:      l[4:],
				Entries:   make([]Entry, 0),
			}

			t.entries[t.Path()] = append(t.entries[t.Path()], entry)
			t.dirs[path.Join(t.Path(), "/", l[4:])] = &entry
			t.dirs[t.Path()].Entries = append(t.dirs[t.Path()].Entries, entry)
			// make sure we have entry
			t.du[t.Path()] += 0
			return nil
		}

		// file
		if !strings.HasPrefix(l, "dir ") {
			parts := strings.Split(l, " ")

			entry := Entry{
				EntryType: "F",
				Name:      strings.TrimSpace(parts[1]),
			}

			if v, err := strconv.Atoi(parts[0]); err == nil {
				t.du[t.Path()] += v
				entry.Size = v
			}

			t.entries[t.Path()] = append(t.entries[t.Path()], entry)
			t.dirs[t.Path()].Entries = append(t.dirs[t.Path()].Entries, entry)
			return nil
		}

	}

	panic("can't process data for command " + t.cmd)
}

func (t *Terminal) Command(cmd string) {
	//fmt.Printf("cmd %q\n", cmd)
	if ok, args := IsCdCommand(cmd); ok {
		t.cmd = "cd"

		if args[0] == "/" {
			// root
			t.cwd = []string{}
			return
		}

		if args[0] == ".." {
			// pop path
			t.cwd = t.cwd[1:]
			return
		}

		t.cwd = append(t.cwd, args[0])
		return
	}

	if ok, _ := IsLsCommand(cmd); ok {
		t.cmd = "ls"
		return
	}

	panic("unknown command " + cmd)
}

func (t *Terminal) DU(path string) int {
	sum := 0

	for k, v := range t.du {
		if strings.HasPrefix(k, path) {
			sum += v
		}
	}
	return sum
}

func (t *Terminal) SumSize(most int) int {
	sum := 0
	for k := range t.dirs {
		du := t.DU(k)
		if du <= most {
			sum += du
		}
	}

	return sum
}

func (t *Terminal) Print2() {
	t.dirs["/"].Print()
}

func (t *Terminal) Print() {
	for k, v := range t.entries {
		fmt.Printf("# %s %d\n", k, t.du[k])
		for _, entry := range v {
			fmt.Printf("> %s\n", entry)
		}
	}

	for k, v := range t.du {
		fmt.Printf("@ %s %d %d\n", k, v, t.DU(k))
	}
}

func IsCommand(l string) bool {
	if len(l) > 1 && l[0] == '$' {
		return true
	}
	return false
}

func IsCdCommand(cmd string) (bool, []string) {
	return strings.HasPrefix(cmd, "cd "), ParseArgs(3, cmd)
}

func IsLsCommand(cmd string) (bool, []string) {
	if cmd == "ls" {
		return true, []string{}
	}

	return strings.HasPrefix(cmd, "ls "), ParseArgs(3, cmd)
}

func ParseArgs(start int, cmd string) []string {
	if len(cmd) <= start {
		return []string{}
	}
	return strings.Split(cmd[start:], " ")
}
