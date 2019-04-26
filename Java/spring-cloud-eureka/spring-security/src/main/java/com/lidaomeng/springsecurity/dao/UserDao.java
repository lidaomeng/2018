package com.lidaomeng.springsecurity.dao;

import com.lidaomeng.springsecurity.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;


public interface UserDao extends JpaRepository<User, Long> {

    User findByUsername(String username);
}

