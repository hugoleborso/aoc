def parse_carriage_return_list(input_path):
    with open(input_path, 'r') as f:
        return f.read().splitlines()
    
def parse_comma_separated_list(input_path):
    with open(input_path, 'r') as f:
        return f.read().split(',')

def parse_carriage_return_list_with_blank_lines(input_path)->list[list[str]]:
    with open(input_path, 'r') as f:
        lines = f.read().splitlines()
    returned_lines = [[]]
    for line in lines:
        if line.strip() == '':
            returned_lines.append([])
        else:
            returned_lines[-1].append(line)
    return returned_lines