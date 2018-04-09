package qwik

import (
	"os"
)

// GetAllFiles retrieves the paths of all files in a directory
// recursively; it does not include the directories themselves
// in the path list
func GetAllFiles(path string) []string {
	var filePaths []string

	filePaths = getAllFilesInDir(path)

	return filePaths
}

func getAllFilesInDir(dir string) []string {
	file, err := os.Open(dir)

	if err != nil {
		return []string{}
	}

	names, err := file.Readdirnames(0)

	if err != nil {
		return []string{}
	}

	var fileList []string

	for _, name := range names {
		filename := dir + string(os.PathSeparator) + name
		fileList = append(fileList, dir+string(os.PathSeparator)+name)
		if stat, err := os.Lstat(filename); err == nil && stat.IsDir() {
			fileList = append(fileList, getAllFilesInDir(dir+string(os.PathSeparator)+name)...)
		}
	}

	return fileList
}
