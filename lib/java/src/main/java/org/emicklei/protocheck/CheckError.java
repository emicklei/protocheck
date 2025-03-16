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

    // withPath returns a copy with the path set.
    public CheckError withPath( java.lang.String path) {
        CheckError cpy = new CheckError(path, check);
        cpy.path = path;
        return cpy;
    }

    // withParentField returns a copy with the path set.
    public CheckError withParentField(String parent, Object key) {        
        CheckError cpy = new CheckError(path, check);
        String path = parent + "[" +key.toString() + "]";
        if (this.path == "") {
            cpy.path = "." + path;
        } else {
            cpy.path = path.substring(0, path.length() - 2) + "." + cpy.path;
        } 
        return cpy;
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