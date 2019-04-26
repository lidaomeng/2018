package com.lidaomeng.springsecurity.service;

import com.lidaomeng.springsecurity.entity.Blog;

import java.util.List;

public interface IBlogService {
    List<Blog> getBlogs();
    void deleteBlog(long id);
}

