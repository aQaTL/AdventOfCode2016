import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

/**
 * @author Maciej on 26.12.2017.
 */
public class Day2 {
	public static void main(String[] args) throws Exception {
		List<String> lines = Files.readAllLines(Paths.get("input.txt"));

		int[][] keypad = {
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
		};

		String code = "";
		Point p = new Point(1, 1);

		for (String steps : lines) {
			for (byte b : steps.getBytes()) {
				switch (b) {
					case 'U':
						p.y = p.y - 1 < 0 ? 0 : p.y - 1;
						break;
					case 'D':
						p.y = p.y + 1 > 2 ? 2 : p.y + 1;
						break;
					case 'R':
						p.x = p.x + 1 > 2 ? 2 : p.x + 1;
						break;
					case 'L':
						p.x = p.x - 1 < 0 ? 0 : p.x - 1;
						break;
				}
			}
			code += keypad[p.y][p.x];
		}
		System.out.printf("Part 1: %s\n", code);

		char[][] wtfKeypad = {
				{0, 0, '1', 0, 0},
				{0, '2', '3', '4', 0},
				{'5', '6', '7', '8', '9'},
				{0, 'A', 'B', 'C', 0},
				{0, 0, 'D', 0, 0},
		};

		code = "";
		p = new Point(0, 2);

		for (String steps : lines) {
			for (byte b : steps.getBytes()) {
				switch (b) {
					case 'U':
						if (p.y > 0 && wtfKeypad[p.y-1][p.x] != 0) {
							p.y--;
						}
						break;
					case 'D':
						if (p.y < 4 && wtfKeypad[p.y+1][p.x] != 0) {
							p.y++;
						}
						break;
					case 'R':
						if (p.x < 4 && wtfKeypad[p.y][p.x+1] != 0) {
							p.x++;
						}
						break;
					case 'L':
						if (p.x > 0 && wtfKeypad[p.y][p.x-1] != 0) {
							p.x--;
						}
						break;
				}
			}
			code += wtfKeypad[p.y][p.x];
		}
		System.out.printf("Part 2: %s\n", code);
	}

	private static class Point {
		public int x, y;

		public Point(int x, int y) {
			this.x = x;
			this.y = y;
		}
	}
}
