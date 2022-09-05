package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Scanner;
import java.util.StringTokenizer;

public class Problem11720 {
    public static void main(String[] args) throws IOException {
        //buffer 동기화
        //stirng 불변
        //buffer builder 변경가능
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int sum = 0;
        int n = Integer.parseInt(br.readLine());

        String s = br.readLine();

        if(s.length() == n) {

            for (int j = 0; j < n; j++) {
                sum += Character.getNumericValue(s.charAt(j));
            }


            System.out.println(sum);

        }


    }
}
