public class CircleQueue {
    int front = 0;
    int rear = 0;
    int[] queue;
    int idx = 0;

    public CircleQueue(int idx){
        queue = new int[idx];
        this.idx = idx;
    }

    public void enQueue(int i){
        if((rear+1)%(idx) != front){
            if(front ==0) {
                rear += 1;
                queue[rear] = i;
            }else{
                queue[front] = i;
            }
        }else{
            System.out.println("isFull");
        }
    }

    public void deQueue(){
        if(front < idx) {
            front += 1;
            System.out.println(queue[front]);
            queue[front] = 0;
        }else{
            front = 0;
        }
    }
}
