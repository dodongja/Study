package array;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem5597 {
    public static void main(String[] args) throws IOException {
        boolean[] arr = new boolean[30];

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        for(int i = 0; i < 28; i++){
            arr[Integer.parseInt(br.readLine()) - 1] = true;
        }

        for(int i = 0; i < arr.length; i++){
            if(arr[i] == false){
                System.out.println(i+1);
            }
        }
    }
}
