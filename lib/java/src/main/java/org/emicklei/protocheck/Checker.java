package org.emicklei.protocheck;

import dev.cel.runtime.*;

import org.emicklei.protocheck.pb.Check;

import dev.cel.common.CelAbstractSyntaxTree;

public final class Checker {
    private static final CelRuntime CEL_RUNTIME = CelRuntimeFactory.standardCelRuntimeBuilder().build();

    org.emicklei.protocheck.pb.Check check;
    CelRuntime.Program program;
    String fieldName; // for field level checks
    boolean isOptional;

    public Checker(java.lang.String id, java.lang.String fail, java.lang.String cel, CelRuntime.Program program,
            String fieldName, boolean isOptional) {
        this.check = Check.newBuilder().setId(id).setFail(fail).setCel(cel).build();
        this.program = program;
        this.fieldName = fieldName;
        this.isOptional = isOptional;
    }

    public static CelRuntime.Program makeProgram(CelAbstractSyntaxTree ast)
            throws dev.cel.runtime.CelEvaluationException {
        return CEL_RUNTIME.createProgram(ast);
    }
}
