import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @author Maciej on 27.12.2017.
 */
public class Day4 {
	public static void main(String[] args) throws Exception {
		List<String> input = Files.readAllLines(Paths.get("input.txt"));

		int sumOfIds = input.
				stream().
				map(Room::new).
				filter(Room::isValid).
				mapToInt(Room::getId).
				sum();

		System.out.printf("Part 1: %d\n", sumOfIds);

		int northpoleId = input.
				stream().
				map(Room::new).
				filter(room -> room.
						decrypt().
						contains("northpole")).
				findFirst().
				get().
				id;

		System.out.printf("Part 2: %d\n", northpoleId);
	}

	private static class Room {
		private static Pattern p = Pattern.compile("(.+)*-(\\d+)\\[(.+)\\]");

		private String name;
		private int id;
		private String checksum;

		public Room(String text) {
			Matcher m = p.matcher(text);
			if (m.matches()) {
				name = m.group(1);
				id = Integer.parseInt(m.group(2));
				checksum = m.group(3);
			}
		}

		public boolean isValid() {
			Map<Character, Integer> letterOccurrence = new TreeMap<>();

			for (char c : name.replaceAll("-", "").toCharArray()) {
				int occurrences = letterOccurrence.getOrDefault(c, 0);
				letterOccurrence.put(c, ++occurrences);
			}

			ArrayList<Map.Entry<Character, Integer>> sorted = new ArrayList<>(letterOccurrence.entrySet());
			sorted.sort((o1, o2) -> {
				if (o1.getValue().intValue() == o2.getValue().intValue()) {
					return o1.getKey() - o2.getKey();
				}
				return o2.getValue() - o1.getValue();
			});

			StringBuffer buf = new StringBuffer(sorted.size());
			for (Map.Entry<Character, Integer> entry : sorted) {
				buf.append(entry.getKey());
			}

			return buf.toString().substring(0, 5).equals(this.checksum);
		}

		public String decrypt() {
			StringBuffer buf = new StringBuffer(name.length());
			String[] pieces = name.split("-");
			for (int i = 0; i < pieces.length; i++) {
				for (char c : pieces[i].toCharArray()) {
					c = (char) ((c - 97 + (id % 26)) % 26 + 97);
					buf.append(c);
				}
				if (i != pieces.length - 1) buf.append(' ');
			}
			return buf.toString();
		}

		public int getId() {
			return id;
		}
	}
}
