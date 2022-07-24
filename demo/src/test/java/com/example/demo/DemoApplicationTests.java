package com.example.demo;


import net.nurigo.java_sdk.api.Message;
import net.nurigo.java_sdk.exceptions.CoolsmsException;

import org.json.simple.JSONObject;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.HashMap;

@SpringBootTest
class DemoApplicationTests {


	@Test
	void contextLoads() {
		//만약에 2022-06-18 ~ 2022-06-20 까지 머무른다 치면 어떻게 검색이 되어야할까요?
		//그럼 20022-06-18/2022-06-19 를 머물러야함
		//그럼 06/18 , 06/19 가능한 날짜를 구해야함

		//일단 그럼 사업자는 날짜마다 방을 등록해야되나??
		//만약에 객실1을 이용할 수 있는 날짜가 20220-06-18 ~ 2022-06-31 일이라치자
		// roomNo  | check_in | check_out
		//  1 	   | 2022-06-18| 2022-06-31 이라고 들어오면 만약에 중간에 방하나가 나가면 어떻게되는가??
		//결국에는 날짜마다 객실 이용한 것이 다 들어가야되는가?
		// 그럼 결국 등록할때 날짜 마다 객식 종류/ 객실 갯수를 등록해줘야함( 일할때 해봤는데 상당히 귀찮고 비효율적이라생각함)
		//하지만 달리 방법이 생각 안남 ㅋ
		// 그럼 결국 db에도 객실 종류/ 날짜/ 갯수 이렇게 들어가줘야하는건가??
		//그럼 상세보기 페이지로 넘어가면 객실 종류의 정보가 들어오고 / 날짜 검색했을때 가능한 객실과 , 방의갯수가 업로드 되면 되는거네
		// 배열로 만든다??
		// 날짜당 이용할 수 있는 방은 그냥 다 하나 하나씩 db저장하나???


		//roomsetting 에는 방갯수 가격 날짜  3가지 요래하면 되려나
		//그러면 날짜가 3박4일이다  첫쨋날 <= date < 마지막날 해서 나온것들
		// 방갯수중최소값으로 해서 가능한 방 구하면될듯??
		// 기본 1박은 그냥 그날짜 방갯수만 가져오면 됨
		//

	}
}
