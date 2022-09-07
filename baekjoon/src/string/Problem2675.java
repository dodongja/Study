package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Problem2675 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int t = Integer.parseInt(br.readLine());
        int r = 0;
        String s;
        StringBuffer answer;

        StringTokenizer st;

        for(int i = 0; i < t; i++){
            st = new StringTokenizer(br.readLine());

            r = Integer.parseInt(st.nextToken());
            s = st.nextToken();

            answer = new StringBuffer();

            for(int j = 0; j < s.length(); j++){

                for(int y = 0; y < r; y++){
                    answer.append(s.charAt(j));
                }
            }

            System.out.println(answer);
        }

    }
}
