package string;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Problem2941 {
    public static void main(String[] args) throws IOException {
        //일단 기본 2개 단위로 짜르고
        //dz=만 예외네
        //ljes=njak
        //글짜 두개 비교 있으면 -> +1
        //3개까지 비교 있으면 -> +1
        // 없으면  +1
        // 후에 다음글자로 넘어가야함
        //비교할 크로아티아 알파벳을 어디에 넣을까??
        //크로아티아 알파벳 비교 메서드를 하나 만드는게 좋을듰?
        // c , d, l, n , s , z
        //입력받고 글자 하나하나 조회
            int numberOfCroatiaAlphabet = 0;
            BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

            char[] input = br.readLine().toCharArray();

            for (int i = 0; i < input.length; i++) {
                try {
                    if (input[i] == 'c') {
                        if (input[i + 1] == '=') {
                            numberOfCroatiaAlphabet++;
                            i++;
                            continue;
                        } else if (input[i + 1] == '-') {
                            numberOfCroatiaAlphabet++;
                            i++;
                            continue;
                        }
                    }

                    if (input[i] == 'd') {
                        if (input[i + 1] == 'z' && input[i + 2] == '=') {
                            numberOfCroatiaAlphabet++;
                            i += 2;
                            continue;
                        } else if (input[i + 1] == '-') {
                            numberOfCroatiaAlphabet++;
                            i++;
                            continue;
                        }
                    }

                    if (input[i] == 'l' || input[i] == 'n') {
                        if (input[i + 1] == 'j') {
                            numberOfCroatiaAlphabet++;
                            i++;
                            continue;
                        }
                    }

                    if (input[i] == 's' || input[i] == 'z') {
                        if (input[i + 1] == '=') {
                            numberOfCroatiaAlphabet++;
                            i++;
                            continue;
                        }
                    }

                    numberOfCroatiaAlphabet++;
                }catch (Exception e){
                    numberOfCroatiaAlphabet++;
                }
        }
        System.out.println(numberOfCroatiaAlphabet);
        }
    }
