#!/usr/bin/env python3

import itertools
import math
import os
import sys

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))

TARGET = 36000000
# TARGET = 200


def is_prime(n):
    if n == 1:
        return False
    return all(n%i != 0 for i in range(2, int(n**0.5)+1))


def get_primes(max):
    result = []
    for n in range(2, max):
        if is_prime(n):
            result.append(n)
    return result


def decompose(Number, primes):
    result = []
    if Number in primes:
        return []
    N = Number
    for i in primes:
        r = 0
        while r == 0:
            q, r = divmod(N, i)
            if r == 0:
                result.append(i)
                N = q
            if N == 1:
                return result
    if N == Number:
        return []   # prime
    if N != 1:
        print(f"WARNING: decomposition not finalized for {Number}, remaining {N}")
        sys.exit(1)
    return result


def run1_first_try(target): 
    '''
    Works but takes ages....
    '''
    primes = []
    house = 0
    best = 0
    while True:
        house += 1

        house_primes = decompose(house, primes)

        if house_primes == [] and house != 1:
            primes.append(house)

        divisors = set()
        if len(house_primes) < 8:
            continue  # skip if there is not a minimum number for divisor
        for L in range(1, len(house_primes)):
            for subset in itertools.combinations(house_primes, L):
                divisors.add(math.prod(subset))
        divisors.add(1)
        divisors.add(house)

        house_gift_count = sum(divisors) * 10
        if house_gift_count > best:
            best = house_gift_count
            print(house, house_gift_count, house_primes)

        if house_gift_count >= target:
            print(house, house_gift_count, divisors)
            return house


def run1(target, primes, current_primes=[], best=0):
    if best == 0:
        best = target
    
    if len(current_primes) > 0: 
        last = current_primes[-1]
    else:
        last = 0

    for p in primes:

        if p < last:
            continue

        next_primes = current_primes + [p]
        house = math.prod(next_primes)
        
        if house > best:
            continue

        divisors = set()
        for L in range(1, len(next_primes)):
            for subset in itertools.combinations(next_primes, L):
                divisors.add(math.prod(subset))
        divisors.add(1)
        divisors.add(house)
        gift_count = sum(divisors) * 10

        if gift_count >= target:
            if house < best:
                best = house
        else:
            best = run1(target, primes, next_primes, best)
    return best


primes = get_primes(50)
result1 = run1(TARGET, primes)
print("Result1 = ", result1)

# =========================================


def run2(target, primes, current_primes=[], best=0):
    if best == 0:
        best = target
    
    if len(current_primes) > 0: 
        last = current_primes[-1]
    else:
        last = 0

    for p in primes:

        if p < last:
            continue

        next_primes = current_primes + [p]
        house = math.prod(next_primes)
        
        if house > best:
            continue

        divisors = set()
        for L in range(1, len(next_primes)):
            for subset in itertools.combinations(next_primes, L):
                possible_divisor = math.prod(subset)
                if house/possible_divisor <= 50:
                    divisors.add(possible_divisor)
        divisors.add(house)
        gift_count = sum(divisors) * 11

        if gift_count >= target:
            if house < best:
                best = house
        else:
            best = run2(target, primes, next_primes, best)
    return best


result2 = run2(TARGET, primes)
print("Result2 = ", result2)
