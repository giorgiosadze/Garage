#!/usr/bin/python
import sys 
import hashlib

string = sys.argv[1]
for i in range(10000000):
    print(i,"->",hashlib.md5(f"iwrupvqb{i}".encode('utf-8')).hexdigest())
