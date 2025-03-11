package org.emicklei.protocheck;

import org.junit.Test;


public class MessageValidatorTest {

    @Test
    public void test()  throws Exception {
        MessageValidator<org.emicklei.protocheck.pb.Check> mv = new MessageValidator<org.emicklei.protocheck.pb.Check>();
        mv.addFieldChecker(Helpers.createChecker());
        org.emicklei.protocheck.pb.Check c = org.emicklei.protocheck.pb.Check.newBuilder().setCel("1==1").build();
        try {
            mv.validate(c);
        } catch (ValidationException ex) {
            for (CheckError e : ex.getErrors()) {
                System.err.println(e);
            } 
            throw ex;
        }
    }
}