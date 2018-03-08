package main

import (

	"io"
	"strings"
	"os"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (mr MyReader) Read(out []byte) (int, error){
	for i := range out {
		out[i] = 'A'
	}
	return len(out), nil
}

type rot13Reader struct {
	r io.Reader
}

func (mr rot13Reader) Read(out []byte) (int, error){
	n, err := mr.r.Read(out)
	if err == io.EOF {
		return 0,io.EOF
	}

	for i := range out {
		if('A' <= out[i] && out[i] <= 'M') || ('a' <= out[i] && out[i] <= 'm'){
			out[i] += 13
		} else if('N' <= out[i] && out[i] <= 'Z') || ('n' <= out[i] && out[i] <= 'z'){
			out[i] -= 13
		} else if out[i] == ' '{
		}
	}
	return n, nil
}

func main() {
	//reader.Validate(MyReader{})
	s := strings.NewReader("Lbh penpxrq gur pbqr")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}