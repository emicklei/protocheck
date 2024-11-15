package org.protocheck;

public class CheckError {
    private String id;
    private String fail;
    private Exception err;

    public CheckError(String id, String fail, Exception err) {
        this.id = id;
        this.fail = fail;
        this.err = err;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getFail() {
        return fail;
    }

    public void setFail(String fail) {
        this.fail = fail;
    }

    public Exception getErr() {
        return err;
    }

    public void setErr(Exception err) {
        this.err = err;
    }
}