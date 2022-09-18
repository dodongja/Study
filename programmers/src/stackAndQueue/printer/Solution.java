package stackAndQueue.printer;

import java.util.LinkedList;
import java.util.Queue;

public class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] p = new int[6];
        p[0] = 1;
        p[1] = 1;
        p[2] = 9;
        p[3] = 1;
        p[4] = 1;
        p[5] = 1;
        int answer = s.solution(p, 0);

        System.out.println(answer);
    }
    public int solution(int[] priorities, int location) {
        //일단 모든 인덱스의 순서를 구해서 location값의 인덱스를 구할 것인가
        //근데 한넘만 순서알면 나머지는 굳이 몰라도 되는게 아닐까??
        //처음에 있던 녀석이 뒤에 숫자 큰녀석이 있으면 맨뒤로넘어감
        //그럼 이런식으로 하면 최종적으로 마지막 인덱스까지 검사 끝나면 순서가 정해지긋네??
        //근데 옮기고나서 원래 인덱스였던걸 어케 하노 ?? 클래스만들어서 하나??

        int answer = 0;
        int index = 0;
        int[] order = new int[priorities.length];

        Queue<Node> q = new LinkedList<>();
        //idx 와 num 넣어주기
        for(int i = 0; i < priorities.length; i++){
           Node n = new Node(i,priorities[i]);
           q.add(n);
        }

        //1 2 3 2
        // 2 3 2 1
        // 3 2 1 2
        // 2 1 2  idx 1
        // 1 2    idx 2
        // 2 1
        // 1       idx 3

        //1 1 9 1 1 1
        //1 9 1 1 1 1
        //
        while(true){
            //q size 1
            if(q.size() == 1){
                order[index] = q.peek().idx;
                break;
            }

            int idx = q.peek().idx;
            int check = 0;

            for(Node n : q){
                if(idx == n.idx){
                    continue;
                }

                if(n.num > q.peek().num){
                    q.add(q.peek());
                    q.remove();
                    check++;
                    break;
                }
            }

            if(check == 0){
                order[index] = q.peek().idx;
                index++;
                q.remove();
            }
        }


        for(int i = 0; i < order.length; i++){
            if(order[i] == location){
                answer = i + 1;
                break;
            }

        }

        return answer;
    }

    public static class Node {
        int idx;
        int num;
        public Node(int i, int j) {
            this.idx = i;
            this.num = j;
        }
    }
}

