import os
from ..utils.parse import parse_carriage_return_list

lines = parse_carriage_return_list(os.path.join(os.path.dirname(__file__), 'input.txt'))

position = 50
password_count = 0

for line in lines:
    direction = line[0]
    number = int(line[1:])
    password_count += number//100
    
    if direction == 'L':
        new_position = position - number%100
    elif direction == 'R':
        new_position = position + number%100
    if position!=0 and (new_position <= 0 or new_position >= 100):
         password_count += 1        

    position = new_position % 100
    


print(password_count)   
