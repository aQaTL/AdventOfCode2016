#![feature(split_ascii_whitespace)]

extern crate regex;
#[macro_use]
extern crate lazy_static;

use std::str::FromStr;

fn main() {
    let mut nodes = std::fs::read_to_string("input.txt")
        .expect("input.txt not found")
        .lines()
        .skip(2)
        .filter_map(|x| x.parse::<Node>().ok())
        .collect::<Vec<Node>>();

    println!("Part 1: {}", part_1(&nodes));
    part_2(&mut nodes);
}

fn part_1(nodes: &Vec<Node>) -> u32 {
    let mut count = 0;
    for node_a in nodes.iter() {
        if node_a.used_size == 0 {
            continue;
        }
        for node_b in nodes.iter() {
            if node_a.x != node_b.x && node_a.y != node_b.y && node_a.used_size <= node_b.avail_size
            {
                count += 1;
            }
        }
    }
    count
}

fn part_2(nodes: &mut Vec<Node>) {
    nodes.sort_by(|x, y| x.y.cmp(&y.y).then(x.x.cmp(&y.x)));
    let mut last_x = 0;
    for node in nodes.iter() {
        if node.x < last_x {
            println!();
        }
        match node.used_size {
            0 => print!("0"),
            s if s >= 100 => print!("|"),
            _ => print!("#"),
        }
        last_x = node.x;
    }
    println!("\nsolve by hand :)");
}

#[derive(Debug)]
struct Node {
    x: usize,
    y: usize,
    size: u32,
    used_size: u32,
    avail_size: u32,
    use_percentage: String,
}

impl FromStr for Node {
    type Err = (String);

    fn from_str(s: &str) -> Result<Self, <Self as FromStr>::Err> {
        let mut elems = s.split_ascii_whitespace();

        lazy_static! {
            static ref RE: regex::Regex =
                regex::Regex::new(r#"/dev/grid/node-x(\d+)-y(\d+)"#).unwrap();
        }

        let caps: regex::Captures = RE.captures(elems.next().unwrap()).unwrap();
        let x = caps.get(1).unwrap().as_str().parse::<usize>().unwrap();
        let y = caps.get(2).unwrap().as_str().parse::<usize>().unwrap();

        let size = elems.next().unwrap();
        let size = size[..(size.len() - 1)].parse::<u32>().unwrap();
        let used_size = elems.next().unwrap().to_string();
        let used_size = used_size[..(used_size.len() - 1)].parse::<u32>().unwrap();
        let avail_size = elems.next().unwrap().to_string();
        let avail_size = avail_size[..(avail_size.len() - 1)].parse::<u32>().unwrap();
        let use_percentage = elems.next().unwrap().to_string();

        Ok(Node {
            x,
            y,
            size,
            used_size,
            avail_size,
            use_percentage,
        })
    }
}
