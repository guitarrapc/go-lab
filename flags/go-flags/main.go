package main

import (
	"bytes"
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

type CreateOption struct {
	Title   string `short:"i" long:"title" value-name:"<title>" description:"The title of an issue"`
	Message string `short:"m" long:"message" value-name:"<message>" description:"The message of an issue"`
}

type ListOption struct {
	Num     int    `short:"n" long:"num" value-name:"<num>" default:"20" default-mask:"20" description:"Limit the number of issue to output."`
	State   string `long:"state" value-name:"<state>" default:"all" default-mask:"all" description:"Print only issue of the state just those that are \"opened\", \"closed\" or \"all\""`
	Scope   string `long:"scope" value-name:"<scope>" default:"all" default-mask:"all" description:"Print only given scope. \"created-by-me\", \"assigned-to-me\" or \"all\"."`
	OrderBy string `long:"orderby" value-name:"<orderby>" default:"updated_at" default-mask:"updated_at" description:"Print issue ordered by \"created_at\" or \"updated_at\" fields."`
	Sort    string `long:"sort"  value-name:"<sort>" default:"desc" default-mask:"desc" description:"Print issue ordered in \"asc\" or \"desc\" order."`
	Search  string `short:"s" long:"search"  value-name:"<search word>" description:"Search issues against their title and description."`
}

type Option struct {
	CreateOption *CreateOption `group:"Create Options"`
	ListOption   *ListOption   `group:"List Options"`
}

func newOptionParser(opt *Option) *flags.Parser {
	opt.CreateOption = &CreateOption{}
	opt.ListOption = &ListOption{}
	parser := flags.NewParser(opt, flags.HelpFlag|flags.PassDoubleDash)
	return parser
}

func main() {
	buf := &bytes.Buffer{}
	var opt Option
	parser := newOptionParser(&opt)
	parser.WriteHelp(buf)
	fmt.Println(buf.String())
}
