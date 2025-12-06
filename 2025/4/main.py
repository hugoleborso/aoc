import os
import numpy as np
from ..utils.parse import parse_carriage_return_list
from ..utils.matrixes import add_matrixes_from_cell #,print_matrix

lines = parse_carriage_return_list(os.path.join(os.path.dirname(__file__), 'input.txt'))

rolls_matrix = np.zeros((len(lines), len(lines[0])), dtype=int)
visible_matrix = np.full((len(lines), len(lines[0])), ".", dtype=str)
rolls_neigbor_count_matrix = np.zeros((len(lines), len(lines[0])), dtype=int)
matrix_increment = np.array([[1,1,1], [1,0,1], [1,1,1]])

removed_rolls = 0
removed_rolls_previous_step = None

for row_index, line in enumerate(lines):
    for column_index, content in enumerate(line):
        if content=="@":
            rolls_matrix[row_index][column_index] = 1

while removed_rolls != removed_rolls_previous_step:
    removed_rolls_previous_step = removed_rolls
    visible_matrix = np.full((len(lines), len(lines[0])), ".", dtype=str)
    rolls_neigbor_count_matrix = np.zeros((len(lines), len(lines[0])), dtype=int)

    for row_index, line in enumerate(rolls_matrix):
        for column_index, content in enumerate(line):
            if content==1:
                visible_matrix[row_index][column_index] = "@"
                add_matrixes_from_cell(rolls_neigbor_count_matrix, matrix_increment, row_index-1, column_index-1)


    for row_index, line in enumerate(rolls_matrix):
        for column_index, content in enumerate(line):
            if content==1:
                if rolls_neigbor_count_matrix[row_index][column_index] < 4:
                    removed_rolls += 1
                    visible_matrix[row_index][column_index] = "x"
                    rolls_matrix[row_index][column_index] = 0
    
    # print_matrix(visible_matrix)

print(removed_rolls)