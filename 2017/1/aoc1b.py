#!/usr/bin/env python3

total = 0
with open('aoc1b.txt','r') as f:
    mys = f.read()

for i in range(len(mys)):
    if i < int(len(mys)/2):
        if mys[i] == mys[i + int(len(mys)/2)]:
            total += int(mys[i])
    else:
        if mys[i] == mys[i - int(len(mys)/2)]:
            total += int(mys[i])

print(total)
