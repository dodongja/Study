package stackAndQueue.truckCrossingTheBridge;

import java.util.LinkedList;
import java.util.Queue;

public class Solution {
    //1초마다 전진전진
    //마지막애가 도착하면 끝이네 [7,4,5,6]
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] i = new int[]{7,4,5,6};
        int w = 10;
        int b = 2;
        System.out.println(s.solution(b,w,i));
    }
    public int solution(int bridge_length, int weight, int[] truck_weights) {
        int answer = 0;

        Queue<Integer> q = new LinkedList<>();
        Queue<Integer> r = new LinkedList<>();
        Queue<Integer> t = new LinkedList<>();
        // 100 100 10
        for(int i = 0; i < truck_weights.length; i++){
            q.add(truck_weights[i]);
        }

        for (int i = 1; ; i++) {
            int sum = 0;
        //올라간 애들 총 무게
            if(!r.isEmpty()) {
                for (int s : r) {
                    sum += s;
                }
            }
        //시간되면 나가세요~
        if(!t.isEmpty()) {
            if (t.peek() + bridge_length == i) {
                sum -= r.peek();
                r.remove();
                t.remove();

                if (r.isEmpty() && q.isEmpty()) {
                    answer = i;
                    break;
                }
            }
        }
            //대기중인애 올려주고 몇초에 올라갔는지
            if(!q.isEmpty()) {
                if (sum + q.peek() <= weight) {
                    r.add(q.poll());
                    t.add(i);
                }
            }

        }


        return answer;
    }
}
