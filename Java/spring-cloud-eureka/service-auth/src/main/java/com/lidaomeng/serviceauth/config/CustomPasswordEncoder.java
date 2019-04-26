package com.lidaomeng.serviceauth.config;

import org.springframework.security.crypto.password.PasswordEncoder;

/**
 * CustomPasswordEncoder
 *
 * @author ldm
 * @date 2019-01-09
 */
public class CustomPasswordEncoder implements PasswordEncoder {

    @Override
    public String encode(CharSequence charSequence) {
        return charSequence.toString();
    }

    @Override
    public boolean matches(CharSequence charSequence, String s) {
        return s.equals(charSequence.toString());
    }
}