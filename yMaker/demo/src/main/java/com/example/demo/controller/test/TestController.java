package com.example.demo.controller.test;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Controller
@RequestMapping("/test")
public class TestController {
	
	 @GetMapping("/main")
     public String domain(){
	        log.info("main");

	        return "/test/test";
	    }

}
