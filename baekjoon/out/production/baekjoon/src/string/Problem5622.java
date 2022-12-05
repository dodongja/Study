package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem5622 {
	public static void main(String[] args) throws IOException  {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

		char[] input = br.readLine().toCharArray();

		int time = 0;
		// 그냥 예제가 숫자 앞에 있는게 많은듯?
		for(int i = 0; i < input.length; i++) {
			if(input[i] >= 'W') {
				time += 10;
			}else if(input[i] >= 'T') {
				time += 9;
			}else if(input[i] >= 'P') {
				time += 8;
			}else if(input[i] >= 'M') {
				time += 7;
			}else if(input[i] >= 'J') {
				time += 6;
			}else if (input[i] >= 'G') {
				time += 5;
			}else if (input[i] >= 'D') {
				time += 4;
			}else if (input[i] >= 'A') {
				time += 3;
			}
		}

		System.out.println(time);


	}
}

