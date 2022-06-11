class A {
    void f(A x) {
        System.out.println("A");
    }
}
class B extends A {
    void f(A x) {
        System.out.println("B");
        if (x != null) x.f(null);
    }
}
class C extends A {
    void f(A x) {
        System.out.println("C");
        if (x != null) x.f(null);
    }
}


public class Test3 {
    public static void main(String args[])
    {
        A x = new A();
        B y = new B();
        C z = new C();
        y.f(x);
        z.f(z);
        y.f(z);
        x.f(y);
    }

}
