#!/usr/bin/env python3


def increment_char(char):
    reset = False
    ascii_val = ord(char) + 1
    if ascii_val == 105 or ascii_val == 111 or ascii_val == 108:
        ascii_val += 1
    if ascii_val == 123:
        ascii_val = 97
        reset = True
    return chr(ascii_val), reset


def increment_password(password):
    index = len(password) - 1
    while True:
        char = password[index]
        new_char, again = increment_char(char)
        password = password[:index] + new_char + password[index+1:]
        index -= 1
        if not again:
            break
    return password


def check_rule1_3_increasing_letters(password):
    for i in range(len(password) - 2):
        if ord(password[i]) == ord(password[i+1]) - 1 and ord(password[i+1]) == ord(password[i+2]) - 1:
            return True
    return False


def check_rule2_2_doubles(password):
    count = 0
    index = 0
    while index < len(password) - 1:
        if password[index] == password[index+1]:
            count += 1
            index += 1
            if count >= 2:
                return True
        index += 1
    return False


def run(password):
    while True:
        password = increment_password(password)
        if check_rule1_3_increasing_letters(password) and check_rule2_2_doubles(password):
            break
    return password


password = "vzbxkghb"
result1 = run(password)
print("Result1 = ", result1)

result2 = run(result1)
print("Result2 = ", result2)
