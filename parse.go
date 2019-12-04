package fakego

import (
	"bufio"
	"errors"
	"github.com/esrrhs/go-engine/src/loggo"
	"os"
	"strconv"
)

func Parse(file string) (err error) {
	var ll *lexerwarpper
	defer func() {
		if r := recover(); r != nil {
			line := 0
			if ll != nil {
				line = ll.yyLexer.(*Lexer).Line()
			}
			switch x := r.(type) {
			case string:
				err = errors.New(file + ":" + strconv.Itoa(line) + ":" + x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	loggo.Debug("start parse " + file)

	f, err := os.OpenFile(file, os.O_RDONLY, os.ModeType)
	if err != nil {
		panic(err)
	}
	fr := bufio.NewReader(f)
	lex := NewLexer(fr)

	mf := myflexer{}

	l := lexerwarpper{
		lex,
		mf,
	}
	ll = &l

	ret := yyParse(l)
	if ret != 0 {
		panic(errors.New("yyParse fail " + strconv.Itoa(ret)))
	}

	loggo.Debug("start parse ok" + file)
	return nil
}
