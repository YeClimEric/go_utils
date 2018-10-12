package go_utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"
)

type ReadLineCallback func(int, string)

/**
读取长行文件
*/
func ReadLongLineFile(filename string, call ReadLineCallback) (bool) {

	ret := true
	count := 0
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				ret = false
				fmt.Println("Read file error!", err)
				break
			}
		}

		count++
		line = strings.TrimSpace(line)
		call(count, line)
	}

	return ret
}

/**
*读取行短的文件
 */
func ReadFileLineByLine(filename string, call ReadLineCallback) (string) {

	count := 0
	f, err := os.OpenFile(filename, os.O_RDONLY, 0)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		count++
		call(count, sc.Text())
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("an error occured:%s\n", err)
	}

	return ""
}
