package org.emicklei.protocheck;

import dev.cel.common.CelAbstractSyntaxTree;
import dev.cel.common.CelValidationException;
import dev.cel.common.types.SimpleType;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelEvaluationException;
import dev.cel.runtime.CelRuntime;
import dev.cel.runtime.CelRuntimeFactory;
import java.util.Map;

public class HelloWorld {
  // Construct the compilation and runtime environments.
  // These instances are immutable and thus trivially thread-safe and amenable to caching.
  private static final CelCompiler CEL_COMPILER =
      CelCompilerFactory.standardCelCompilerBuilder().addVar("my_var", SimpleType.STRING).build();
  private static final CelRuntime CEL_RUNTIME =
      CelRuntimeFactory.standardCelRuntimeBuilder().build();

  public void run() throws CelValidationException, CelEvaluationException {
    // Compile the expression into an Abstract Syntax Tree.
    CelAbstractSyntaxTree ast = CEL_COMPILER.compile("my_var + '!'").getAst();

    // Plan an executable program instance.
    CelRuntime.Program program = CEL_RUNTIME.createProgram(ast);

    // Evaluate the program with an input variable.
    String result = (String) program.eval(Map.of("my_var", "Hello World"));
    System.out.println(result); // 'Hello World!'
  }
}