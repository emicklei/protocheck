package protocheck;

ublic class Checker {
    private Check check;
    private CelProgram program;

    public Checker(Check check, CelProgram program) {
        this.check = check;
        this.program = program;
    }

    public static Checker newChecker(String id, String message, String cel, CelProgram program) {
        Check check = new Check(id, message, cel);
        return new Checker(check, program);
    }

    // Getters and setters for check and program
    public Check getCheck() {
        return check;
    }

    public void setCheck(Check check) {
        this.check = check;
    }

    public CelProgram getProgram() {
        return program;
    }

    public void setProgram(CelProgram program) {
        this.program = program;
    }
}