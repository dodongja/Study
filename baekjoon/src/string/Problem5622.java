package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem5622 {
		public static void main(String[] args) throws IOException  {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		
		String str = br.readLine().toUpperCase();
		
		int time = 0;
		
		for(int i = 0; i < str.length(); i++) {
			if(str.charAt(i) >= 86) {
				time += 10;
			}else if(str.charAt(i) >= 83) {
				time += 9;
			}else if(str.charAt(i) >= 80) {
				time += 8;
			}else if(str.charAt(i) >= 77) {
				time += 7;
			}else if(str.charAt(i) >= 74) {
				time += 6;
			}else if (str.charAt(i) >= 71) {
				time += 5;
			}else if (str.charAt(i) >= 68) {
				time += 4;
			}else if (str.charAt(i) >= 65) {
				time += 3;
			}else if (str.charAt(i) == '1') {
				time += 2;
			}else {
				time += 11;
			}
		}
		
		System.out.println(time);
		
		
		}
}		
