google-vision-ocr
====

Search image file and extract OCR text using Google Cloud Vision API


# Installation

Install google-vision-ocr and required packages using `go get` command:

```bash
$ go get github.com/evalphobia/google-vision-ocr/...
```


# Usage

## root command

```bash
$ google-vision-ocr
Commands:

  help      show help
  prepare   Create csv list file from --output dir
  ocr       Call Goole Cloud Vision API and get OCR result
```

## prepare command

```bash
$ google-vision-ocr help prepare
Create csv list file from --output dir

Options:

  -h, --help                    display help information
  -i, --input                  *image dir path --input='/path/to/image_dir'
  -o, --output[=./output.csv]  *output TSV file path --output='./output.csv'
  -t, --type[=jpg,png,gif]      comma separate file extensions --type='jpg,png,gif'
```

```bash
$ google-vision-ocr prepare -i ./Pictures/ -o ./output.csv

$ cat ./output.csv
./Pictures/1.png
./Pictures/2.png
```

## ocr command

```bash
$ google-vision-ocr help ocr
Call Goole Cloud Vision API and get OCR result

Options:

  -h, --help           display help information
  -f, --file          *image list file --file='/path/to/dir/output.tsv'
  -p, --parallel[=2]   parallel number --parallel=2
```

```bash
$ cat ./output.csv
./Pictures/1.png
./Pictures/2.png

$ google-vision-ocr ocr -f ./output.csv
exec: ./Pictures/1.png	abc
exec: ./Pictures/2.png	父母
```
