package com.example.spring_service1;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import java.util.List;
import java.util.Optional;

@Service
public class ProductService {

    @Autowired
    private ProductRepository productRepository;

    @Autowired
    private CategoryClient categoryClient;

    @Autowired
    private KafkaProducer kafkaProducer;

    public Product createProduct(Product product) {
        // Validate category
        categoryClient.getCategory(product.getCategoryId());
        Product savedProduct = productRepository.save(product);
        kafkaProducer.sendMessage(savedProduct);
        return savedProduct;
    }

    public Optional<Product> getProduct(Long id) {
        Optional<Product> product = productRepository.findById(id);
        product.ifPresent(p -> p.setCategoryId(categoryClient.getCategory(p.getCategoryId()).getId()));
        return product;
    }

    public List<Product> getProductsByCategoryId(Long categoryId) {
        return productRepository.findByCategoryId(categoryId);
    }
}