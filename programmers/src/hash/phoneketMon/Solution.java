package hash.phoneketMon;

import java.util.HashSet;
import java.util.Set;

class Solution {
    public int solution(int[] nums) {
        int answer = 0;
        int n = nums.length/2;

        Set<Integer> monster = new HashSet<>();

        for(int i = 0; i < nums.length; i++){
            monster.add(nums[i]);
        }
        //monseter.size() 3이지만 n=2 이기 때문에 2
        if(n >= monster.size()){
            return answer = monster.size();
        }else{
            return answer = n;
        }
    }
}
