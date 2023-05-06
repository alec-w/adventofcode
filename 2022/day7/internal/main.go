package internal

import (
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

type Dir struct {
	Name   string
	Dirs   []*Dir
	Files  []File
	Parent *Dir
}

type File struct {
	Name string
	Size int
}

func (d *Dir) Size() int {
	size := 0
	for _, dir := range d.Dirs {
		size += dir.Size()
	}
	for _, file := range d.Files {
		size += file.Size
	}
	return size
}

func (d *Dir) Sizes() []int {
	sizes := []int{d.Size()}
	for _, dir := range d.Dirs {
		sizes = append(sizes, dir.Sizes()...)
	}
	return sizes
}

func DirFromFile(fileName string) (*Dir, error) {
	lines, err := helpers.FileToStrings(fileName)
	if err != nil {
		return nil, err
	}

	lines = lines[1:]
	currentDir := &Dir{
		Name: "/",
	}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "$ ls" {
			i++
			for ; i < len(lines) && lines[i][0] != '$'; i++ {
				line = lines[i]
				parts := strings.Split(line, " ")
				if parts[0] == "dir" {
					currentDir.Dirs = append(currentDir.Dirs, &Dir{Name: parts[1], Parent: currentDir})
				} else {
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						return nil, err
					}
					currentDir.Files = append(currentDir.Files, File{Name: parts[1], Size: size})
				}
			}
			i--
			continue
		}
		if line == "$ cd .." {
			currentDir = currentDir.Parent
			continue
		}
		dirName := line[5:]
		for _, dir := range currentDir.Dirs {
			if dir.Name == dirName {
				currentDir = dir
				break
			}
		}
	}

	for ; currentDir.Parent != nil; currentDir = currentDir.Parent {
	}

	return currentDir, nil
}
