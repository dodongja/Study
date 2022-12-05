import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Problem3003 {
    public static void main(String[] args) throws IOException {

        int[] peace = new int[]{1, 1, 2, 2, 2, 8};
        double[] answer = new double[6];
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        StringTokenizer st;

        st = new StringTokenizer(br.readLine()," ");

        for(int i = 0; i < peace.length; i++){
            if(i != 5) {
                System.out.print(Math.round(peace[i] - Double.parseDouble(st.nextToken())) + " ");
            }else{
                System.out.print(Math.round(peace[i] - Double.parseDouble(st.nextToken())));
            }
        }
    }
}
