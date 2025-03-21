package org.emicklei.protocheck;

import java.util.List;
import java.util.Map;
import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.Arrays;

import org.emicklei.protocheck.pb.CheckError;

public final class MessageValidator<M extends com.google.protobuf.GeneratedMessage> {
    private List<Checker> fieldCheckers = new ArrayList<Checker>();
    private List<Checker> messageCheckers = new ArrayList<Checker>();

    public void addFieldChecker(Checker checker) {
        fieldCheckers.add(checker);
    }

    public void addMessageChecker(Checker checker) {
        messageCheckers.add(checker);
    }

    public List<CheckError> validate(M message, ValidationOption... options) {
        List<CheckError> errors = new ArrayList<CheckError>();
        Map<String, com.google.protobuf.GeneratedMessage> envMap = Map.of("this", message);
        for (Checker checker : messageCheckers) {
            evalChecker(checker, envMap, errors);
        }
        for (Checker checker : fieldCheckers) {
            if (checker.isEnabledFunc != null) {
                final boolean isEnabled = checker.isEnabledFunc.apply(message);
                if (!isEnabled) {
                    continue;
                }
            }
            final boolean isSet = checker.isSetFunc.apply(message);
            
            // skip unset?
		    if (!isSet && Arrays.asList(options).contains(ValidationOption.FieldsSetOnly)) {
    			continue;
		    }            

            // if not set but optional, skip
            if (!isSet && checker.isOptional) {
                continue;
            }            
            evalChecker(checker, envMap, errors);
        }
        return errors;
    }

    private void evalChecker(
            Checker checker,
            Map<String, com.google.protobuf.GeneratedMessage> envMap,
            List<CheckError> errors) {
        try {
            if (!(boolean) checker.program.eval(envMap)) {
                CheckError err = CheckError.newBuilder()
                        .setFail(checker.check.getFail())
                        .setId(checker.check.getId())
                        .setPath(checker.fieldName)
                        .build();
                errors.add(err);
            }
        } catch (Exception ex) {
            CheckError err = CheckError.newBuilder()
                    .setFail("invalid validation expresssion for "+checker.fieldName)
                    .setId("exception")
                    .setPath(checker.fieldName)
                    .build();
            errors.add(err);
        }
    }
}