package array.hIndex;

import java.util.*;

class Solution {
    public int solution(int[] citations) {
        int answer = 0;

        int[] sortedDesc = Arrays.stream(citations)
                .boxed()
                .sorted(Collections.reverseOrder())
                .mapToInt(Integer::intValue)
                .toArray();

        int len = citations.length;

        for(int i = 0; i < len; i++){
            if(i+1 > sortedDesc[i]){
                answer = i;
                break;
            }

        }

        return answer;
    }
}
