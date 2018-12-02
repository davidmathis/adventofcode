#!/usr/bin/python

import collections

letter = ''
output = ''

with open('6.txt','r') as f:
    data = [line.rstrip() for line in f.readlines()]
for i in range(0,len(data[0])):
    for word in data:
        letter += word[i]
    output += (collections.Counter(letter).most_common()[-1][0])
    letter = ''
print output
