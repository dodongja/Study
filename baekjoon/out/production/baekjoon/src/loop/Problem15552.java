package loop;

import java.io.*;
import java.util.StringTokenizer;

public class Problem15552 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));
        int t = Integer.parseInt(br.readLine());
        int[] answer = new int[t];

        StringTokenizer st;

        for(int i = 0; i < t; i++){
            st = new StringTokenizer(br.readLine()," ");
            answer[i] = Integer.parseInt(st.nextToken()) + Integer.parseInt(st.nextToken());
        }

        for(int i = 0; i < t; i++){
            bw.write(answer[i] + "\n");
        }

        bw.close();
    }
}
