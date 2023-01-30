import { createReadStream } from "node:fs";
import { createInterface } from "node:readline/promises";
import { map_to_num } from "./common.js";

const line = createInterface({
  input: createReadStream("./input.txt", { encoding: "utf-8" }),
});

let group = [];

const findBadge = (group) => {
  let mappedGroup = group.map((rucksack) => new Set(rucksack));
  let val = 0;
  for (let item of mappedGroup[0]) {
    if (mappedGroup[1].has(item) && mappedGroup[2].has(item)) {
      val = map_to_num(item);
      break;
    }
  }
  return val;
};

let sum = 0;

line.on("line", (line) => {
  group.push(line.trim());
  if (group.length >= 3) {
    // aggregate group
    sum += findBadge(group);
    // clear the array once done
    group.splice(0);
  } 
});

line.on("close", () => console.log(sum));
