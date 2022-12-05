package loop;

import java.io.*;
import java.util.StringTokenizer;

public class Problem11022 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        int t = Integer.parseInt(br.readLine());
        String[] answer = new String[t];

        StringTokenizer st;

        for(int i = 0; i < t; i++){
            st = new StringTokenizer(br.readLine()," ");
            int a = Integer.parseInt(st.nextToken());
            int b = Integer.parseInt(st.nextToken());
            answer[i] = "Case #" + (i+1) + ": " +
                    a + " + " + b + " = " +
                    (a + b);
        }

        for(int i = 0; i < t; i++){
            bw.write(answer[i] + "\n");
        }

        bw.close();
    }
}
