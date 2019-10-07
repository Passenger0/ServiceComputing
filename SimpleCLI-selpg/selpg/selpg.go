package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

var startPage, endPage, pageLen int
var formFeed bool
var outputPath string
var programName string
var inputFileName string

var cmd *exec.Cmd
var out io.WriteCloser //use to write to or close I/O devices
var in io.ReadCloser   //use to read or close I/O device
var buff *bufio.Reader

/*SomeInit function init the binding params*/
func SomeInit() {
	//bind params and read values
	pflag.IntVarP(&startPage, "start page", "s", 0, "The start page of  the file.")
	pflag.IntVarP(&endPage, "end page", "e", 0, "The end page of the file.")
	pflag.IntVarP(&pageLen, "page length", "l", 72, "Number of lines in one page.")
	pflag.BoolVarP(&formFeed, "form feed", "f", false, "The page will end until meet a form feed.")
	pflag.StringVarP(&outputPath, "output file", "d", "", "Path of the destinate device")
	pflag.Usage = func() { //output how to use this tools
		fmt.Printf("Usage of %s:\n", programName)
		fmt.Printf("-s\t--Start page of the file,must be chosed,default 0.\n-e\t--End page of the file,must be chosed,default 0.\n-l\t--Lines of one page,optional,default 72.\n-f\t--Page will end with form feed,default false.\n-d\t--Output to which device,default null.\n")
		//pflag.PrintDefaults()
	}
}

/*IsArgsLegal function detect if the values of the params are legal*/
func IsArgsLegal() { //judge if there is more than 3 args for the format of "program -sNumber -eNumber"
	if len(os.Args) < 3 || os.Args[1] != "-s" || os.Args[3] != "-e" {
		fmt.Printf("\nError: must be the format as ./CLI -s Number -e Number\n ")
		pflag.Usage()
		os.Exit(1)
	}
	IntMax := 1<<32 - 1
	if startPage < 1 || startPage > IntMax || endPage < 1 || endPage > IntMax || endPage < startPage {
		fmt.Printf("Error: the page number is not legal.\n ")
		pflag.Usage()
		os.Exit(2)
	}
	if pageLen < 1 || pageLen > IntMax {
		fmt.Printf("Error: the lines in one page is not legal.\n")
		pflag.Usage()
		os.Exit(3)
	}
	if pflag.NArg() == 1 { //get non-flag param, only input filename is permitted
		_, err := os.Stat(pflag.Args()[0])
		if err != nil && os.IsNotExist(err) {
			fmt.Printf("Error: the file %s is not exist.\n", pflag.Args()[0])
			pflag.Usage()
			os.Exit(4)
		}
		inputFileName = pflag.Args()[0]
	}
}

/*DeviceToInAndOut function decides which device to input or output*/
func DeviceToInAndOut() {

	if inputFileName == "" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(inputFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(6)
		}
	}
	buff = bufio.NewReader(in)

	if outputPath == "" {
		out = os.Stdout
	} else {
		var err error
		out, err = os.OpenFile(outputPath, os.O_CREATE|os.O_RDWR, 777) //let cli output to this file
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(7)
		}
	}
}

/*ReadAndWrite function complete the work of input and output*/
func ReadAndWrite() {
	var pageNum int
	if formFeed {
		pageNum = 1
		for {
			line, err := buff.ReadString('\f')
			if err != nil && err != io.EOF {
				fmt.Println(err)
				os.Exit(9)
			}
			line += "\f"
			fmt.Println(line)
			if err == io.EOF {
				break
			}
			pageNum++
			if pageNum >= startPage && pageNum <= endPage {
				_, err := out.Write([]byte(line))
				if err != nil {
					fmt.Println(err)
					os.Exit(10)
				}
			}
		}
	} else {
		lineNum := 1
		pageNum = 1
		for {
			line, _, err := buff.ReadLine()
			if err != nil && err != io.EOF {
				fmt.Println(err)
				os.Exit(9)
			}
			if err == io.EOF {
				break
			}
			lineNum++
			if lineNum > pageLen {
				lineNum = 1
				pageNum++
			}
			if pageNum >= startPage && pageNum <= endPage {
				_, err := out.Write(append([]byte(line), []byte("\n")...))
				if err != nil {
					fmt.Println(err)
					os.Exit(10)
				}
			}
		}
	}
	in.Close()
	out.Close()
	if pageNum < startPage || pageNum < endPage {
		fmt.Printf("Error: the page index is not in the range.\n")
	}
}

func main() {
	programName = os.Args[0] //get the name of the program
	SomeInit()               //init params
	pflag.Parse()
	IsArgsLegal()
	DeviceToInAndOut()
	ReadAndWrite()
}
