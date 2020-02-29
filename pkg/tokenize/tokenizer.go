package tokenize

import "fmt"

var (
	Delimeters = []string{""}
)

type Tokenizer struct {
	pos int
}

type Token struct {
	token string
}

func (t *Tokenizer) Parse(str string) []Token {
	tokens := []Token{}

	for {
		token := t.NextToken(str)
		fmt.Println("token: ", token)

		if token.token == "" {
			break
		}

		tokens = append(tokens, token)
	}

	// fmt.Println("char: ", string(char))
	// if string(char) != " " {
	// 	token = token + string(char)
	// 	fmt.Println("token: ", token)
	// }

	// if string(char) == " " || i == (len(str)-1) {
	// 	tokens = append(tokens, token)
	// 	token = ""
	// }

	fmt.Println("tokens: ", tokens)

	return tokens
}

func (t *Tokenizer) NextToken(str string) Token {
	fmt.Println("t.pos: ", t.pos)

	str = str[t.pos:]

	var token string
	for i, char := range str {
		fmt.Println("char: ", string(char))
		if string(char) != " " {
			t.pos++
			token = token + string(char)
		}

		if string(char) == " " {
			t.pos++
			return Token{
				token: token,
			}
		}

		if i == (len(str) - 1) {
			return Token{
				token: token,
			}
		}
	}

	return Token{}
}
