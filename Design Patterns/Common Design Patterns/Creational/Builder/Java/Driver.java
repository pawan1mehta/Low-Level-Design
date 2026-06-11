package Builder.Java;

class Computer {
    
    private String CPU;
    private String RAM;
    private Boolean isGraphicsCardEnabled;
    private Boolean isBluetoothEnabled;

    public Computer(ComputerBuilder builder) {
        this.CPU = builder.CPU;
        this.RAM = builder.RAM;
        this.isGraphicsCardEnabled = builder.isGraphicsCardEnabled;
        this.isBluetoothEnabled = builder.isBluetoothEnabled;
    }

    public static class ComputerBuilder {
        private String CPU;
        private String RAM;
        private Boolean isGraphicsCardEnabled;
        private Boolean isBluetoothEnabled;


        public ComputerBuilder setCPU(String CPU) {
            this.CPU = CPU;
            return this;
        }

        public ComputerBuilder setRAM(String RAM) {
            this.RAM = RAM;
            return this;
        }

        public ComputerBuilder setGraphicsCardEnabled(Boolean isGraphicsCardEnabled) {
            this.isGraphicsCardEnabled = isGraphicsCardEnabled;
            return this;
        }

        public ComputerBuilder setsBluetoothEnabled(Boolean isBluetoothEnabled) {
            this.isBluetoothEnabled = isBluetoothEnabled;
            return this;
        }

        public Computer build() {
            return new Computer(this);
        }
    }

    public String toString() {
        return "Computer [CPU=" + CPU + ", RAM=" + RAM + ", isGraphicsCardEnabled=" + isGraphicsCardEnabled
                + ", isBluetoothEnabled=" + isBluetoothEnabled + "]";
    }
}

class Driver {

    public static void main(String[] args) {
        Computer computer = new Computer.ComputerBuilder()
        .setCPU("Intel i7")
        .setRAM("16GB")
        .setGraphicsCardEnabled(true)
        .setsBluetoothEnabled(true)
        .build();

        System.out.println(computer);
    }
}