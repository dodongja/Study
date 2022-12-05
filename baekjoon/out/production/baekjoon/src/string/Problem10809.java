package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem10809 {
    public static void main(String[] args) throws IOException {
        //애초에 int 배열에 다 -1 넣어놓고
        // 있는 애들만 카운트해서 넣어주는방법이 이쑤나
        //        System.out.println(Character.getNumericValue('a')); // 10
        //        System.out.println(Character.getNumericValue('A')); // 10
        //        System.out.println(Character.getNumericValue('z')); // 35
        //        System.out.println(Character.getNumericValue('Z')); // 35

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));


        String s = br.readLine();
        int count = 0;

        //영어
        String alphabet = "abcdefghijklmnopqrstuvwxyz";

        for(int i = 0; i < alphabet.length(); i++) {
            for (int j = 0; j < s.length(); j++) {
                if (alphabet.charAt(i) == s.charAt(j)) {
                    break;
                } else {
                    count++;
                }
            }

            if(count == s.length()){
                System.out.print(-1 + " ");
            }else{
                System.out.print(count + " ");
            }

            count = 0;
        }
    }
}
