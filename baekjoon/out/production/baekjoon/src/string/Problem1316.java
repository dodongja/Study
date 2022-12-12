package string;

import java.io.*;
import java.util.HashSet;
import java.util.Set;

public class Problem1316 {
    /** 하나씩 앞뒤로 비교하다가 바뀌면 map에 넣는데
     *  false뜨면 오류
     */
    public static void main(String[] args) throws IOException {
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        Set<Character> check = new HashSet<>();
        int count = 0;

        int n;

        while(true){
            try {
                n = Integer.parseInt(br.readLine());
                break;
            }catch (Exception e) {
                bw.write("숫자만 가능합니다");
                continue;
            }
        }

        for(int i = 0; i < n; i++) {
            char[] charArray =  br.readLine().toCharArray();

            for(int j = 0; j < charArray.length; j++) {
                 try {
                     if (charArray[j] != charArray[j + 1]) {
                         if (!check.add(charArray[j])) {
                             System.out.println("dup");
                             break;
                         } else {
                             check.add(charArray[j]);
                         }
                     }
                 } catch (IndexOutOfBoundsException e){
                     System.out.println("per");
                     count++;
                     break;
                 }

            }

        }


        bw.write(String.valueOf(count));
        bw.flush();
        bw.close();

    }
}
