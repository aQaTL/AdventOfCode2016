fn main() {
	let mut ranges = std::fs::read_to_string("input.txt").expect("No input file found")
		.lines()
		.map(|x| {
			let mut split = x.split("-");
			(split.next().unwrap().parse::<u32>().unwrap(),
			 split.next().unwrap().parse::<u32>().unwrap())
		})
		.collect::<Vec<(u32, u32)>>();
	ranges.sort_by(|x, y| x.0.cmp(&y.0));

	println!("Part 1: {}", part_1(&ranges));
	println!("Part 2: {}", part_2(&ranges));
}

fn part_1(ranges: &Vec<(u32, u32)>) -> u32 {
	let mut addr = 0u32;
	for range in ranges {
		if addr < range.0 {
			return addr;
		}
		addr = range.1 + 1
	}
	return addr;
}

fn part_2(ranges: &Vec<(u32, u32)>) -> u32 {
	let (mut addr, mut total, mut idx) = (0u32, 0u32, 0usize);
	while addr < 0xff_ff_ff_ff {
		let range = match ranges.get(idx) {
			Some(r) => r,
			None => break,
		};

		if addr < range.0 {
			total += 1;
			addr += 1;
		} else if addr >= range.0 && addr <= range.1 {
			if range.1 == 0xff_ff_ff_ff {
				break;
			}
			addr = range.1 + 1;
			idx += 1;
		} else {
			idx += 1;
		}
	}
	return total;
}
