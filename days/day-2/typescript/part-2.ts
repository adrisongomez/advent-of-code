import { Entry, runTask } from "./utils";

const value: { [key: string]: { [key: string]: number } } = {
  A: {
    X: 3,
    Y: 4,
    Z: 8,
  },
  B: {
    X: 1,
    Y: 5,
    Z: 9,
  },
  C: {
    X: 2,
    Y: 6,
    Z: 7,
  },
};

function computeTotalScore(entries: Entry[]): bigint {
  return entries.reduce((totalScore, [oponent, mine]) => {
    if (!oponent || !mine) return totalScore;
    const score = value[oponent][mine];
    return totalScore + BigInt(score);
  }, BigInt(0));
}

runTask(computeTotalScore, { test: false });
