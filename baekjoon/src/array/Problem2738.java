package array;


import java.util.Arrays;
import java.util.Scanner;

public class Problem2738 {
	public static void main(String[] args) throws Exception {
		Scanner sc = new Scanner(System.in);

		int n = sc.nextInt();
		int m = sc.nextInt();

		int[][] a = new int[n][m];
		int[][] b = new int[n][m];
		int[][] answer = new int[n][m];

		a = arr(a, sc);
		b = arr(b, sc);

		answer = add(a, b, answer);

		for(int i = 0; i < answer.length; i++) {
			for(int j = 0; j < answer[i].length; j++) {
				System.out.print(answer[i][j] + " ");
			}
			System.out.println();
		}
	}

	public static int[][] arr(int[][] arr, Scanner sc) {
		for(int i = 0; i < arr.length; i++) {
			for(int j = 0; j < arr[i].length; j++) {
				arr[i][j] = sc.nextInt();
			}
		}

		return arr;
	}

	public static int[][] add(int[][] a, int[][] b, int[][] answer ) {
		for(int i = 0; i < a.length; i++) {
			for(int j = 0; j < a[i].length; j++) {
				answer[i][j] = a[i][j] + b[i][j];
			}
		}

		return answer;
	}


}
