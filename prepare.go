package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mkideal/cli"
)

// prepare command
type prepareT struct {
	cli.Helper
	Input  string `cli:"*i,input" usage:"image dir path --input='/path/to/image_dir'"`
	Output string `cli:"*o,output" usage:"output TSV file path --output='./output.csv'" dft:"./output.csv"`
	Type   string `cli:"t,type" usage:"comma separate file extensions --type='jpg,png,gif'" dft:"jpg,png,gif"`
}

var prepare = &cli.Command{
	Name: "prepare",
	Desc: "Create csv list file from --output dir",
	Argv: func() interface{} { return new(prepareT) },
	Fn:   execPrepare,
}

var types map[string]struct{}

func execPrepare(ctx *cli.Context) error {
	argv := ctx.Argv().(*prepareT)

	types = make(map[string]struct{})
	for _, s := range strings.Split(argv.Type, ",") {
		types["."+strings.TrimSpace(s)] = struct{}{}
	}

	f, err := NewFileHandler(argv.Output)
	if err != nil {
		return err
	}

	list := getFilesFromDir(argv.Input)
	return f.WriteAll(list)
}

func getFilesFromDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			paths = append(paths, getFilesFromDir(filepath.Join(dir, fileName))...)
			continue
		}

		if _, ok := types[filepath.Ext(fileName)]; ok {
			paths = append(paths, filepath.Join(dir, fileName))
		}
	}

	return paths
}
