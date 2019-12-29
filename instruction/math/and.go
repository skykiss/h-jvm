package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x7e iand     将栈顶两int型数值作“按位与”并将结果压入栈顶
//0x7f land     将栈顶两long型数值作“按位与”并将结果压入栈顶

type IAnd struct {
	base.NoOperandsInstruction
}

func (i *IAnd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 & val2)
}

type LAnd struct {
	base.NoOperandsInstruction
}

func (i *LAnd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 & val2)
}
