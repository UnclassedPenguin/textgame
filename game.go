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
  printSlow("You wake up and become aware. You're not quite sure what happened, and your head\nfeels a bit cloudy. Your eyes slowly start to open...")
  s()
  dashLine()
  s()
  printSlow("What's your name?")
  fmt.Print(" > ")
  fmt.Scan(&name)
  return name
}

func startArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  description1 := "You find yourself in the middle of a forest. The trees surrounding you are tall and the canopy is thick, blocking nearly all the sunlight from coming through. There is an axe leaning up against a tree.\nYou can go south or west."
  description2 := "You find yourself in the middle of a forest. The trees surrounding you are tall and the canopy is thick, blocking nearly all the sunlight from coming through.\nYou can go south or west."

  i := inv("?")
  if contains("axe", i) {
    printSlow(description2)
  } else {
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
        printSlow("You pick up the axe. It's a nice heavy American felling axe.")
        inv("axe")
      }
    } else if userchoice == "look" {
      s()
      // Checks inventory, if you have axe in your inventory prints description without axe. 
      // Otherwise prints the description that mentions the axe
      i := inv("?")
      if contains("axe", i) {
        printSlow(description2)
      } else {
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
    description = "There is a little clearing in the trees here with a small pond, fed by a natural spring, which has a stream leading out of it to the south.  To the north it looks like there is a path, but with a large log blocking the way.\nYou can go east or south."
  } else {
    description = "There is a little clearing in the trees here with a small pond, fed by a natural spring, which has a stream leading out of it to the south.  To the north there is a path you cleared, with a large log split in half on either side.\nYou can go north, east or south."
  }

  s()
  printSlow(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] || userchoice != validDirections[2] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "north" {
      s()
      i := inv("?")
      // if user has axe and log is still there
      if contains("axe", i) && event["log"] == true {
        printSlow("You use your axe to clear the log and travel north.")
        event["log"] = false
        nwArea()
        // if user has axe and log is not there
      } else if contains("axe", i) && event["log"] == false {
        printSlow("You travel north.")
        nwArea()
        // if user has already cleared the log, dropped the axe back in startArea 
        // and comes back. So log not there, and doesn't have axe.
      } else if event["log"] == false {
        printSlow("You travel north.")
        nwArea()
        // if user doesn't have axe and the log is still there
      } else {
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
      printSlow("You look in the pond. There are some small fish swimming around.")
    } else if userchoice == "fish" {
      s()
      printSlow("You say hi to the fish in the pond, but they don't seem interested in being friends.")
    } else if userchoice == "axe" && event["log"] == true {
      i := inv("?")
      // if user has the axe, and the log is still there
      if contains("axe", i){
        s()
        printSlow("You use your axe to clear the log and travel north.")
        event["log"] = false
        nwArea()
      // if user doesn't have axe and log is still there
      } else {
        s()
        printSlow("There is a log blocking the way! If only you had a way to clear it...")
      }
      // if user already cleared log
    } else if userchoice == "axe" && event["log"] == false {
      s()
      printSlow("You already cleared the log, there's no need to use the axe.")
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
      printSlow("You go south.")
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

  description := "The forest has cleared here, and there is a stream running from the north. The banks of the stream are covered in rocks. Don't slip!\nYou can go north or east."

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
    } else if userchoice == "stream" {
      s()
      printSlow("You bend down and use your hands to cup some water and drink it.")
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
  description1 := "You find yourself in clear grasslands. The land starts to rise gently to the east, with mountains in the distance. It looks like there is a rope laying in the grass\nYou can go north, east, or west."
  description2 := "You find yourself in clear grasslands. The land here starts to rise gently, with mountains in the distance.\nYou can go north, east, or west."

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
        printSlow("You drop the rope.")
        indexOfRope := indexOf("rope", inventory)
        if indexOfRope != -1 {
          inventory = append(inventory[:indexOfRope], inventory[indexOfRope+1:]...)
        }
      } else {
        printSlow("You pick up the Rope. It seems heavy enough to support your weight.")
        inv("rope")
      }
    } else if userchoice == "look" {
      s()
      // Checks inventory, if you have rope in your inventory prints description without rope. 
      // Otherwise prints the description that mentions the rope
      i := inv("?")
      if contains("rope", i) {
        printSlow(description2)
      } else {
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

  // count for attempts at climbing cliff without rope.
  count := 0

  description := "The terrain has turned mountainous. There is a cliff to the north. You *might* be able to climb it...\nYou can go west."

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
        printSlow("You use the rope to climb the cliff.")
        eArea()
      // This is if you don't have a rope. Gives you a small change of
      // climbing the cliff. Currently 1/20 chance. Too small? Too big? 
      } else {
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
    } else if userchoice == "rope" {
      i := inv("?")
      if contains("rope", i) {
        s()
        printSlow("You use the rope to climb the cliff.")
        eArea()
      } else {
        s()
        printSlow("What rope?")
      }
    } else if userchoice == "mountain" {
      s()
      printSlow("To the east and south are mountains for miles.")
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

  description := "You are on a high plateau. To the south is a cliff, and further south than that you can see large mountain ranges in the distance. \nYou can go north or south."

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

func neArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string

  description1 := "You enter another forest area. Pines are surrounding you on all sides. There is a path to the west, but when you look closer you see there is a monster standing there, blocking your path. \nYou can go west or south."
  description2 :="You enter another forest area. Pine trees are surrounding you on all sides. The monster you have slain is laying to the side of the path heading west.\nYou can go west or south."

  if event["monster"] {
    s()
    printSlow(description1)
  } else {
    s()
    printSlow(description2)
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
      eArea()
    } else if userchoice == "west" {
      i := inv("?")
      if contains("sword", i) {
        monsterFight()
      } else {
        s()
        printSlow("I don't think you can fight the monster without a sword...")
      }
    } else if userchoice == "sword" {
      i:= inv("?")
      if contains("sword", i) {
        monsterFight()
      } else {
        s()
        printSlow("What sword?")
      }
    } else if userchoice == "look" {
      if event["monster"] {
        s()
        printSlow(description1)
      } else {
        s()
        printSlow(description2)
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
        //fmt.Println("--------diag---------------")
        //fmt.Println("Random Number:", rn)
        //fmt.Println("User choice:", userchoice)
        //fmt.Println("Damage:", damage)
        //fmt.Println("Monster Health:", monsterHealth)
        //fmt.Println("--------diag---------------")
        if monsterHealth <= 0 {
          s()
          printSlow("You Defeated the monster!")
          s()
          printSlow("You go west.")
          event["monster"] = false
          nArea()
        }
      } else {
        s()
        printSlow("Invalid Number!")
      }
    }
  } else {
    s()
    printSlow("You go west.")
    nArea()
  }
}

func nArea() {
  var validDirections = [2]string{"north", "east"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "The forest clears and you find yourself in a field of wildflowers. Purple, blue, yellow and red as far as the eye can see. \nYou can go north or east."
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
    } else if userchoice == "flower" {
      s()
      printSlow("You pick a flower and smell it.")
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
  printSlow("You win!")
  printSlow("Hope you had fun!")
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
  printSlow("Hope you enjoy the game, " + name + ". Good luck!\nIf you get stuck, try 'help'.")
  s()
  dashLine()
  s()
  startArea()
}
