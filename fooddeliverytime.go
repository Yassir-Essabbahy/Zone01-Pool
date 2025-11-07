package piscine

type food struct {
	order    string
	preptime int
}

func FoodDeliveryTime(order string) int {
	b := food{order: "burger", preptime: 15}
	c := food{order: "chips", preptime: 10}
	n := food{order: "nuggets", preptime: 12}
	if order == b.order {
		return b.preptime
	} else if order == c.order {
		return c.preptime
	} else if order == n.order {
		return n.preptime
	} else {
		return 404
	}
}
