//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) Text Adventure 2022
//
// Author: Tyler(UnclassedPenguin)
// GitHub: https://github.com/UnclassedPenguin/textgame/
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------


package main

import (
  "fmt"
  "os"
  "math/rand"
  "time"
  "strings"
  "flag"
  "strconv"
)


//-----------------------------------------------------------------------------
// Functions start
//-----------------------------------------------------------------------------

func help() {
  dashLine()
  fmt.Println("Help: ")
  fmt.Println("To move a direction, simply type the direction you want to go.")
  fmt.Println("If there is an item, just try typing its name to pick it up.")
  fmt.Println("Type 'inv' to check whats in your inventory.")
  fmt.Println("Type 'look' to check your surroundings.")
  fmt.Println("Type 'exit' to exit the game.")
  dashLine()
}

//I'd like to make a few phrases here and randomly pick one to say, just
//for some variety.
func cantGo() {
  rn := randNumber(4)
  switch rn {
    case 0:
      s()
      printSlow("I'm sorry, That way is blocked.")
    case 1:
      s()
      printSlow("I'm sorry, you can't go that way.")
    case 2:
      s()
      printSlow("Doesn't seem to be a path in that direction.")
    case 3:
      s()
      printSlow("It's not possible to go that way.")
    default:
      s()
      printSlow("How'd you get here?")
    }
}

// Prints the text character by character.
func printSlow(str string) {
  if slowMode {
    stringSplit := strings.Split(str, "")
    for _, l := range stringSplit {
      if l != " " {
        fmt.Print(l)
        time.Sleep(20 * time.Millisecond)
      } else {
        fmt.Print(l)
        time.Sleep(50 * time.Millisecond)
      }
    }
    fmt.Print("\n")
  } else {
    fmt.Println(str)
  }
}

// Adds an item to the inventory, or just returns the inventory
func inv(item string) []string {
  if item == "?" {
    return inventory
  } else {
    inventory = append(inventory, item)
    return inventory
  }
}

// Check if item is contained in slice (inventory)
func contains(str string, s []string) bool {
  for _, v := range s {
    if v == str {
      return true
    }
  }
  return false
}

// Get index of item in slice
func indexOf(element string, data []string) (int) {
  for k, v := range data {
    if element == v {
      return k
    }
  }
  return -1    //not found.
}

// Just prints a separator
func dashLine() {
  fmt.Println("--------------------------------------------------------------------------------")
}

// Get a single random number
func randNumber(max int) int {
  rand.Seed(time.Now().UnixNano())
  rn := rand.Intn(max)
  return rn
}

//s for give me some (s)pace
func s() {
  fmt.Print("\n")
}

func exit(i int) {
  fmt.Println("Thanks For Playing!")
  os.Exit(i)
}


//-----------------------------------------------------------------------------
// Functions end
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Areas start
//-----------------------------------------------------------------------------

func intro() string{
  //NEEDS WORK. Make more descriptive/fun
  var name string
  s()
  fmt.Println("Welcome to UnclassedPenguin TextAdventure!")
  fmt.Println("Author: Tyler(UnclassedPenguin)")
  fmt.Println("Github: https://github.com/UnclassedPenguin/textgame/")
  s()
  dashLine()
  s()
  fmt.Println("This is the intro. It certaintly needs some work...")
  s()
  dashLine()
  s()
  fmt.Println("What's your name?")
  fmt.Print(" > ")
  fmt.Scan(&name)
  return name
}

func startArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  description1 := "You find yourself in the middle of a forest. There is an axe leaning up against a tree.\nYou can go south or west."
  description2 := "You find yourself in the middle of a forest.\nYou can go south or west."

  i := inv("?")
  if contains("axe", i) {
    s()
    printSlow(description2)
  } else {
    s()
    printSlow(description1)
  }

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "south" {
      s()
      printSlow("You go south.")
      sArea()
    } else if userchoice == "west" {
      s()
      printSlow("You go west.")
      wArea()
    } else if userchoice == "axe" {
      s()
      i := inv("?")
      if contains("axe", i) {
        printSlow("You drop the axe.")
        indexOfAxe := indexOf("axe", inventory)
        if indexOfAxe != -1 {
          inventory = append(inventory[:indexOfAxe], inventory[indexOfAxe+1:]...)
        }
      } else {
        s()
        printSlow("You pick up the axe. It's a nice heavy American felling axe.")
        inv("axe")
      }
    } else if userchoice == "look" {
      s()
      // Checks inventory, if you have axe in your inventory prints description without axe. 
      // Otherwise prints the description that mentions the axe
      i := inv("?")
      if contains("axe", i) {
      s()
        printSlow(description2)
      } else {
      s()
        printSlow(description1)
      }
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
      s()
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option, or try 'help'\n")
    }
  }
}

