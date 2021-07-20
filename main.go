package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"strconv"

	// _ "os"
	// package Not Used Underscore

	"github.com/mactsouk/go/simpleGitHub"
)

// printing p.50
func printing() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Print(v1, v2, v3, v4)
	fmt.Println()
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)

	// S 시리즈 함수(fmt.Srpint(), fmt.Sprintln(), fmt.Sprintf())
	// 지정된 포맷에 따라 스트링(문자열)을 생성할 때 사용

	// F 시리즈 함수(fmt.Fprint(), fmt.Fprintln(), fmt.Fprintf())
	// is.Writer로 파일을 쓸때 사용
}

// stdOUT p.52
func stdOUT() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}

// stdIN p.56
//lint:ignore U1000 input
func stdIN() {
	// var f *os.File
	// f = os.Stdin
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}

// cla p.58
func cla() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		return
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

// logFiles p.67
func logFiles() {
	// 로그 수준(logging level)
	// debug, info, notice, warning, err, crit, alert, emerg

	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging int Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Will you see this?")
}

// logFatal p.71
//lint:ignore U1000 Fatal
func logFatal() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Fatal(sysLog)
	fmt.Println("Will you see this")
}

// logPanic p.72
//lint:ignore U1000 Panic
func logPanic() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Panic(sysLog)
	fmt.Println("Will you see this")
}

func main() {
	fmt.Println("This is a sample Go program!")
	fmt.Println(simpleGitHub.AddTwo(5, 6))

	printing()
	stdOUT()
	// stdIN()
	cla()
	logFiles()
	// logFatal()
	// logPanic()
}
