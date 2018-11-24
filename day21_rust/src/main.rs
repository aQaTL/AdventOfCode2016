use std::str::FromStr;

fn main() {
	let steps = std::fs::read_to_string("input.txt").
		expect("input.txt not found").
		lines().
		filter_map(|x| x.parse::<Op>().ok()).
		collect::<Vec<Op>>();

	println!("Part 1: {}", scramble(&steps, "abcdefgh"));
	println!("Part 2: {}", unscramble(&steps, "fbgdceah"));
}

#[derive(Debug)]
enum Op {
	SwapPos(usize, usize),
	SwapLetter(char, char),
	RotateBasedOn(char),
	RotateLeft(usize),
	RotateRight(usize),
	ReversePos(usize, usize),
	MovePos(usize, usize),
}

impl FromStr for Op {
	type Err = String;

	fn from_str(s: &str) -> Result<Self, <Self as FromStr>::Err> {
		let chars = s.chars().collect::<Vec<char>>();
		if s.starts_with("swap position") {
			Ok(Op::SwapPos(chars[14].to_digit(10).unwrap() as usize,
						   chars[30].to_digit(10).unwrap() as usize))
		} else if s.starts_with("swap letter") {
			Ok(Op::SwapLetter(chars[12], chars[26]))
		} else if s.starts_with("rotate based on") {
			Ok(Op::RotateBasedOn(chars[35]))
		} else if s.starts_with("rotate left") {
			Ok(Op::RotateLeft(chars[12].to_digit(10).unwrap() as usize))
		} else if s.starts_with("rotate right") {
			Ok(Op::RotateRight(chars[13].to_digit(10).unwrap() as usize))
		} else if s.starts_with("reverse") {
			Ok(Op::ReversePos(chars[18].to_digit(10).unwrap() as usize,
							  chars[28].to_digit(10).unwrap() as usize))
		} else if s.starts_with("move") {
			Ok(Op::MovePos(chars[14].to_digit(10).unwrap() as usize,
						   chars[28].to_digit(10).unwrap() as usize))
		} else {
			Err(format!("Unknown op: {}", s))
		}
	}
}

fn scramble(steps: &Vec<Op>, password: &str) -> String {
	let mut password: Vec<char> = password.chars().collect();

	for step in steps {
		match step {
			Op::SwapPos(x, y) => password.swap(*x, *y),
			Op::SwapLetter(x, y) => {
				let idx_x = password.iter().enumerate().find(|e| *e.1 == *x).unwrap().0;
				let idx_y = password.iter().enumerate().find(|e| *e.1 == *y).unwrap().0;
				password.swap(idx_x, idx_y);
			}
			Op::RotateBasedOn(x) => {
				let mut idx_x = password.iter().enumerate().find(|e| *e.1 == *x).unwrap().0;
				if idx_x >= 4 {
					idx_x += 1;
				}
				let len = password.len();
				password.rotate_right((idx_x + 1) % len);
			}
			Op::RotateLeft(x) => password.rotate_left(*x),
			Op::RotateRight(x) => password.rotate_right(*x),
			Op::ReversePos(x, y) => (*x..(*x + ((*y - *x + 1) / 2))).enumerate().
				for_each(|(i, idx)| password.swap(idx, *y - i)),
			Op::MovePos(x, y) => {
				let l = password.remove(*x);
				password.insert(*y, l)
			}
		}
	}

	password.iter().collect::<String>()
}

fn unscramble(steps: &Vec<Op>, password: &str) -> String {
	let mut password: Vec<char> = password.chars().collect();

	for step in steps.iter().rev() {
		match step {
			Op::SwapPos(x, y) => password.swap(*x, *y),
			Op::SwapLetter(x, y) => {
				let idx_x = password.iter().enumerate().find(|e| *e.1 == *x).unwrap().0;
				let idx_y = password.iter().enumerate().find(|e| *e.1 == *y).unwrap().0;
				password.swap(idx_x, idx_y);
			}
			Op::RotateBasedOn(x) => {
				let rotate_based_on = |password: &mut Vec<char>, x: char| {
					let mut idx_x = password.iter().enumerate().find(|e| *e.1 == x).unwrap().0;
					if idx_x >= 4 {
						idx_x += 1;
					}
					let len = password.len();
					password.rotate_right((idx_x + 1) % len);
				};
				let mut unscrambled = password.clone();
				let mut rescrambled = password.clone();
				loop {
					unscrambled.rotate_left(1);
					rescrambled.copy_from_slice(&unscrambled);
					rotate_based_on(&mut rescrambled, *x);
					if rescrambled == password {
						password = unscrambled;
						break;
					}
				}
			}
			Op::RotateLeft(x) => password.rotate_right(*x),
			Op::RotateRight(x) => password.rotate_left(*x),
			Op::ReversePos(x, y) => (*x..(*x + ((*y - *x + 1) / 2))).enumerate().
				for_each(|(i, idx)| password.swap(idx, *y - i)),
			Op::MovePos(x, y) => {
				let l = password.remove(*y);
				password.insert(*x, l);
			}
		}
	}

	password.iter().collect::<String>()
}
