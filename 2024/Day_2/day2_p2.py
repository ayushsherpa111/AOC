def issafe(elements):
    # check asc or dsc
    diff = [x - y for x, y in zip(elements, elements[1:])]
    is_aesc = all(elements[i] < elements[i + 1] for i in range(0, len(elements) - 1))
    is_desc = all(elements[i] > elements[i + 1] for i in range(0, len(elements) - 1))
    is_in_range = all(1 <= abs(i) <= 3 for i in diff)
    return is_in_range and (is_aesc or is_desc)


with open("./input.txt") as file:
    safe = 0
    for line in file:
        inp_parsed = list(map(int, line.strip().split(" ")))

        if issafe(inp_parsed):
            safe += 1
        else:
            for i in range(len(inp_parsed)):
                if issafe([*inp_parsed[:i], *inp_parsed[i + 1 :]]):
                    safe += 1
                    break
    print(safe)
