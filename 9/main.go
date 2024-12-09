package main

import (
	"bufio"
	"fmt"
	"os"
)

type blockKind int

const (
	blockKindFreeSpace blockKind = iota
	blockKindUsed
)

type input []byte
type block struct {
	kind blockKind
	id   int
}
type diskLayout []block

var freeSpace = block{
	kind: blockKindFreeSpace,
	id:   -1,
}

func decodeInput(input input) (layout diskLayout) {
	layout = make(diskLayout, 0)
	blockId := 0

	for k, v := range input {
		var currentBlock block
		if k%2 == 0 {
			currentBlock = block{
				kind: blockKindUsed,
				id:   blockId,
			}
			blockId++
		} else {
			currentBlock = freeSpace
		}
		for i := 0; i < int(v); i++ {
			layout = append(layout, currentBlock)
		}
	}
	return layout
}

func checksum(layout diskLayout) (result int64) {
	for position, b := range layout {
		if b.kind == blockKindUsed {
			result += int64(position * b.id)
		}
	}
	return result
}

func solve1(input input) int64 {
	layout := decodeInput(input)
	movable := len(layout) - 1
	nextFree := 0
	for movable > nextFree {
		for layout[movable].kind == blockKindFreeSpace {
			movable--
		}
		for layout[nextFree].kind == blockKindUsed {
			nextFree++
		}
		if movable <= nextFree {
			break
		}
		layout[nextFree] = layout[movable]
		layout[movable] = freeSpace
	}
	return checksum(layout)
}

type freeBlock struct {
	pos int
	len int
}

type fileBlock struct {
	pos int
	len int
	id  int
}

type fileSystem struct {
	freeSpace []freeBlock
	files     []fileBlock
}

func parseFilesystem(input input) fileSystem {
	fileId := 0
	fs := fileSystem{
		freeSpace: make([]freeBlock, 0),
		files:     make([]fileBlock, 0),
	}
	pos := 0
	for k, v := range input {
		l := int(v)
		if k%2 == 0 {
			// file
			fs.files = append(fs.files, fileBlock{
				pos: pos,
				len: l,
				id:  fileId,
			})
			fileId++
		} else {
			// free space
			if l == 0 {
				continue
			}
			fs.freeSpace = append(fs.freeSpace, freeBlock{
				pos: pos,
				len: l,
			})
		}
		pos += l
	}
	return fs
}

func (fs *fileSystem) deFragment() {
	for fileId := len(fs.files) - 1; fileId > 1; fileId-- {
		for freeId := 0; freeId < len(fs.freeSpace); freeId++ {
			if fs.freeSpace[freeId].pos > fs.files[fileId].pos {
				// do not move files forward
				break
			}
			if fs.files[fileId].len <= fs.freeSpace[freeId].len {
				fs.files[fileId].pos = fs.freeSpace[freeId].pos
				fs.freeSpace[freeId].len -= fs.files[fileId].len
				fs.freeSpace[freeId].pos += fs.files[fileId].len
				break
			}
		}
	}
}

func (fs *fileSystem) checksum() int64 {
	var res int64 = 0
	for _, file := range fs.files {
		for i := 0; i < file.len; i++ {
			res += int64(file.id * (file.pos + i))
		}
	}
	return res
}

func solve2(input input) int64 {
	fs := parseFilesystem(input)
	fs.deFragment()
	return fs.checksum()
}

func readInput(theirs bool) (result input) {
	fileName := "input.txt"
	if theirs {
		fileName = "theirs.txt"
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	result = make(input, len(line))
	for k, v := range []byte(line) {
		result[k] = v - '0'
	}
	return result
}

func main() {
	data := readInput(false)
	fmt.Println(solve1(data))
	fmt.Println(solve2(data))
}
