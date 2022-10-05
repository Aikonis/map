package mapq

var Err int64
var Index int64

// E→TE'
func E() {

	if Err == 0 {
		//T()
		E1()
	}
}

// E'→+TE'| -TE' |ε

func E1() {
	//if Err == 0 && Index != len() {
	//	Index += 1
	//	if Result_Lex[0][Index] in [11, 12] {
	//	if Index != len(Result_Lex[0]) - 1{
	//	T()
	//	E1()
	//	}else{
	//	Index = len(Result_Lex[0])
	//	}
	//
	//	}else if Result_Lex[0][Index] != 24{
	//	Err = 1
	//	}
	//}

}

// Parser 语法分析器
type Parser struct {
	l *Lexer
}

// 你的递归下降分析代码（如果你使用递归下降的话
// func (p *Parser) boolexp() (node Node, err error) {
// 	panic("not implemented")
// }

// func (p *Parser) boolean() (node Node, err error) {
// 	panic("not implemented")
// }
// 别的分析函数
// 。。。。。

// Parse 生成ast
func (p *Parser) Parse(str string) (n Node, err error) {

	panic("not implemented")
}
