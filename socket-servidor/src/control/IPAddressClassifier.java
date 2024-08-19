package control;

public class IPAddressClassifier {

    private static final String IP_ERROR = "Invalid IP address\n";

    static class IPClass {
        String ipClass;
        int min;
        int max;

        IPClass(String ipClass, int min, int max) {
            this.ipClass = ipClass;
            this.min = min;
            this.max = max;
        }
    }


    public static String addressIP(String message) {
        // Define the classes and their ranges
        IPClass[] classes = new IPClass[]{
                new IPClass("A", 0, 127),
                new IPClass("B", 128, 191),
                new IPClass("C", 192, 223),
                new IPClass("D", 224, 239)
        };

        String[] ipAddress = message.split("\\.");
        if (ipAddress.length <3){
            return IP_ERROR;
        }

        int firstByte;
        try {
            firstByte = Integer.parseInt(ipAddress[0]);
        } catch (NumberFormatException e) {
            return IP_ERROR;
        }


        for (IPClass ipClass : classes){
            if (firstByte >= ipClass.min && firstByte <= ipClass.max){
                return String.format("classe: %s", ipClass.ipClass);
            }
        }

        return IP_ERROR;
    }
}
