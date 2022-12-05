package function;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem1065 {
    //어떤 양의 정수 X의 각 자리가 등차수열을 이룬다면, 그 수를 한수라고 한다. 등차수열은 연속된 두 개의 수의 차이가 일정한 수열을 말한다.
    // N이 주어졌을 때, 1보다 크거나 같고, N보다 작거나 같은 한수의 개수를 출력하는 프로그램을 작성하시오.

    public static void main(String[] args) throws IOException {

        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int n = Integer.parseInt(br.readLine());
        int answer;
        int tmp;
        int placeValue1 = 0;
        int placeValue2;
        int minus = 0;
        int cnt;
        boolean check;

        //일단 2자리수까지는 전부 한수 인거같음
        if(n < 100){
            System.out.println(n);
        } else {
            answer = 99;

            for (int i = 100; i <= n; i++) {
                tmp = i;
                cnt = 0;
                check = true;
                while (check) {
                    if (cnt == 0) {
                        placeValue1 = tmp % 10;
                        tmp /= 10;
                        cnt++;
                        continue;
                    } else {
                        placeValue2 = tmp % 10;
                        if (cnt == 2) {
                            if (placeValue1 - placeValue2 != minus) {
                                check = false;
                                continue;
                            } else {
                                cnt = 1;
                            }
                        }
                        minus = placeValue1 - placeValue2;
                        placeValue1 = placeValue2;
                        cnt++;
                    }
                    tmp /= 10;

                    if (tmp == 0) {
                        answer++;
                        check = false;
                    }
                }
            }

            System.out.println(answer);
        }
    }
}
