import { readFile } from "node:fs/promises";

let data = await readFile("./input.txt", { encoding: "utf8" });
data = data.trim().split("\n");

let symbol_re = /\*/gi;
let part_num_re = /([0-9]+)/gi;
let parts = [];
let numbers_per_line = {};

data.forEach((line, index) => {
  let symbols = [...line.matchAll(symbol_re)].map((match) => [
    match[1],
    match.index + 1,
  ]);
  let numbers = [...line.matchAll(part_num_re)].map((match) => [
    match[1],
    match.index + 1,
  ]);
  numbers_per_line[index] = numbers;

  if (index != data.length - 1) {
    let next_line = [...data[index + 1].matchAll(part_num_re)].map((match) => [
      match[1],
      match.index + 1,
    ]);
    numbers_per_line[index + 1] = next_line;
  }
  let gear_ratio = {};
  for (let [_, symb_index] of symbols) {
    let sym_range = [symb_index - 1, symb_index, symb_index + 1];
    gear_ratio?.[`${index},${symb_index}`] ??
      (gear_ratio[`${index},${symb_index}`] = []);
    if (index != 0) {
      gear_ratio[`${index},${symb_index}`].push(
        ...numbers_per_line[index - 1]
          .filter(
            ([num, idx]) =>
              sym_range.includes(idx + num.length - 1) ||
              sym_range.includes(idx)
          )
          .map((v) => +v[0])
      );
    }
    gear_ratio[`${index},${symb_index}`].push(
      ...numbers_per_line[index]
        .filter(
          ([num, idx]) =>
            sym_range.includes(idx + num.length - 1) || sym_range.includes(idx)
        )
        .map((v) => +v[0])
    );
    if (index != data.length - 1) {
      gear_ratio[`${index},${symb_index}`].push(
        ...numbers_per_line[index + 1]
          .filter(
            ([num, idx]) =>
              sym_range.includes(idx + num.length - 1) ||
              sym_range.includes(idx)
          )
          .map((v) => +v[0])
      );
    }
  }
  for (const [_, array] of Object.entries(gear_ratio)) {
    if (array.length >= 2) {
      parts.push(array[0] * array[1]);
    }
  }
});
console.log(parts.reduce((acc, curr) => acc+curr));
