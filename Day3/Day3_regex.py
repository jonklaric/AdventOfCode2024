import os
import re

# The regex pattern from before
pattern = r"mul\(\d{1,3},\d{1,3}\)"

def read_file(file):
    with open(file, "r") as inputfile:
        data = inputfile.read()
    return data

def day3_part1(input_file):
    test_string = read_file(input_file)
    matches_findall = re.findall(pattern, test_string)
    matches_findall = [x[4:-1].split(",") for x in matches_findall]
    matches_findall = [int(numpairs[0])*int(numpairs[1]) for numpairs in matches_findall]
    return sum(matches_findall)

test_input_file = os.path.dirname(__file__) + "\\test_input.txt"
print(day3_part1(test_input_file))

my_input_file = os.path.dirname(__file__) + "\\my_input.txt"
print(day3_part1(my_input_file))

