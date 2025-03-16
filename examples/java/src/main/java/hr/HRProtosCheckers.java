// Code generated by protoc-gen-protocheck. DO NOT EDIT.

package hr;

import java.util.List;

import org.emicklei.protocheck.Checker;
import org.emicklei.protocheck.CheckError;
import org.emicklei.protocheck.MessageValidator;
import org.emicklei.protocheck.ValidationException;

import dev.cel.common.CelValidationException;
import dev.cel.common.types.SimpleType;
import dev.cel.common.types.StructTypeReference;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelEvaluationException;
import dev.cel.runtime.CelRuntime.Program;
import dev.cel.parser.CelStandardMacro;

public final class HRProtosCheckers {
    private HRProtosCheckers() {}
    private static MessageValidator<Person.Health> person_healthValidator= new MessageValidator<Person.Health>();
    
    static { 
        try {
            CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                    .addMessageTypes(Person.Health.getDescriptor())
                    .addVar("this", StructTypeReference.create(Person.Health.getDescriptor().getFullName()))
                    .setStandardEnvironmentEnabled(true)
                    .setStandardMacros(CelStandardMacro.STANDARD_MACROS)
                    .setResultType(SimpleType.BOOL)
                    .build();
            { // Weight
                String expr = "this.weight > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","weight in kg must be positive",expr,prog,"Weight",false);            
                person_healthValidator.addFieldChecker(checker);
            }
            { // AvgHartRate
                String expr = "this.avg_hart_rate > 0.0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","average heart rate must be positive",expr,prog,"AvgHartRate",false);            
                person_healthValidator.addFieldChecker(checker);
            }
        } catch (CelEvaluationException ex) {
            System.err.println("CelEvaluationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());

        } catch (CelValidationException ex) {
            System.err.println("CelValidationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());
        }
    }

    private static List<CheckError> collectValidationErrors(Person.Health x) {
        List<CheckError> errors = person_healthValidator.validate(x);        
        return errors;
    }

    public static void validate(Person.Health x) throws ValidationException {
        if (x == null) { return; }
        List<CheckError> errors = collectValidationErrors(x);
        if (!errors.isEmpty()) {
            throw new ValidationException(errors);
        }
    }
    private static MessageValidator<Person> personValidator= new MessageValidator<Person>();
    
    static { 
        try {
            CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                    .addMessageTypes(Person.getDescriptor())
                    .addVar("this", StructTypeReference.create(Person.getDescriptor().getFullName()))
                    .setStandardEnvironmentEnabled(true)
                    .setStandardMacros(CelStandardMacro.STANDARD_MACROS)
                    .setResultType(SimpleType.BOOL)
                    .build();	
	        { // person_invariant
                String expr = "size(this.name + this.surname) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("person_invariant","name and surname cannot be empty",expr,prog,"",false);            
                personValidator.addMessageChecker(checker);
            }	
	        { // person_health_weight_invariant
                String expr = "this.health.weight <= 300";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("person_health_weight_invariant","weight cannot be larger than 300",expr,prog,"",false);            
                personValidator.addMessageChecker(checker);
            }
            { // Name
                String expr = "size(this.name) > 1";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","name must be longer than 1",expr,prog,"Name",false);            
                personValidator.addFieldChecker(checker);
            }
            { // MiddleName
                String expr = "size(this.middle_name) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","middle name (if set) cannot be empty",expr,prog,"MiddleName",true);            
                personValidator.addFieldChecker(checker);
            }
            { // Surname
                String expr = "size(this.surname) > 1";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","surname must be longer than 1",expr,prog,"Surname",false);            
                personValidator.addFieldChecker(checker);
            }
            { // BirthDate
                String expr = "this.birth_date.getFullYear() > 2000";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("check_birth_date","[this.birth_date.getFullYear() > 2000] is false",expr,prog,"BirthDate",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Email
                String expr = "this.email.matches('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$')";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("email","email is not valid",expr,prog,"Email",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Phone
                String expr = "this.phone.matches('^[0-9]{3}-[0-9]{3}-[0-9]{4}$')";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","phone is not valid",expr,prog,"Phone",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Nicknames
                String expr = "size(this.nicknames) > 0 && this.nicknames.all(x,size(x)>0)";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","at least one nickname is required",expr,prog,"Nicknames",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Nicknames
                String expr = "this.nicknames.all(x,size(x)>0)";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","nickname cannot be empty",expr,prog,"Nicknames",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Pets
                String expr = "size(this.pets) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","at least one Pet is required",expr,prog,"Pets",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Attributes
                String expr = "size(this.attributes) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","at least one attribute is required",expr,prog,"Attributes",false);            
                personValidator.addFieldChecker(checker);
            }
            { // Favorites
                String expr = "size(this.favorites) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","at least one favorite is required",expr,prog,"Favorites",false);            
                personValidator.addFieldChecker(checker);
            }
        } catch (CelEvaluationException ex) {
            System.err.println("CelEvaluationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());

        } catch (CelValidationException ex) {
            System.err.println("CelValidationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());
        }
    }

    private static List<CheckError> collectValidationErrors(Person x) {
        List<CheckError> errors = personValidator.validate(x);	         
        for (CheckError each : collectValidationErrors(x.getHealth())) { // Health
            errors.add(each.withPath(".Health"));
        }
        // Pets
        List<Pet> list = x.getPetsList();
        for (int i=0;i<list.size();i++) {
            for (CheckError err : collectValidationErrors(list.get(i))) {
                errors.add(err.withParentField("Pets", i));
            }
        }
        // Favorites
        for (java.util.Map.Entry<java.lang.String, Pet> entry : x.getFavoritesMap().entrySet()) {
            for (CheckError err : collectValidationErrors(entry.getValue())) {
                errors.add(err.withParentField("Favorites", entry.getKey()));
            }
        }        
        return errors;
    }

    public static void validate(Person x) throws ValidationException {
        if (x == null) { return; }
        List<CheckError> errors = collectValidationErrors(x);
        if (!errors.isEmpty()) {
            throw new ValidationException(errors);
        }
    }
    private static MessageValidator<Pet> petValidator= new MessageValidator<Pet>();
    
    static { 
        try {
            CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                    .addMessageTypes(Pet.getDescriptor())
                    .addVar("this", StructTypeReference.create(Pet.getDescriptor().getFullName()))
                    .setStandardEnvironmentEnabled(true)
                    .setStandardMacros(CelStandardMacro.STANDARD_MACROS)
                    .setResultType(SimpleType.BOOL)
                    .build();
            { // Kind
                String expr = "this.kind == 'cat' || this.kind == 'dog'";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("pet1","only dog or cat is allowed",expr,prog,"Kind",false);            
                petValidator.addFieldChecker(checker);
            }
            { // Name
                String expr = "size(this.name) > 0";
                Program prog = Checker.makeProgram(compiler.compile(expr).getAst());
                Checker checker = new Checker("","name cannot be empty",expr,prog,"Name",false);            
                petValidator.addFieldChecker(checker);
            }
        } catch (CelEvaluationException ex) {
            System.err.println("CelEvaluationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());

        } catch (CelValidationException ex) {
            System.err.println("CelValidationException: " + ex.getMessage());
            throw new RuntimeException(ex.getCause());
        }
    }

    private static List<CheckError> collectValidationErrors(Pet x) {
        List<CheckError> errors = petValidator.validate(x);        
        return errors;
    }

    public static void validate(Pet x) throws ValidationException {
        if (x == null) { return; }
        List<CheckError> errors = collectValidationErrors(x);
        if (!errors.isEmpty()) {
            throw new ValidationException(errors);
        }
    }
}
