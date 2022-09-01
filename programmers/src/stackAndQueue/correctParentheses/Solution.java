package stackAndQueue.correctParentheses;

import java.util.Stack;

class Solution {
    boolean solution(String s) {
        boolean answer = true;

        Stack<Character> st = new Stack<>();

        if(s.charAt(s.length()-1) == '('){
            answer = false;
            return answer;
        }

        for(int i = 0; i < s.length(); i++){
            if(s.charAt(i) == ')'){
                if(st.isEmpty()){
                    answer = false;
                    return answer;
                }

                st.pop();
            }else{
                st.push('(');

                if(st.size() > s.length()/2){
                    answer = false;
                    return answer;
                }
            }
        }

        if(!st.isEmpty()){
            answer = false;
            return answer;
        }


        return answer;
    }
}
