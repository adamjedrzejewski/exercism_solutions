public class Lasagna {
    public int expectedMinutesInOven() {
        return 40;
    }
 
    public int remainingMinutesInOven(int timeInOven) {
        return expectedMinutesInOven() - timeInOven;
    }
 
    public int preparationTimeInMinutes(int layers) {
        return 2 * layers;
    }
 
    public int totalTimeInMinutes(int layers, int timeInOven) {
        return preparationTimeInMinutes(layers) + timeInOven;
    }
}