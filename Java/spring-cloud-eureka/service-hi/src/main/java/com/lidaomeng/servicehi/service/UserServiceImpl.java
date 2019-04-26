package com.lidaomeng.servicehi.service;

import com.lidaomeng.servicehi.config.CustomPasswordEncoder;
import com.lidaomeng.servicehi.dao.UserDao;
import com.lidaomeng.servicehi.domain.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl {

//    private static final BCryptPasswordEncoder ENCODER = new BCryptPasswordEncoder();

    /**
     * Todo: 必须加！2019-01-09
     */
    private static final CustomPasswordEncoder ENCODER = new CustomPasswordEncoder();

    @Autowired
    private UserDao userDao;


    public User create(String username, String password) {

        User user = new User();
        user.setUsername(username);
        String hash = ENCODER.encode(password);
        user.setPassword(hash);
        return userDao.save(user);
    }
}

