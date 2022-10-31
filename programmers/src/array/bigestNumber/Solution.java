package array.bigestNumber;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Solution {
    public static void main(String[] args) {
        int[] arr = new int[]{3, 30, 34, 5, 9};
        StringBuilder answer = new StringBuilder();
        List<String> str = new ArrayList<>();

        for(int i = 0; i < arr.length; i++){

           str.add(String.valueOf(arr[i]));
        }

        Collections.sort(str, Collections.reverseOrder());



        System.out.println(answer);


    }
    public String solution(int[] numbers) {
        String answer = "";
        return answer;
    }
}
