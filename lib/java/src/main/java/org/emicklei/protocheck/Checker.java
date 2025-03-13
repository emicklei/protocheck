package org.emicklei.protocheck;

import dev.cel.runtime.*;

import java.util.function.Function;

import dev.cel.common.CelAbstractSyntaxTree;

public final class Checker {
    private static final CelRuntime CEL_RUNTIME = CelRuntimeFactory.standardCelRuntimeBuilder().build();

    Check check;
    CelRuntime.Program program;
    String fieldName; // for field level checks
    boolean isOptional;
    Function<?,Object> fieldAccess;

    public Checker(java.lang.String id, java.lang.String fail, java.lang.String cel, CelRuntime.Program program,
            String fieldName, boolean isOptional) {
        this.check = new Check(id,fail,cel);
        this.program = program;
        this.fieldName = fieldName;
        this.isOptional = isOptional;
    }

    public void setFieldAccess(Function<Object,Object> fieldAccess) {
        this.fieldAccess = fieldAccess;
    }

    public static CelRuntime.Program makeProgram(CelAbstractSyntaxTree ast)
            throws dev.cel.runtime.CelEvaluationException {
        return CEL_RUNTIME.createProgram(ast);
    }
}
