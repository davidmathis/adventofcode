#!/usr/bin/python 

with open("3.txt","r") as f:
    mydata = [txt.strip() for txt in f.readlines()]

found = 0

for row in mydata:
    data = " ".join(row.split()).replace(" ",",").split(",")
    values = [int(i) for i in data]
    values.sort()
    if (values[0]+min(values[1],values[2])) > max(values[1],values[2]):
        found += 1

print "Triangles Found: %i" % found
