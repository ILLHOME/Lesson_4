package entities

import "math/rand"

type Cat struct {
	Entity
}

func (c *Cat) Eat(amount uint32) {
	if c.houseResources.Food < amount {
		amount = c.houseResources.Food
	}
	c.Satiety += amount * 2
	c.houseResources.CatFoodDecrease(amount)
}

func (c *Cat) Sleep() {
	c.Satiety -= SATIETY_REDUCTION
}

func (c *Cat) TearWall() {
	c.Satiety -= SATIETY_REDUCTION
	c.houseResources.DirtIncrease(DIRTINESS_INCREASE_FROM_TEARING_WALL)
	c.houseResources.Totals.WallTeared++
}

func (c *Cat) CatSimulation() {
	c.Check()

	if c.IsAlive {
		if c.Satiety < 90 && c.houseResources.CatFood > 0 {
			c.Eat(10)
			c.toDo = "Eat"
		} else {
			switch rand.Intn(2) {
			case 0:
				{
					c.Sleep()
					c.toDo = "Sleep"
				}
			case 1:
				{
					c.TearWall()
					c.toDo = "TearWall"
				}
			}
		}
	}
}
