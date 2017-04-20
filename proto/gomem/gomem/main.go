package main

// TODO: List
//     : implemetation methods
//     : get Title, push *Gomem,
//     : convert JSON
//     : interactive repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

type Gomem struct {
	Title   string
	Content string
	Tags    []string

	// TODO: imple WriteDate string
	//WriteDate string // use time.Parse(time.layout, Gomem.WriteDate)
}

type Gomems []*Gomem

func (g *Gomems) String() string {
	if g == nil {
		return ""
	}
	var str string
	for _, x := range *g {
		if x == nil {
			continue
		}
		str += "---\n"
		str += fmt.Sprintf("Title: \"%s\"\n", x.Title)
		str += fmt.Sprintf("Content: \"%s\"\n", x.Content)
		str += fmt.Sprintf("Tags:\n")
		for _, t := range x.Tags {
			str += fmt.Sprintf("  - \"%s\"\n", t)
		}
	}
	return str
}

func (g *Gomems) at(index int) (*Gomem, error) {
	if g == nil {
		return nil, fmt.Errorf("Gomems.at: *Gomems nil")
	}
	if len(*g) <= index {
		return nil, fmt.Errorf("Gomems.at: invalid index")
	}
	return (*g)[index], nil
}

func gomemNew(title string, content string, tags []string) *Gomem {
	return &Gomem{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
}

func repl(g *Gomems) {
	fmt.Print("repl:>")
	for sc := bufio.NewScanner(os.Stdin); sc.Scan(); {
		if sc.Err() != nil {
			log.Fatalf("repl:%v", sc.Err())
		}
		// TODO: implement commands?
		switch s := sc.Text(); s {
		case "exit", "quit", "q":
			fmt.Println("see you later :)")
			return
		case "get":
			fmt.Printf("%q\n", get())
		default:
			fmt.Println(s)
		}
		fmt.Print("repl:>")
	}
}

// mock
func get() *Gomem {
	return &Gomem {
		Title:   "hello",
		Content: "hello from gops a new instance",
		Tags:    []string{"gops", "new", "hello", "todo"},
	}
}

func main() {
	log.SetPrefix("gomem: ")

	repl(new(Gomems))

	gops := new(Gomems)
	*gops = append(*gops, &Gomem{
		Title:   "hello",
		Content: "hello from gops a new instance",
		Tags:    []string{"gops", "new", "hello", "todo"},
	})
	v, err := gops.at(0)
	if err != nil {
		log.Fatal(err)
	}
	*gops = append(*gops, v)
	fmt.Println(gops, len(*gops))
}
