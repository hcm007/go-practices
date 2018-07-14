package display

import "gopl.io/ch7/eval"

func Example_expr() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
	// Output:
	// Display e (eval.call):
}

func Example_slice(){
	Display("slice",[]*int{new(int),nil})
}

func Example_nilInterface(){

}