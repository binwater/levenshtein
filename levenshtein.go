package levenshtein

import (
	"math"
)

//Levenshtein 两个字符串的编辑距离
func Levenshtein(a, b string) float64 {
	r1 := []rune(a)
	r2 := []rune(b)
	len1 := len(r1)
	len2 := len(r2)

	matrix := make([][]float64, 0, len1+1)

	for i := 0; i < len1+1; i++ {
		dlist := make([]float64, 0, len2+1)
		if i == 0 {
			for j := 0; j < len2+1; j++ {
				dlist = append(dlist, float64(j))
			}
		} else {
			dlist = append(dlist, float64(i))
		}
		matrix = append(matrix, dlist)
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			var c float64
			if r1[i-1] != r2[j-1] {
				c = 1
			}

			val := math.Min(math.Min(matrix[i-1][j-1]+c, matrix[i][j-1]+1), matrix[i-1][j]+1)
			matrix[i] = append(matrix[i], val)
		}
	}

	return matrix[len1][len2]
}

//SimilarDegree 字符串相似度
func SimilarDegree(a, b string) float64 {
	return 1 - Levenshtein(a, b)/math.Max(float64(len([]rune(a))), float64(len([]rune(b))))
}
