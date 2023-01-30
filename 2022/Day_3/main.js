import { createReadStream } from "node:fs";
import * as readline from "node:readline/promises";
import { map_to_num } from "./common.js";

const rl = readline.createInterface({
  input: createReadStream("./input.txt", { encoding: "utf8" }),
});

let total_score = 0;

const rucksack_sort = (items) => {
  items = items.split("").map(map_to_num);
  let [first_half, second_half] = [
    items.slice(0, items.length / 2),
    items.slice(items.length / 2),
  ];
  let first_set = new Set(first_half);
  let second_set = new Set(second_half);
  second_set.forEach((val) => {
    total_score += first_set.has(val) ? val : 0;
  });
};

rl.on("line", rucksack_sort);
rl.on("close", () => {
  console.log(total_score);
});
