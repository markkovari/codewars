import java.util.*;
import java.util.stream.Collectors;

public class WeightSort {

    public static class Weight implements Comparable<Weight> {
        public String value;
        public int added;

        public Weight(final String input) {
            this.value = input;
            this.added = Arrays.stream(input.split(""))
                    .map(Integer::parseInt)
                    .reduce(0, (acc, curr) -> acc + curr);
        }

        public String getValue() {
            return this.value;
        }

        @Override
        public int compareTo(Weight o) {
            return (this.added == o.added) ? this.value.compareTo(o.value) : this.added - o.added;
        }

    }

    public static String orderWeight(final String strng) {
        if (strng.isEmpty()) {
            return "";
        }
        List<Weight> weights = Arrays.stream(strng.split(" "))
                .map(WeightSort.Weight::new)
                .collect(Collectors.toList());
        Collections.sort(weights);
        return String.join(" ", weights.stream().map(WeightSort.Weight::getValue).toList());
    }

    public static void main(String[] args) {
        System.out.println(WeightSort.orderWeight("56 65 74 100 99 68 86 180 90"));
    }
}