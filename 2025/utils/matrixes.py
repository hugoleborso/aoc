import numpy as np

def add_matrixes_from_cell(base_matrix:np.ndarray[int], increment_matrix:np.ndarray[int], row_index:int, column_index:int):
    base_matrix_vertical_range = slice(max(0, row_index), max(min(row_index + increment_matrix.shape[0], base_matrix.shape[0]), 0))
    base_matrix_horizontal_range = slice(max(0, column_index), max(min(column_index + increment_matrix.shape[1], base_matrix.shape[1]), 0))

    increment_matrix_vertical_range = slice(max(0, -row_index), min(-row_index + base_matrix.shape[0], increment_matrix.shape[0]))
    increment_matrix_horizontal_range = slice(max(0, -column_index), min(-column_index + base_matrix.shape[1], increment_matrix.shape[1]))

    base_matrix[base_matrix_vertical_range, base_matrix_horizontal_range] += increment_matrix[increment_matrix_vertical_range, increment_matrix_horizontal_range]

def print_matrix(matrix:np.ndarray):
    print("-"*matrix.shape[1])
    for row in matrix:
        print("".join(str(cell) for cell in row))
    print("-"*matrix.shape[1])
    print("")