with open("./input.txt") as input_file:
    total_score = 0
    for line in input_file:
        card = list(
            map(
                lambda x: x.strip().split(),
                line[9:].split("|"),
            )
        )
        win, score = set(card[0]), set(card[1])
        wins = win.intersection(score)
        if len(wins) > 0:
            print(wins, card, total_score)
            total_score += 1
            total_score += pow(2, len(wins) - 1)
    print(total_score)
