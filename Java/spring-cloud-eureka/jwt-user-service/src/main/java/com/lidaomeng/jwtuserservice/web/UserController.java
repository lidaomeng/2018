package com.lidaomeng.jwtuserservice.web;

import com.lidaomeng.jwtuserservice.dto.UserLoginDTO;
import com.lidaomeng.jwtuserservice.entity.User;
import com.lidaomeng.jwtuserservice.service.UserServiceDetail;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    UserServiceDetail userServiceDetail;

    @PostMapping("/register")
    public User postUser(@RequestParam("username") String username, @RequestParam("password") String password) {
        //参数判断，省略
        return userServiceDetail.insertUser(username, password);
    }

    @PostMapping("/login")
    public UserLoginDTO login(@RequestParam("username") String username, @RequestParam("password") String password) {
        //参数判断，省略
        return userServiceDetail.login(username, password);
    }
}
