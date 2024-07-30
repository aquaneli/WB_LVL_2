package main

import (
	"flag"
	"fmt"
	"regexp"
)

type cutArgs struct {
	args          flags
	formatStrings []string
}

type flags struct {
	f []int
	d string
	s bool
}

func main() {
	parseFlags()
}

func parseFlags() (*cutArgs, error) {
	fFlag := flag.String("f", "", "выбрать поля (колонки)")
	// dFlag := flag.String("d", "", "использовать другой разделитель")
	// sFlag := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// if len(flag.Args()) < 2 {
	// 	return &cutArgs{}, errors.New("file not specified\nusage: ")
	// }

	// fConvInt, err := strconv.Atoi(*fFlag)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// path := flag.Args()[1:]
	

	// parts := strings.Split(*fFlag, ",")
	re := regexp.MustCompile(`^(\d+)\s*-\s*(\d+)$`)
	matches := re.FindStringSubmatch(*fFlag)

	// for _, v := range parts {

	// }

	fmt.Println(matches)

	return &cutArgs{}, nil
}
