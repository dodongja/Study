package heap.moreSpicy;
import java.util.*;

public class Solution {
    public static void main(String[] args) {
        int[] s = new int[6];
        s = new int[]{1,2,3,9,10,12};

        Solution sl = new Solution();
        int k = sl.solution(s,7);

        System.out.println(k);
    }

    public int solution(int[] scoville, int K) {
        int answer = 0;
        int mix = 0;
        //처음 첫번째랑 두번째 섞음
        //결국 배열에 섞은 두종류 없애고 새로운 종류를 넣어야 함
        //배열 사용은 어렵겠지??
        //1 2 3 9 10 12
        //5 3 9 10 12 -> 3 5 9 10 12
        PriorityQueue<Integer> pq = new PriorityQueue();

        for(int i = 0; i < scoville.length; i++){
            pq.add(scoville[i]);
        }

        while(true){
            //1개일때 k 이상일 경우
            //1개일때 k 이하일 경우 -1
            // 1,2,3,9,10,12
            if(pq.peek() >= K){
                return answer;
            }

            if(answer == scoville.length-1){
                return -1;
            }

            if(scoville.length > 1) {
                mix = pq.poll() + (pq.poll() * 2);

                answer++;

                pq.add(mix);
            }
        }

    }
}
