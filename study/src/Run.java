public class Run {
    public static void main(String[] args) {
        /*LineQueue lineQueue = new LineQueue(5);
        lineQueue.enQueue(5);
        lineQueue.enQueue(6);
        lineQueue.enQueue(6);
        lineQueue.enQueue(6);
        lineQueue.enQueue(6);
        lineQueue.deQueue();
        lineQueue.enQueue(6);



        System.out.println(lineQueue.queue[0]);
        System.out.println(lineQueue.queue[1]);
        System.out.println(lineQueue.queue[2]);
        System.out.println(lineQueue.queue[3]);
        System.out.println(lineQueue.queue[4]);*/

        CircleQueue circleQueue = new CircleQueue(5);
        circleQueue.enQueue(5);
        circleQueue.enQueue(5);
        circleQueue.enQueue(5);
        circleQueue.enQueue(5);
        circleQueue.enQueue(5);
        circleQueue.deQueue();
        circleQueue.enQueue(6);

        System.out.println(circleQueue.queue[0]);
        System.out.println(circleQueue.queue[1]);
        System.out.println(circleQueue.queue[2]);
        System.out.println(circleQueue.queue[3]);
        System.out.println(circleQueue.queue[4]);
    }
}
