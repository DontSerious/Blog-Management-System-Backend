package utils

import (
	"Bishe/be/kitex_gen/edit"
	"Bishe/be/pkg/constants"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

var length = len(constants.EditDirectory)

// 检查是否为文件夹
func ReadDirectory(path string) (*edit.DataNode, error) {
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
func buildFileNode(path string) (*edit.DataNode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	node := &edit.DataNode{
		Title:  info.Name(),
		IsLeaf: !info.IsDir(),
		Key:    path[length:], // 变为相对路径
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
func readChildren(dirPath string) ([]*edit.DataNode, error) {
	var dirNodes []*edit.DataNode
	var fileNodes []*edit.DataNode

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

		if childNode.IsLeaf {
			fileNodes = append(fileNodes, childNode)
		} else {
			dirNodes = append(dirNodes, childNode)
		}
	}

	// Sort the folder nodes by key
	sort.Slice(dirNodes, func(i, j int) bool {
		return dirNodes[i].Key < dirNodes[j].Key
	})

	// Concatenate the folder nodes and file nodes
	return append(dirNodes, fileNodes...), nil
}
