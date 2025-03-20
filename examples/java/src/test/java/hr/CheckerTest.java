package hr;

import org.junit.Test;

import org.emicklei.protocheck.pb.*;

public class CheckerTest {

    @Test
    public void test() throws Exception {
        Person p = Person.newBuilder()
                .setName("")
                .build();
        for (CheckError e : HRProtosCheckers.validate(p)) {
            System.err.println(e);
        }
    }
}