use std::collections::{HashMap, HashSet};

pub fn find_repeat_item(rucksacks: &str) -> usize {
    let rucksacks: Vec<String> = rucksacks
        .split("\n")
        .map(|x| String::from(x.trim()))
        .collect();
    return rucksacks
        .into_iter()
        .map(handle_find_item_in_rucksack)
        .sum();
}

pub fn find_share_item(groupsacks: &str) -> usize {
    let groupsacks: Vec<String> = groupsacks
        .replace("\n\n", "\n")
        .split("\n")
        .map(|x| String::from(x))
        .collect();
    return groupsacks
        .chunks(3)
        .map(|x| x.join("\n"))
        .into_iter()
        .map(handl_share_item_in_groupsack)
        .sum();
}

fn handl_share_item_in_groupsack(groupsack: String) -> usize {
    let rucksacks: Vec<String> = groupsack
        .split("\n")
        .map(|x| String::from(x.trim()))
        .collect();
    let condition = rucksacks.len() - 1;
    let mut found: HashMap<char, usize> = HashMap::new();

    for rucksack in rucksacks.into_iter() {
        let unique_chars: HashSet<char> = HashSet::from_iter(rucksack.chars());
        for x in unique_chars {
            let counter = found.get(&x);
            if let Some(value) = counter {
                if *value == condition {
                    return get_priority_value(x);
                } else {
                    found.insert(x, value + 1);
                }
            } else {
                found.insert(x, 1);
            }
        }
    }
    return 0;
}

fn handle_find_item_in_rucksack(rucksack: String) -> usize {
    let (first_compartment, second_comparment) = rucksack.split_at(rucksack.len() / 2);
    let values = first_compartment
        .trim()
        .chars()
        .filter(|x| second_comparment.contains(&x.to_string()[..]))
        .map(get_priority_value);
    // to remove duplicates values
    let values: HashSet<usize> = HashSet::from_iter(values);
    return values.into_iter().sum();
}

fn get_priority_value(value: char) -> usize {
    let value = value as u32;
    return match value {
        66..=90 => (value - 38) as usize,
        97..=122 => (value - 96) as usize,
        _ => 0,
    };
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn day_3_part_1() {
        let input = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
        assert_eq!(find_repeat_item(input), 157);
    }

    #[test]
    fn day_3_part_2() {
        let input = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg";
        assert_eq!(find_share_item(input), 18);
    }
}
