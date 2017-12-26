import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.Set;

public class Main {
	public static void main(String[] args) throws Exception {
		String[] steps = Files.readAllLines(Paths.get("input.txt")).get(0).split(", ");

		Point pos = new Point(0, 0);
		Point partTwoPos = null;

		Set<Point> locations = new HashSet<>();

		for (String step : steps) {
			pos.direction = step.charAt(0) == 'L' ?
					(pos.direction + 1) % 4 :
					pos.direction - 1 < 0 ? 3 : pos.direction - 1;

			int stepsForward = Integer.parseInt(step.substring(1));
			for (int i = 0; i < stepsForward; i++) {
				pos.move();
				if (partTwoPos == null && locations.stream().anyMatch(point -> point.equals(pos))) {
					partTwoPos = new Point(pos);
				} else {
					locations.add(new Point(pos));
				}
			}
		}

		Point start = new Point(0, 0);
		System.out.printf("Part 1: %d\nPart 2: %d\n",
				manhattanDistance(start, pos),
				manhattanDistance(start, partTwoPos));
	}

	public static int manhattanDistance(Point a, Point b) {
		return Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
	}

	private static class Point {
		public int x, y;
		public int direction;

		public Point(int x, int y, int direction) {
			this.x = x;
			this.y = y;
			this.direction = direction;
		}

		public Point(int x, int y) {
			this(x, y, 0);
		}

		public Point(Point p) {
			this(p.x, p.y, p.direction);
		}

		public void move() {
			switch (direction) {
				case 0:
					x++;
					break;
				case 1:
					y++;
					break;
				case 2:
					x--;
					break;
				case 3:
					y--;
					break;
			}
		}

		@Override
		public boolean equals(Object obj) {
			return obj instanceof Point &&
					((Point) obj).x == this.x &&
					((Point) obj).y == this.y;
		}
	}
}

