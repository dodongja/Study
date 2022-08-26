import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem2438 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());

        String star = "*";
        String answer = "";
        for(int i = 0; i < n; i++){
            answer += star;
            System.out.println(answer);
        }
    }
}