import java.util.HashSet;
import java.util.Set;
import java.util.StringTokenizer;

public class Solution {
    //한 번에 한명의 유저신고
    //신고 횟수 제한 x , 계속 신고 가능, 동일한유저 횟수 1회로 처리
    //k번이상 신고된 유저는 게시판 이용 정지, 신고한 유저에게 정지 사실 메일로 발송
    //마지막에 한꺼번에 정지 시키고 , 정지 메일 발송
    //유저 목록 muzi, frodo, apeach, neo
    //k는 신고 횟수
    //이용자 아이디 id_list
    //신고한 이용자의 id정보 report
    //각 유저별로 처리결과 메일을 받은 횟수 배열에 담아 return
    public int[] solution(String[] id_list, String[] report, int k) {
        int[] answer = {};
        int[] reportNum ={};
        answer = new int[id_list.length];
        reportNum = new int[id_list.length];
        //중복 걸러내기??
        Set<String> set = new HashSet<String>();
        for(int i = 0; i < report.length; i++){
            set.add(report[i]);
        }

        for(int i = 0; i < id_list.length; i++){
                for(String s : set) {
                    String[] split = s.split("\\s+");

                    if(id_list[i].equals(split[1])){
                        reportNum[i]++;
                    }
                }
        }

        for(int i = 0; i < id_list.length; i++){
            for(String s : set) {
                String[] split = s.split("\\s+");

                if(id_list[i].equals(split[0])){
                    for(int y = 0; y < id_list.length; y++){
                        if(id_list[y].equals(split[1])){
                            if(reportNum[y] >= 2){
                                answer[i]++;
                            }
                        }
                    }
                }
            }
        }

        for(int i = 0; i < answer.length; i++){
            System.out.println(answer[i]);
        }

        return answer;
    }
}


