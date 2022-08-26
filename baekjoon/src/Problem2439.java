import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem2439 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        int one = 1;
        String plusStar = "*";
        String star = "*";


        for(int j = 0; j < n; j++) {
            String answer = "";
            String space = " ";
            for (int i = one; i < n; i++) {
                answer += space;
            }
            answer += star;
            star += plusStar;
            one++;
            System.out.println(answer);
        }
    }
}
