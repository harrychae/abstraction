# Install 
 go get github.com/harrychae/abstraction
 go get github.com/harrychae/abstraction/game

# Sample
{
	package main

	import (
		"github.com/harrychae/abstraction"
		composition "github.com/harrychae/abstraction/game"
	)

	func main() {
		hillary := new(abstraction.Hillary)
		trump := new(abstraction.Trump)

		h := abstraction.Campaign{hillary}
		t := abstraction.Campaign{trump}

		h.Slogan() // "Stronger together."
		t.Slogan() // "Make America great again."

		abstraction.SaySlogan(hillary) // "Stronger together."
		abstraction.SaySlogan(trump)   // "Make America great again."

		abstraction.SaySlogan(h) // "Stronger together."
		abstraction.SaySlogan(t) // "Make America great again."

		amy := composition.Amy{
			Human: composition.Human{
				FirstName: "Amy",
				LastName:  "Chen",
				CanSwim:   true,
			},
		}

		alan := composition.Alan{
			Human: composition.Human{
				FirstName: "Alan",
				LastName:  "Chen",
				CanSwim:   false,
			},
		}

		amy.Name()  // "Hello! My name is Amy Chen"
		amy.Swim()  // "I can swim!"
		alan.Name() // "Hello! My name is Alan Chen"
		alan.Swim() // "I can't swim"
	}
}
