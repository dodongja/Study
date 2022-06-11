class Stack {
    private int[] s;
    private int top;
    private void initialize() {
        s = new int[10];
        for(int i = 0; i < 10; i++)
            s[i] = -1;
    }
    Stack() {
        initialize();
        top = 1;
    }
    void push(int x) {
        top++;
        top++;
        s[top] = x;
    }
    void pop() {
        top--;
    }
    void print() {
        for(int i = 0; i < 10; i++) {
            System.out.println("s["+i+"] = "+s[i]);
        }
    }
}

public class Test3 {
    public static void main(String args[])
    {
        Stack p = new Stack();

        p.push(1);
        p.push(2);
        p.push(3);
        p.pop();
        p.pop();
        p.pop();
        p.push(4);

        p.print();
    }

}
