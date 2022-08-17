export function age(planet: string, seconds: number): number {
    var denom: number = 1.0;
    switch (planet) {
      case 'mercury':
        denom = 0.2408467;
        break;
      case 'venus':
        denom = 0.61519726;
        break;
      case 'earth':
        denom = 1.0;
        break;
      case 'mars':
        denom = 1.8808158;
        break;
      case 'jupiter':
        denom = 11.862615;
        break;
      case 'saturn':
        denom = 29.447498;
        break;
      case 'uranus':
        denom = 84.016846;
        break;
      case 'neptune':
        denom = 164.79132;
        break;
    }
    return Math.round(seconds / denom / 31557600 * 100) / 100;
  }
  