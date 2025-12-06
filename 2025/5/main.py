import os
from ..utils.parse import  parse_carriage_return_list_with_blank_lines

lines = parse_carriage_return_list_with_blank_lines(os.path.join(os.path.dirname(__file__), 'input.txt'))
fresh_ingredients_ranges_str = lines[0]
available_ingredients = [int(nb) for nb in lines[1]]
fresh_ingredients_start_end_pairs = [[int(nb) for nb in ing_range.split('-')] for ing_range in fresh_ingredients_ranges_str]
fresh_ingredients_start_end_pairs_in_order = [[min(pair),max(pair)] for pair in fresh_ingredients_start_end_pairs]


number_of_fresh_ingredients = 0

for available_ingredient in available_ingredients:
    for start, end in fresh_ingredients_start_end_pairs_in_order:
        if start <= int(available_ingredient) <= end:
            number_of_fresh_ingredients += 1
            break
print(number_of_fresh_ingredients)

number_of_potentially_fresh_ingredients = 0
fresh_ingredients_start_end_pairs_sorted = sorted(fresh_ingredients_start_end_pairs_in_order, key=lambda element: (element[0], element[1]))
previous_end = fresh_ingredients_start_end_pairs_sorted[0][0]-1

for index, fresh_ingredient_range in enumerate(fresh_ingredients_start_end_pairs_sorted):
    start, end = fresh_ingredient_range
    if end <= previous_end:
        # print(f"Skipping range {fresh_ingredient_range} as it is fully covered by previous_end {previous_end}")
        continue

    # print(f"Considering range {fresh_ingredient_range} with previous_end {previous_end}")

    considered_start = max(previous_end + 1, start)
    # print(f"  Considered range is from {considered_start} to {end}")

    number_of_potentially_fresh_ingredients += end - considered_start + 1
    # print(f"  Current nb of potentially fresh ingredients: {number_of_potentially_fresh_ingredients}")
    previous_end = end

print(number_of_potentially_fresh_ingredients)