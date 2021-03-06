package main

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	combineBacktracking(n, k, 1, &result, &temp)
	return result
}

func combineBacktracking(n, k, index int, result *[][]int, temp *[]int) {
	if len(*temp) == k {
		new := make([]int, k)
		for i := 0; i < k; i++ {
			new[i] = (*temp)[i]
		}
		*result = append(*result, new)
		return
	}

	if (n-index+1)+len(*temp) < k {
		return
	}

	for i := index; i <= n; i++ {
		*temp = append(*temp, i)
		combineBacktracking(n, k, i+1, result, temp)
		*temp = (*temp)[:len(*temp)-1]
	}
}
