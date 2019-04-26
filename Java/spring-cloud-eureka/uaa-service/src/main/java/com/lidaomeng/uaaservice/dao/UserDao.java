package com.lidaomeng.uaaservice.dao;


import com.lidaomeng.uaaservice.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;


public interface UserDao extends JpaRepository<User, Long> {

    User findByUsername(String username);
}
