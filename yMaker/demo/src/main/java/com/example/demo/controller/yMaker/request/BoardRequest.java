package com.example.demo.controller.yMaker.request;

import lombok.Data;

@Data
public class BoardRequest {
    private String id;
    private String name;
    private String gender;
    private String country;
    private String city;
    private String startDate;
    private String endDate;
}

