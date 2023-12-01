use std::collections::HashMap;

fn main() {
    let input: &str = include_str!("./input.txt");
    println!("{:?}", part1(input));
    println!("{:?}", part2(input));
}

fn part1(input: &str) -> i32 {
    let mut sum: i32 = 0;
    let lines: Vec<&str> = input.split("\n").collect::<Vec<&str>>();
    for i in 0..lines.len() {
        for c in lines[i].chars() {
            if c.is_numeric() {
                let mut a = 0;
                a += 10 * c.to_digit(10).unwrap() as i32;
                sum += a;
                break;
            } else {
                continue;
            }
        }

        for c in lines[i].chars().rev() {
            if c.is_numeric() {
                let mut b: i32 = 0;
                b += c.to_digit(10).unwrap() as i32;
                sum += b;
                break;
            } else {
                continue;
            }
        }
    }

    return sum;
}

fn part2(input: &str) -> i32 {
    let map: HashMap<&str, i32> = HashMap::from([
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
    ]);

    let mut sum: i32 = 0;
    let lines: Vec<&str> = input.split("\n").collect::<Vec<&str>>();
    for i in 0..lines.len() {
        let mut a: i32 = 0;
        let mut b: i32 = 0;

        // iterate forwards
        'outer: for (k, c) in lines[i].chars().enumerate() {
            if c.is_numeric() {
                a += 10 * c.to_digit(10).unwrap() as i32;
                sum += a;
                break;
            } else {
                let mut word: String = "".to_owned();
                for j in k..lines[i].len() {
                    let chr = lines[i].as_bytes()[j] as char;
                    if chr.is_alphabetic() {
                        let next_char: &str = &chr.to_string();
                        word.push_str(next_char);
                        if map.contains_key(word.as_str()) {
                            let val = map.get(word.as_str()).unwrap().to_owned();
                            a = 10 * val;
                            sum += a;
                            break 'outer;
                        }
                    }
                }
            }
        }
        // iterate backwards
        'outer: for (k, c) in lines[i].chars().rev().enumerate() {
            if c.is_numeric() {
                b += c.to_digit(10).unwrap() as i32;
                sum += b;
                break;
            } else {
                let mut word: String = "".to_owned();
                for j in k..lines[i].len() {
                    let chr = lines[i].chars().rev().collect::<String>().as_bytes()[j] as char;
                    if chr.is_alphabetic() {
                        let next_char: &str = &chr.to_string();
                        word.push_str(next_char);
                        if map.contains_key(word.chars().rev().collect::<String>().as_str()) {
                            let val = map
                                .get(word.chars().rev().collect::<String>().as_str())
                                .unwrap()
                                .to_owned();
                            b = val;
                            sum += b;
                            break 'outer;
                        }
                    }
                }
            }
        }
    }

    return sum;
}
