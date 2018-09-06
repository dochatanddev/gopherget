package main

import "flag"
import "os"
import "fmt"
import "io"
import "github.com/prologic/go-gopher"

var (
	uri string
)

func get(uri string) (resp *gopher.Response) {
	resp, _ = gopher.Get(uri)
	return
}

func handle(r *gopher.Response) (answer string) {
	switch r.Type {
	case gopher.DIRECTORY:
		for _, item := range r.Dir.Items {
			switch item.Type {
			case gopher.DIRECTORY:
				{
					answer += fmt.Sprintf("dir: (gopher://%s:%d/1%s) %s\n", item.Host, item.Port, item.Selector, item.Description)
				}
			case gopher.INFO:
				{
					answer += fmt.Sprintf("info: %s\n", item.Description)
				}
			case gopher.HTML:
				{
					answer += fmt.Sprintf("html: %s %s\n", item.Description, item.Selector)
				}
			default:
				{
					answer += fmt.Sprintf("[%s] (gopher://%s:%d/0%s) %s\n", item.Type.String(), item.Host, item.Port, item.Selector, item.Description)
				}
			}

		}
	case gopher.FILE:
		{
			io.Copy(os.Stdout, r.Body)
			answer = ""
		}
	default:
		answer = fmt.Sprintf("Unknown response type %s; idk what you get", r.Type)
	}
	return
}

func init() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println(`Usage: gopherget "gopher://some.uri"`)
		os.Exit(1)
	}
	uri = args[0]
}

func main() {
	response := get(uri)
	parsed := handle(response)
	fmt.Println(parsed)
}
