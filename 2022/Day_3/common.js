export const map_to_num = (char) =>
  97 <= char.charCodeAt(0) && char.charCodeAt(0) <= 122
    ? char.charCodeAt(0) - 96
    : 27 + (char.charCodeAt(0) - 65);

