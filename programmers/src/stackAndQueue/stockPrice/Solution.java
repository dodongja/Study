package stackAndQueue.stockPrice;

import java.util.LinkedList;
import java.util.Queue;

public class Solution {
    public int[] solution(int[] prices) {
        int[] answer = {};
        answer = new int[prices.length];
        int idx = 0;

        Queue<Integer> q = new LinkedList<>();

        for(int i = 0; i < prices.length; i++){
            q.add(prices[i]);
        }

        while(!q.isEmpty()){
            int a = q.poll();
            int count = 0;

            for(int i : q){
                if(a > i){
                    break;
                }else{
                    count++;
                }
            }

            answer[idx] = count;
            idx++;

        }

        return answer;
    }
}
