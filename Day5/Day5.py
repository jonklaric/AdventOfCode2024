from copy import copy
import numpy as np
import os


def read_input(input_file_location:str):
    with open(input_file_location, "r") as file:
        input_str = file.read()
    rules = {}
    manuals = []
    for row in input_str.split("\n"):
        if "|" in row:
            rule_page_preceeds, rule_page_follows = [int(x) for x in row.split("|")]
            if rule_page_preceeds in rules:
                rules[rule_page_preceeds].append(rule_page_follows)
            else:                
                rules[rule_page_preceeds] = [rule_page_follows]
        elif "," in row:        
            manuals.append({int(x):i for i, x in enumerate(row.split(","))})
        else:
            continue
    return rules, manuals

print("Welcome to day 5 part 1")
rules, manuals = read_input(os.path.join(os.path.dirname(__file__), "my_input.txt"))
# print(rules)
# print(manuals)
passing_manuals = []
middle_page_sum = 0
for manual in manuals:
    manual_passes = True
    for page in manual.keys():
        if page in rules:
            for rule_page in rules[page]:
                if rule_page not in manual:
                    continue
                if manual[page] > manual[rule_page]:
                    manual_passes = False
                    print(manual, "fails due to rule", "{}|{}".format(page, rule_page))
                    break
        if not manual_passes:
            break
    if manual_passes:
        passing_manuals.append(manual)
        if len(manual)%2==0:
            print(len(manual))
        #print(manual[int(len(manual)/2)])
        middle_page_sum += list(manual.keys())[int(len(manual)/2)]
#print(passing_manuals)
print(middle_page_sum)

print("Welcome to day 5 part 2")
rules, manuals = read_input(os.path.join(os.path.dirname(__file__), "tet_input.txt"))
# print(rules)
# print(manuals)
passing_manuals = []
middle_page_sum = 0
for manual in manuals:
    manual_passes = True
    i = 0
    while i < len(manual):
        page = list(manual.keys())[i]
        if page in rules:
            for rule_page in rules[page]:
                if rule_page not in manual:
                    continue
                if manual[page] > manual[rule_page]:
                    manual_passes = False
                    print(manual, "fails due to rule", "{}|{}.".format(page, rule_page),"Switching...")
                    manual[page], manual[rule_page] = manual[rule_page], manual[page]
                    i = 0
                    manual = dict(sorted(manual.items(), key=lambda item: item[1]))
                    break
        i += 1
    if not manual_passes:
        final_manual = dict(sorted(manual.items(), key=lambda item: item[1]))
        print(final_manual)
        page_num = [x for x in manual.keys() if manual[x] == int(len(manual)/2)][0]
        #middle_page_num = final_manual
        print("middle page num", page_num)
        middle_page_sum += page_num
print(middle_page_sum)