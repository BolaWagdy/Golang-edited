package calculator
func Calc(op1 int, operator byte,op2 int)int{
	switch operator {
	case '+':
		return op1 + op2;
	case '-':
		return op1- op2;
	case '*':
		return op1 * op2;
	case '/':
		return op1 /op2;
	default:
		return 0;
	}
}