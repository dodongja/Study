package com.example.demo.filter;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

import static org.springframework.util.MimeTypeUtils.APPLICATION_JSON_VALUE;

@Slf4j
public class CustomAuthenticationFilter extends UsernamePasswordAuthenticationFilter {
    private final AuthenticationManager authenticationManager;

    public CustomAuthenticationFilter(AuthenticationManager authenticationManager) {
        this.authenticationManager = authenticationManager;
    }

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response) throws AuthenticationException {
        log.info("request" + request.getClass());
        String username = request.getParameter("username");
        String password = request.getParameter("password");
        log.info("Username is: {}", username); log.info("Password is: {}", password);
        //request에서 정보가져와서(username, password) 인증 토큰에 전달
        UsernamePasswordAuthenticationToken authenticationToken = new UsernamePasswordAuthenticationToken(username, password);
        //그후 인증매니저 불러서 사용자를 인증하도록 함
        return authenticationManager.authenticate(authenticationToken);

    }

    //로그인 성공하면 access token 이랑 refresh token 발행
   @Override
    protected void successfulAuthentication(HttpServletRequest request, HttpServletResponse response, FilterChain chain, Authentication authentication) throws IOException, ServletException {
       //로그인한 사용자 정보를 가져와 json토큰 생성 할 수 있다
        User user = (User)authentication.getPrincipal();
       //JSON 웹 토큰과 리프레쉬토큰에 서명하는데 사용할 알고리즘
       Algorithm algorithm = Algorithm.HMAC256("secret".getBytes());
       String access_token = JWT.create()
               .withSubject(user.getUsername()) //토큰 제목 - 토큰에서 사용자에 대한 식별 값이 됨
               .withExpiresAt(new Date(System.currentTimeMillis()+ 10 * 60 * 1000)) // 토큰 만료 시간 10분
               .withIssuer(request.getRequestURL().toString()) //토근 발급자
               //Payload 부분에는 토큰에 담을 정보가 들어있습니다.
               // 여기에 담는 정보의 한 ‘조각’ 을 클레임(claim) 이라고 부르고,
               // 이는 name / value 의 한 쌍으로 이뤄져있습니다
               .withClaim("roles",user.getAuthorities()
                                        .stream().map(GrantedAuthority::getAuthority)
                                        .collect(Collectors.toList()))
               .sign(algorithm);

       String refresh_token = JWT.create()
               .withSubject(user.getUsername()) //토큰 제목 - 토큰에서 사용자에 대한 식별 값이 됨
               .withExpiresAt(new Date(System.currentTimeMillis()+ 30 * 60 * 1000)) // 토큰 만료 시간 10분
               .withIssuer(request.getRequestURL().toString()) //토근 발급자
               .sign(algorithm);

       /*response.setHeader("access_token",access_token);
       response.setHeader("refresh_token",refresh_token);*/
       Map<String, String> tokens = new HashMap<>();
       tokens.put("access_token",access_token);
       tokens.put("refresh_token",refresh_token);
       response.setContentType(APPLICATION_JSON_VALUE);
       new ObjectMapper().writeValue(response.getOutputStream(), tokens);
    }
}
