#!/usr/bin/python

grid = [[1,2,3,],[4,5,6],[7,8,9]]
x,y = 1,1

def findNum(data,x,y):
    for i in range(0,len(data)):
        if data[i] in ("U","D"):
            if data[i] == "U":
                x -= 1
            else:
                x += 1
        else: 
            if data[i] == "L":
                y -= 1
            else:
                y += 1
        if y > 2:
            y = 2
        elif y < 0:
            y = 0
        if x > 2:
            x = 2
        elif x < 0:
            x = 0
    return x,y

with open("2a.txt","r") as f:
    mydata = [txt.strip() for txt in f.readlines()]
print findNum(mydata[0],x,y)
x, y = findNum(mydata[0],x,y)
print findNum(mydata[1],x,y)
x, y = findNum(mydata[1],x,y)
print findNum(mydata[2],x,y)
x, y = findNum(mydata[2],x,y)
print findNum(mydata[3],x,y)
x, y = findNum(mydata[3],x,y)
print findNum(mydata[4],x,y)
#print [grid[findNum(mydata[i])[1]][findNum(mydata[i])[0]] for i in range(len(mydata))]

