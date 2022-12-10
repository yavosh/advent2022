package terminal

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

type Terminal struct {
	cwd     []string
	handles map[string]Entry // path -> entry
	dirs    map[string]*Entry
	cmd     string
}

func New() *Terminal {
	t := Terminal{
		cwd:     make([]string, 0),
		handles: make(map[string]Entry, 0),
		dirs:    make(map[string]*Entry, 0),
	}

	t.dirs["/"] = &Entry{
		EntryType: "D",
		Path:      "/",
		Name:      "/",
	}

	t.handles["/"] = Entry{
		EntryType: "D",
		Path:      "/",
		Name:      "/",
	}

	return &t
}

func (t *Terminal) Path() string {
	return "/" + strings.Join(t.cwd, "/")
}

func (t *Terminal) DU(p string) (int, error) {
	entry, ok := t.handles[p]
	if !ok {
		return 0, fmt.Errorf("du invalid path %s", p)
	}

	if !entry.IsDir() {
		return 0, fmt.Errorf("du invalid path, not dir %s %s", entry.EntryType, p)
	}

	sum := 0
	for _, child := range t.handles {
		if child.IsFile() {
			if strings.HasPrefix(child.Path, entry.Path) {
				sum += child.Size
			}
		}
	}

	return sum, nil
}

func (t *Terminal) Read(l string) error {
	if IsCommand(l) {
		t.Command(l[2:])
		return nil
	}

	if t.cmd == "ls" {

		// dir
		if strings.HasPrefix(l, "dir ") {
			entry := Entry{
				EntryType: "D",
				Path:      path.Join(t.Path(), l[4:]),
				Name:      l[4:],
			}

			t.handles[entry.Path] = entry
			t.dirs[path.Join(t.Path(), "/", l[4:])] = &entry
			return nil
		}

		// file
		if !strings.HasPrefix(l, "dir ") {
			parts := strings.Split(l, " ")

			entry := Entry{
				EntryType: "F",
				Name:      strings.TrimSpace(parts[1]),
				Path:      path.Join(t.Path(), parts[1]),
			}

			if exist, ok := t.handles[entry.Path]; ok {
				panic(fmt.Errorf("duplicate file %#v %#v", exist, entry))
			}

			if v, err := strconv.Atoi(parts[0]); err == nil {
				entry.Size = v
			} else {
				return fmt.Errorf("could not parse size for file %s", parts[0])
			}

			t.handles[entry.Path] = entry
			return nil
		}

	}

	return fmt.Errorf("can't process data for command %q", t.cmd)
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

func (t *Terminal) SumSizeUpTo(most int) (int, error) {
	sum := 0
	for _, entry := range t.handles {
		if !entry.IsDir() {
			continue
		}

		size, err := t.DU(entry.Path)
		//fmt.Printf("SumSizeUpTo %s %s %d\n", entry.Path, entry.EntryType, size)
		if err != nil {
			return 0, err
		}

		if size <= most {
			sum += size
		}
	}

	return sum, nil
}

func (t *Terminal) Print() {

	sum := 0
	files := 0
	dirs := 0

	for _, v := range t.handles {
		fmt.Printf("%s \n", v)

		if v.IsFile() {
			files++
		}

		if v.IsDir() {
			dirs++
		}

		sum += v.Size
	}

	fmt.Printf("%d entries, %d size, %d files, %d dirs\n", len(t.handles), sum, files, dirs)
	tot, err := t.DU("/")
	fmt.Printf("du '/' %d %v\n", tot, err)
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
