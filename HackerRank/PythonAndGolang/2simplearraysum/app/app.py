#!/bin/python3
'''
Given an array of integers, find the sum of its elements
'''

import math
import os

import random
import re
import sys

#
# Complete the 'simpleArraySum' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY ar as parameter.
# Given an array of integers, find the sum of its elements

def simpleArraySum(ar):
	sum = 0
    # Write your code here
	for element in range(len(ar)):
		sum += ar[element]
	return (sum)

if __name__ == '__main__':

		os.environ['OUTPUT_PATH'] = 'junk.txt'	
		fptr = open(os.environ['OUTPUT_PATH'], 'w')

		ar_count = int(input().strip())

		ar = list(map(int, input().rstrip().split()))

		result = simpleArraySum(ar)

		fptr.write(str(result) + '\n')


		fptr.close()


# answer
'''
for element in range(len(ar)):
        sum += ar[element]
    return (sum)
'''