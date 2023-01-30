selection_score = {
    "X": "loss",
    "Y": "draw",
    "Z": "win",
}

rock_paper_scissor_score = {
    "A": {"loss": 3, "win": 2, "draw": 1},
    "B": {"loss": 1, "win": 3, "draw": 2},
    "C": {"loss": 2, "win": 1, "draw": 3},
}


outcome_score = {"win": 6, "draw": 3, "loss": 0}

score = 0


def conditions(cpu, outcome):
    return (
        rock_paper_scissor_score[cpu][selection_score[outcome]]
        + outcome_score[selection_score[outcome]]
    )


with open("./input.txt") as file:
    for line in file:
        [cpu, outcome] = line.strip().split(" ")
        print(cpu, outcome, selection_score[outcome], conditions(cpu, outcome))
        score += conditions(cpu, outcome)
print(score)
