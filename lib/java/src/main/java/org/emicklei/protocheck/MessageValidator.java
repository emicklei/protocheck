package org.emicklei.protocheck;

import java.util.List;
import java.util.Map;
import java.util.ArrayList;

public final class MessageValidator<M extends com.google.protobuf.GeneratedMessage> {
    private List<Checker> fieldCheckers = new ArrayList<Checker>();
    private List<Checker> messageCheckers = new ArrayList<Checker>();

    public void addFieldChecker(Checker checker) {
        fieldCheckers.add(checker);
    }

    public void addMessageChecker(Checker checker) {
        messageCheckers.add(checker);
    }

    public void validate(M message, ValidationOption... options) throws ValidationException {
        List<CheckError> errors = new ArrayList<CheckError>();
        Map<String, com.google.protobuf.GeneratedMessage> envMap = Map.of("this", message);
        for (Checker checker : messageCheckers) {
            evalChecker(checker, envMap, errors);
        }        
        for (Checker checker : fieldCheckers) {
            evalChecker(checker, envMap, errors);
        }
        if (errors.size() > 0) {
            throw new ValidationException(errors);
        }
    }

    private void evalChecker(
            Checker checker,
            Map<String, com.google.protobuf.GeneratedMessage> envMap,
            List<CheckError> errors) {
        try {
            if (!(boolean) checker.program.eval(envMap)) {
                CheckError err = new CheckError(checker.fieldName, checker.check);
                errors.add(err);
            }
        } catch (Exception ex) {
            CheckError err = new CheckError(checker.fieldName, checker.check);
            err.setException(ex);
            errors.add(err);
        }
    }
}