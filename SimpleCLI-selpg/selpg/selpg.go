package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag" //引入pflag并重命名为flag
)

//Arguments 存储参数
type Arguments struct {
	startPage   int    //读取开始的页数
	endPage     int    //读取结束的页数
	pathToRead  string //读取数据的路径
	pathToWrite string //写入数据的路径
	pageLines   int    //读取的数据以多少行作为一页
	sepWithFeed bool   //是否已换页符作为分页依据，与pageLines互斥

}

//var fread,fwrite *os.File //文件读写流
//var args Arguments
var source, target string

func checkError(err error, wrong string) { //处理错误
	if err != nil { //若有错，则输出错误信息
		fmt.Fprintf(os.Stderr, "\nError:%s:", wrong)
		os.Exit(0)
	}
}
func writePath(pathToWrite string) io.WriteCloser { //找到写入的路径
	cmd := exec.Command("lp", "-d"+pathToWrite) //利用lp打开管道
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fwrite, err := cmd.StdinPipe()
	checkError(err, "StdinPipe() open failed!")
	return fwrite
}

//读取数据并将数据写到指定路径

func write(fwrite interface{}, toWrite string) {
	var writeError error
	if isStd, ok := fwrite.(*os.File); ok { //判断是否为标准输出
		_, writeError = fmt.Fprintf(isStd, "%s", toWrite) //是则写入，返回写入信息
	} else if isPipe, ok := fwrite.(io.WriteCloser); ok { //判断是否为管道
		_, writeError = isPipe.Write([]byte(toWrite))
	} else { //两者都不是则表明发生错误
		fmt.Fprintf(os.Stderr, "\nError:Fstream to write is invalid.Please check your arguments[-d]! ")
		os.Exit(8)
	}
	checkError(writeError, "Error happend when writing!.") //判断写入过程中是否发生错误
}
func readAndWrite(fwrite interface{}, fread *os.File, start int, end int, lines int, feed bool) {
	pageCount := 0                   //所读的页所在的页数
	lineCount := 0                   //当前页所读的行数
	buffer := bufio.NewReader(fread) //存储已读取数据的缓冲区

	var aline, apage string //读取的一行/页数据
	var err error           //检测是否出错
	fmt.Println("We have read:")
	if feed { //如果是按照“\f”分页读取
		for true {
			apage, err = buffer.ReadString('\f')            //每次读一页
			pageCount++                                     //更新页数
			if (pageCount >= start) && (pageCount <= end) { //判断是否可写
				write(fwrite, apage)
			}
			if err == io.EOF || pageCount > end { //先写入文件再判断是否为文件末尾，避免最后一页或者一行因EOF而无法写入。
				break
			}
		}
	} else {
		pageCount = 1 //第一次读确定为第一页
		for true {
			aline, err = buffer.ReadString('\n')
			lineCount++            //每次读行数加一
			if lineCount > lines { //行数超过一页的行数，则行数归一，页数加一
				pageCount++
				lineCount = 1
			}
			if (pageCount >= start) && (pageCount <= end) { //判断是否可写
				write(fwrite, aline)
			}
			if err == io.EOF || pageCount > end {
				break
			}
		}
	}

	if pageCount < start { //假如文件页数不够开始写
		checkError(err, "StartPage greater than the number of total pages!Nothing to write! ")
	} else if pageCount < end { //假如文件页数达不到所要写的页数
		checkError(err, "EndtPage greater than the number of total pages!Incomplete read and write!")
	}

	//fmt.Printf("\nWriting complete!\nSource: %s\t\t\tTarget:%s",source,target)
}

func run(args *Arguments) { //开始运行指令
	fread := os.Stdin
	source = "os.stdin"
	if len(args.pathToRead) > 0 { //args.pathToRead  == ""  不行
		var err error
		fread, err = os.Open(args.pathToRead) //打开该路径
		if err != nil {                       //失败则报错退出
			fmt.Fprintf(os.Stderr, "Error:File open failed! Please check the filename!")
			os.Exit(0)
		}
		source = args.pathToRead
	}

	/*fwrite := os.Stdout
	if args.pathToWrite != "" {//   args.pathToWrite  == nil
		fwrite = writePath(args.pathToWrite)
	}
	readAndWrite(fwrite, fread, args.startPage, args.endPage, args.pageLines, args.sepWithFeed)*/

	if len(args.pathToWrite) == 0 { //读取数据，如果没有指定输出路径，则输出到stdout
		target = "os.Stdout"
		readAndWrite(os.Stdout, fread, args.startPage, args.endPage, args.pageLines, args.sepWithFeed)
	} else { //否则输出到指定路径
		target = args.pathToWrite
		readAndWrite(writePath(args.pathToWrite), fread, args.startPage, args.endPage, args.pageLines, args.sepWithFeed)
	}
}
func checkArgs(args *Arguments) { //检查命令参数是否正确，若有错则报错并退出
	fmt.Println("Start to check the command!")
	if (args.startPage == 0) || (args.endPage == 0) { //参数不可少或者为0
		fmt.Fprintf(os.Stderr, "\nError:StartPage and endPage are necessary and should be greater than 0!")
		os.Exit(0)
	} else if args.startPage > args.endPage { //开始页不可大于结束页
		fmt.Fprintf(os.Stderr, "\nError:The startPage can't be bigger than the endPage!")
		os.Exit(0)
	} else if (args.sepWithFeed == true) && (args.pageLines != 72) { //以分页符分页和页行数不可同时出现（根据参考资料链接4）
		fmt.Fprintf(os.Stderr, "\nError:The command -l and -f are exclusive!")
		os.Exit(0)
	} else if args.pageLines <= 0 { //页行数不能为0
		fmt.Fprintf(os.Stderr, "\nError:The pageLines should be greater than 0! ")
		os.Exit(0)
	}
	fmt.Println("Command valid! Start to run!")
}

func readArgs(args *Arguments) { //读取命令行参数
	//绑定需要读取的，带前缀的参数
	flag.IntVarP(&(args.startPage), "startPage", "s", 0, "Start page to read")
	flag.IntVarP(&(args.endPage), "endPage", "e", 0, "End page to read")
	flag.StringVarP(&(args.pathToWrite), "pathToWrite", "d", "", "Where to write the read data")
	flag.IntVarP(&(args.pageLines), "pageLines", "l", 72, "Number of lines of one page")
	flag.BoolVarP(&(args.sepWithFeed), "sepWithFeed", "f", false, "Seperates pages with \"\\f\"")
	//函数用法
	flag.Usage = func() {
		fmt.Println("Usage of selpg:")
		fmt.Println("[NECESSARY]")
		fmt.Println("-s\t--Start page to read, default 0.")
		fmt.Println("-e\t--End page to read, default 0.")
		fmt.Println("[OPTIONAL]")
		fmt.Println("-d\t--Where to write the read data, default stdout.")
		fmt.Println("-l\t--Number of lines of one page, default 72. Exclusive with -f")
		fmt.Println("-f\t--Seperates pages with \"\\f\", default false.")
		fmt.Println("File to read.  -Just input the name of the file in your command.Default stdin")
	}
	//解析命令行参数
	flag.Parse()
	//假如还有剩下的参数,传给pathToRead
	if len(flag.Args()) > 0 {
		args.pathToRead = string(flag.Args()[0])
	}
}

func main() {
	var args Arguments
	args.pathToRead = "" //路径预先初始化为空
	readArgs(&args)      //读取参数
	checkArgs(&args)     //检查参数是否合理
	run(&args)           //执行命令
}
