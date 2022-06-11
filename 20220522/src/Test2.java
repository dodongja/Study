public class Test2 {
    public static void main(String args[])
    {
        int x = 1;
        for (int i = 0; i < 4; i++)
            x = x << 1;
        x = x | 8;
        x = x >> 2;
        x = x & 15;
        System.out.println("x = "+x);
    }

}
