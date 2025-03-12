package org.emicklei.protocheck;

import com.google.common.collect.ImmutableList;

import dev.cel.common.types.ProtoMessageTypeProvider;
import dev.cel.common.types.SimpleType;
import dev.cel.compiler.CelCompiler;
import dev.cel.compiler.CelCompilerFactory;
import dev.cel.runtime.CelRuntime.Program;
import pb.BikeOuterClass.Bike;

public class Helpers {

    public static Checker createChecker() throws Exception {
        ProtoMessageTypeProvider prov = new ProtoMessageTypeProvider(ImmutableList.of(Bike.getDescriptor()));
        CelCompiler compiler = CelCompilerFactory.standardCelCompilerBuilder()
                .setTypeProvider(prov)
                .addVar("this", prov.findType("tests.Bike").get())
                .setStandardEnvironmentEnabled(true)
                .setResultType(SimpleType.BOOL)
                .build();
        Program prog = Checker.makeProgram(compiler.compile("size(this.brand) > 0").getAst());
        Checker checker = new Checker("brand_id", "brand must not be empty", "size(this.brand) > 0", prog, "cel",
                false);
        return checker;
    }
}
