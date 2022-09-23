package main

import (
  "fmt"
  "os"
  "math/rand"
  "time"
)

func intro() string{
  //NEEDS WORK. Make more descriptive/fun
  var name string
  fmt.Println("This is the intro. It certaintly needs some work... Whats your name?")
  fmt.Print(" > ")
  fmt.Scan(&name)
  return name
}

func help() {
  fmt.Println("----------------------------------------------------------------")
  fmt.Println("Help: ")
  fmt.Println("To move a direction, simply type the direction you want to go.")
  fmt.Println("If there is an item, just try typing its name to pick it up.")
  fmt.Println("Type 'inv' to check whats in your inventory.")
  fmt.Println("Type 'look' to check your surroundings.")
  fmt.Println("Type 'exit' to exit the game.")
  fmt.Println("----------------------------------------------------------------")
}

func exit(i int) {
  fmt.Println("Thanks For Playing!")
  os.Exit(i)
}

func cantGo() {
  //I'd like to make a few phrases here and randomly pick one to say, just
  //for some variety.
  rand.Seed(time.Now().UnixNano())
  rn := rand.Intn(2)
  switch rn {
    case 0:
      s()
      fmt.Println("I'm sorry, That way is blocked")
    case 1:
      s()
      fmt.Println("I'm sorry, you can't go that way")
    default:
      s()
      fmt.Println("I'm sorry, you can't go that way")
    }
}

func inv(item string) []string {
  if item == "?" {
    return inventory
  } else {
    inventory = append(inventory, item)
    return inventory
  }
}

// Check if item is contained in slice (inventory)
func contains(s []string, str string) bool {
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


//s for give me some (S)pace
func s() {
  fmt.Print("\n")
}

// Area Template
//func Area() {
  //var validDirections = [2]string{"south", "west"}
  //var userchoice string
  //description := ""
  //fmt.Println(description)
  //fmt.Println("You can go", validDirections[0], "or", validDirections[1])

  //for userchoice != validDirections[0] || userchoice != validDirections[1] {
    //fmt.Print(" > ")
    //fmt.Scan(&userchoice)
    //if userchoice == "south" {
      //fmt.Println("You go south")
    //} else if userchoice == "west" {
      //fmt.Println("You go west")
    //} else if userchoice == "north" {
      //fmt.Println("I'm sorry, That way is blocked")
    //} else if userchoice == "east" {
      //fmt.Println("I'm sorry, That way is blocked")
    //} else if userchoice == "look" {
      //fmt.Println(description)
    //} else if userchoice == "inv" {
      //i := inv("?")
      //fmt.Println(i)
    //} else if userchoice == "help" {
      //help()
    //} else if userchoice == "exit" {
      //exit(0)
    //} else {
      //fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    //}
  //}
//}

func startArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  description1 := "You find yourself in the middle of a forest. There is an axe leaning up against a tree.\nYou can go south or west."
  description2 := "You find yourself in the middle of a forest.\nYou can go south or west."

  i := inv("?")
  if contains(i, "axe") {
    fmt.Println(description2)
  } else {
    fmt.Println(description1)
  }

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      s()
      fmt.Println("You go south")
      sArea()
    } else if userchoice == "west" {
      s()
      fmt.Println("You go west")
      wArea()
    } else if userchoice == "north" || userchoice == "east" {
      cantGo()
    } else if userchoice == "axe" {
      s()
      i := inv("?")
      if contains(i, "axe") {
        fmt.Println("You drop the axe")
        indexOfAxe := indexOf("axe", inventory)
        if indexOfAxe != -1 {
          inventory = append(inventory[:indexOfAxe], inventory[indexOfAxe+1:]...)
        }
      } else {
        s()
        fmt.Println("You pick up the axe. It's a nice heavy American logging axe.")
        inv("axe")
      }
    } else if userchoice == "look" {
      s()
      // Checks inventory, if you have axe in your inventory prints description without axe. 
      // Otherwise prints the description that mentions the axe
      i := inv("?")
      if contains(i, "axe") {
      s()
        fmt.Println(description2)
      } else {
      s()
        fmt.Println(description1)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option, or try 'help'\n")
    }
  }
}

func wArea() {
  var validDirections = [3]string{"south", "east", "south"}
  var userchoice string
  description := "There is a small pond here, and a log to the north. You can go east or south."
  fmt.Println(description)
  pond := false

  for userchoice != validDirections[0] || userchoice != validDirections[1] || userchoice != validDirections[2] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      s()
      fmt.Println("You go south.")
    } else if userchoice == "east" {
      s()
      fmt.Println("You go east.")
      startArea()
    } else if userchoice == "north" {
      s()
      i := inv("?")
      if contains(i, "axe") {
        s()
        fmt.Println("You use your axe to clear the log and travel north")
        nwArea()
      } else {
        s()
        fmt.Println("There is a log blocking the way! If only you had a way to clear it...")
      }
    } else if userchoice == "west" {
      s()
      cantGo()
    } else if userchoice == "pond" {
      s()
      pond = true
      fmt.Println("You look in the pond. There are some small fish swimming around.")
    } else if userchoice == "fish" {
      if pond == true {
        s()
        fmt.Println("You say hi to the fish, but they don't seem interested in being friends.")
      } else {
        s()
        fmt.Println("I'm sorry I don't understand fish. Please enter a valid option, or try 'help'")
      }
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option, or try 'help'\n")
    }
  }
}

