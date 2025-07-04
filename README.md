
# Dungeons & Dragons 5e Encounter Generator Reforged
A CLI tool that helps DMs generate random encounters for their players. Using the 2024 experience budget rules!

I have made a tool similiar to this in Python, that is why I named this one Reforged! [Link here](https://github.com/kyoukyuubi/dnd-encounter-gen)

This is a fan project, I am in now way affiliated with Wizards of the coast or D&D Beyond. 


## Installation

Head over to the [releases](https://github.com/kyoukyuubi/dnd-encounter-gen-reforged/releases/tag/Release) page and download the latest release. Make sure that the exe and the json folder and filers are inside the same folder. It will throw and error which can useually be found inside json/logs folder. 

If you don't see a logs folder, do not worry. The Program will make that on it's on, if an error has occurred.
## Usage/Examples
Open the *dnd-encounter-gen.exe* making sure the included json folder is in the same place the exe is and run something like this:

```
level 5
party-size 6
generate
```

Which could give you something like so:
```
Encounter Creatures:
===================
Name:   Bandit
Type:   Humanoid
Book:   Monster Manual 2024 (page 27)
Amount: 1
--------------------------------
Name:   Awakened Shrub
Type:   Plant
Book:   Monster Manual 2024 (page 23)
Amount: 1
--------------------------------
Name:   Pixie
Type:   Fey
Book:   Monster Manual 2024 (page 244)
Amount: 1
--------------------------------
Name:   Unicorn
Type:   Celestial
Book:   Monster Manual 2024 (page 313)
Amount: 1
--------------------------------
Name:   Dire Wolf
Type:   Beast
Book:   Monster Manual 2024 (page 352)
Amount: 1
--------------------------------
Name:   Swarm of Rats
Type:   Beast
Book:   Monster Manual 2024 (page 370)
Amount: 1
--------------------------------
Total Creatures: 6
Total XP Budget Used: 2135
```

Use the "help" inside the software to see what commands are available. You can even use the "list" command to see what you can set the different filters to.

Make sure to close the software with the "exit" command, since it saves your filters in a json within the json/settings folder.
## Features

- Party Size
- Level 1-20
- Difficulty
- Different Planes of Existence
- Different Environments
- Minimum Experience
- Maximum Creatures
- User defined sources used for generation
- Supports Homebrew Creatures, Planes, Enviroments and Difficulties
    - User can add their own by editing the included Json files
## Editing/Adding creature, plane, type, source or difficulty

if you want to add your own homebrew things to the generator all you have to do is edit the json! I will go into detail how to do so.

note: Ignore the code errors on some of the code snippets on this page. JSON files do not support comments, but I added them to help make sense what goes where!

### Adding/editing Creatures    

If you want to add your own creature to one of the included books, open the json in any editor and add to the json, the structure are as follows:
```json
{
    "Creatures": [
      {
        "Name": "Leviathan",
        "Type": "Elemental",
        "Exp": 25000,
        "Environment": ["Underwater", "Coastal"],
        "Plane": [],
        "Book": "Mordenkainen's Tome of Foes",
        "Page": 198
      }
    ]
  }
```

the "Creatures" are important to have at the top! If you just want to add 1 or more creatures, do it like so:
```json
{
    "Creatures": [
      {
        "Name": "Leviathan",
        "Type": "Elemental",
        "Exp": 25000,
        "Environment": ["Underwater", "Coastal"],
        "Plane": [],
        "Book": "Mordenkainen's Tome of Foes",
        "Page": 198
      },
      {
        "Name": "My Custom Creature", // the name
        "Type": "Creature Type (make sure type is in the types.json)", // the type
        "Exp": 25000, // the exp! whole numbers only
        "Environment": ["Underwater", "Coastal"], // the Environments/Habitats, make sure it looks like this. leave it as [] if if the creature is in any environment.
        "Plane": [], // planes here, looks like the same as Environment, leave it as [] if the creature is in any plane.
        "Book": "My own book", // which book, you can have multiple books in 1 json, but try to split them to multiple files for the source filter to work.
        "Page": 198 // which page, is displayed if the creature is picked, whole numbers only
      }
    ]
  }
```

If you have your own book of creatures, you can make a new json file. Just save the name in the source.json read on for how to do so.

Make sure the Environment, Plane and Type matches what you put in the environments.json, planes.json and types.json files. Else it won't select properly.

### Adding/editing Sources

The layout of the sources.json is as follows:
```json
{
    "sources": [
        {
            "name": "Monster Manual 2024",
            "filename": "mm2024_creatures"
        },
        {
            "name": "Mordenkainens Tome of Foes",
            "filename": "mtof_creatures"
        }
    ]
}
```

If you want to add your own book, make sure you have a file. Let's say that you have one called `custom.json`, do the following:
```json
{
    "sources": [
        {
            "name": "Monster Manual 2024",
            "filename": "mm2024_creatures"
        },
        {
            "name": "Mordenkainens Tome of Foes",
            "filename": "mtof_creatures"
        },
        {
            "name": "My Custom Book", // the name you need to type to select this filter
            "filename": "custom" // the file it looks for, without the .json
        }
    ]
}
```

### Adding/editing Environments

The layout of the environments.json is as follows(this is a snippet of the whole file):

```json
{
    "environments": [
        {
            "name": "Arctic"
        },
        {
            "name": "Coastal"
        }
    ]
}
```

If you want to add your own environment do the following:

```json
{
    "environments": [
        {
            "name": "Arctic"
        },
        {
            "name": "Coastal"
        },
        {
            "name": "My Own Environment" // the name for the environment
        }
    ]
}
```

### Adding/editing Types

It works the same was as environments does!

### Adding/editing Planes

The layout of the planes.json is as follows (here is a snippet of it):

```json
{
    "planes": [
        {
            "name": "Feywild",
            "category": "Material Realms"
        },
        {
            "name": "Material Plane",
            "category": "Material Realms"
        },
        {
            "name": "Shadowfell",
            "category": "Material Realms"
        },
        {
            "name": "Hades",
            "category": "Outer Planes",
            "sub-category": "Lower Planes"
        }
    ]
}
```

If you want to add your own, do as follows: 

```json
{
    "planes": [
        {
            "name": "Feywild",
            "category": "Material Realms"
        },
        {
            "name": "Material Plane",
            "category": "Material Realms"
        },
        {
            "name": "Shadowfell",
            "category": "Material Realms"
        },
        {
            "name": "Hades",
            "category": "Outer Planes",
            "sub-category": "Lower Planes"
        },
        {
            "name": "Custom Plane", // the name of the plane
            "category": "My Own Category", // the category of the plane
            "sub-category": "My Own Sub Category" // the sub category for the plane. Note: This one is optional
        },
        {
            "name": "Custom Plane No Sub Category",
            "category": "My Own Category", // the category of the plane, this one does not have a sub category.
        }
    ]
}
```

### Adding/editing the level table

If you want to edit the leveling table to add your own difficulty, it might be a little tricky. The `experience_table.json` has the following structure:

```json
{
    "1": { "Low": 50, "Moderate": 75, "High": 100 },
    "2": { "Low": 100, "Moderate": 150, "High": 200 },
    "3": { "Low": 150, "Moderate": 225, "High": 400 },
    "4": { "Low": 250, "Moderate": 375, "High": 500 },
    "5": { "Low": 500, "Moderate": 750, "High": 1100 },
    "6": { "Low": 600, "Moderate": 1000, "High": 1400 },
    "7": { "Low": 750, "Moderate": 1300, "High": 1700 },
    "8": { "Low": 1000, "Moderate": 1700, "High": 2100 },
    "9": { "Low": 1300, "Moderate": 2000, "High": 2600 },
    "10": { "Low": 1600, "Moderate": 2300, "High": 3100 },
    "11": { "Low": 1900, "Moderate": 2900, "High": 4100 },
    "12": { "Low": 2200, "Moderate": 3700, "High": 4700 },
    "13": { "Low": 2600, "Moderate": 4200, "High": 5400 },
    "14": { "Low": 2900, "Moderate": 4900, "High": 6200 },
    "15": { "Low": 3300, "Moderate": 5400, "High": 7800 },
    "16": { "Low": 3800, "Moderate": 6100, "High": 9800 },
    "17": { "Low": 4500, "Moderate": 7200, "High": 11700 },
    "18": { "Low": 5000, "Moderate": 8700, "High": 14200 },
    "19": { "Low": 5500, "Moderate": 10700, "High": 17200 },
    "20": { "Low": 6400, "Moderate": 13200, "High": 22000 }
}
```

So to add your own difficulty, you need to add it a few times. Like so

```json
{
    "1": { "Low": 50, "Moderate": 75, "High": 100, "Super Easy": 10},
    "2": { "Low": 100, "Moderate": 150, "High": 200, "Super Easy": 20 },
    "3": { "Low": 150, "Moderate": 225, "High": 400, "Super Easy": 30 },
    "4": { "Low": 250, "Moderate": 375, "High": 500, "Super Easy": 40 },
    "5": { "Low": 500, "Moderate": 750, "High": 1100, "Super Easy": 50 },
    "6": { "Low": 600, "Moderate": 1000, "High": 1400, "Super Easy": 60 },
    "7": { "Low": 750, "Moderate": 1300, "High": 1700, "Super Easy": 70 },
    "8": { "Low": 1000, "Moderate": 1700, "High": 2100, "Super Easy": 80 },
    "9": { "Low": 1300, "Moderate": 2000, "High": 2600, "Super Easy": 90 },
    "10": { "Low": 1600, "Moderate": 2300, "High": 3100, "Super Easy": 100 },
    "11": { "Low": 1900, "Moderate": 2900, "High": 4100, "Super Easy": 110 },
    "12": { "Low": 2200, "Moderate": 3700, "High": 4700, "Super Easy": 120 },
    "13": { "Low": 2600, "Moderate": 4200, "High": 5400, "Super Easy": 130 },
    "14": { "Low": 2900, "Moderate": 4900, "High": 6200, "Super Easy": 140 },
    "15": { "Low": 3300, "Moderate": 5400, "High": 7800, "Super Easy": 150 },
    "16": { "Low": 3800, "Moderate": 6100, "High": 9800, "Super Easy": 160 },
    "17": { "Low": 4500, "Moderate": 7200, "High": 11700, "Super Easy": 170 },
    "18": { "Low": 5000, "Moderate": 8700, "High": 14200, "Super Easy": 180 },
    "19": { "Low": 5500, "Moderate": 10700, "High": 17200, "Super Easy": 190 },
    "20": { "Low": 6400, "Moderate": 13200, "High": 22000, "Super Easy": 1200 }
}
```

As you can see, I added `Super Easy` as a difficulty. To attempt to explain the structure a bit more it's like this:

```json
{
    // "level": {"difficulty name": base Exp as a whole number}
    "1": { "Low": 50, "Moderate": 75, "High": 100, "Super Easy": 10}
}
```

Other than that it follows the normal json syntax! 

## Support

If something happens you want to report/need help with. Please open an [issue](https://github.com/kyoukyuubi/dnd-encounter-gen-reforged/issues).
## Feedback

If you want to leave any feedback, you can do so in [discussions](https://github.com/kyoukyuubi/dnd-encounter-gen-reforged/discussions)! If you have any questions feel free to do so in there. There is a category for it.


## Credits

Icon is made using Canva!