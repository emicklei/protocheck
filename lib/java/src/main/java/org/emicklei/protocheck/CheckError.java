package org.emicklei.protocheck;

import org.emicklei.protocheck.pb.Check;

public class CheckError {
    java.lang.String path; // path to the field that failed, if empty then the field is part of the root
                           // message
   Check check; // the check that failed
    java.lang.Exception exception; // set if the check failed due to an error

    public CheckError(java.lang.String path, Check check) {
        this.path = path;
        this.check = check;
    }

    public Check getCheck() {
        return check;
    }

    public void setException(java.lang.Exception e) {
        this.exception = e;
    }

    public String toString() {
        return "path=" + path +
                ",id=" + check.getId() +
                ",fail=" + check.getFail() +
                ",err=" + (exception == null ? ""
                        : " (" + exception.getMessage() + ")");

    }
}