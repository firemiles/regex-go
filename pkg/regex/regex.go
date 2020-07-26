package regex

//Regexp ...
type Regexp struct {
	nfa *State
}

//Compile ...
func Compile(pattern string) (*Regexp, error) {
	se, err := SuffixExpression(pattern)
	if err != nil {
		return nil, err
	}
	nfa, err := buildNFA(se)
	if err != nil {
		return nil, err
	}
	return &Regexp{
		nfa: nfa,
	}, nil
}

//SuffixExpression ...
//- 创建输出队列 output 和运算符栈 ops
//- 依次读取输入字符串中每一个字符 ch
// - 如果 ch 是普通字符，追加到 output
// - 如果 ch 是运算符，只要 ops 栈顶的运算符优先级不低于 ch，依次出栈并追加到 output，最后将 ch 入栈 ops
// - 如果 ch 是“(”，入栈 ops
// - 如果 ch 是“)”，只要 ops 栈顶不是“(”，依次出栈并追加到 output
//- 将 ops 中运算符依次出栈追加到 output
//- 返回 output
// 4. * ? +
// 3. .
// 2. |
// 1. (
func SuffixExpression(exp string) (token []string, err error) {
	ops := Stack{}
	var str []byte
	for i := 0; i < len(exp); i++ {
		c := exp[i]
		switch c {
		case ')':
			for {
				op, err := ops.Pop()
				if err != nil {
					return nil, err
				}
				if op == '(' {
					break
				}
				token = append(token, op)
			}
		case '*', '?', '+', '.', '|', '(':
			for {
				if ops.Len() == 0 {
					break
				}
				op, _ := ops.Top()
				if opsPriority(c) > opsPriority(op) {
					break
				}
				op, _ = ops.Pop()
				token = append(token, string(op))
			}
			ops.Push(c)
		default:
			str = append(str, c)
		}
	}
	for ops.Len() > 0 {
		op, _ := ops.Pop()
		token = append(token, string(op))
	}
	return token, nil
}

func opsPriority(ops byte) int {
	switch ops {
	case '*', '?', '+':
		return 4
	case '.':
		return 3
	case '|':
		return 2
	case '(':
		return 1
	default:
		return 0
	}
}

func buildNFA(token []byte) (*State, error) {
	root := State{}
	for 

}

//Find ...
func (re *Regexp) Find(b string) string {
	return ""
}

//Match ...
func (re *Regexp) Match(s string) (bool, error) {
	return false, nil
}
