class StringConverter {
    private String s[];
    StringConverter(String x,String y) {
        s = new String[2];
        s[0] = x;
        s[1] = y;
    }
    void convert() {
        StringBuffer b = new StringBuffer(s[0]);
        for (int i = 0; i < b.length(); i++)
            b.setCharAt(i,(char)(b.charAt(i)+1));
        s[0] = b.toString();

        s[1] = s[1] + s.length;
    }
    void print() {
        System.out.println("s[0] = "+s[0]);
        System.out.println("s[1] = "+s[1]);
    }
}

public class Test4 {
    public static void main(String args[])
    {
        String x = "kim";
        String y = "lee";
        StringConverter p = new StringConverter(x,y);
        x = "kwon";
        y = "park";
        p.convert();
        p.print();
    }

}
