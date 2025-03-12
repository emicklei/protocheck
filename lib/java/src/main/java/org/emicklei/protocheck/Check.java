package org.emicklei.protocheck;

public class Check {
    String id;
    String fail;
    String cel;
    public Check(String id, String fail, String cel) {
        this.id = id;
        this.fail = fail;
        this.cel = cel;
    }
    public String getId() {
        return id;
    }
    public String getFail() {
        return fail;
    }
    public String getCel() {
        return cel;
    }
}
