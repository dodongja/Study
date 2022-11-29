package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.StringTokenizer;

public class Problem2908 {
    public static void main(String[] args) throws IOException {
    	BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    	
    	StringTokenizer st = new StringTokenizer(br.readLine()," ");
    	
    	StringBuffer x = new StringBuffer(st.nextToken());
    	StringBuffer y = new StringBuffer(st.nextToken());
    	
    	int a = Integer.parseInt(x.reverse().toString());
    	int b = Integer.parseInt(y.reverse().toString());
    	
    	if(a > b) {
    		System.out.println(a);
    	} else {
    		System.out.println(b);
    	}
    }
}
