package main


func rotate(matrix [][]int)  {
	var n int = len(matrix);
	for i := 0; i < n; i++ {
		for j := 0; j < n; i++ {
			swap(&matrix[i][j], &matrix[n-i][n-j]);
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; i++ {
			swap(&matrix[i][j], &matrix[n-i][j]);
		}
	}
}

func swap(x *int, y *int) {
	var temp int;
	temp  = *y;
	*y = *x;
	*x = temp;
}