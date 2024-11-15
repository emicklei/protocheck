package org.protocheck;

import java.util.List;

public class ValidationError extends Exception {
    private List<CheckError> errors;

    public ValidationError(List<CheckError> errors) {
        this.errors = errors;
    }

    @Override
    public String getMessage() {
        StringBuilder b = new StringBuilder();
        String plural = errors.size() > 1 ? "s" : "";
        b.append(String.format("%d error%s occurred:\n", errors.size(), plural));
        for (CheckError each : errors) {
            String fail = each.getFail();
            if (fail == null || fail.isEmpty()) {
                fail = "id=" + each.getId();
            }
            if (each.getErr() == null) {
                b.append(String.format("\t* %s\n", fail));
            } else {
                b.append(String.format("\t* %s: %s\n", fail, each.getErr().getMessage()));
            }
        }
        return b.toString();
    }
}