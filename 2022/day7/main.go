// https://adventofcode.com/2022/day/7

package main

import (
	"2022/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := input.GetStringInput("input.txt")

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}

type Dir struct {
	Name   string
	Parent *Dir
	Dirs   []Dir
	Files  []File
}
type File struct {
	Parent *Dir
	Name   string
	Size   int64
}

func parseInput(in []string) Dir {
	var root *Dir = new(Dir)
	var curr *Dir
	root.Name = "/"
	curr = root
	for i := 2; i < len(in); i++ {
		// fmt.Printf("CMD: %s, CURR: %v \n", in[i], curr.name)
		split := strings.Fields(in[i])
		if split[0] == "$" { // handle commands
			if split[1] == "cd" {
				if split[2] == ".." {
					// fmt.Printf("\tCD FROM: %+v TO %+v\n", (*curr).name, (*curr).parent.name)
					curr = curr.Parent
				} else {
					// fmt.Printf("\tCD TO: %+v\n", split[2])
					// find directory in curr dir with matching name and "cd" to it
					for j := 0; j < len(curr.Dirs); j++ {
						if curr.Dirs[j].Name == split[2] {
							// fmt.Printf("\tcurr before: %v", curr.name)
							curr = &curr.Dirs[j]
							// fmt.Printf("\tcurr after: %v\n", curr.name)
							break
						}
					}
				}
			}
			// implicit else "ls", ignore and move to next line for directory listing
		} else { // process "ls" output
			// add new directory
			if split[0] == "dir" {
				newDir := Dir{Name: split[1], Parent: curr}
				curr.Dirs = append(curr.Dirs, newDir)
				// fmt.Printf("i: %d, NewDir: %+v\n", i, newDir)
			} else { // add new file
				s, _ := strconv.Atoi(split[0])
				newFile := File{Parent: curr, Name: split[1], Size: int64(s)}
				curr.Files = append(curr.Files, newFile)
				// fmt.Printf("i: %d, NewFile: %+v\n", i, newFile)
			}
		}
	}
	return *root
}

func printFs(in Dir, depth int) {
	fmt.Printf("%s- %s (dir, size=%d)\n", strings.Repeat(" ", depth), in.Name, dirSize(in))
	if len(in.Dirs) > 0 {
		for _, d := range in.Dirs {
			printFs(d, depth+2)
		}
	}
	for _, f := range in.Files {
		fmt.Printf("%s - %s (file, size=%d)\n", strings.Repeat(" ", depth), f.Name, f.Size)
	}
}

func dirSize(in Dir) int64 {
	var s int64
	if len(in.Dirs) > 0 {
		for _, d := range in.Dirs {
			s += dirSize(d)
		}
	}
	for _, f := range in.Files {
		s += f.Size
	}
	return s
}

func listDirectories(in Dir) []Dir {
	var ret []Dir
	ret = append(ret, in)
	for _, d := range in.Dirs {
		ret = append(ret, listDirectories(d)...)
	}
	return ret
}

func part1(in []string) int64 {
	fs := parseInput(in)
	// printFs(fs, 0)
	dirs := listDirectories(fs)
	var total int64
	for _, d := range dirs {
		t := dirSize(d)
		if t < 100000 {
			total += t
		}
	}
	return total
}

func deleteDirectory(in *Dir, target string) {
	for i := 0; i < len(in.Dirs); i++ {
		if in.Dirs[i].Name == target {
			fmt.Printf("DELETING: %v\n", in.Dirs[i].Name)
			fmt.Printf("LEN: %v\n", len(in.Dirs))
			fmt.Printf("BEFORE: %v\n", in.Dirs)
			if len(in.Dirs) == 1 {
				in.Dirs = []Dir{}
			} else {
				in.Dirs[i] = in.Dirs[len(in.Dirs)-1]
				in.Dirs = in.Dirs[:len(in.Dirs)-1]
			}
			fmt.Printf("LEN: %v\n", len(in.Dirs))
			fmt.Printf("AFTER: %v\n", in.Dirs)
			return
		}
		deleteDirectory(&in.Dirs[i], target)
	}
}

func part2(in []string) int64 {
	fs := parseInput(in)
	totalSize := dirSize(fs)
	free := 70000000 - totalSize
	needed := 30000000 - free
	// fmt.Printf("needed: %v\n", needed)
	dirs := listDirectories(fs)
	targetSize := totalSize
	targetDir := fs
	for _, d := range dirs {
		t := dirSize(d)
		// find smallest dir to delete that frees enough space
		if t < targetSize && t > needed {
			targetSize = t
			targetDir = d
		}
	}
	// fmt.Printf("DELETE: %s\n", targetDir.Name)
	// delete dir
	// deleteDirectory(&fs, targetDir.Name)
	return dirSize(targetDir)
}
