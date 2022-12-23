package main

import (
	"flag"
	"os"
	"strings"

	"github.com/karalef/retextaigo"
)

var (
	paraphrase = flag.Bool("p", true, "paraphrase text")
	summarize  = flag.Bool("s", false, "summarize text")
	extend     = flag.Bool("e", false, "extend text")
	count      = flag.Uint("c", 1, "max number of paraphrase results")
	maxLength  = flag.Uint("l", 150, "max length of the summarized text")
	input      = flag.String("f", "", "input file")
	output     = flag.String("o", "", "output file")
)

func init() {
	flag.Parse()

	var parSpec, sumSpec, extSpec bool
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "s":
			sumSpec = true
		case "p":
			parSpec = true
		case "e":
			extSpec = true
		}
	})
	if (sumSpec && *summarize || extSpec && *extend) && !parSpec {
		*paraphrase = false
	}

	isTrue := 0
	if *paraphrase {
		isTrue++
	}
	if *summarize {
		isTrue++
	}
	if *extend {
		isTrue++
	}
	if isTrue != 1 {
		exit("specify one of -p, -s or -e")
	}
}

var client = retextaigo.New(nil)

func main() {
	available, err := client.IsAvailable()
	check(err)
	if !available {
		println("service is not available")
		return
	}

	source := strings.Join(flag.Args(), " ")
	if source == "" && *input != "" {
		d, err := os.ReadFile(*input)
		check(err)
		source = string(d)
	}
	if len(source) == 0 {
		exit("specify input")
	}

	println("waiting...\n")

	var result string
	switch {
	case *paraphrase:
		result = doParaphrase(source)
	case *summarize:
		result = doSummarization(source)
	case *extend:
		result = doExtension(source)
	}

	out := os.Stdout
	if *output != "" {
		out, err = os.OpenFile(*output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
		check(err)
		defer out.Close()
	}

	out.WriteString(result + "\n")
}

func doParaphrase(src string) string {
	task, err := client.Paraphrase(src)
	check(err)
	res, err := task.Wait()
	checkResult(err, res.Successful)
	n := int(*count)
	if n > len(res.Result) {
		n = len(res.Result)
	}
	return strings.Join(res.Result[:n], "\n\n")
}

func doSummarization(src string) string {
	task, err := client.Summarize(src, int(*maxLength))
	check(err)
	res, err := task.Wait()
	checkResult(err, res.Successful)
	return res.Result
}

func doExtension(src string) string {
	task, err := client.Extension(src)
	check(err)
	res, err := task.Wait()
	checkResult(err, res.Successful)
	return res.Result.Complete()
}

func checkResult(err error, successful bool) {
	check(err)
	if !successful {
		exit("unseccessful")
	}
}

func check(err error) {
	if err != nil {
		exit(err.Error())
	}
}

func exit(s string) {
	println(s)
	os.Exit(1)
}
