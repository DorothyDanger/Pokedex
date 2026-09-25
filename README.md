This is my guided learning project from BOOT.DEV on a CLI Pokedex.
You can explore areas, find out which Pokemon are where and try to catch them!
Caught Pokemon can be inspected.


I make usage of PokeAPI for location and Pokemon information.
A cache is used to store: 
  1. location area information to speed up the map and mapb command.
  2. area information about the Pokemon found with the explore command.
  3. Pokemon data so attempting to re-catch a Pokemon is not slow.

There are 8 commands in total:
  exit
    - exits the Pokedex
  help
    - displays a list of commands and what they do
  map
    - displays the next 20 locations
  mapb
    - displays the previous 20 locations
  explore 
    - takes an argument for one location and returns a list of Pokemon found in the location
  catch
    - takes an argument of a Pokemon name and attempts to catch it
  inspect
    - displays the name, weight, height, stats, type of a caught Pokemon
  pokedex
    - lists all Pokemon caught in the Pokedex
