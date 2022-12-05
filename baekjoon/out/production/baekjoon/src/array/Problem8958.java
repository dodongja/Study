package array;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem8958 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        int score = 0;
        int totalScore = 0;

        for(int i = 0; i < n; i++){
            String s = br.readLine();
            for(int j = 0; j < s.length(); j++){
                if(s.charAt(j) == 'O'){
                    score++;
                    totalScore += score;
                }else{
                    score = 0;
                }
            }
            System.out.println(totalScore);
            score = 0;
            totalScore = 0;
        }
    }
}
