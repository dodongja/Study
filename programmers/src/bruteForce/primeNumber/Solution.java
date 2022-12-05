package bruteForce.primeNumber;

import java.util.HashSet;
import java.util.Set;

public class Solution {
    //일단 numbers로 숫자 나올 수 있는 조합 부터 찾고
    // 소수 판단해야 할듯?
    // 1 2 3
    // 12 13
    // 21 23
    // 31 32
    //123 132 213 231 321 312

    //1 2
    //12 21

    //1 2 3 4
    //12 13 14
    //23 21 24
    //31 32 34
    //41 42 43

    //123 124
    //132 134
    //143 142

    //213 214
    //231 234
    //241 243
    public int solution(String numbers) {
        int answer = 0;

        Set<String> combination = new HashSet<>();
        String tmp;

        String[] nums = new String[numbers.length()];

        for(int i = 0; i < nums.length; i++){
            nums[i] = String.valueOf(numbers.charAt(i));
        }
        int count = 1;
        while(true){

            for(int i = 0; i < numbers.length(); i++){

            }
        }



        return answer;
    }
}
