package com.lidaomeng.servicehi.dao;

import com.lidaomeng.servicehi.domain.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserDao extends JpaRepository<User, Long> {

    User findByUsername(String username);
}
