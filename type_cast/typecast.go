package main

type Gender interface {
	Run()
}
type Male struct {
	Name string
}
func (this Male) Run() {
	println(this.Name)
}

type Female struct {
	Name string
}
func (this Female) Run() {
	println(this.Name)
}

type PointedFemale struct {
	Name string
}
func (this *PointedFemale) Run() {
	println(this.Name)
}

func main() {
	m := Male{Name: "hoge"}
	var i Gender = m // copied structure
	i.Run()
	m.Name = "fuga"
	i.Run()

	a := []Gender{Male{Name: "Male"}, Female{Name: "Female"}}
	for _, i := range a {
		i.Run()
	}

	p := PointedFemale{Name: "pointer"}
	var f Gender = &p
	f.Run()
	p.Name = "changed"
	f.Run()
}
