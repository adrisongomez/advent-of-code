import { readFile } from "node:fs/promises";

export type Entry = [string, string];

export const ChoiceValue: { [key: string]: number } = {
  X: 1,
  Y: 2,
  Z: 3,
};

export const WinCombination: { [key: string]: string | undefined } = {
  X: "C",
  Y: "A",
  Z: "B",
};

export const DrawCombination: { [key: string]: string | undefined } = {
  X: "A",
  Y: "B",
  Z: "C",
};

export function getWinScore(oponent: string, mine: string): number {
  const draw = DrawCombination[mine];
  if (draw && draw == oponent) {
    return 3;
  }
  const value = WinCombination[mine];
  if (value && value == oponent) {
    return 6;
  }
  return 0;
}

type RunTaksOptions = {
  test?: boolean;
};

export async function runTask(
  computeTotalScore: (entries: Entry[]) => bigint,
  opt: RunTaksOptions = { test: false }
) {
  let entries: Entry[];
  if (opt.test) {
    entries = [
      ["A", "Y"],
      ["B", "X"],
      ["C", "Z"],
    ];
  } else {
    const path = process.env.PROJECT_ROOT + "/days/day-2/input.txt";
    const file = await readFile(path, "utf-8");
    const lines = file.split("\n");

    entries = lines.map((line) => {
      const [openentChoice, myChoice] = line.split(" ");
      return [openentChoice, myChoice] satisfies Entry;
    });
  }

  const totalScore = computeTotalScore(entries);
  console.log(totalScore);
}
