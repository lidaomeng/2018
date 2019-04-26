package com.lidaomeng.jwtuserservice.service;


import com.lidaomeng.jwtuserservice.client.hystrix.AuthServiceHystrix;
import com.lidaomeng.jwtuserservice.config.CustomPasswordEncoder;
import com.lidaomeng.jwtuserservice.dao.UserDao;
import com.lidaomeng.jwtuserservice.dto.UserLoginDTO;
import com.lidaomeng.jwtuserservice.entity.JWT;
import com.lidaomeng.jwtuserservice.entity.User;
import com.lidaomeng.jwtuserservice.exception.UserLoginException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;


@Service
public class UserServiceDetail implements UserDetailsService {

    @Autowired
    private UserDao userRepository;
    @Autowired
    AuthServiceHystrix authServiceHystrix;

    @Autowired
    CustomPasswordEncoder customPasswordEncoder;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        return userRepository.findByUsername(username);
    }

    public User insertUser(String username, String password) {
        User user = new User();
        user.setUsername(username);
        user.setPassword(customPasswordEncoder.encode(password));
        return userRepository.save(user);
    }

    public UserLoginDTO login(String username, String password) {
        User user = userRepository.findByUsername(username);
        if (null == user) {
            throw new UserLoginException("error username");
        }
        if (!customPasswordEncoder.matches(password, user.getPassword())) {
            throw new UserLoginException("error password");
        }
        // 获取token
        JWT jwt = authServiceHystrix.getToken("Basic dXNlci1zZXJ2aWNlOjEyMzQ1Ng==", "password", username, password);
        // 获得用户菜单
        if (jwt == null) {
            throw new UserLoginException("error internal");
        }
        UserLoginDTO userLoginDTO = new UserLoginDTO();
        userLoginDTO.setJwt(jwt);
        userLoginDTO.setUser(user);
        return userLoginDTO;

    }
}
