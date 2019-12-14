package parser

import (
	"github.com/arata-nvm/Solitude/lexer"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct{
		input string
		expected string
	}{
		{"0", "Int(0)"},
		{"42", "Int(42)"},

		{"4 + 4", "Infix(Int(4) + Int(4))"},
		{"4 - 4", "Infix(Int(4) - Int(4))"},
		{"4 * 4", "Infix(Int(4) * Int(4))"},
		{"4 / 4", "Infix(Int(4) / Int(4))"},
		{"4 == 4", "Infix(Int(4) == Int(4))"},
		{"4 != 4", "Infix(Int(4) != Int(4))"},
		{"4 < 4", "Infix(Int(4) < Int(4))"},
		{"4 <= 4", "Infix(Int(4) <= Int(4))"},
		{"4 > 4", "Infix(Int(4) > Int(4))"},
		{"4 >= 4", "Infix(Int(4) >= Int(4))"},

		{"4 + 4 * 4", "Infix(Int(4) + Infix(Int(4) * Int(4)))"},
		{"4 * 4 + 4", "Infix(Infix(Int(4) * Int(4)) + Int(4))"},

		{"var hoge = 1", "var Ident(hoge) = Int(1)"},
		{"var fuga = hoge * 2 + 2", "var Ident(fuga) = Infix(Infix(Ident(hoge) * Int(2)) + Int(2))"},

		{"return 0", "return Int(0)"},
		{"return hoge", "return Ident(hoge)"},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		actual := p.ParseProgram().Inspect()
		checkParserErrors(t, p)

		if actual != test.expected {
			t.Fatalf("tests[%d] - expected=%q, got=%q", i, test.expected, actual)
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.Errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(p.Errors))

	for _, msg := range p.Errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
