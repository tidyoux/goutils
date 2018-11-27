package config

import (
	"fmt"
	"testing"
)

const (
	dataDesc = "this is the data description."
)

func assert(t *testing.T, ok bool, msg string) {
	if !ok {
		t.Fatal(msg)
	}
}

func TestPair(t *testing.T) {
	role := NewPair("role").
		Add(NewString(dataDesc)).
		Add(NewPair("name").Add(NewString("role1"))).
		Add(NewPair("level").Add(NewNumber("1"))).
		Add(NewPair("movable").Add(NewBool(true))).
		Add(NewPair("tags").Add(NewString("robot")).Add(NewString("npc"))).
		Add(NewPair("items").
			Add(NewPair("").
				Add(NewPair("name").Add(NewString("item\"1\""))).
				Add(NewPair("price").Add(NewNumber("1.5")))).
			Add(NewPair("").
				Add(NewPair("name").Add(NewString("'item'2"))).
				Add(NewPair("price").Add(NewNumber("0.3")))))

	fmt.Println(role.Format(0))

	desc, _ := role.String("", 0)
	assert(t, desc == dataDesc, "data description not equal")

	name, _ := role.String("", "name")
	assert(t, name == "role1", "role name not equal")

	level, _ := role.Int64(0, "level")
	assert(t, level == 1, "role level not equal")

	movable, _ := role.Bool(false, "movable")
	assert(t, movable == true, "role movable not equal")

	tag1, _ := role.String("", "tags", 0)
	assert(t, tag1 == "robot", "role tag1 not equal")

	tag2, _ := role.String("", "tags", 1)
	assert(t, tag2 == "npc", "role tag2 not equal")

	item1Name, _ := role.String("", "items", 0, "name")
	assert(t, item1Name == "item\"1\"", "role item1 name not equal")

	item1Price, _ := role.Float64(0, "items", 0, "price")
	assert(t, item1Price == 1.5, "role item1 price not equal")

	item2Name, _ := role.String("", "items", 1, "name")
	assert(t, item2Name == "'item'2", "role item2 name not equal")

	item2Price, _ := role.Float64(0, "items", 1, "price")
	assert(t, item2Price == 0.3, "role item2 price not equal")
}

const data = `(role
	"this is the data description."

	; base data
	(name "role1")
	(level 1)
	(movable true)
	(tags "robot" "npc")

	(items ; item data
		((name "item\"1\"")
			(price 1.5))
		((name "'item'2")
			(price 0.3))
	)
)`

func TestParser(t *testing.T) {
	c := New()
	err := c.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(c.Format(0))

	role, err := c.Pair.Pair("role")
	if err != nil {
		t.Fatal(err)
	}

	desc, _ := role.String("", 0)
	assert(t, desc == dataDesc, "data description not equal")

	name, _ := role.String("", "name")
	assert(t, name == "role1", "role name not equal")

	level, _ := role.Int64(0, "level")
	assert(t, level == 1, "role level not equal")

	movable, _ := role.Bool(false, "movable")
	assert(t, movable == true, "role movable not equal")

	tag1, _ := role.String("", "tags", 0)
	assert(t, tag1 == "robot", "role tag1 not equal")

	tag2, _ := role.String("", "tags", 1)
	assert(t, tag2 == "npc", "role tag2 not equal")

	item1Name, _ := role.String("", "items", 0, "name")
	assert(t, item1Name == "item\"1\"", "role item1 name not equal")

	item1Price, _ := role.Float64(0, "items", 0, "price")
	assert(t, item1Price == 1.5, "role item1 price not equal")

	item2Name, _ := role.String("", "items", 1, "name")
	assert(t, item2Name == "'item'2", "role item2 name not equal")

	item2Price, _ := role.Float64(0, "items", 1, "price")
	assert(t, item2Price == 0.3, "role item2 price not equal")
}