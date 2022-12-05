package array;

import java.io.*;
import java.util.Arrays;
import java.util.StringTokenizer;

public class Problem10818 {
    public static void main(String[] args) throws IOException {
        //속도가 엄청 느림
        // 다른 사람꺼 보니깐 InputReader이라는것과
        //Math.min(a,b)
        //Math.max(a,b)를 씀 a와 b를 비교해서 큰값을 가져오는 것이라고함 ㅎㅎ
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        int n = Integer.parseInt(br.readLine());

        int[] answer = new int[n];

        StringTokenizer st = new StringTokenizer(br.readLine()," ");
        for(int i = 0; i < n; i++){
            answer[i] = Integer.parseInt(st.nextToken());
        }
        //sort가 메모리 많이 잡아먹는듯?
        Arrays.sort(answer);

        bw.write(answer[0] + " " + answer[n-1]);

        bw.close();

    }
}
