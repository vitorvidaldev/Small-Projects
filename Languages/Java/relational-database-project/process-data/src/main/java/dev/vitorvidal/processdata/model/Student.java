package dev.vitorvidal.processdata.model;

import java.time.LocalDate;
import java.util.List;

import jakarta.persistence.Entity;
import jakarta.persistence.Table;

@Entity
@Table(name = "student")
public class Student {
    private String cpf;
    private String name;
    private String email;
    private List<String> phones;
    private LocalDate birthDate;
    private LocalDate admissionDate;
    private Address address;
    private String registration;
    private List<Course> courses;
    private University university;
}
