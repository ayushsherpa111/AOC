const fs = require("fs/promises");

fs.readFile("./input.txt").then((buffer) => {
  const key_len = 14;
  for (let i = 0; i < buffer.length - key_len; i++) {
    let st = new Set(buffer.subarray(i, i+key_len));
    if (st.size == key_len) {
      console.log(i + key_len) ;
      break;
    }
  }
});
