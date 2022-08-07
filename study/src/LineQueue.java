public class LineQueue {
    int front = 0;
    int rear = 0;
    int[] queue;

    public LineQueue(int idx){
        queue = new int[idx];
    }

    public void enQueue(int i){
        if(queue.length != rear){
            queue[rear] = i;
            rear += 1;
        }else{
            System.out.println("isFull");
        }

    }

    public void deQueue(){
        System.out.println(queue[front]);
        queue[front] = 0;
        front += 1;

    }

}
