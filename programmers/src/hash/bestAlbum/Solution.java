package hash.bestAlbum;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();

        String[] g = {"classic", "pop", "classic", "classic", "pop"};
        int[] p = {500, 600, 150, 800, 2500};

        int[] a = new int[4];

        a = s.solution(g,p);

        System.out.println(a[0]);
        System.out.println(a[1]);
        System.out.println(a[2]);
        System.out.println(a[3]);



    }
    //각 장르별 전체 재생 수 구함
    //제일 많이 재생 된 장르 순서 구함
    //장르별로 가장 많이 재생된 곡 2개 구함
    //재생 숫자 같은경우 순서 대로
    // 장르에 속한곡 하나라면 한곡만
    public int[] solution(String[] genres, int[] plays) {
        int[] answer = {};

        String[] ranking;
        //answer 길이 때문에 하나 만들었음
        Map<Integer, Integer> solution = new HashMap<>();

        int idx = 0;
        //idx에 무슨 음악 장르인지 체크하기 위해 만들었음
        Map<Integer, String> music = new HashMap<>();

        for(int i = 0; i < genres.length; i++){
                music.put(i , genres[i]);
        }

        //인기 있는 장르 랭킹
        ranking = total(genres,plays);

        //sort써서 해가지고 i = ranking.length
        for(int i = ranking.length; i > 0; i--){
            //순위랑 그 순위 인덱스
            int first = 0;
            int second = 0;
            int firstIdx = 0;
            int secondIdx = 0;

            for(Map.Entry<Integer, String> m : music.entrySet()) {
                //각 idx별 장르랑 순위에 맞는 idx에서 1,2위 정하는거
                if(m.getValue().equals(ranking[i-1])){
                    if(first < plays[m.getKey()]){
                        if(first > second){
                            second = first;
                            secondIdx = firstIdx;
                        }
                        first = plays[m.getKey()];
                        firstIdx = m.getKey();

                    }else if(second < plays[m.getKey()]){
                        second = plays[m.getKey()];
                        secondIdx = m.getKey();

                    }
                }
            }
            //1곡만 있는 경우때문에 또 이래 들어옴
            if(second != 0){
                solution.put(idx, firstIdx);
                idx++;
                solution.put(idx, secondIdx);
                idx++;
            }else{
                solution.put(idx, firstIdx);
                idx++;
            }
        }
        // answer 길이 정해줌
        answer = new int[solution.size()];
        //길이 정하고 넣어줌줌
        for(Map.Entry<Integer , Integer> s : solution.entrySet()) {
            answer[s.getKey()] = s.getValue();
        }

        return answer;
    }

    public String[] total (String[] genres, int[] plays) {
        //제일 좋은건 클래식 0 1 2
        // 팝 0 1 2 이런식으로 뭘 만드는게 좋은거 같은데데

        Map<String, Integer> totalPlays = new HashMap<>();
        //각 장르별 재생수 구해서
        for(int i = 0; i < genres.length; i++){
            if(totalPlays.get(genres[i]) == null){
                totalPlays.put(genres[i], plays[i]);
            }else{
                totalPlays.put(genres[i], totalPlays.get(genres[i])+ plays[i]);
            }
        }

        String[] ranking = new String[totalPlays.size()];
        int[] playsOrder = new int[totalPlays.size()];

        //재생수만 배열에 넣어가지고
        int idx = 0;
        for(Map.Entry<String , Integer> t : totalPlays.entrySet()) {

            playsOrder[idx] = t.getValue();
            idx++;
        }
        // 정렬해서
        Arrays.sort(playsOrder);

        //System.out.println("playOrder" + playsOrder[0]);
        //장르별 순위 ranking에 넣어줌
        for(Map.Entry<String , Integer> t : totalPlays.entrySet()) {
            for(int i = 0; i < playsOrder.length; i++) {
                if (playsOrder[i] == t.getValue()) {
                    ranking[i] = t.getKey();
                }
            }

        }
        //System.out.println("ranking" + ranking[0]);
        return ranking;
    }
}
