package cli

import (
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path"
	"text/template"
)

type Service struct{}

var ServiceInstance = &Service{}

func (s *Service) GetList() []string {
	return []string{"1", "2", "3"}
}

type TemplateContext struct {
	Name string
}

func (s *Service) CreateModule(name string) error {
	dir := path.Join("modules", name)
	fmt.Println("dir", dir)
	if fileutil.IsExist(dir) {
		return errors.New("文件夹已存在")
	}
	creatDirErr := os.Mkdir(dir, 0755)
	if creatDirErr != nil {
		return creatDirErr
	}

	modNames := []string{"controller", "service", "model"}
	for _, modName := range modNames {
		filePath := path.Join(dir, name+"_"+modName+".go")
		tempFile, _ := template.ParseFiles("modules/cli/template/" + modName + ".tpl")
		writeFileErr := os.WriteFile(filePath, []byte(""), 0755)
		if writeFileErr != nil {
			return writeFileErr
		}
		file, _ := os.OpenFile(filePath, os.O_RDWR, 0755)
		err := tempFile.Execute(file, TemplateContext{Name: name})
		if err != nil {
			return err
		}
	}
	return nil
}
