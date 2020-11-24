package com.example.demo;

import com.example.demo.model.Counter;
import com.example.demo.repository.CounterRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class MainController {
    @Autowired
    public CounterRepository counterRepository;
    @GetMapping("init")
    public ResponseEntity<String> init(){
        Counter counter = new Counter();
        counter.setCounter(0);
        counterRepository.save(counter);
        return new ResponseEntity<String>("Success",HttpStatus.OK);
    }
    @GetMapping("count")
    public ResponseEntity<String> count(){
            String pass = System.getenv("pass");
            String user = System.getenv("user");
            Counter counter = (Counter) counterRepository.getOne(1l);
            counter.setCounter(counter.getCounter()+1);
            counterRepository.save(counter);
            System.out.println(user + pass);
            return new ResponseEntity<String>("User from env:" + user + " Pass from env: " + pass + "Counter: " + counter.getCounter(), HttpStatus.OK);
        }


}
