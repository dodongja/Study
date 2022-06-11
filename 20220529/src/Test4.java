class Pan {
    char data[][];
    Pan() {
        data = new char[5][5];
        for(int i = 0; i < 5; i++)
            for(int j = 0; j < 5; j++)
                data[i][j] = 'X';
    }
    void f() {
        for(int i = 0; i < 5; i++)
            for(int j = 0; j < 5; j++)
                if (i != j) data[i][j] = 'Y';
        for(int i = 0; i < 5; i++) {
            data[i][4-i] = 'Z';
        }
    }
    void print() {
        for(int i = 0; i < 5; i++) {
            for(int j = 0; j < 5; j++) {
                System.out.print(data[i][j]);
            }
            System.out.println();
        }
    }
}


public class Test4 {
    public static void main(String args[])
    {
        Pan p = new Pan();
        p.f();
        p.print();
    }

}
