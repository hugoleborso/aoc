import os
from ..utils.parse import parse_carriage_return_list

lines = parse_carriage_return_list(os.path.join(os.path.dirname(__file__), 'input.txt'))

joltage_sum = 0

for line in lines:
    bank_joltages = [int(x) for x in line]
    for i in range(11,0, -1):
        bestjoltage = max(bank_joltages[:-i])
        best_joltage_index = bank_joltages.index(bestjoltage)
        joltage_sum += bestjoltage*10**i
        bank_joltages = bank_joltages[best_joltage_index+1:]
    bestjoltage = max(bank_joltages)
    joltage_sum += bestjoltage

print(joltage_sum)   
