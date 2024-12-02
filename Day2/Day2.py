import os
from functools import reduce
from operator import mul

def read_file(file):
    with open(file, "r") as inputfile:
        data = inputfile.readlines()
        data = [row.replace("\n", "") for row in data]
        data = [row.split(" ") for row in data]
        data = [[int(x) for x in row] for row in data]
    return data

def acceptable_diff(element, prev_element):
    abs_diff = abs(element - prev_element)
    if abs_diff >= 1 and abs_diff <= 3:
        return True
    else:
        return False

def isReportSafe(row):
    if row[1] > row[0]:
        increasing = True
    elif row[1] < row[0]:
        increasing = False
    else:
        # first two elements are the same
        return False
    for i in range(1, len(row)):
        if not acceptable_diff(row[i], row[i-1]):
            return False
        if increasing and row[i] <= row[i-1]:
            return False
        elif not increasing and row[i] >= row[i-1]:
            return False
        else:
            continue
    return True
    
def day1_part1(filepath):
    load_data = read_file(filepath)
    # initialize number of safe reports
    num_safe_reports = 0
    increasing = True
    for report in load_data:
        if isReportSafe(report):
            num_safe_reports += 1
    return num_safe_reports

print("Welcome to AoC 2024 day 2!")
test_input_file = os.path.dirname(__file__) + "\\test_input.txt"
print("Test input:", test_input_file)
print("Test result:", day1_part1(test_input_file))

my_input_file = os.path.dirname(__file__) + "\\my_input.txt"
print("My input file:", my_input_file)
print("My result:", day1_part1(my_input_file))

######### PART 2 #################
## so one approach is WHEN an unsafe row is found, check if removing the element that made it unsafe OR the previous one would still make it unsafe.

def problemDampener(row, i):
    for i in range(len(row)):
        temp_row = row.copy()
        temp_row.pop(i)
        if isReportSafe(temp_row):
            return True
    return False

def recheckedIsReportSafe(row):
    if row[1] > row[0]:
        increasing = True
    elif row[1] < row[0]:
        increasing = False
    else:
        # first two elements are the same
        return problemDampener(row, 1)
    for i in range(1, len(row)):
        if not acceptable_diff(row[i], row[i-1]):
            return problemDampener(row, i)
        if increasing and row[i] <= row[i-1]:
            return problemDampener(row, i)
        elif not increasing and row[i] >= row[i-1]:
            return problemDampener(row, i)
        else:
            continue
    return True

def day1_part2(filepath):
    load_data = read_file(filepath)
    # initialize number of safe reports
    num_safe_reports = 0
    increasing = True
    for report in load_data:
        if recheckedIsReportSafe(report):
            num_safe_reports += 1
    return num_safe_reports

print("Test input:", test_input_file)
print("Test result:", day1_part2(test_input_file))
print("My result:", day1_part2(my_input_file))