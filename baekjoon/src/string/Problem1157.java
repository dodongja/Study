package string;

import java.io.*;
import java.util.HashMap;
import java.util.Map;

public class Problem1157 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        String s = br.readLine().toUpperCase();
        char answer = ' ';
        int checkNum = 0;

        Map<Character, Integer> answers = new HashMap<>();

        for(int i = 0; i < s.length(); i++){
            if(answers.get(s.charAt(i)) == null){
                answers.put(s.charAt(i), 1);
            }else{
                answers.put(s.charAt(i), answers.get(s.charAt(i)) + 1);
            }
        }

        for(Map.Entry<Character, Integer> a: answers.entrySet()){
            if(a.getValue() > checkNum){
                answer = a.getKey();
                checkNum = a.getValue();
            }else if(a.getValue() == checkNum){
                answer = '?';
            }
        }
        
        bw.write(answer);
        bw.close();



    }

}
