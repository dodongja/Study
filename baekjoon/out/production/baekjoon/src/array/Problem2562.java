package array;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem2562 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int maxNum = 0;
        int order = 0;

        for(int i = 0; i < 9; i++){
            int num = Integer.parseInt(br.readLine());

            if(maxNum < num){
                maxNum = num;
                order = i + 1;
            }
        }

        System.out.println(maxNum);
        System.out.println(order);

    }
}
