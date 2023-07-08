import { ChoiceValue, Entry, getWinScore, runTask } from "./utils";

function computeTotalScore(entries: Entry[]): bigint {
  return entries.reduce((totalScore, entry) => {
    const [openentChoice, myChoice] = entry;
    const winValue = getWinScore(openentChoice, myChoice);
    const choiceValue = ChoiceValue[myChoice] ?? 0;
    return totalScore + BigInt(winValue) + BigInt(choiceValue);
  }, BigInt(0));
}


runTask(computeTotalScore);

