package com.example.spring_service1;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
public class KafkaProducer {

    private static final String TOPIC = "product_events";

    @Autowired
    private KafkaTemplate<String, Object> kafkaTemplate;

    public void sendMessage(Object message) {
        this.kafkaTemplate.send(TOPIC, message);
    }
}