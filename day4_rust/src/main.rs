extern crate regex;
#[macro_use]
extern crate lazy_static;

use std::cmp::Ordering;
use std::collections::HashMap;
use std::env;
use std::fs;
use std::str::FromStr;

fn main() {
    let filename = env::args_os()
        .nth(1)
        .expect("No input file")
        .into_string()
        .unwrap();
    let input = fs::read_to_string(&filename).expect(&format!("Failed to read {}", &filename));

    let valid_rooms: Vec<Room> = input
        .split('\n')
        .filter(|s| !s.is_empty())
        .filter_map(|s| s.parse::<Room>().ok())
        .filter(|x| x.is_valid())
        .collect();

    let part_1: u32 = valid_rooms.iter().map(|x| x.sector_id).sum();
    println!("Part 1: {}", part_1);

    let part_2 = valid_rooms
        .iter()
        .find(|x| x.decrypt().contains("northpole"))
        .expect("Failed to find northpole");
    println!("Part 2: {}", part_2.sector_id);
}

struct Room {
    letters: String,
    sector_id: u32,
    checksum: String,
}

impl Room {
    fn is_valid(&self) -> bool {
        let mut letters = HashMap::new();
        for c in self.letters.chars() {
            if c != '-' {
                *letters.entry(c).or_insert(0) += 1;
            }
        }

        let mut char_occurences = Vec::with_capacity(letters.len());
        for (k, v) in letters.iter() {
            char_occurences.push((k, v));
        }

        char_occurences.sort_by(|a, b| {
            let cmp = b.1.cmp(a.1);
            if cmp == Ordering::Equal {
                return a.0.cmp(b.0);
            }
            return cmp;
        });

        self.checksum == char_occurences
            .iter()
            .take(5)
            .map(|x| x.0)
            .collect::<String>()
    }

    fn decrypt(&self) -> String {
        let mut name = self.letters.clone().into_bytes();
        for b in name.iter_mut() {
            if *b == '-' as u8 {
                *b = ' ' as u8;
            } else {
                *b = (*b - 97u8 + (self.sector_id % 26) as u8) % 26 + 97u8;
            }
        }
        return String::from_utf8(name).unwrap();
    }
}

impl FromStr for Room {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        use regex::Regex;
        lazy_static! {
            static ref RE: Regex = Regex::new(r"(.+)*-(\d+)\[(.+)\]").unwrap();
        }
        let groups = RE.captures(s);
        let groups = groups.unwrap();
        let letters = groups.get(1).unwrap().as_str().to_string();
        let sector: u32 = groups.get(2).unwrap().as_str().parse().unwrap();
        let checksum = groups.get(3).unwrap().as_str().to_string();

        Ok(Room {
            letters: letters,
            sector_id: sector,
            checksum: checksum,
        })
    }
}
