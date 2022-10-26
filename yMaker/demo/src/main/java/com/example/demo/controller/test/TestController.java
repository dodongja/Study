package com.example.demo.controller.test;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import com.example.demo.controller.yMaker.BoardController;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Controller
@RequestMapping("/")
public class TestController {
	
	 @GetMapping("/main")
	    public String domain(){
	        log.info("main");

	        return "/test";
	    }

}
