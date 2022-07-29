package image

func mulMatrix(a [][]int, b [][]int, m int, n int, k int) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		resI := make([]int, k)
		for j := 0; j < n; j++ {
			for kk := 0; kk < k; kk++ {
				resI[kk] += a[i][j] * b[j][kk]
			}
		}
		res[i] = resI
	}
	return res
}
