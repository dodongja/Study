package heap.diskCotroller;

import java.util.ArrayList;

public class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] jobs = new int[][]{{0,3},{3,5},{4,2}};
        // 3 + 5 + 2 = 3
        // 3 + 5 + 6 = 14
        // 3 + 2 + 8 = 13
        System.out.println(s.solution(jobs));
    }
    public int solution(int[][] jobs) {
        //작업진행중
        //다음 작업은 작업 후 길이랑 요청 시간 이랑 - 해서 재일 짧은거해야함?

        //처음에 작업이 시작되고
        //그다음에 작업 가능한 것들중에서
        //제일 작업시간 짧은애 실행해야함
        int answer = 0;
        int finishTime = 0;
        int sum = 0;


        ArrayList<Job> list = new ArrayList<>();
        ArrayList<Job> wait = new ArrayList<>();

        addList(jobs, list);
        arrayList(list);

        //[[0, 3], [1, 9], [2, 6]]
        while (true) {
            for(Job j : list){
                if(j.start <= finishTime){
                    wait.add(j);
                }
            }


            if(!wait.isEmpty()) {
                int count = 0;
                int check = 0;
                int idx = 0;

                for (Job j : wait) {
                    list.remove(j);
                    int time = finishTime + j.jobLength - j.start;
                    if (check == 0) {
                        check = time;
                    }

                    if (check > time) {
                        check = time;
                        idx = count;
                    }

                    count++;
                }

                finishTime += wait.get(idx).jobLength;
                sum += finishTime - wait.get(idx).start;
                wait.remove(idx);
            }else{
                if(list.size()>0){
                    if(list.get(0).start > finishTime){
                        finishTime = list.get(0).start + list.get(0).jobLength;
                    }else{
                        finishTime += list.get(0).jobLength;
                    }

                    sum += finishTime - list.get(0).start;
                    list.remove(0);
                }else{
                    System.out.println(sum);
                    answer = sum / jobs.length;
                    return answer;
                }
            }


        }



    }

    public ArrayList<Job> addList(int[][] jobs, ArrayList<Job> list) {
        // jobs에 들어오는 시작시간이 순서대로 안들어오는경우가 있음
        for(int i = 0; i < jobs.length; i++) {
            Job job = new Job(jobs[i][0], jobs[i][1]);
            list.add(job);
        }
        return list;
    }

    public ArrayList<Job> arrayList(ArrayList<Job> list) {
        boolean check = true;
        int count = 0;

        while (check) {
            for (int i = 0; i < list.size(); i++) {
                if(count == list.size() - 1) {
                    check = false;
                    break;
                }
                //[0,3] , [4,5] , [8,8]
                //
                Job a = list.get(i);
                Job b = list.get(i + 1);

                if(a.start > b.start) {
                    list.set(i, b);
                    list.set(i + 1, a);
                }

                count++;
            }
        }

        return list;
    }
}

class Job{
    int start;
    int jobLength;

    public Job(int start, int jobLength){
        this.start = start;
        this.jobLength = jobLength;
    }
}

// 처음 1 5  위치는 6  걸린시간 위치 - request
// 위치 3 걸린시간 3
//



