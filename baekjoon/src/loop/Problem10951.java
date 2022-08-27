package loop;

import java.io.*;
import java.util.Scanner;
import java.util.StringTokenizer;

public class Problem10951 {
    public static void main(String[] args) throws IOException {
       Scanner sc = new Scanner(System.in);
       int[] answer = new int[5];

        int a;
        int b;

        for(int i = 0; i < 5; i++){
            a = sc.nextInt();
            b = sc.nextInt();

            answer[i] = a + b;
        }

        for(int i = 0; i < 5; i++){
            System.out.println(answer[i]);
        }


    }
}
