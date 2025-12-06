import os
from ..utils.parse import parse_comma_separated_list

lines = parse_comma_separated_list(os.path.join(os.path.dirname(__file__), 'input.txt'))
invalid_ids_sum = 0

for line in lines:
    first_id_str, last_id_str = line.split('-')
    first_id = int(first_id_str)
    last_id = int(last_id_str)

    for id in range(first_id, last_id + 1):
        str_id = str(id)
        id_len = len(str_id)
        for seq_len in range(1, id_len//2+1):
            if id_len % seq_len != 0:
                continue
            sequences = [str_id[seq_index*seq_len:(seq_index+1)*seq_len] for seq_index in range(id_len//seq_len)]
            if len(set(sequences)) == 1:
                invalid_ids_sum += id
                break

print(invalid_ids_sum)
        
