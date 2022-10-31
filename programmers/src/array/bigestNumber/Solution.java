package array.bigestNumber;

import java.util.Arrays;
import java.util.Collections;

public class Solution {
    public static void main(String[] args) {
        int[] arr = new int[]{1,22,11,33,5,521 ,53};
        String[] str = new String[7];
        for(int i = 0; i < 7; i++){
            str[i] = String.valueOf(arr[i]);
        }
        Arrays.sort(str, Collections.reverseOrder());
        for(int i = 0; i < 7; i++){
            System.out.println(str[i]);
        }
    }
    public String solution(int[] numbers) {
        String answer = "";
        return answer;
    }
}
