package loop;

import java.io.*;
import java.util.ArrayList;
import java.util.StringTokenizer;

public class Problem10952 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        StringTokenizer st;
        ArrayList<Integer> answers = new ArrayList<>();

        while(true){
            st = new StringTokenizer(br.readLine()," ");
            int answer = Integer.parseInt(st.nextToken()) + Integer.parseInt(st.nextToken());

            if(answer != 0){
                answers.add(answer);
            }else{
                break;
            }
        }

        for(int i = 0; i < answers.size(); i++){
            bw.write(answers.get(i) + "\n");
        }

        bw.close();

    }
}
