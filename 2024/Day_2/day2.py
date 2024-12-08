def isSafe(parsedReport):
    for i in range(0, len(parsedReport) - 1):
        if abs(parsedReport[i] - parsedReport[i + 1]) not in [1, 2, 3]:
            return i
    return -1


def isAscending(inputList):
    for index in range(len(inputList) - 1):
        if inputList[index] <= inputList[index + 1]:
            continue
        else:
            return index


def isDescending(inputList):
    for index in range(len(inputList) - 1):
        if inputList[index] >= inputList[index + 1]:
            continue
        else:
            return index


with open("./input.txt", "r") as lab_reports:
    safe = 0
    for report in lab_reports:
        reportParsed = list(map(lambda x: int(x), report.split(" ")))

        only_dec = True
        only_inc = True
        is_diff = True
        has_skipped = False
        for i in range(0, len(reportParsed) - 1):
            diff = reportParsed[i] - reportParsed[i + 1]

            if not(has_skipped) and only_inc and diff < 0:
                has_skipped = True
                continue

            if not(has_skipped) and only_dec and diff > 0:
                has_skipped = True
                continue

            if diff > 0:
                only_dec = False

            if diff < 0:
                only_inc = False


            if abs(diff) not in [1, 2, 3]:
                if not(has_skipped):
                    has_skipped = True
                    continue
                is_diff = False
                break

        if is_diff and (only_inc or only_dec):
            safe += 1
        
    print(safe)
