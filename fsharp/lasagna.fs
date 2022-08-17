module LuciansLusciousLasagna

let layerPreparationTime = 2

let expectedMinutesInOven = 40

let remainingMinutesInOven inOven = expectedMinutesInOven - inOven

let preparationTimeInMinutes layers = layers * layerPreparationTime

let elapsedTimeInMinutes layers timeInOven = preparationTimeInMinutes layers + timeInOven