func wArea() {
  var validDirections = [3]string{"south", "east", "south"}
  var userchoice string
  var description string

  if event["log"] {
    description = "There is a small pond here, fed by a natural spring, with a stream leading out of it to the south.  To the north it looks like there is a path, but with a large log blocking the way.\nYou can go east or south."
  } else {
  description = "There is a small pond here, fed by a natural spring, with a stream leading out of it to the south.  To the north there is a path you cleared, with a large log split in half on either side.\nYou can go north, east or south."
  }

  s()
  printSlow(description)
  pond := false

  for userchoice != validDirections[0] || userchoice != validDirections[1] || userchoice != validDirections[2] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      i := inv("?")
      if contains("axe", i) {
        s()
        printSlow("You use your axe to clear the log and travel north.")
        event["log"] = false
        nwArea()
      } else {
        s()
        printSlow("There is a log blocking the way! If only you had a way to clear it...")
      }
    } else if userchoice == "east" {
      s()
      printSlow("You go east.")
      startArea()
    } else if userchoice == "south" {
      s()
      printSlow("You go south.")
      swArea()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "pond" {
      s()
      pond = true
      printSlow("You look in the pond. There are some small fish swimming around.")
    } else if userchoice == "fish" {
      if pond == true {
        s()
        printSlow("You say hi to the fish, but they don't seem interested in being friends.")
      } else {
        s()
        printSlow("I'm sorry I don't understand 'fish'. Please enter a valid option, or try 'help'")
      }
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func nwArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  description1 := "There are tall trees all around you. The sun gleams through a few of the trees. Is that something shiny behind that tree? It almost looks like it could be a sword...\nYou can only go south."
  description2 := "There are tall trees all around you. The sun gleams through a few of the trees.\nYou can only go south."
  i := inv("?")
  if contains("sword", i) {
    s()
    printSlow(description2)
  } else {
    s()
    printSlow(description1)
  }

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "south" {
      s()
      printSlow("You go south")
      wArea()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "sword" {
      i := inv("?")
      if contains("sword", i) {
        s()
        printSlow("You drop the sword, but why would you want to do that?")
        indexOfSword := indexOf("sword", inventory)
        if indexOfSword != -1 {
          inventory = append(inventory[:indexOfSword], inventory[indexOfSword+1:]...)
        }
      } else {
        s()
        printSlow("You go and look behind the trees. Sure enough, there is a long sword laying in the grass. Who would have dropped this?! You pick it up.")
        inv("sword")
      }
    } else if userchoice == "look" {
      i := inv("?")
      if contains("sword", i) {
        s()
        printSlow(description2)
      } else {
        s()
        printSlow(description1)
      }
    } else if userchoice == "inv" {
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func swArea() {
  var validDirections = [2]string{"north", "east"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "There is a stream running from the north. The banks are covered in rocks. Don't slip!\nYou can go north or east."
  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      printSlow("You go north.")
      wArea()
    } else if userchoice == "east" {
      s()
      printSlow("You go east.")
      sArea()
    } else if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func sArea() {
  var validDirections = [3]string{"north", "east", "west"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description1 := "This is a smelly area. Don't breath too deep! It looks like there is a rope on the ground\nYou can go north, east, or west."
  description2 := "This is a smelly area. Don't breath too deep!\nYou can go north, east, or west."

  i := inv("?")
  if contains("rope", i) {
    s()
    printSlow(description2)
  } else {
    s()
    printSlow(description1)
  }


  for userchoice != validDirections[0] || userchoice != validDirections[1] || userchoice != validDirections[2] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      printSlow("You go north.")
      startArea()
    } else if userchoice == "east" {
      s()
      printSlow("You go east.")
      seArea()
    } else if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      s()
      printSlow("You go west.")
      swArea()
    } else if userchoice == "rope" {
      s()
      i := inv("?")
      if contains("rope", i) {
        printSlow("You drop the rope")
        indexOfRope := indexOf("rope", inventory)
        if indexOfRope != -1 {
          inventory = append(inventory[:indexOfRope], inventory[indexOfRope+1:]...)
        }
      } else {
        s()
        printSlow("You pick up the Rope. It seems heavy enough to support your weight.")
        inv("rope")
      }
    } else if userchoice == "look" {
      s()
      // Checks inventory, if you have rope in your inventory prints description without rope. 
      // Otherwise prints the description that mentions the rope
      i := inv("?")
      if contains("rope", i) {
      s()
        printSlow(description2)
      } else {
      s()
        printSlow(description1)
      }
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func seArea() {
  var validDirections = [2]string{"north", "west"}
  var userchoice string
  count := 0

  //THIS DESCRIPTION NEEDS WORK
  description := "This is SE AREA. There is a cliff to the north. You *might* be able to climb it...\nYou can go west."
  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      i := inv("?")
      // If you have the rope, you are guaranteed to climb the cliff. 
      if contains("rope", i) {
        s()
        printSlow("You use the rope to climb the cliff.")
        eArea()
      // This is if you don't have a rope. Gives you a small change of
      // climbing the cliff. Currently 1/20 chance. Too small? Too big? 
      } else {
        s()
        if count < 5 {
          printSlow("You decide to try free climbing the cliff...")
          s()
          rn := randNumber(20)
          switch rn {
            case 7:
              printSlow("You used your skill to successfully climb the cliff!")
              eArea()
            default:
              count += 1
              printSlow("You failed to climb the cliff, and fell to the bottom! Ouch!")
          }
        } else {
          printSlow("You're a little beat up from failed attempts. Maybe take a break, and try again later.")
        }
      }
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      s()
      printSlow("You go west.")
      sArea()
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func eArea() {
  var validDirections = [2]string{"north", "south"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is E AREA. \nYou can go north or south."
  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      printSlow("You go north.")
      neArea()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "south" {
      s()
      printSlow("You go south.")
      seArea()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

//THIS AREA HAS THE MONSTER. NEED TO WRITE IT OUT
func neArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is NE AREA. \nYou can go west or south."
  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      cantGo()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "south" {
      s()
      printSlow("You go south.")
      eArea()
    } else if userchoice == "west" {
      i := inv("?")
      if contains("sword", i) {
        monsterFight()
      } else {
        printSlow("I don't think you can fight the monster without a sword...")
      }
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func monsterFight() {
  if event["monster"] {
    var userchoice int
    var damage int
    //var stringDamage string
    monsterHealth := 20
    s()
    printSlow("You're gonna fight the monster")

    for monsterHealth > 0 {
      rn := randNumber(10) + 1
      s()
      printSlow("Pick a number between 1 and 10")
      fmt.Print(" > ")
      fmt.Scan(&userchoice)
      if userchoice <= 10 && userchoice >= 0 {
        if userchoice > rn {
          damage = userchoice - rn
        } else {
          damage = rn - userchoice
        }
        monsterHealth -= damage
        s()
        printSlow("You deal " + strconv.Itoa(damage) + " damage to the monster.")
        fmt.Println("--------diag---------------")
        fmt.Println("Random Number:", rn)
        fmt.Println("User choice:", userchoice)
        fmt.Println("Damage:", damage)
        fmt.Println("Monster Health:", monsterHealth)
        fmt.Println("--------diag---------------")
        if monsterHealth <= 0 {
          s()
          printSlow("You Defeated the monster!")
          event["monster"] = false
          nArea()
        }
      } else {
        s()
        printSlow("Invalid Number!")
      }
    }
  } else {
    printSlow("You go west")
    nArea()
  }
}

func nArea() {
  var validDirections = [2]string{"north", "east"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is N AREA. \nYou can go north or east."
  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      printSlow("You go north.")
      exitArea()
    } else if userchoice == "east" {
      s()
      printSlow("You go east.")
      neArea()
    } else if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "look" {
      s()
      printSlow(description)
    } else if userchoice == "inv" {
      s()
      i := inv("?")
      fmt.Println(i)
    } else if userchoice == "help" {
      s()
      help()
    } else if userchoice == "exit" {
      s()
      exit(0)
    } else {
      s()
      printSlow("I'm sorry I don't understand '" + userchoice + "'. Please enter a valid option\n")
    }
  }
}

func exitArea() {
  s()
  printSlow("Congratulations, " + name + "!")
  printSlow("YOU WIIIIIINNNNNNN!")
  printSlow("Hope you had fun! Bye!")
  exit(0)
}


//-----------------------------------------------------------------------------
// Areas end
//-----------------------------------------------------------------------------


//-----------------------------------------------------------------------------
// Global variables start
//-----------------------------------------------------------------------------

var name string
var slowMode bool

// Global inventory
var inventory = []string{}

// Global event tracker
// ask, Does it exist? if so, true. 
var event = map[string]bool {
  "log":true,
  "monster":true,
}

//-----------------------------------------------------------------------------
// Global variables end
//-----------------------------------------------------------------------------

func main() {

  flag.BoolVar(&slowMode, "s", false, "Print out the text in slow mode")
  flag.Parse()

  name = intro()
  s()
  printSlow("Hope you enjoy the game, " + name + ". Good luck!\nIf you get stuck, try 'help'.\n")
  startArea()
}
