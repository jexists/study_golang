// https://school.programmers.co.kr/learn/courses/30/lessons/12950?language=javascript
// 행렬의 덧셈
// 문제 설명
// 행렬의 덧셈은 행과 열의 크기가 같은 두 행렬의 같은 행, 같은 열의 값을 서로 더한 결과가 됩니다. 2개의 행렬 arr1과 arr2를 입력받아, 행렬 덧셈의 결과를 반환하는 함수, solution을 완성해주세요.

// 제한 조건
// 행렬 arr1, arr2의 행과 열의 길이는 500을 넘지 않습니다.

// My Code
package main

import "fmt"

func main() {
	// solution([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}})
	solution([][]int{{1}, {2}}, [][]int{{3}, {4}})
	// solution([[1,2],[2,3]]	[[3,4],[5,6]]	)
}

func solution(arr1 [][]int, arr2 [][]int) [][]int {

	// fmt.Println(arr1)
	// fmt.Println(arr2)
	res := make([][]int, len(arr1))
	fmt.Println(res)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Println(i)
			fmt.Println(j)
			fmt.Println(arr1[i][j])
			res[i] = append(res[i], arr1[i][j]+arr2[i][j])
		}
	}

	fmt.Println(res)
	return res
}

// 다른사람 코드 ========================

func solution1(arr1 [][]int, arr2 [][]int) [][]int {
	for i := range arr1 {
		for j := range arr1[i] {
			arr1[i][j] += arr2[i][j]
		}
	}

	return arr1
}

func solution2(arr1 [][]int, arr2 [][]int) [][]int {

	row := len(arr1)
	col := len(arr1[0])

	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return result
}

func solution3(arr1 [][]int, arr2 [][]int) [][]int {
	ret := make([][]int, len(arr1))
	for i := range ret {
		ret[i] = make([]int, len(arr1[i]))
		for j := range ret[i] {
			ret[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	return ret
}

func solution4(arr1 [][]int, arr2 [][]int) [][]int {
	r := make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		r[i] = make([]int, len(arr1[i]))
		for j := 0; j < len(arr1[i]); j++ {
			r[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return r
}
