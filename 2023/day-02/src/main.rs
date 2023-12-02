use sscanf::sscanf;
use std::cmp::max;
use std::collections::HashMap;

fn main() {
    let input: &str = include_str!("../input.txt");
    assert!(part1(input) == 2447);
    assert!(part2(input) == 56322);
}

fn parse(s: String) -> Vec<Vec<HashMap<String, usize>>> {
    let mut games: Vec<Vec<HashMap<String, usize>>> = vec![];
    let lines: Vec<&str> = s.split("\n").collect();
    for line in lines {
        let l: Vec<&str> = line.split(": ").collect();
        let groups: Vec<&str> = l[1].split("; ").collect();
        let mut game: Vec<HashMap<String, usize>> = vec![];
        for group in groups {
            let colors: Vec<&str> = group.split(", ").collect();
            let mut c: HashMap<String, usize> = HashMap::new();
            for color in colors {
                let parsed = sscanf!(color, "{} {}", usize, str).unwrap();
                c.insert(parsed.1.to_string(), parsed.0);
            }
            game.push(c);
        }
        games.push(game)
    }

    return games;
}

fn is_possible(game: &Vec<HashMap<String, usize>>, cubes: HashMap<String, usize>) -> bool {
    for cube in game {
        for color in cube.keys() {
            if cubes.get(color).unwrap() < cube.get(color).unwrap() {
                return false;
            }
        }
    }

    return true;
}

fn part1(input: &str) -> i32 {
    let mut sum: i32 = 0;
    let cubes: HashMap<String, usize> = HashMap::from([
        ("red".to_string(), 12.to_owned()),
        ("green".to_string(), 13.to_owned()),
        ("blue".to_string(), 14.to_owned()),
    ]);
    let games = parse(input.to_string());

    for (i, game) in games.iter().enumerate() {
        if is_possible(game, cubes.clone()) {
            sum += i as i32 + 1;
        }
    }

    return sum;
}

fn min_cubes(game: &Vec<HashMap<String, usize>>) -> usize {
    let mut product = 1;

    let res = game.iter().fold(HashMap::new(), |mut acc, val| {
        for (color, count) in val.iter() {
            let m = max(acc.get(color).unwrap_or_else(|| &0), count);
            acc.insert(color, *m);
        }

        return acc;
    });

    for count in res.values() {
        product *= count;
    }

    return product;
}

fn part2(input: &str) -> usize {
    let games = parse(input.to_string());
    return games.iter().fold(0, |acc, val| acc + min_cubes(val));
}
