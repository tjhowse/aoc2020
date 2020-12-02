#!/usr/bin/python3

a = []

with open('input','r') as f:
    for line in f:
        a.append(int(line))
def day1_part1(a):
    for i in a:
        for j in a:
            if i+j == 2020:
                print("i: {}, j: {}".format(i,j))
                print(i*j)
                return

def day1_part2(a):
    for i in a:
        for j in a:
            for k in a:
                if i+j+k == 2020:
                    print("i: {}, j: {}, k: {}".format(i,j,k))
                    print(i*j*k)
                    return

day1_part2(a)