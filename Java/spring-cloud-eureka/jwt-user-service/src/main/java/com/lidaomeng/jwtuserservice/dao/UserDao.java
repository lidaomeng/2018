package com.lidaomeng.jwtuserservice.dao;


import com.lidaomeng.jwtuserservice.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserDao extends JpaRepository<User, Long> {

	User findByUsername(String username);

}
