package hr;

import org.emicklei.protocheck.Checker;
import org.emicklei.protocheck.MessageValidator;
import org.emicklei.protocheck.ValidationException;

import com.google.common.collect.ImmutableList;

import dev.cel.common.CelValidationException;
import dev.cel.common.types.ProtoMessageTypeProvider;
import dev.cel.common.types.SimpleType;
import dev.cel.common.types.StructTypeReference;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelEvaluationException;
import dev.cel.runtime.CelRuntime.Program;

public final class PersonCheckers {
    private static MessageValidator<Person> personValidator;

    // never create an instance
    private PersonCheckers() {
    }

    static {
        personValidator = new MessageValidator<Person>();
        try {
            // https://github.com/google/cel-java/blob/main/common/src/main/java/dev/cel/common/types/ProtoMessageTypeProvider.java
            ProtoMessageTypeProvider prov = new ProtoMessageTypeProvider(ImmutableList.of(Person.getDescriptor()));
            CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                    .setTypeProvider(prov)
                    .addVar("this", StructTypeReference.create(Person.getDescriptor().getFullName()))
                    // .addVar("this", prov.findType("golang.Person").get())
                    .setStandardEnvironmentEnabled(true)
                    .setResultType(SimpleType.BOOL)
                    .build();
            Program conditionProgram = Checker.makeProgram(compiler.compile("size(this.name) > 1").getAst());
            // Program fieldAccessProgram =
            // Checker.makeProgram(compiler.compile("this.name").getAst());
            Checker checker = new Checker("name_id", "name must be longer than 1",
                    "size(this.name) > 1",
                    conditionProgram, "name", false);
            //checker.setFieldAccess((Person p) -> p.getName());

            personValidator.addFieldChecker(checker);
        } catch (CelEvaluationException ex) {
            System.err.println("CelEvaluationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());

        } catch (CelValidationException ex) {
            System.err.println("CelValidationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());
        }
    }

    public static void validate(Person x) throws ValidationException {
        personValidator.validate(x);
    }
}