package hr;

import org.junit.Test;

import org.emicklei.protocheck.*;

public class CheckerTest {

    @Test
    public void test()  throws Exception {
        Person p = Person.newBuilder().setName("").build();
        try {
            HRProtosCheckers.validate(p);
        } catch (ValidationException ex) {
            System.err.println(ex.getMessage());
            for (CheckError e : ex.getErrors()) {
                System.err.println(e);
            } 
        }
    }
}