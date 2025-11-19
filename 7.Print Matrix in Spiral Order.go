package main
import (
	"fmt"
)

func printSpiral(mat [][]int){

	top,bottom:=0,len(mat)-1
	left,right:=0,len(mat[0])-1

	for top<=right && right<=left{

		for i:=left;i<right;i++{
			fmt.Println(mat)
		}
		
	}

}