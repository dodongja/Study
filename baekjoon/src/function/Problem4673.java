package function;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Set;

public class Problem4673 {
    public static void main(String[] args) {
        //1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97
        // 이 숫자에서 공식을 구할 것인가
        // 각 숫자의 생성자 구하는 법을 구할것인가
        // d(n)값을 구해서 1~ n까지 값에서 제거하는식으로해야되나? 비효율적인거같은디
        // 1 ~ n까지 가면서 셀프넘버인 애들을 list에 넣을수 있으면 그게 베스트인거 같음
        //근데 그 숫자가 셀프넘버인지 어떻게 아는지 생각해봐야할 듯

        //생성자로 구해진 값을 set에 넣어주고
        // i랑를 set에 add할때 true가 나오면 셀프넘버?로 가보자

        //boolean배열 만들어서 0~9999
        //sum값 index-1 넣어준곳 true고
        // false인곳은 idx + 1해서 이 값을 출력해주는구나 허허허 그런 방법이!
        //so sexy한 님들이 많네

        final int N = 10000;

        Set<Integer> numbers = new HashSet<>();

        ArrayList<Integer> selfNumbers = new ArrayList<>();

        for(int i = 1; i <= N; i++){

            int digit = digit(i);

            int x = 10;
            int sum = 0;

            //이곳이 문제네
            sum =  dn(i,digit);

            numbers.add(sum);

            boolean check = numbers.contains(i);

            if(check == false){
                selfNumbers.add(i);
            }
        }

        for(int x = 0; x < selfNumbers.size(); x++){
            System.out.println(selfNumbers.get(x));
        }
    }
    //i가 몇자리 수인지 구함
    public static int digit(int i){

        int x;
        int digit = 1;
        int n = 1;

        while(true){
            x = i/n;

            if(x < 10){
                break;
            }

            n *= 10;
            digit++;
        }

        return digit;
    }
    //d(n)의 값 구함
    public static int dn(int i, int digit) {

        int tenOfPower;

        if(digit != 0){
            tenOfPower = (int)(Math.pow(10, digit-1));
        }else{
            //1자리숫자들
            tenOfPower = 1;
        }

        int remain = i;
        int sum = 0;

        //이곳이 문제네
        for(int j = 0; j < digit; j++){
            sum += remain /tenOfPower;
                    // 21 임 2 + 1 /10 %10
                    //111 임 1 + 1 + 1  /100 = 1 %100 = 11 /10 = 1 %10 = 1

            if(tenOfPower != 1) {
                remain %= tenOfPower;

                tenOfPower /= 10;
            }
        }

        sum += i;

        return sum;
    }
}
