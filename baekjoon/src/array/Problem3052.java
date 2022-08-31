package array;

import java.util.HashSet;
import java.util.Scanner;

public class Problem3052 {
    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);
        HashSet<Integer> remains = new HashSet<>();

        for(int i = 0; i < 10; i++){
            int a = sc.nextInt();

            remains.add(a%42);
        }

        System.out.println(remains.size());
    }
}
