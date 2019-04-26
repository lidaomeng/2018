package com.lidaomeng.serviceauth.service;

import com.lidaomeng.serviceauth.config.CustomPasswordEncoder;
import com.lidaomeng.serviceauth.dao.UserDao;
import com.lidaomeng.serviceauth.domain.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl implements UserService {

    //private final Logger log = LoggerFactory.getLogger(getClass());

    //private static final BCryptPasswordEncoder ENCODER = new BCryptPasswordEncoder();

    /**
     * Todo: 必须加！2019-01-09
     */
    private static final CustomPasswordEncoder ENCODER = new CustomPasswordEncoder();

    @Autowired
    private UserDao userDao;

    @Override
    public void create(User user) {

        User existing = userDao.findByUsername(user.getUsername());
        //Assert.isNull(existing, "user already exists: " + user.getUsername());

        String hash = ENCODER.encode(user.getPassword());
        user.setPassword(hash);
        userDao.save(user);

    }
}

