package metaspace

import (
	"fmt"
	"h-jvm/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class     *Class
	constants []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp) // 常量池的大小
	constants := make([]Constant, cpCount)
	cp := &ConstantPool{class: class, constants: constants}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			constants[i] = intInfo.Value()
		}
	}
	return cp
}

func (s *ConstantPool) GetConstant(index uint) Constant {
	if c := s.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
