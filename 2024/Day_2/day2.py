def isSafe(parsedReport):
    for i in range(0, len(parsedReport) - 1):
        if abs(reportParsed[i] - reportParsed[i + 1]) not in [1, 2, 3]:
            return False
    return True


with open("./input.txt", "r") as lab_reports:
    safe = 0
    for report in lab_reports:
        reportParsed = list(map(lambda x: int(x), report.split(" ")))
        if isSafe(reportParsed) and (
            sorted(reportParsed) == reportParsed
            or sorted(reportParsed, reverse=True) == reportParsed
        ):
            safe += 1
    print(safe)
