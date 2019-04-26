package com.lidaomeng.eurekaribbonone;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.eureka.EnableEurekaClient;
import org.springframework.cloud.netflix.hystrix.EnableHystrix;
import org.springframework.cloud.netflix.hystrix.dashboard.EnableHystrixDashboard;

/**
 * 负载均衡 Ribbon
 *
 * @author lidaomeng
 */
@EnableEurekaClient
@SpringBootApplication
@EnableHystrix
@EnableHystrixDashboard
public class EurekaRibbonOneApplication {

    public static void main(String[] args) {
        SpringApplication.run(EurekaRibbonOneApplication.class, args);
    }

}
