package aoc22

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/thijzert/advent-of-code/ch"
)

func Dec07a(ctx ch.AOContext) error {
	root, err := readDirectoryStructure(ctx, "inputs/2022/dec07.txt")
	if err != nil {
		return err
	}

	ctx.Printf("\n%s", root)

	under100k := 0
	var addDirSize func(dir *directory) int
	addDirSize = func(dir *directory) int {
		rv := 0
		for _, d := range dir.childDirectories {
			rv += addDirSize(d)
		}
		for _, s := range dir.files {
			rv += s
		}
		if rv <= 100000 {
			under100k += rv
		}
		return rv
	}
	ctx.Printf("Total size: %d", addDirSize(root))

	ctx.FinalAnswer.Print(under100k)
	return nil
}

type directory struct {
	parent           *directory
	childDirectories map[string]*directory
	files            map[string]int
}

func newDir(parent *directory) *directory {
	return &directory{
		parent:           parent,
		childDirectories: make(map[string]*directory),
		files:            make(map[string]int),
	}
}

func (*directory) indentString(s string) string {
	lines := strings.Split(strings.TrimRight(s, "\n "), "\n")
	if len(lines) == 0 {
		return ""
	}
	return "  " + strings.Join(lines, "\n  ") + "\n"
}
func (d *directory) String() string {
	rv := ""
	for name, dir := range d.childDirectories {
		rv += "- " + name + " (dir)\n" + d.indentString(dir.String())
	}
	for name, size := range d.files {
		rv += fmt.Sprintf("- %s (file, %d)\n", name, size)
	}
	if d.parent == nil {
		rv = "- / (dir)\n" + d.indentString(rv)
	}
	return rv
}

func (d *directory) Size() int {
	return d.sizeCached(make(map[*directory]int))
}
func (d *directory) sizeCached(m map[*directory]int) int {
	rv := 0
	for _, s := range d.files {
		rv += s
	}
	for _, c := range d.childDirectories {
		if s, ok := m[c]; ok {
			rv += s
		} else {
			rv += c.sizeCached(m)
		}
	}
	m[d] = rv
	return rv
}

func readDirectoryStructure(ctx ch.AOContext, resourceName string) (*directory, error) {
	cmds, err := ctx.DataLines(resourceName)
	if err != nil {
		return nil, err
	}

	root := newDir(nil)
	cwd := root
	for _, cmd := range cmds {
		if cmd[0:2] == "$ " {
			cmd = cmd[2:]
			if cmd == "cd /" {
				cwd = root
			} else if cmd == "cd .." {
				cwd = cwd.parent
			} else if len(cmd) > 3 && cmd[:3] == "cd " {
				dirName := cmd[3:]
				if dir, ok := cwd.childDirectories[dirName]; ok {
					cwd = dir
				} else {
					dir := newDir(cwd)
					cwd.childDirectories[dirName] = dir
					cwd = dir
				}
			} else if cmd == "ls" {
			} else {
				return nil, errors.Wrapf(errNotImplemented, "command '%s' unknown", cmd)
			}
		} else {
			var name, isDir string
			var size int

			if _, err := fmt.Sscanf(cmd, "%d %s", &size, &name); err == nil {
				cwd.files[name] = size
			} else if _, err := fmt.Sscanf(cmd, "%s %s", &isDir, &name); err == nil && isDir == "dir" {
				if _, ok := cwd.childDirectories[name]; !ok {
					cwd.childDirectories[name] = newDir(cwd)
				}
			} else {
				return nil, errors.Wrapf(errNotImplemented, "line '%s' of unknown format", cmd)
			}
		}
	}

	ctx.Print(len(cmds))
	return root, nil
}

func Dec07b(ctx ch.AOContext) error {
	root, err := readDirectoryStructure(ctx, "inputs/2022/dec07.txt")
	if err != nil {
		return err
	}

	totalDiskSize := 70000000
	spaceNeeded := 30000000

	m := make(map[*directory]int)
	freeSpace := totalDiskSize - root.sizeCached(m)
	toFree := spaceNeeded - freeSpace
	ctx.Printf("free space: %d; to free up: %d", freeSpace, toFree)

	best := totalDiskSize
	for _, freed := range m {
		if freed >= toFree && freed < best {
			best = freed
		}
	}

	ctx.FinalAnswer.Print(best)
	return nil
}
