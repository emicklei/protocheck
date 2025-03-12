package org.emicklei.protocheck;

import java.util.List;

public class ValidationException extends RuntimeException {
    java.util.List<CheckError> errors = new java.util.ArrayList<CheckError>();

    public ValidationException(List<CheckError> errors) {
        this.errors = errors;
    }
    public List<CheckError> getErrors() {
        return errors;
    }
}