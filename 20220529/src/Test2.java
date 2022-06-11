class List {
    static class Node {
        public Node next;
        public int data;
        public Node(int data) {
            this.next = this;
            this.data = data;
        }
    }
    Node head;
    public List() {
        head = null;
    }
    public void add(int data) {
        Node tmp = new Node(data);
        if (head != null) head.next = tmp;
        head = tmp;
    }
    public void print(int n) {
        Node tmp = head;
        for(int i = 0; i < n; i++) {
            System.out.println(tmp.data);
            tmp = tmp.next;
        }
    }
}


public class Test2 {
    public static void main(String args[])
    {
        List x = new List();
        x.add(1);
        x.add(2);
        x.add(3);
        x.add(4);
        x.add(6);
        x.print(8);
    }

}
