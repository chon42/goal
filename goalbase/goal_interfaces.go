package goalbase

import (
	"reflect"
)

type IPrintf interface {
	Printf(format string, args ...any) (int, error)
}

type IName interface {
	Name() string
}
type IType interface {
	TypeName() string
	Type() reflect.Type
	ValueOf(rtype reflect.Type) any
}
type IValue interface {
	Value() any
}
type IValueRange interface {
	ValueRange() string
	ShouldTest() bool
	TestInRange(val any) func(val any) bool
}

type IArg interface {
	IName
	IType
	IValue
}
type IAlgo interface {
	IName
	Takes() []IArg
	Returns() []IArg
	Effects() map[string]string
	Does() string
	AlgoDesc() string
	Func() func(any) any
}

type IMathPrimeUnary interface {
	Calc(int64) (int64, error)
}

type IMathPrimeBinary interface {
	Calc(int64, int64) (int64, error)
}
type IMathPrimeBinBin interface {
	Calc(int64, int64) (int64, int64, error)
}

type IMathPrimeBSliceBin interface {
	Calc([]int64) ([]int64, error)
}

type IMathPrimeFUnary interface {
	Calc(float64) (float64, error)
}

type IMathPrimeFBinary interface {
	Calc(float64, float64) (float64, error)
}

type IMathSliceFFloat interface {
	Calc([]float64) (float64, error)
}

type IMathSliceFScliceF interface {
	Calc([]float64) ([]float64, error)
}

type Instruction interface {
	Do(IEnviron)
}

type IEnviron interface {
}

type IMathAlgLong interface {
	CalcLong([]Instruction,
		[]uint64, []int64, []float64, []complex128) (uint64,
		[]int64, []float64, []complex128, error)
}
type reg = uintptr
type Index struct {
	Type  byte
	Index uint16
}

const (
	Unused int = iota << 1
	Int
	Uint
	Float
	Complex
	BigInt
	BigFloat
	String
	Bytes
)

type IMathAlg interface {
	Calc([]Instruction,
		[]reg, []reg) ([]reg, []reg, error)
}

type IMath interface {
	IMathAlgLong
	IMathAlg
}

type IMathBig interface {
	Calc([]Instruction, []Index) ([]uint32, error)
}
