package org.emicklei.protocheck;

import org.junit.Test;

import pb.BikeOuterClass.Bike;
import org.emicklei.protocheck.pb.CheckError;


public class MessageValidatorTest {

    @Test
    public void test()  throws Exception {
        MessageValidator<Bike> mv = new MessageValidator<Bike>();
        mv.addFieldChecker(Helpers.createChecker());
        Bike b = Bike.newBuilder().setBrand("brand").build();
        for (CheckError e :  mv.validate(b)) {
            System.err.println(e);
        } 
    }
}