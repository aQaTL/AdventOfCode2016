fn main() {
    let elves = 3004953;
    // let elves = 5;

    println!("{}", part_1(elves));
    println!("{}", part_2(elves));
}

fn part_1(elves: usize) -> usize {
    let mut circle: Vec<u8> = std::iter::repeat(1u8).take(elves).collect();

    let mut idx = 0;
    for _ in 0..elves - 1 {
        while circle[idx] == 0 {
            idx = (idx + 1) % elves;
        }
        idx = (idx + 1) % elves;
        while circle[idx] == 0 {
            idx = (idx + 1) % elves;
        }
        circle[idx] = 0;
    }

    return circle.iter().position(|x| *x != 0).unwrap() + 1;
}

fn part_2(elves: usize) -> usize {
    0
}
