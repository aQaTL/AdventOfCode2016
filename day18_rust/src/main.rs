fn main() {
    let input = "^.....^.^^^^^.^..^^.^.......^^..^^^..^^^^..^.^^.^.^....^^...^^.^^.^...^^.^^^^..^^.....^.^...^.^.^^.^";
    println!("{}", solve(input, 40));
    println!("{}", solve(input, 400000));
}

fn solve(input: &str, rows: u64) -> u64 {
    //true -> safe, false -> trap
    let mut row: Vec<bool> = input.chars().map(|x| x == '.').collect();
    let mut safe_tiles = row.iter().filter(|x| **x).count() as u64;
    let mut next_row: Vec<bool> = Vec::with_capacity(row.len());

    for _ in 0..rows - 1 {
        next_row.clear();
        for i in 0..row.len() {
            let prev = if i == 0 {
                true
            } else {
                *row.get(i - 1).unwrap_or(&true)
            };
            let next = *row.get(i + 1).unwrap_or(&true);

            let tile_safe = prev == next;
            if tile_safe {
                safe_tiles += 1;
            }
            next_row.push(tile_safe);
        }
        row.copy_from_slice(&next_row[..]);
    }

    return safe_tiles;
}