func nwArea() {
  var validDirections = [2]string{"south", "west"}
  var userchoice string
  description1 := "There are tall trees all around you. The sun gleams through a few of the trees. Is that something shiny behind that tree? It almost looks like it could be a sword...\nYou can only go south."
  description2 := "There are tall trees all around you. The sun gleams through a few of the trees.\nYou can only go south."
  i := inv("?")
  if contains(i, "sword") {
    s()
    fmt.Println(description2)
  } else {
    s()
    fmt.Println(description1)
  }

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      s()
      fmt.Println("You go south")
      wArea()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "north" {
      cantGo()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "sword" {
      i := inv("?")
      if contains(i, "sword") {
        s()
        fmt.Println("You drop the sword, but why would you want to do that?")
        indexOfSword := indexOf("sword", inventory)
        if indexOfSword != -1 {
          inventory = append(inventory[:indexOfSword], inventory[indexOfSword+1:]...)
        }
      } else {
        s()
        fmt.Println("You go and look behind the trees. Sure enough, there is a long sword laying in the grass. Who would have dropped this?! You pick it up.")
        inv("sword")
      }
    } else if userchoice == "look" {
      i := inv("?")
      if contains(i, "sword") {
        s()
        fmt.Println(description2)
      } else {
        s()
        fmt.Println(description1)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option, or try 'help'\n")
    }
  }
}

func swArea() {
  var validDirections = [2]string{"north", "east"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is a rocky area. Don't slip!\nYou can go north or east."
  s()
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "north" {
      s()
      fmt.Println("You go north.")
      wArea()
    } else if userchoice == "east" {
      s()
      fmt.Println("You go east.")
      sArea()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    }
  }
}

func sArea() {
  var validDirections = [3]string{"north", "east", "west"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is a smelly area. Don't breath too deep!\nYou can go north, east, or west."
  s()
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] || userchoice != validDirections[2] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      s()
      fmt.Println("You go west.")
      swArea()
    } else if userchoice == "north" {
      s()
      fmt.Println("You go north.")
      startArea()
    } else if userchoice == "east" {
      s()
      fmt.Println("You go east.")
      seArea()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    }
  }
}

func seArea() {
  var validDirections = [2]string{"north", "west"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is SE AREA. \nYou can go north or west."
  s()
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      s()
      fmt.Println("You go west.")
      sArea()
    } else if userchoice == "north" {
      s()
      fmt.Println("You go north.")
      //NEED TO CREATE THIS AREA
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    }
  }
}

func eArea() {
  var validDirections = [2]string{"north", "south"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is E AREA. \nYou can go north or south."
  s()
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      s()
      fmt.Println("You go south.")
      seArea()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "north" {
      s()
      fmt.Println("You go north.")
      neArea()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
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
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      s()
      fmt.Println("You go south.")
      eArea()
    } else if userchoice == "west" {
      s()
      fmt.Println("MOOOOONSTEEEER")
    } else if userchoice == "north" {
      cantGo()
    } else if userchoice == "east" {
      cantGo()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    }
  }
}

func nArea() {
  var validDirections = [2]string{"north", "east"}
  var userchoice string
  //THIS DESCRIPTION NEEDS WORK
  description := "This is NE AREA. \nYou can go north or east."
  s()
  fmt.Println(description)

  for userchoice != validDirections[0] || userchoice != validDirections[1] {
    fmt.Print(" > ")
    fmt.Scan(&userchoice)
    if userchoice == "south" {
      cantGo()
    } else if userchoice == "west" {
      cantGo()
    } else if userchoice == "north" {
      s()
      fmt.Println("You go north.")
      exitArea()
    } else if userchoice == "east" {
      s()
      fmt.Println("You go east.")
      neArea()
    } else if userchoice == "look" {
      s()
      fmt.Println(description)
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
      fmt.Print("I'm sorry I don't understand ", userchoice, ". Please enter a valid option\n")
    }
  }
}

func exitArea() {
  s()
  fmt.Println("YOU WIIIIIINNNNNNN!")
  fmt.Println("Hope you had fun! Bye!")
}


var inventory = []string{}

func main() {

  // User choices
  //var verb, noun string

  // Inventory
  //var inv = []string{}

  name := intro()
  fmt.Print("Hope you enjoy the game, ", name, ". Good luck!\nIf you get stuck, try 'help'.\n\n")
  startArea()
}
