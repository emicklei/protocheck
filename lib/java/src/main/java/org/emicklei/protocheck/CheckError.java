package org.emicklei.protocheck;

public class CheckError {
    java.lang.String path; // path to the field that failed, if empty then the field is part of the root
                          // message
    java.lang.String id; // optional id of the check that failed
    java.lang.String fail; // optional message to display
    java.lang.Exception exception; // set if the check failed due to an error

    public CheckError(java.lang.String path, java.lang.String id, java.lang.String fail) {
        this.path = path;
        this.id = id;
        this.fail = fail;
    }
    public void setException(java.lang.Exception e) {
        this.exception = e;
    }
    public String toString() {
        return "path="+ path + ",id=" + id + ",fail=" + fail + ",err=" + (exception == null ? "" : " (" + exception.getMessage() + ")");
    }
}