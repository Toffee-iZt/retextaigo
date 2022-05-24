package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Toffee-iZt/retextaigo"
)

var (
	paraphrase = flag.Bool("p", true, "paraphrase text")
	summarize  = flag.Bool("s", false, "summarize text")
	count      = flag.Uint("c", 1, "max paraphrase results count")
	maxLength  = flag.Int("l", 150, "max summarization length")
	input      = flag.String("f", "", "input file")
	output     = flag.String("o", "", "output file")
)

var client = retextaigo.NewClient(&http.Client{})

func main() {
	flag.Parse()

	var parSpec, sumSpec bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "s" {
			sumSpec = true
		} else if f.Name == "p" {
			parSpec = true
		}
	})
	if sumSpec && *summarize && !parSpec {
		*paraphrase = false
	}
	if *paraphrase == *summarize {
		flag.Usage()
		os.Exit(1)
	}

	available, err := client.IsAvailable()
	check(err)
	if !available {
		println("service is unavailable")
		return
	}

	var source = strings.Join(flag.Args(), " ")
	if *input != "" {
		d, err := os.ReadFile(*input)
		check(err)
		source = string(d)
	}
	if len(source) == 0 {
		println("specify input")
		os.Exit(1)
	}

	println("waiting\n")

	var result string
	if *paraphrase {
		result = doParaphrase(source)
	} else {
		result = doSummarization(source)
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
	status, q, err := client.QueueParaphrase(src, "")
	check(err)
	checkStatus(status)
	for {
		status, c, err := client.QueueCheckParaphrase(q.TaskID)
		if !checkQueue(status, err, c.Ready, c.Successful) {
			time.Sleep(time.Second)
			continue
		}
		n := int(*count)
		if n > len(c.Result) {
			n = len(c.Result)
		}
		return strings.Join(c.Result[:n], "\n\n")
	}
}

func doSummarization(src string) string {
	status, q, err := client.QueueSummarization(src, *maxLength, "")
	check(err)
	checkStatus(status)
	for {
		status, c, err := client.QueueCheckSummarization(q.TaskID)
		if !checkQueue(status, err, c.Ready, c.Successful) {
			time.Sleep(time.Second)
			continue
		}
		return c.Result
	}
}

func checkQueue(status string, err error, ready, successful bool) bool {
	check(err)
	checkStatus(status)
	if !ready {
		return false
	}
	if !successful {
		println("unseccessful")
		os.Exit(1)
	}
	return true
}

func check(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func checkStatus(status string) {
	if status != retextaigo.StatusOK {
		println("invalid status:", status)
		os.Exit(1)
	}
}
