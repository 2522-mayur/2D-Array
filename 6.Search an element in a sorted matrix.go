package main
import (
	"fmt"
)

func searchElement(mat [][]int,target int){

//brute force approach --> tc=O(n*n)
	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			if mat[i][j]==target{
				fmt.Println(i,j)
			}
		}
	}

	//optimized approach --> tc=O(m+n)

	r,c:=0,len(mat[0])-1
	for r<len(mat) && c>=0{
		if mat[r][c]==target{
			fmt.Println("found at ",r,c)
			return
		}else if mat[r][c] > target{
			c--
		}else{
			r++
		}
	}
}