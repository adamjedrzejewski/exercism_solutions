class Lasagna
{
    
    public int ExpectedMinutesInOven() => 40;
    
    public int RemainingMinutesInOven(int spentInOven)
        => ExpectedMinutesInOven() - spentInOven;
    
    public int PreparationTimeInMinutes(int layers)
        => 2 * layers;
    
    public int ElapsedTimeInMinutes(int layers, int spentInOven)
        => PreparationTimeInMinutes(layers) + spentInOven;
}