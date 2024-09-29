import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

import java.util.stream.Collectors;

public class Order {

    public static void main(String[] args) {
        String input = "4of Fo1r pe6ople g3ood th5e the2";
        System.out.println(order(input));
    }

    static class Word implements Comparable<Word> {
        public String word;
        public Integer place;

        public Word(String word, Integer place) {
            this.word = word;
            this.place = place;
        }

        @Override
        public int compareTo(Order.Word o) {
            return this.place.compareTo(o.place);
        }

        public static Word fromStr(final String from) {
            return new Word(from, getNumber(from));
        }

        @Override
        public String toString() {
            return this.word;
        }

        public String getValue() {
            return this.word;
        }
    }

    public static String order(final String words) {
        ArrayList<String> asd = new ArrayList<String>();
        for (String element : words.split(" ")) {
            asd.add(element);
        }
        List<Word> wordsC = asd.stream().map(Word::fromStr).collect(Collectors.toList());
        Collections.sort(wordsC);
        return wordsC.stream().map(Word::getValue).collect(Collectors.joining(" "));
    }

    public static Integer getNumber(final String input) {
        for (String c : input.split("")) {
            try {
                return Integer.parseInt(c);
            } catch (Exception e) {
                // return null;
            }
        }
        return null;
    }
}