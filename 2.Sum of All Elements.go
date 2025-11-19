package main

import(
	"fmt"
)

func sum(mat [][]int)int{

	sum:=0
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			sum+=mat[i][j]
		}
	}
	fmt.Println(sum)
	return sum
	

}