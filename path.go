package web

//Path represents http path eg /user
type Path string

func (path Path) String() string {
	return string(path)
}
