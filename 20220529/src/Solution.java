public class Solution {
    public static void main(String[] args) {
        String new_id = "..##..#>asfdlh$WV0sdㅏㅗㄴㅁ이ㅏㅁㄴㅇㅁㄴㅇ..";
        Solution sl = new Solution();
        System.out.println(sl.solution(new_id));
    }

    public String solution(String new_id) {
        String answer = new_id;
        final int MAX_LENGTH = 15;
        final int MIN_LENGTH = 2;

        //소문자
        answer = lowerCase(answer);
        System.out.println("lower case: " + answer);
        //알파벳 소문자, 숫자, 빼기(-), 밑줄(_), 마침표(.)를 제외한 모든 문자를 제거합니다.
        answer = delete(answer);
        System.out.println("delete:" + answer);
        //마침표(.)가 2번 이상 연속된 부분을 하나의 마침표(.)로 치환합니다.
        answer = deleteDuplication(answer);
        System.out.println("deleteDuplication:" + answer);
        //new_id에서 마침표(.)가 처음이나 끝에 위치한다면 제거합니다.
        answer = deleteDot(answer);
        System.out.println("deleteDot:" + answer);
        //new_id가 빈 문자열이라면, new_id에 "a"를 대입합니다.
        answer = ifEmpty(answer);
        System.out.println("ifEmpty:" + answer);
        //new_id의 길이가 16자 이상이면, new_id의 첫 15개의 문자를 제외한 나머지 문자들을 모두 제거합니다.
        //만약 제거 후 마침표(.)가 new_id의 끝에 위치한다면 끝에 위치한 마침표(.) 문자를 제거합니다.
        answer = checkLengthOver15(answer,MAX_LENGTH);
        answer = deleteDot(answer);
        System.out.println("checkLengthOver15:" + answer);
        //new_id의 길이가 2자 이하라면, new_id의 마지막 문자를 new_id의 길이가 3이 될 때까지 반복해서 끝에 붙입니다.
        answer = checkLengthDown2(answer,MIN_LENGTH);
        System.out.println("checkLengthDown2:" + answer);


        return answer;
    }

    public String lowerCase(String answer){
        return answer.toLowerCase();
    }

    public String delete(String answer) {
        return answer.replaceAll("[^a-z0-9-_.]","");
    }

    public String deleteDuplication(String answer){
        for(int i = 0; i < answer.length(); i++){
            if(i+1 < answer.length()) {
                //i , i+1이 같으면 i 제거
                if (answer.charAt(i) == '.') {
                    if (answer.charAt(i) == answer.charAt(i + 1)) {
                        StringBuffer delete = new StringBuffer(answer);
                        delete.deleteCharAt(i);
                        //StringBuffer -> String
                        answer = delete.toString();
                        i--;
                    }
                }
            }
        }
        return answer;
    }

    public String deleteDot(String answer){
        StringBuffer delete = new StringBuffer(answer);
        if(answer.charAt(0) == '.'){
            delete.deleteCharAt(0);
            answer = delete.toString();

        }
        if(answer.length() > 1){
            if(answer.charAt(answer.length() - 1) == '.') {
                delete.deleteCharAt(answer.length() - 1);
                answer = delete.toString();
            }
        }
        return answer;
    }

    public String ifEmpty(String answer) {
        if(answer.equals("")){
            answer = "a";
        }
        return answer;
    }

    public String checkLengthOver15(String answer, int max_length){
        if(answer.length() > max_length){
            StringBuffer delete = new StringBuffer(answer);
            delete.delete(max_length, answer.length());

            answer = delete.toString();
        }

        return answer;
    }

    public String checkLengthDown2(String answer, int min_length){

        while(true){
            if(answer.length() <= min_length){
                answer = answer + answer.charAt(answer.length() - 1);
                continue;
            }
            break;
        }
        return answer;
    }
}
