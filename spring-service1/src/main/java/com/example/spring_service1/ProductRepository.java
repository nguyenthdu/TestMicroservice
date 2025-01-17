package com.example.spring_service1;



import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface ProductRepository extends JpaRepository<Product, Long> {
	    List<Product> findByCategoryId(Long categoryId);

}