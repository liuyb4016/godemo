package test

import "fmt"

func TestPointerV1(){
	var b int = 9
	fmt.Printf("thi is a pointer value:%x\n",&b)

	var c *int
	c = &b
	fmt.Printf("*ip 变量的值: %d\n", *c )

}

const MAX  = 3

func TestPointerV2(){
	var c *int
	if(c!=nil){
		fmt.Printf("*ip 变量的值: %d\n", *c )
	}else{
		fmt.Println("*c是空指针")
	}

	a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }
}


func TestPointerV3(){
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	ptr = &a

	pptr = &ptr

	 /* 获取 pptr 的值 */
   fmt.Printf("变量 a = %d\n", a )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

func swap(x *int,y *int){
	var temp int
	temp = *x
	*x = *y
	*y =temp
}

func TestPointerV4(){
	/* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前 a 的值 : %d\n", a )
   fmt.Printf("交换前 b 的值 : %d\n", b )

   /* 调用函数用于交换值
   * &a 指向 a 变量的地址
   * &b 指向 b 变量的地址
   */
   swap(&a, &b);

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}
