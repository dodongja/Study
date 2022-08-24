import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Problem25304 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int tmp = 0;
        int X = Integer.parseInt(br.readLine());
        int N = Integer.parseInt(br.readLine());

        StringTokenizer st;

        for(int i = 0; i < N; i++) {
            st = new StringTokenizer(br.readLine(), " ");
            tmp += Integer.parseInt(st.nextToken()) * Integer.parseInt(st.nextToken());
        }

        if(X == tmp){
            System.out.println("Yes");
        }else{
            System.out.println("No");
        }
    }
}
