package composition 

import "fmt"

type Human struct {
    FirstName   string
    LastName    string
    CanSwim     bool
}

// Amy는 Human 타입으로 임베딩 되어있으며 따라서 Human의 메서드 집합에 속하는 메서드를 실행할 수 있음
type Amy struct {
    Human
}

// Alan 또한 Human 타입으로 임베딩 되어있음
type Alan struct {
    Human
}

func (h *Human) Name() {
    
    fmt.Printf("Hello! My name is %v %v", h.FirstName, h.LastName)
}

func (h *Human) Swim() {
    
    if h.CanSwim == true {
        fmt.Println("I can swim!")
    } else {
        fmt.Println("I can not swim.")
    }
}
