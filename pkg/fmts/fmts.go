package fmts

import (
	"io"
	"os"
	"strings"

	gconcat "github.com/jeffotoni/concat"
)

//Stdout func
func Stdout(strs ...interface{}) {
	str := gconcat.Build(strs...)
	io.Copy(os.Stdout, strings.NewReader(str))
}

//Concat contaquena
func Concat(strs ...interface{}) string {
	return gconcat.Build(strs...)
}

//Println printa com n\
func Println(strs ...interface{}) {
	str := gconcat.Build(strs...)
	str = gconcat.Build(str, "\n")
	io.Copy(os.Stdout, strings.NewReader(str))
}

//Print printa
func Print(strs ...interface{}) {
	str := gconcat.Build(strs...)
	str = gconcat.Build(str, "\n")
	io.Copy(os.Stdout, strings.NewReader(str))
}
