
public class VaporCase {
    public class Solution {
        public static String vaporcode(String s) {
            return String.join("  ", s.replaceAll(" ", "").toUpperCase().split(""));
        }
    }
}
