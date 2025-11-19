package main 

import(
"fmt"
)

func rowWiseSum(mat [][]int ){
	tempArr:=[]int{}
	

	for i:=0;i<len(mat);i++{
		rsum:=0
		for j:=0;j<len(mat[i]);j++{
			rsum+=mat[i][j]
		}
		tempArr=append(tempArr,rsum)
	}

	fmt.Println(tempArr)

	//columnwise sum

	temp:=[]int{}

	for k:=0;k<len(mat[0]);k++{
		csum:=0
		for l:=0;l<len(mat);l++{
			csum+=mat[l][k]
		}
		temp=append(temp,csum)

	}
	fmt.Println(temp)
}