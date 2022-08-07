package hash.phoneList;

import java.util.*;

class Solution {
    public static void main(String[] args) {
        String i = "dong";

        System.out.println(i.length());
        System.out.println(i.substring(0,i.length()));
        System.out.println(i.substring(0,1));

        Set<String> s = new HashSet<>();
        s.add("a");
        s.add("b");
        System.out.println(s.size());

        Solution solution = new Solution();

        String[] b = new String[5];
        b[0]="12";
        b[1]="123";
        b[2]="1235";
        b[3]="567";
        b[4]="1111";

        Arrays.sort(b);
        System.out.println(b[0]);
        System.out.println(b[1]);
        System.out.println(b[2]);
        System.out.println(b[3]);
        System.out.println(b[4]);

    }
    public boolean solution(String[] phone_book) {
        boolean answer = true;

        Set<String> k = new HashSet<>();

        Arrays.sort(phone_book);

        while(answer) {
            int j = 0;
            for (int i = j; i < phone_book.length; i++) {
                if(k.size() >1 && !k.add(phone_book[i].substring(0, phone_book[j].length()))) {
                    System.out.println("??");
                    answer = false;
                    return answer;
                }
                j++;
            }
        }
        return answer;
    }
}
