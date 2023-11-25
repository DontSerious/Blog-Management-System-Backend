package utils

import (
	"Bishe/be/kitex_gen/edit"
	"fmt"
	"os"
	"path/filepath"
)

// ReadDirectory reads a directory and constructs a FileNode tree
func ReadDirectory(path string) (*edit.FileNode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	return buildFileNode(path)
}

// buildFileNode recursively builds a FileNode tree for the given directory
func buildFileNode(path string) (*edit.FileNode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := &edit.FileNode{
		Name:  info.Name(),
		IsDir: info.IsDir(),
	}

	if info.IsDir() {
		children, err := readChildren(path)
		if err != nil {
			return nil, err
		}
		node.Children = children
	}

	return node, nil
}

// readChildren reads the children of a directory and constructs FileNode for each child
func readChildren(dirPath string) ([]*edit.FileNode, error) {
	var children []*edit.FileNode

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		childPath := filepath.Join(dirPath, file.Name())
		childNode, err := buildFileNode(childPath)
		if err != nil {
			return nil, err
		}

		children = append(children, childNode)
	}

	return children, nil
}
