package hr;

import dev.cel.common.CelAbstractSyntaxTree;
import dev.cel.common.CelValidationException;
import dev.cel.common.types.SimpleType;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelEvaluationException;
import dev.cel.runtime.CelRuntime;
import dev.cel.runtime.CelRuntimeFactory;
import java.util.Map;

import org.junit.Test;

public class RegexTest {
  // Construct the compilation and runtime environments.
  // These instances are immutable and thus trivially thread-safe and amenable to
  // caching.
  private static final CelCompiler CEL_COMPILER = CelCompilerFactory.standardCelCompilerBuilder()
      .addVar("my_var", SimpleType.STRING).build();
  private static final CelRuntime CEL_RUNTIME = CelRuntimeFactory.standardCelRuntimeBuilder().build();

  @Test
  public void test() throws CelValidationException, CelEvaluationException {
    // Compile the expression into an Abstract Syntax Tree.
    CelAbstractSyntaxTree ast = CEL_COMPILER.compile("my_var.matches('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$')").getAst();

    // Plan an executable program instance.
    CelRuntime.Program program = CEL_RUNTIME.createProgram(ast);

    // Evaluate the program with an input variable.
    Boolean result = (Boolean) program.eval(Map.of("my_var", "john.doe@mars.com"));
    if (!result) {
      throw new CelEvaluationException("Expected true");
    }
  }
}