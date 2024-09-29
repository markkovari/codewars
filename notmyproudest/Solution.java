import java.util.Arrays;

public class Solution {
    public static String capitalizeFirstLetter(String str) {
        return (str == null || str.length() == 0) ? str : str.substring(0, 1).toUpperCase() + str.substring(1);
    }

    public static String toCamelCase(String s) {
        if (s.length() == 1) {
            return String.valueOf(s.charAt(0)).toUpperCase();
        }
        var first = s.split("[-_]")[0];
        return first + Arrays.stream(s.split("[-_]"))
                .skip(1)
                .reduce("", (acc, curr) -> acc + capitalizeFirstLetter(curr));
    }

    public static void main(String[] args) {
        String input = "The_Stealth_Warrior";
        System.out.println(toCamelCase(input));
    }
}
