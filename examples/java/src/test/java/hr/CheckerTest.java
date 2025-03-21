package hr;

import org.junit.Test;

import com.google.protobuf.Timestamp;

import hr.Person.Health;

import static org.junit.Assert.fail;

import org.emicklei.protocheck.pb.*;

public class CheckerTest {

    @Test
    public void testValid() throws Exception {
        Person p = createValidPerson();
        java.util.List<CheckError> e = HRProtosCheckers.validate(p);
        if (e != null) {
            fail(e.toString());
        }
    }

    @Test
    public void testInvalidName() throws Exception {
        Person p = createValidPerson();
        p = Person.newBuilder(p).setName("").build();
        java.util.List<CheckError> e = HRProtosCheckers.validate(p);
        if (e == null) {
            fail("Expected error for empty name");
        }
        if (e.size() != 1) {
            fail("Expected one error for empty name");
        }
    }

    private Person createValidPerson() {
        Pet cat = Pet.newBuilder().setKind("cat").setName("harry").build();
        Person.Builder pb = Person.newBuilder()
                .setName("John")
                .setMiddleName("Que")
                .setSurname("Doe")
                .addNicknames("johnny")
                .setEmail("john.doe@mars.com") 
                .setPhone("012-345-7890")
                .setBirthDate(Timestamp.newBuilder()
                        .setSeconds(1234567890)
                        .setNanos(123456789)
                        .build())
                .addPets(cat)
                .setHealth(Health.newBuilder()
                        .setWeight(45)
                        .setAvgHartRate(80));
        pb.putAttributes("color", "black");
        pb.putFavorites("cat", cat);
        Person p = pb.build();
        System.err.println(p);
        return p;
    }
}