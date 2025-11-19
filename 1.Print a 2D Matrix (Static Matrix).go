package main

import (
	"fmt"
)

func matrix(){
// 	var r,c int
// fmt.Print("entrr rows : ")
// fmt.Scan(&r)
// fmt.Print("enter cols : ")
// fmt.Scan(&c)

// matrixx:=make([][]int,r)
// for i:=0;i<r;i++{
// 	matrixx[i]=make([]int,c)
// 	for j:=0;j<c;j++{
// 		fmt.Scan(&matrixx[i][j])
// 	}
// }

// fmt.Println(matrixx)

mat:=[][]int{
	{1,2,3},
	{4,5,6},
	{7,8,9},
}
fmt.Println("len",len(mat))
fmt.Println("col len",len(mat[0]))
for i:=0;i<len(mat);i++{
	for j:=0;j<len(mat[i]);j++{
		fmt.Print(mat[i][j]," ")
	}
	fmt.Println()
}


}