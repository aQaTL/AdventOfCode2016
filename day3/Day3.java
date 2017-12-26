import sun.security.krb5.internal.PAData;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;

/**
 * @author Maciej on 26.12.2017.
 */
public class Day3 {
	public static void main(String[] args) throws Exception {
		String input = new String(Files.readAllBytes(Paths.get("input.txt")));

		long count = Arrays.stream(input.split("\n")).
				map(s -> s.trim().split(" +")).
				map(values -> Arrays.stream(values).mapToInt(Integer::parseInt).toArray()).
				filter(sides ->
						sides[0] + sides[1] > sides[2] &&
								sides[0] + sides[2] > sides[1] &&
								sides[1] + sides[2] > sides[0]).
				count();

		System.out.printf("Part 1: %d\n", count);

		int[] sides = Arrays.stream(
				input.
						trim().
						replaceAll(" +", " ").
						replaceAll("\n", "").
						split(" ")).
				mapToInt(Integer::parseInt).toArray();

		count = 0;
		for (int i = 0; i < sides.length; i += 9) {
			for (int j = 0; j < 3; j++) {
				int a = sides[i + j];
				int b = sides[i + j + 3];
				int c = sides[i + j + 6];
				if (a + b > c && a + c > b && b + c > a) {
					count++;
				}
			}

		}

		System.out.printf("Part 2: %d\n", count);
	}
}
