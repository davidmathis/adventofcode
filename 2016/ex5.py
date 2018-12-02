#!/usr/bin/python

import md5
import re

data = "wtnhxymk"
pw = {i:"z" for i in range(8)}
i = 0

while sum(z != "z" for z in pw.values()) < 8:
    while md5.new(data + str(i)).hexdigest()[0:5] != "00000":
        i += 1
    pw_hash = (md5.new(data + str(i)).hexdigest()[5:7])
    if re.match("[0-7]",pw_hash[0]):
        if pw[int(pw_hash[0])] == "z":
            pw[int(pw_hash[0])] = pw_hash[1]
    i += 1

print ''.join(pw.values())
