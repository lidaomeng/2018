package com.lidaomeng.jwtuserservice.client.hystrix;


import com.lidaomeng.jwtuserservice.client.AuthServiceClient;
import com.lidaomeng.jwtuserservice.entity.JWT;
import org.springframework.stereotype.Component;


@Component
public class AuthServiceHystrix implements AuthServiceClient {
    @Override
    public JWT getToken(String authorization, String type, String username, String password) {
        return null;
    }
}
