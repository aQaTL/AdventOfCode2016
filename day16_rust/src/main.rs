use std::thread;

fn main() {
    let input_p1 = "10011111011011001";
    let input_p2 = input_p1.clone();

    let p1 = thread::spawn(move || {
        println!("{}", part_1(&input_p1, 272));
    });
    let p2 = thread::spawn(move || {
        println!("{}", part_1(&input_p2, 35651584));
    });

    p1.join().unwrap();
    p2.join().unwrap();
}

fn part_1(input: &str, length: usize) -> String {
    let mut a: Vec<bool> = Vec::with_capacity(length * 2);
    input.chars().for_each(|x| a.push(x == '1'));

    while a.len() < length {
        a.push(false);
        for i in (0..a.len() - 1).rev() {
            let x = !a[i];
            a.push(x);
        }
    }

    a.truncate(length);
    while a.len() % 2 == 0 {
        let mut checksum_idx = 0;
        for i in (0..a.len()).step_by(2) {
            a[checksum_idx] = a[i] == a[i + 1];
            checksum_idx += 1;
        }
        let new_len = a.len() / 2;
        a.truncate(new_len);
    }

    return a
        .iter()
        .map(|x| if *x { '1' } else { '0' })
        .collect::<String>();
}
