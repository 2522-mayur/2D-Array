package main 

import (
	"fmt"
)

func minMax(mat [][]int){
min,max:=mat[0][0],mat[0][0]

for i:=0;i<len(mat);i++{
	for j:=0;j<len(mat[i]);j++{
		if mat[i][j]<min{
			min=mat[i][j]
		}else if mat[i][j]>max{
			max=mat[i][j]
		}
	}
}

fmt.Println(min,max)



}