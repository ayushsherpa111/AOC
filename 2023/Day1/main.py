word_to_num = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

import re

with open("./input.txt") as input_value:
    calibration_sum = 0
    for line in input_value:
        nums_in_word = {}
        for word, num in word_to_num.items():
            for index in [match.start() for match in re.finditer(rf"{word}", line)]:
                nums_in_word[index] = word_to_num[word]
        for index, digit in filter(lambda x: x[1].isnumeric(), enumerate(line)):
            nums_in_word[index] = digit
        calibration_sum += int(nums_in_word[min(nums_in_word)]+nums_in_word[max(nums_in_word)])
    print(calibration_sum) 
    print(nums_in_word)
