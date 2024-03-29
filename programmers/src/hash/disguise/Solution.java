package hash.disguise;

import java.util.HashMap;
import java.util.Map;

public class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();

        String[][] clothes = {{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}};


        s.solution(clothes);

    }
    public int solution(String[][] clothes) {
        int answer = 1;
         Map<String, Integer> m = new HashMap<>();

        for(int i = 0; i < clothes.length; i++){
            if(m.get(clothes[i][1]) == null){
                m.put(clothes[i][1], 1);
            }else{
                m.put(clothes[i][1], m.get(clothes[i][1])+ 1);
            }
        }

        System.out.println(m);
        System.out.println(m.values());
        for (String key : m.keySet()) {
            answer *= m.get(key)+1;
        }

        return answer;
    }
    //face  , headgear, eyewear 
    
    // 2 1 = 5
    // 2 2 = 8
    //3 = 3
    //3 2 1 = 23
    // 4 2 2 = 44
    // 4 2 3 1 = 119

    // *1
    // 2*1 + 1*1 + 2*1
    //

    //head 3-1 3-2 3-3    3
    // eye 2-1 2-2        2

    // 1 - 5
    //2 - 6
    // 11
    // 5 + 3*2

    //face 3
    //3

    //face - 3   3-1 3-2 3-3
    //head - 2   2-1 2-2
    //eye - 1    1-1
    // 1
    // 6개

    //2
    // 3*2 + 2*1 + 3*1    -   6 +2 +3 = 11개

    //3
    // 3-1 2개
    // 3-2 2개
    // 3-3 2개
    // 6개

    //23개

    //face - 4   4-1 4-2 4-3 4-4
    //head - 2   2-1 2-2
    //eye - 2    1-1 1-2

    // 1
    // 8개

    //2
    //4*2 + 4*2 + 2*2  = 20개

    //3
    //4-1 2-1 1-1
    //4-1 2-1 1-2
    //4-1 2-2 1-1
    //4-1 2-2 1-2

    // 16개

    // 8개 + 20+ 16 = 44개


    //1종류 = 종류 갯수 만큼
    //2종류 = (1종류 + 2종류) + (1종류*2종류)
    //3종류 = (1종류+2종류+3종류) + (1종류*2종류) + (1종류*3종류) + (2종류*3종류) + 1종류 * 2종류 * 3종류
    //4종류 = (1 + 2 + 3 + 4) +

    //eye , nose , head, arm
    // eye - 1-1 1-2 1-3 1-4   //4
    // nose - 2-1 2-2        //2
    // head - 3-1 3-2 3-3    //3             3 2 3 1  19가지
    // arm - 4-1          //1

    // 1 - 4+ 2+ 3+ 1           10개

    // 2 - 4*2 + 4*3 + 4*1 + 2*3 + 2*1 + 3*1   8 + 12 + 4 +6 + 2 + 3 = 35
    // 1-1 2-1
    // 1-1 2-2
    // 1-1 3-1
    // 1-1 3-2
    // 1-1 3-3
    // 1-1 4-1
    // *4 = 24

    //2-1 3-1
    //2-1 3-2
    //2-1 3-3
    //2-1 4-1
    //*2 = 8

    //3-1
    //*3 = 3


    //1-1 2-1 3-1
    //1-1 2-1 3-2      1 2 3   1 2 4  1 3 4  2 3 4
    //1-1 2-1 3-3
    //1-1 2-1 4-1
    //1-2 2-1 3-1
    //1-2 2-1 3-2
    //1-2 2-1 3-3
    //1-2 2-1 4-1
    //1-3 2-1 3-1
    //1-3 2-1 3-2
    //1-3 2-1 3-3
    //1-3 2-1 4-1
    //1-4 2-1 3-1
    //1-4 2-1 3-2
    //1-4 2-1 3-3
    //1-4 2-1 4-1
    //16 *2  = 32


    //2-1 3-1 4-1
    //2-1 3-2 4-1
    //2-1 3-3 4-1
    //2-2 3-1 4-1
    //2-2 3-2 4-1
    //2-2 3-3 4-1
    // 6

    //1-1 3-1 4-1
    //1-1 3-2 4-1
    //1-1 3-3 4-1
    // * 4 = 12


    //34

    //45 + 34 = 79

    //1-1 2-1 3-1 4-1
    //1-1 2-1 3-2 4-1
    //1-1 2-1 3-3 4-1
    //1-1 2-2 3-1 4-1
    //1-1 2-2 3-2 4-1
    //1-1 2-2 3-3 4-1

    // 6*4 = 24

    //79 + 24 = 103 + 16 = 119



    // 3 - 3
    // 2 1  - 5
    // 3 2 1 - 23   3 * 2 *1 = 6 * 3 18 + 5
    // 4 2 2 - 44


    // 1 - 3 2  1 = 6
    // 2  3-1/2-1 3-1/2-2 3-1/1-1
}
