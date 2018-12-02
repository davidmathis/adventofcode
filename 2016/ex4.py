#!/usr/bin/python 

import string
import re

all_letters = string.lowercase
valid, valid_sector = 0, 0

with open("4.txt","r") as f:
    mydata = [txt.strip() for txt in f.readlines()]

for line in mydata:
    letters, checksum = line.replace("]","").split("[")
    
    for letter in checksum:
        if (letters.count(letter) >= letters.count(checksum[min(4,checksum.index(letter)+1)])): 
            if (checksum.index(letter) == 4):
                valid_sector += int(','.join(re.findall('\d+',letters)))
                valid += 1
            elif checksum.index(letter) < checksum.index(checksum[checksum.index(letter)+1]):
                continue
        else:
            break

print "There are %i valid lines and sector ID is: %i" % (valid,valid_sector)
