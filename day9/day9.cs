using System;

class Day9
{
    static void Main()
    {
        string input = System.IO.File.ReadAllText("input.txt");

        int decompressed = 0;
        for (int i = 0; i < input.Length; i++)
        {
            if (input[i] == '(')
            {
                var buf = new System.Text.StringBuilder();
                i++;
                for (;input[i] != ')';)
                {
                    buf.Append(input[i]);
                    i++;
                }

                var group = buf.ToString().Split('x');
                int length = Convert.ToInt32(group[0]);
                int repeat = Convert.ToInt32(group[1]);

                i += length;
                decompressed += length * repeat;
            }
            else
            {
                decompressed++;
            }
        }

        Console.WriteLine("Part 1: {0}", decompressed);
    }
}