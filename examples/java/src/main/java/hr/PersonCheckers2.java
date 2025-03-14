package hr;

import org.emicklei.protocheck.Checker;
import org.emicklei.protocheck.MessageValidator;
import org.emicklei.protocheck.ValidationException;

import dev.cel.common.CelValidationException;
import dev.cel.common.types.SimpleType;
import dev.cel.common.types.StructTypeReference;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelEvaluationException;
import dev.cel.runtime.CelRuntime.Program;

public final class PersonCheckers2 {
    private static MessageValidator<Person> personValidator= new MessageValidator<Person>();

    // never create an instance
    private PersonCheckers2() {
    }

    static { 
        try {
            CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                    .addMessageTypes(Person.getDescriptor())
                    .addVar("this", StructTypeReference.create(Person.getDescriptor().getFullName()))
                    .setStandardEnvironmentEnabled(true)
                    .setResultType(SimpleType.BOOL)
                    .build();
            Program prog = Checker.makeProgram(compiler.compile("size(this.name) > 1").getAst());
            Checker checker = new Checker("name_id", "name must be longer than 1",
                    "size(this.name) > 1",
                    prog, "name", false);

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