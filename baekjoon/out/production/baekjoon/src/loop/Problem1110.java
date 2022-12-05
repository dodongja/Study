package loop;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem1110 {
    public static void main(String[] args) throws IOException {
        //2+6 = 8이다. 새로운 수는 68이다
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        int x = n;
        int count = 0;

        while(true){
            count++;

            if((n/10) == 0){
                n = n * 11;
            }else{
                n = ((n%10)*10) + ((n/10)+(n%10))%10;
            }

            if(x == n){
                break;
            }
        }

        System.out.println(count);
    }
}
