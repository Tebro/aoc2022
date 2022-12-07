package day7

import (
	day6 "aoc/cmd/day_6"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TreeNode[T any] struct {
	parent   *TreeNode[T]
	children []*TreeNode[T]
	Data     T
}

func CreateTree[T any](data T) *TreeNode[T] {
	return &TreeNode[T]{
		parent:   nil,
		children: []*TreeNode[T]{},
		Data:     data,
	}
}

func (tn *TreeNode[T]) AddChild(data T) *TreeNode[T] {
	child := &TreeNode[T]{
		parent:   tn,
		children: []*TreeNode[T]{},
		Data:     data,
	}
	tn.children = append(tn.children, child)
	return child
}

func (tn *TreeNode[T]) FindChild(fn func(T) bool) *TreeNode[T] {
	for _, c := range tn.children {
		if fn(c.Data) {
			return c
		}
	}
	return nil
}

func pathToParent[T any](node *TreeNode[T], fn func(T) string) []string {
	if node.parent == nil {
		return []string{fn(node.Data)}
	}
	return append(pathToParent(node.parent, fn), fn(node.Data))
}

func (tn *TreeNode[T]) PathToParent(sep string, fn func(T) string) string {
	path := pathToParent(tn, fn)
	return strings.Join(path, sep)
}

type directory struct {
	name      string
	fileNames *day6.Set[string]
	fileSizes map[string]int
}

func createDir(name string) directory {
	return directory{
		name:      name,
		fileNames: day6.CreateSet[string](),
		fileSizes: make(map[string]int),
	}
}

func (d directory) fileSizeSum() int {
	sum := 0
	for _, v := range d.fileSizes {
		sum += v
	}
	return sum
}

func directoryNodeRecursiveSize(n *TreeNode[directory]) int {
	total := n.Data.fileSizeSum()
	for _, c := range n.children {
		total += directoryNodeRecursiveSize(c)
	}
	return total
}

func allSizesOfDirNodes(n *TreeNode[directory]) []int {
	if len(n.children) == 0 {
		return []int{n.Data.fileSizeSum()} // could use the directoryNodeRecusiveSize here as well... readability?
	}

	childValues := [][]int{{directoryNodeRecursiveSize(n)}}
	for _, c := range n.children {
		childValues = append(childValues, allSizesOfDirNodes(c))
	}
	res := []int{}
	for _, l := range childValues {
		res = append(res, l...)
	}
	return res
}

func getStandardPath(current *TreeNode[directory]) string {
	return current.PathToParent("/", func(d directory) string { return d.name })
}

func buildTree(lines []string) (*TreeNode[directory], error) {
	root := CreateTree(createDir("/"))
	current := root

	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if line == "$ ls" {
				continue
			} else {
				parts := strings.Split(line, " ")

				if parts[1] != "cd" {
					return nil, fmt.Errorf("unsupported command detected: %v", parts[1])
				}

				dirName := parts[2]
				if dirName == "/" {
					current = root
				} else if dirName == ".." {
					current = current.parent
				} else {
					found := current.FindChild(func(d directory) bool { return d.name == dirName })
					if found == nil {
						return nil, fmt.Errorf("got a nil child when navigating from: %s to %s", getStandardPath(current), dirName)
					}
					current = found
				}
			}
		} else {
			if strings.HasPrefix(line, "dir") {
				parts := strings.Split(line, " ")
				if current.FindChild(func(d directory) bool { return d.name == parts[1] }) == nil {
					current.AddChild(createDir(parts[1]))
				}
			} else {
				parts := strings.Split(line, " ")
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, fmt.Errorf("could not parse file size, path: %s, line: %s, err: %w", getStandardPath(current), line, err)
				}
				if current.Data.fileNames.Has(parts[1]) { // file exists, did it change? update size anyway
					current.Data.fileSizes[parts[1]] = size
				} else { // New file
					current.Data.fileNames.Add(parts[1])
					current.Data.fileSizes[parts[1]] = size
				}
			}
		}
	}
	return root, nil
}

func part1(lines []string) (int, error) {

	root, err := buildTree(lines)
	if err != nil {
		return 0, err
	}

	allSizes := allSizesOfDirNodes(root)

	sum := 0
	for _, s := range allSizes {
		if s <= 100000 {
			sum += s
		}
	}

	return sum, nil
}

func part2(lines []string) (int, error) {
	root, err := buildTree(lines)
	if err != nil {
		return 0, err
	}

	allSizes := allSizesOfDirNodes(root)

	rootSize := directoryNodeRecursiveSize(root)
	fsSize := 70000000

	requiredSize := 30000000
	needToDelete := requiredSize - (fsSize - rootSize)

	bigEnough := []int{}
	for _, ds := range allSizes {
		if ds >= needToDelete {
			bigEnough = append(bigEnough, ds)
		}
	}

	sort.Ints(bigEnough)

	return bigEnough[0], nil
}
