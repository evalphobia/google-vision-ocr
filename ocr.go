package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/evalphobia/google-api-go-wrapper/config"
	"github.com/evalphobia/google-api-go-wrapper/vision"
	"github.com/mkideal/cli"
)

// ocr command
type ocrT struct {
	cli.Helper
	File     string `cli:"*f,file" usage:"image list file --file='/path/to/dir/output.tsv'"`
	Parallel int    `cli:"p,parallel" usage:"parallel number --parallel=2" dft:"2"`
}

var ocr = &cli.Command{
	Name: "ocr",
	Desc: "Call Goole Cloud Vision API and get OCR result",
	Argv: func() interface{} { return new(ocrT) },
	Fn:   execOCR,
}

func execOCR(ctx *cli.Context) error {
	argv := ctx.Argv().(*ocrT)
	maxReq := make(chan struct{}, argv.Parallel)

	cli, err := vision.New(config.Config{})
	if err != nil {
		return err
	}

	f, err := NewFileHandler(argv.File)
	if err != nil {
		return err
	}

	mu := sync.Mutex{}
	var wg sync.WaitGroup

	lines := f.ReadAll()
	for i, line := range lines {
		cols := strings.Split(line, "\t")
		if len(cols) > 1 {
			continue
		}

		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()
			maxReq <- struct{}{}

			img, err := ioutil.ReadFile(line)
			if err != nil {
				panic(err)
			}

			fmt.Println("exec: " + line)
			resp, err := cli.Text(img)
			if err != nil {
				panic(err)
			}

			body := strings.Replace(strings.Join(resp.TextResult(), " "), "\n", " ", -1)
			lines[i] = line + "\t" + body

			mu.Lock()
			defer mu.Unlock()
			f.WriteAll(lines)

			<-maxReq
		}(i, line)
	}

	wg.Wait()
	return nil
}
