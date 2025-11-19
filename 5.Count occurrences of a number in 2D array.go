package main


import(
	"fmt"
)


func countOccurences(mat [][]int){
	freq:=make(map[int]int)

	for i:=0;i<len(mat);i++{
		for j:=0;j<len(mat[i]);j++{
			freq[mat[i][j]]++
		}
	}

	for i,val:=range freq{
		fmt.Println(i,"->",val)
	}
}