package com.lidaomeng.jwtuserservice.dto;

import com.lidaomeng.jwtuserservice.entity.JWT;
import com.lidaomeng.jwtuserservice.entity.User;


public class UserLoginDTO {

    private JWT jwt;

    private User user;

    public JWT getJwt() {
        return jwt;
    }

    public void setJwt(JWT jwt) {
        this.jwt = jwt;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
