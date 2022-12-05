package array;

public class Problem15596 {


    public long sum(int[] a){

        long ans = 0;
        int n = a.length;

        for(int i = 0; i < n; i++){
            ans += a[i];
        }

        return ans;
    }
}
