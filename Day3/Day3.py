import os
import numpy as np

def read_file(file):
    with open(file, "r") as inputfile:
        data = inputfile.read()
    return data

def find_all_occurences(string, substring):
    locations = []
    i = 0
    while i < len(string) and i >= 0:
        temp_string = string[i:]
        found_location = temp_string.find(substring)
        if found_location >= 0:
            locations.append(i + found_location)
            i = found_location + i + 4
        else:
            break            
    return locations

def day3_part1(input_file_location):
    running_sum = 0
    text = read_file(input_file_location)
    locations = find_all_occurences(text, "mul(")
    print(locations)
    for loc in locations:
        search_string = text[loc+4:loc+13]
        if ")" not in search_string or "," not in search_string:
            continue
        search_string = search_string[:search_string.find(")")]
        #print(search_string.split(",", 1))
        numbers = search_string.split(",", 1)
        if all([chars.isnumeric() for chars in numbers]):
            running_sum += np.prod([int(chars) for chars in numbers])
        else:
            continue        
    return running_sum


def day3_part2(input_file_location):
    running_sum = 0
    text = read_file(input_file_location)
    do_locations = find_all_occurences(text, "do()")
    print("do_locations", do_locations)
    do_not_locations = find_all_occurences(text, "don't()")
    print("do_not_locations", do_not_locations)
    do_or_dont = True
    do_dont_list = []
    for i in range(len(text)):
        if len(do_locations) > 0 and i==do_locations[0]:
            do_or_dont = True
            do_locations.pop(0)
        elif len(do_not_locations) > 0 and i == do_not_locations[0]:
            do_or_dont = False
            do_not_locations.pop(0)
        do_dont_list.append(do_or_dont)
    text = "".join([text[i] if do_dont_list[i] else "*" for i in range(len(text)) ])
    #print(text)
    locations = find_all_occurences(text, "mul(")
    #print("mul locations", locations)
    for loc in locations:
        search_string = text[loc+4:loc+13]
        if ")" not in search_string or "," not in search_string:
            continue
        search_string = search_string[:search_string.find(")")]
        print(search_string.split(",", 1))
        numbers = search_string.split(",", 1)
        if all([chars.isnumeric() for chars in numbers]):
            running_sum += np.prod([int(chars) for chars in numbers])
        else:
            continue        
    return running_sum



# test_input_file = os.path.dirname(__file__) + "\\test_input.txt"
# print(read_file(test_input_file))
# print(day3_part1(test_input_file))
my_input_file = os.path.dirname(__file__) + "\\my_input.txt"
# print(day3_part1(my_input_file))
test_input_file = os.path.dirname(__file__) + "\\test_input_part2.txt"
print(day3_part2(test_input_file))

print(day3_part2(my_input_file))
