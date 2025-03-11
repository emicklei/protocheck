package org.emicklei.protocheck;

import org.junit.Test;

import hr.Person;
import hr.PersonCheckers;

public class CheckerTest {

    @Test
    public void test()  throws Exception {
        Person p = Person.newBuilder().setName("").build();
        try {
            PersonCheckers.validate(p);
        } catch (ValidationException ex) {
            for (CheckError e : ex.getErrors()) {
                System.err.println(e);
            } 
        }
    }
}