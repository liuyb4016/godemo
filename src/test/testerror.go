package test

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string{
	strFormat:=`
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat,de.dividee)
}

func Divide(varDivide int,varDivider int) (result int,errorMsg string){
	if varDivider==0{
		dData := DivideError{
			dividee:varDivide,
			divider:varDivider,
		}
		errorMsg = dData.Error()
		return
	}else{
		return varDivide/varDivider,""
	}
}

func TestError(){
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

