package com.example.spring_service1;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;


@RestController
@RequestMapping("/products")
public class ProductController {

    @Autowired
    private ProductService productService;

    @PostMapping
    public Product createProduct(@RequestBody Product product) {
        return productService.createProduct(product);
    }

    @GetMapping("/{id}")
    public Optional<Product> getProduct(@PathVariable Long id) {
        return productService.getProduct(id);
    }
      @GetMapping("/category/{categoryId}")
    public List<Product> getProductsByCategoryId(@PathVariable Long categoryId) {
        return productService.getProductsByCategoryId(categoryId);
    }
}