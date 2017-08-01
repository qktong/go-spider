package file

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type DirFile struct {
	Path     string
	FileName string
}

func ListDir(dirPth string, suffix string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func GetFileName(fullFilename string) (filenameOnly string) {
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

func ListDir2(dirPth string, suffix string) (files []DirFile, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, DirFile{dirPth + PthSep + fi.Name(), fi.Name()})
		}
	}
	return files, nil
}
