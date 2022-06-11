public class Solution {
    public static void main(String[] args) {
        int[] numbers = {7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2};
        String hand = "left";
        Solution s = new Solution();
        String result = s.solution(numbers, hand);
        System.out.println(result);

    }

    //왼손 -> 시작 * 오른손 -> 시작 #
    //상하좌우 4가지 방향으로만 이동 가능
    //1,4,7 입력시 왼손
    //3,6,9 입력 오른손
    //2, 5 , 8 , 0 입력은 가까운 손가락 이용
    // 거리 같으면 오른손잡이 오른손 , 왼손잡이 왼손
    //순서대로 번호가 담긴 배열 numbers , hand는 어느 손 잡이인가
    // 각 번호를 누를 손가락이 왼손인지 오른손인지 나타내는 연속된 문자열 형태로 return
    //[4][3]배열
    //거리를 구해서 가까운 쪽 구해야겠네
    //한칸에 +1
    //각 위치에서 거기까지 +1해서 거리 누가큰지 구해서 누르기
    public String solution(int[] numbers, String hand) {
        String answer = "";
        answer = start(numbers,hand,answer);

        return answer;
        }

    public String start (int[] numbers, String hand,String answer) {
        //휴대폰 번호 위치
        int[][] mobile = new int[4][3];
        int sum = 0;
        for (int i = 0; i < mobile.length - 1; i++) {
            for (int j = 0; j < mobile[i].length; j++) {
                mobile[i][j] += sum + 1;
                sum = mobile[i][j];
            }
        }
        //시작 위치
        int leftPointX = 3;
        int leftPointY = 0;
        int rightPointX = 3;
        int rightPointY = 2;
        //이동 위치
        int movePointX = 0;
        int movePointY = 0;

        for (int i = 0; i < numbers.length; i++) {
            for (int j = 0; j < mobile.length; j++) {
                for (int k = 0; k < mobile[j].length; k++) {
                    System.out.println("start");
                    if (j == 3 && k == 0 || j == 3 && k == 2) {
                        continue;
                    }

                    if (numbers[i] == mobile[j][k]) {
                        System.out.println("checkMovePoint");
                        movePointX = j;
                        movePointY = k;

                        if (numbers[i] == 1 || numbers[i] == 4 || numbers[i] == 7) {
                            System.out.println("check 1,4,7");
                            answer += "L";
                            leftPointX = movePointX;
                            leftPointY = movePointY;
                        } else if (numbers[i] == 3 || numbers[i] == 6 || numbers[i] == 9) {
                            System.out.println("check 3,6,9");
                            answer += "R";
                            rightPointX = movePointX;
                            rightPointY = movePointY;
                        } else {
                            System.out.println("check else");
                            int leftMoving = Math.abs((leftPointX - movePointX)) +  Math.abs((leftPointY - movePointY));
                            int rightMoving = Math.abs((rightPointX - movePointX)) +  Math.abs((rightPointY - movePointY));
                            System.out.println(leftMoving);
                            System.out.println(rightMoving);
                            if (leftMoving < rightMoving) {
                                System.out.println("check win left");
                                answer += "L";
                                leftPointX = movePointX;
                                leftPointY = movePointY;
                            } else if (leftMoving > rightMoving) {
                                System.out.println("check win right");
                                answer += "R";
                                rightPointX = movePointX;
                                rightPointY = movePointY;
                            } else {
                                if (hand.equals("left")) {
                                    System.out.println("check draw but left");
                                    answer += "L";
                                    leftPointX = movePointX;
                                    leftPointY = movePointY;
                                } else {
                                    System.out.println("check draw but right");
                                    answer += "R";
                                    rightPointX = movePointX;
                                    rightPointY = movePointY;
                                }
                            }
                        }

                    }
                }
            }
        }

        return answer;
    }
}
//numbers가 14578로 주어진경우
            //왼손거리 오른손거리 구해서 가까운거 비교해서 answer에 넣으면 되네
            //만약 7을 누르면 [3][0] -> [2][0]  1칸  [3][2] -> [2][0]  3칸  3-2 =1 //3-2=1 2-0 = 2
            //만약 5를 누르면 [3][0] -> [1][1] 3칸 [3][2] -> [1][1]  3칸    3-1 =2 0 -1 = 1 // 3-1 =2 2-1= 1
            //만약 3을 누르면 [3][0] -> [0][2] 5칸 [3][2] -> [0][2] 3칸     3-0 = 3 0-2 = 2// 3-0= 3 2-2=0
            //각 배열 수 -한거 더한만큼 이동하는 칸임




