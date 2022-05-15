public class Run {
    public static void main(String args[]){
        String[] id_list = {"muzi", "frodo", "apeach","neo"};
        String[] report = {"muzi frodo","apeach frodo","frodo neo","muzi neo","apeach muzi"};
        int k = 2;

        String[] id_list2 = {"con", "ryan"};
        String[] report2 = {"ryan con", "ryan con","ryan con","ryan con"};
        int k2 = 3;

        Solution st = new Solution();
        st.solution(id_list,report,k);
        st.solution(id_list2,report2,k2);
    }
}
