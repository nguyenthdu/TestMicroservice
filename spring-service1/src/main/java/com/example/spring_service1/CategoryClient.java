package com.example.spring_service1;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

@FeignClient(name = "category-service", url = "http://localhost:8081")
public interface CategoryClient {

    @GetMapping("/categories/{id}")
    Category getCategory(@PathVariable("id") Long id);
}