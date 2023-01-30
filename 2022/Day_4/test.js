const fs = require("fs");

let buffer = fs.readFileSync("./input.txt").toString();
buffer = buffer
  .trimEnd()
  .split("\n")
  .map((i) => i.split(",").map((i) => i.split("-").map((i) => parseInt(i))));

// console.log(buffer[buffer.length-1]);
console.log(
  buffer.reduce((acc, curr) => {
    let set1 = new Set(
      Array.from({ length: curr[0][1] - curr[0][0] + 1 }).map(
        (_, i) => i + curr[0][0]
      )
    );
    let set2 = new Set(
      Array.from({ length: curr[1][1] - curr[1][0] + 1 }).map(
        (_, i) => i + curr[1][0]
      )
    );
    console.log(set1);
    console.log(set2);
    for (let item of set1) {
      if (set2.has(item)) {
        // console.log(item, curr);
        acc += 1;
        break;
      }
    }
    return acc;
  }, 0)
);
