package org.emicklei.protocheck;

import org.emicklei.protocheck.pb.Check;

import com.google.common.collect.ImmutableList;

import dev.cel.common.types.ProtoMessageTypeProvider;
import dev.cel.common.types.SimpleType;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelRuntime.Program;

public class Helpers {

    public static Checker createChecker() throws Exception {
        ProtoMessageTypeProvider prov = new ProtoMessageTypeProvider(ImmutableList.of(Check.getDescriptor()));
        CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                .setTypeProvider(prov)
                .addVar("this", prov.findType("check.Check").get())
                .setStandardEnvironmentEnabled(true)
                .setResultType(SimpleType.BOOL)
                .build();
        Program prog = Checker.makeProgram(compiler.compile("size(this.cel) > 0").getAst());
        Checker checker = new Checker("cel_id", "cel must not be empty", "size(this.cel) > 0", prog, "cel",
                false);
        return checker;
    }
}
