import os

test_input = """3   4
4   3
2   5
1   3
3   9
3   3"""

print(test_input)

def read_input(input_str:str) -> list:
    list1 = []
    list2 = []
    for row in input_str.split("\n"):
        values = row.split("   ")
        a = values[0]
        b = values[1]
        list1.append(int(a))
        list2.append(int(b))
    assert len(list1) == len(list2)
    return [list1, list2]

def day1_part1(input_str:str) -> int:
    list1, list2 = read_input(input_str)
    list1.sort()
    list2.sort()
    return sum([abs(a-b) for a,b in zip(list1,list2)])

def day1_part2(input_str: str) -> int:
    list1, list2 = read_input(input_str)
    list2_counts = {}
    for elem in list2:
        if elem in list2_counts:
            continue
        list2_counts[elem] = list2.count(elem)
    running_sum = 0
    for elem in list1:
        if elem not in list2_counts:
            continue
        running_sum += elem*list2_counts[elem]
    return running_sum

print("DAY 1: PART 1")
print("Test Result:", day1_part1(test_input))
with open(os.path.dirname(__file__) + "\\MyInput.txt", "r") as myinputfile:
    data = myinputfile.readlines()
my_input = "".join(data)
print("My Result:", day1_part1(my_input))

print("DAY 1: PART 2")
print("Test Result:", day1_part2(test_input))
print("My Result:", day1_part2(my_input))