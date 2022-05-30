package com.example.demo.domain;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.Collection;

import static javax.persistence.GenerationType.AUTO;

@Entity
@Data
@NoArgsConstructor //파라미터가 없는 기본 생성자를 생성
@AllArgsConstructor //모든 필드 값을 파라미터로 받는 생성자를 만듦
public class User {

    //@ManyToMany 말고 양방향 맵핑
    @Id
    @GeneratedValue(strategy = AUTO)
    private Long id;
    private String name;
    //임의의 사용자 이름, 이메일도 String으로 될 수 있음
    private String username;
    private String password;
    //유저 정보 로드할 때 동시에 데이터베이스에 있는 role도 로드 할 것이다
    @ManyToMany(fetch = FetchType.EAGER)
    private Collection<Role> roles = new ArrayList<>();
}
