package org.emicklei.protocheck;

import org.junit.Test;

import pb.BikeOuterClass.Bike;


public class MessageValidatorTest {

    @Test
    public void test()  throws Exception {
        MessageValidator<Bike> mv = new MessageValidator<Bike>();
        mv.addFieldChecker(Helpers.createChecker());
        Bike b = Bike.newBuilder().setBrand("brand").build();
        try {
            mv.validate(b);
        } catch (ValidationException ex) {
            for (CheckError e : ex.getErrors()) {
                System.err.println(e);
            } 
            throw ex;
        }
    }
}