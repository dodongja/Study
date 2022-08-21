package stackAndQueue.functionDevelopment;

import java.util.ArrayList;
import java.util.List;

public class Solution {
    // 진도율 + 진도속도 * x >= 100일때 배포 해야함
    //첫번째 진도 100이 되었을때 뒤에 애들도 100이면 같이 배포
    // 최대 x는 100

    public int[] solution(int[] progresses, int[] speeds) {
        int[] answer = {};


        //7 3 9
        //5 10 1 1 20 1
        //뒤에꺼랑 비교해서 뒤에께 작으면 + 1 아니면 넘어간다
        //고로 작업완료 되는 기간 구해서 비교하면 되나?
        final int MAX_DAYS = 99;
        int workload = progresses.length;
        int[] complete = new int[workload];


        for (int i = 0; i < workload; i++) {
            for (int j = 1; j < MAX_DAYS; j++) {

                if ((progresses[i] + speeds[i] * j) >= 100) {
                    complete[i] = j;
                    break;
                }
            }
        }

        System.out.println(complete);

        int idx = 0;
        int distribute = 1;
        List<Integer> solution = new ArrayList<>();

       for(int i = 1; i < complete.length; i++) {
           if (complete[i] <= complete[idx]) {
               distribute++;
           } else {
               solution.add(distribute);
               idx = i;
               distribute = 1;
           }

           if (i == complete.length - 1) {
               solution.add(distribute);
           }
       }

       answer = new int[solution.size()];

       for(int i = 0; i < solution.size(); i++){
           answer[i] = solution.get(i);
       }


        return answer;
    }
}
