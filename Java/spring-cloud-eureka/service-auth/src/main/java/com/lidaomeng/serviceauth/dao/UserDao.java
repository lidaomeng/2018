package com.lidaomeng.serviceauth.dao;

import com.lidaomeng.serviceauth.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserDao extends JpaRepository<User, Long> {

    User findByUsername(String username);
}
