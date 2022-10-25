package array;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;


public class Problem2566 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int big = 0;
        String answer = "";
        List<String> ls = new ArrayList<>();

        for(int i = 0; i < 9; i++){
            for(int j = 0; j < 9; j++){
                 int a = sc.nextInt();
                 if(a > big){
                     ls.clear();
                     big = a;
                     answer = (i+1) + " " + (j+1);
                     ls.add(answer);
                 }else if( a == big){
                     answer = (i+1) + " " + (j+1);
                     ls.add(answer);
                 }
            }
        }
        if(ls.size() > 1) {
            int ran = (int) (Math.random() * ls.size());
            System.out.println(big);
            System.out.print(ls.get(ran));
        }else {
            System.out.println(big);
            System.out.print(ls.get(0));
        }


    }
}
