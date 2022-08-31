package array;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Problem1546 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        float max = 0;
        int[] score = new int[n];
        float tmp = 0;

        StringTokenizer st = new StringTokenizer(br.readLine());

        for(int i = 0; i < n; i++){
            score[i] = Integer.parseInt(st.nextToken());
            if(max < score[i]){
                max = score[i];
            }
        }

        for(int i = 0; i < n; i++){
            tmp += score[i]/max*100;
        }

        System.out.println(tmp/n);



    }
}
