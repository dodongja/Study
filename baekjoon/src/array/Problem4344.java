package array;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;

public class Problem4344 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        StringTokenizer st;

        float sum = 0;
        int[] score;
        float pass = 0;
        float[] result = new float[n];

        for(int i = 0; i < n; i++){
            st = new StringTokenizer(br.readLine(), " ");
            int num = Integer.parseInt(st.nextToken());
            score = new int[num];

            for(int j = 0; j < num; j++){
                score[j] =  Integer.parseInt(st.nextToken());
                sum += score[j];
            }

            float avg = sum/num;

            for(int k = 0; k < num; k++){
                if(avg < score[k]){
                    pass++;
                }
            }

            result[i] = Math.round(pass/num*100*1000)/1000f;
            
            sum = 0;
            pass = 0;
        }

        for(int i = 0; i < n; i++){
            System.out.printf("%.3f%%\n", result[i]);
        }

    }
}
