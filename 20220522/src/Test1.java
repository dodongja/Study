public class Test1 {
    public static void main(String args[])
    {
        System.out.println("answer = "+what(4));
    }
    static int what(int n)
    {
        if (n == 0) return 1;
        else if (n == 1) return 1;
        else if (n > 0) return n + what(n-2) + 1;
        else {
            System.err.println("Unexpected minus argument.");
            System.exit(-1);
            return -1;
        }
    }

}
