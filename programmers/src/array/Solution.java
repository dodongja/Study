package array;

import java.util.Arrays;

public class Solution {
    public int[] solution(int[] array, int[][] commands) {
        int[] answer = {};
        answer = new int[commands.length];
        int answerIdx = 0;
            //[1, 5, 2, 6, 3, 7, 4]	[[2, 5, 3], [4, 4, 1], [1, 7, 3]]	[5, 6, 3]
        for(int i = 0; i < commands.length; i++){
            int length = commands[i][1] -  commands[i][0] + 1;
            int[] arr = new int[length];
            int idx = 0;
            for(int j = commands[i][0] - 1; j < commands[i][1]; j++) {
                arr[idx] = array[j];
                idx++;
            }
            System.out.println(arr);
            Arrays.sort(arr);
            answer[answerIdx] = arr[commands[i][2] - 1];
            answerIdx++;
        }

            return answer;
    }
}
