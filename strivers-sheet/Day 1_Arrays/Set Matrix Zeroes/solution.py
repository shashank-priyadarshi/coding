def matrix_search(arr, row, col):
    #     iterate over rows
    #     iterate over columns
    #     get index if element equals zero
    zero_row = list()
    zero_col = list()
    for r in range(row):
        for c in range(col):
            if (arr[r][c] == 0):
                zero_row.append(r)
                zero_col.append(c)
    return (matrix_set(arr, row, col, set(zero_row), set(zero_col)))


def matrix_set(arr, row, col, rows, cols):
    #     iterate over rows
    #     iterate over cols
    #     set col as zero if r in rows or c in cols
    for r in range(row):
        for c in range(col):
            if (r in rows) or (c in cols):
                arr[r][c] = 0
    return arr


(row, col) = (int(input()), int(input()))
arr = [[int(input()) for x in range(col)]for y in range(row)]
print(arr)
print(matrix_search(arr, row, col))
