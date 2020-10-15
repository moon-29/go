package main

import "fmt"

   func main()  {
	var a float32
	var c float32
	var b string
	fmt.Println("请第一次输入")
	fmt.Scanln(&a)
	fmt.Println("请输入符号")
	fmt.Scanln(&b)
	fmt.Println("请第二次输入")
	fmt.Scanln(&c)

	switch b {
	case "+":
		fmt.Printf("%f + %f= %.2f\n",a,c,a+c)
	case "-":
        fmt.Printf("%f + %f= %.2f\n",a,c,a-c)
	case "*":
		fmt.Printf("%f * %f= %.2f\n",a,c,a*c)
	case "/":
		fmt.Printf("%f / %f= %.2f\n",a,c,a/c)
	default:
		fmt.Println("kill me")

	}

}