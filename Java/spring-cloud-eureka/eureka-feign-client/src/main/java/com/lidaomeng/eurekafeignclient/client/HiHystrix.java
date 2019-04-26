package com.lidaomeng.eurekafeignclient.client;

import org.springframework.stereotype.Component;

@Component
public class HiHystrix implements EurekaClientFeign {
    @Override
    public String sayHiFromClientEureka(String name) {
        return "[Feign] Hi," + name + ",sorry,error!";
    }
}
