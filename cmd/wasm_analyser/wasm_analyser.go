package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

type pkgWithSize struct {
	pkg  string
	size uint64
}

type pkgsWithSize []pkgWithSize

func (p pkgsWithSize) Len() int {
	return len(p)
}

func (p pkgsWithSize) Less(i int, j int) bool {
	return p[i].size > p[j].size
}

func (p pkgsWithSize) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	funcRegex := regexp.MustCompile(`[a-f0-9]+ func\[[0-9]+\] <([^>]+)>`)
	instrRegex := regexp.MustCompile(`^ ([a-f0-9]+):`)

	var inFile *os.File
	if len(os.Args) < 2 {
		inFile = os.Stdin
	} else {
		inFile, _ = os.Open(os.Args[1])
		defer inFile.Close()
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	sizeByPkg := make(map[string]uint64)

	var start, current uint64
	var currentFunc string
	for scanner.Scan() {
		line := scanner.Text()
		m := funcRegex.FindAllStringSubmatch(line, -1)
		if len(m) > 0 {
			if start != 0 {
				sizeByPkg[getPackageFromFunc(currentFunc)] += current - start
				start = 0
			}
			funcName := m[0][1]
			currentFunc = funcName
			continue
		}

		m = instrRegex.FindAllStringSubmatch(line, -1)
		if len(m) > 0 {
			pos, err := strconv.ParseInt(m[0][1], 16, 64)
			if err != nil {
				log.Fatalf("unable to parse: %s", m[0][1])
			}
			if start == 0 {
				start = uint64(pos)
			}
			current = uint64(pos)
		}
	}
	if start != 0 {
		sizeByPkg[getPackageFromFunc(currentFunc)] += current - start
	}

	var res pkgsWithSize
	var total uint64
	for pkg, size := range sizeByPkg {
		res = append(res, pkgWithSize{
			pkg:  pkg,
			size: size,
		})
		total += size
	}
	sort.Sort(res)
	fmt.Printf("Total size: %s\n", humanize.Bytes(total))
	for _, p := range res {
		fmt.Printf("Package %s is %s\n", p.pkg, humanize.Bytes(p.size))
	}
}

func getPackageFromFunc(f string) string {
	return strings.Split(f, ".")[0]
}