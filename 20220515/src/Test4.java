class TreeNode {
    int val;
    TreeNode one = null;
    TreeNode two = null;
    TreeNode three = null;
    TreeNode(int n,int data[]) {
        val = data[n-1];
        n--;
        if (n > 0) one = new TreeNode(n,data);
        n--;
        if (n > 0) two = new TreeNode(n,data);
        n--;
        if (n > 0) three = new TreeNode(n,data);
    }
    void traverse() {
        if (one != null) one.traverse();

/* ÀÌ ºÎºÐÀº ³ªÁß¿¡ ½É½ÉÇ®ÀÌ·Î ÇØ º¼ °Í.
//    if (two != null) two.traverse();
//    if (three != null) three.traverse();
*/

        System.out.println(val);
    }
}

public class Test4 {
    public static void main(String args[])
    {
        int data[] = { 1, 2, 3, 4, 5 };
        TreeNode t = new TreeNode(5,data);
        t.traverse();
    }

}
