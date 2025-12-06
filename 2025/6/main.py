from functools import reduce
import operator
import os
from ..utils.parse import parse_carriage_return_list

lines = parse_carriage_return_list(os.path.join(os.path.dirname(__file__), 'input.txt'))



grand_total = 0
number_of_operations = len(lines[0].split())
multiplications = []
additions = []
operations = [[] for _ in range(number_of_operations)]

for line in lines:
    for operation_index, content in enumerate(line.split()):
        if content == '+':
            # print(f"Adding {' + '.join([str(x) for x in operations[operation_index]])} to grand total")
            grand_total += sum(operations[operation_index])
            additions.append(operations[operation_index])
        elif content == '*':
            # print(f"Adding {' * '.join([str(x) for x in operations[operation_index]])} to grand total")
            grand_total += reduce(operator.mul, operations[operation_index], 1)
            multiplications.append(operations[operation_index])
        else:
            operations[operation_index].append(int(content))

print(grand_total)

second_grand_total = 0
number_of_characters = len(lines[0])
current_operation = []
current_nbr = ''
current_operation_type = None
for column in range(number_of_characters):
    column_is_empty = True
    for line in lines:
        character = line[column]
        if character==" ":
            continue
        elif character in ['+', '*']:
            current_operation_type = character
            column_is_empty = False
        else:
            current_nbr += character
            column_is_empty = False
    if current_nbr!='':
        current_operation.append(int(current_nbr))
        current_nbr = ''
    if column_is_empty:
        if current_operation_type == '+':
            second_grand_total += sum(current_operation)
        elif current_operation_type == '*':
            second_grand_total += reduce(operator.mul, current_operation, 1)
        current_operation = []
        current_operation_type = None
        current_nbr = ''
        
# Final operation      
if current_operation_type == '+':
    second_grand_total += sum(current_operation)
elif current_operation_type == '*':
    second_grand_total += reduce(operator.mul, current_operation, 1)


print(second_grand_total)
