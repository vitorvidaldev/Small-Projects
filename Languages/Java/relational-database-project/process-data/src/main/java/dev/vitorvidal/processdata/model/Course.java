package dev.vitorvidal.processdata.model;

import java.util.List;

import jakarta.persistence.Entity;
import jakarta.persistence.Table;

@Entity
@Table(name = "course")
public class Course {
    private String name;
    private String code;
    private String description;
    private List<Professor> professors;
    private List<Student> students;
    private List<Subject> subjects;
}
