package main

type Parent struct{}

func (p Parent) say() string { return "parent" }
func (p Parent) Print()      { println(p.say()) }
func (p Parent) Hello()      { println(p.say()) }

type Child struct{ Parent }

func (c Child) say() string { return "child" }

// Print() は Parent の処理を実行する
func (c Child) Hello() { println(c.say()) }

func main() {
	p := Parent{}
	c := Child{}
	p.Hello()
	c.Hello()
	c.Print()
}
