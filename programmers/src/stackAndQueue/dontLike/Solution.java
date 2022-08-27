package stackAndQueue.dontLike;

import java.util.ArrayList;

public class Solution {
    public int[] solution(int []arr) {
        int[] answer = {};
        ArrayList<Integer> array = new ArrayList<>();

        array.add(arr[0]);

        for(int i = 1; i < arr.length; i++){
            if(arr[i] != array.get(array.size()-1)){
                array.add(arr[i]);
            }
        }
        answer = new int[array.size()];

        for(int i = 0; i < array.size(); i++){
            answer[i] = array.get(i);
        }

        return answer;
    }

}
