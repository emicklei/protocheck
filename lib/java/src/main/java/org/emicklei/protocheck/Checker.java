package org.emicklei.protocheck;

import java.util.function.Function;

import dev.cel.common.CelAbstractSyntaxTree;
import dev.cel.runtime.CelRuntime;
import dev.cel.runtime.CelRuntimeFactory;
import org.emicklei.protocheck.pb.Check;
import org.emicklei.protocheck.pb.CheckError;

public final class Checker {
    private static final CelRuntime CEL_RUNTIME = CelRuntimeFactory.standardCelRuntimeBuilder().build();

    Check check;
    CelRuntime.Program program;
    String fieldName; // for field level checks
    boolean isOptional;

    public Checker(java.lang.String id, java.lang.String fail, java.lang.String cel, CelRuntime.Program program,
            String fieldName, boolean isOptional) {
        this.check = Check.newBuilder()
                .setId(id)
                .setFail(fail)
                .setCel(cel)
                .build();
        this.program = program;
        this.fieldName = fieldName;
        this.isOptional = isOptional;
    }

    public static CelRuntime.Program makeProgram(CelAbstractSyntaxTree ast)
            throws dev.cel.runtime.CelEvaluationException {
        return CEL_RUNTIME.createProgram(ast);
    }

    // withPath returns a copy with the path set.
    public static CheckError withPath(CheckError err, java.lang.String path) {
        if (err.getPath() != "") {
            path = path + "." + err.getPath();
        }
        return CheckError.newBuilder()
                .setPath(path)
                .setFail(err.getFail())
                .setId(err.getId())
                .build();
    }

    // withParentField returns a copy with the path set.
    public static CheckError withParentField(CheckError err, String parent, Object key) {
        CheckError.Builder b = CheckError.newBuilder();
        b.setFail(err.getFail());
        b.setId(err.getId());
        String path = parent + "[" + key.toString() + "]";
        if (err.getPath() == "") {
            b.setPath(path);
        } else {
            b.setPath(path.substring(0, path.length() - 2) + "." + err.getPath());
        }
        return b.build();
    }
}
