package preprocess

import (
	"bytes"
)

type preprocesser struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func (p *preprocesser) peekChar() byte {
	if p.readPosition >= len(p.input) {
		return 0
	}
	return p.input[p.readPosition]
}

func (p *preprocesser) readChar() byte {
	if p.readPosition >= len(p.input) {
		p.ch = 0
	} else {
		p.ch = p.input[p.readPosition]
	}
	p.position = p.readPosition
	p.readPosition++
	return p.ch
}

func newPreprocesser(input string) *preprocesser {
	p := &preprocesser{input: input}
	p.readChar()
	return p
}

func Preprocess(input string) string {
	var out bytes.Buffer
	p := newPreprocesser(input)

	//コメント削除
	for p.ch != 0 {
		if p.ch == '/' && p.peekChar() == '/' {
			p.readChar()
			for {
				p.readChar()
				if p.ch == '\n' || p.ch == 0 {
					break
				}
			}
		}
		if p.ch != 0 {
			out.WriteByte(p.ch)
		}
		p.readChar()
	}

	return out.String()
}
