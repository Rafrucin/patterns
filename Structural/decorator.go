package patterns

import "fmt"

//Attach additional responsibilities to an object dynamically keeping the same interface.
//Decorators provide a flexible alternative to subclassing for extending functionality.


type IPizza interface {
    getPrice() int
}

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
    return 15
}

type TomatoTopping struct {
    pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 7
}

type CheeseTopping struct {
    pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 10
}

func Decorator() {

    pizza := &VeggieMania{}

    //Add cheese topping
    pizzaWithCheese := &CheeseTopping{
        pizza: pizza,
    }

    //Add tomato topping
    pizzaWithCheeseAndTomato := &TomatoTopping{
        pizza: pizzaWithCheese,
    }

    fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}