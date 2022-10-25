package array;

import java.util.Scanner;

public class Problem10807 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int count = 0;
        int v;
        int num = sc.nextInt();
        int[] arr = new int[num];

        for(int i = 0; i < num; i++){
            arr[i] = sc.nextInt();
        }

        v = sc.nextInt();
        for(int i = 0; i < arr.length; i++){
            if(arr[i] == v){
                count++;
            }
        }

        System.out.println(count);
    }
}
