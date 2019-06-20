package config

import (
	"fmt"
	"testing"
)

const (
	dataDesc = `This is the data description.
This is the second line.`
)

func assert(t *testing.T, ok bool, msg string) {
	if !ok {
		t.Fatal(msg)
	}
}

func TestPair(t *testing.T) {
	role := NewPair("role").
		Add(NewString(dataDesc)).
		Add(NewPair("name").AddString("role1")).
		Add(NewPair("level").AddNumber("1")).
		Add(NewPair("movable").AddBool(true)).
		Add(NewPair("tags").AddString("robot").AddString("np\\c")).
		Add(NewPair("items").
			Add(NewPair("").
				Add(NewPair("name").AddString("item\"1\"")).
				Add(NewPair("price").AddNumber("1.5"))).
			Add(NewPair("").
				Add(NewPair("name").AddString("'item'2")).
				Add(NewPair("price").AddNumber("12'000.5"))))

	fmt.Println(role.Format(0))

	desc, _ := role.String("", 0)
	assert(t, desc == dataDesc, "data description not equal")

	name, _ := role.String("", "name")
	assert(t, name == "role1", "role name not equal")

	level, _ := role.Int64(0, "level")
	assert(t, level == 1, "role level not equal")

	movable, _ := role.Bool(false, "movable")
	assert(t, movable == true, "role movable not equal")

	role.SetBool(false, "movable")
	movable, _ = role.Bool(true, "movable")
	assert(t, movable == false, "role setted movable not equal")

	tag1, _ := role.String("", "tags", 0)
	assert(t, tag1 == "robot", "role tag1 not equal")

	tag2, _ := role.String("", "tags", 1)
	assert(t, tag2 == "np\\c", "role tag2 not equal")

	item1Name, _ := role.String("", "items", 0, "name")
	assert(t, item1Name == "item\"1\"", "role item1 name not equal")

	item1Price, _ := role.Float64(0, "items", 0, "price")
	assert(t, item1Price == 1.5, "role item1 price not equal")

	item2Name, _ := role.String("", "items", 1, "name")
	assert(t, item2Name == "'item'2", "role item2 name not equal")

	item2Price, _ := role.Float64(0, "items", 1, "price")
	assert(t, item2Price == 12000.5, "role item2 price not equal")
}

var data = fmt.Sprintf(`
"%s"

(version "0.1.1")

(3.14 "pi")

(role
	; this is a comment.
	(name "role1")
	(level 1)
	(movable true)
	(tags "robot" "np\c")

	(items ; another comment.
		((name "item\"1\"")
			(price 1.5))
		((name "'item'2")
			(price 12'000.5))
	)
)`, dataDesc)

func TestParser(t *testing.T) {
	c := New()
	err := c.Parse(data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(c.Format())

	desc, _ := c.String("", 0)
	assert(t, desc == dataDesc, "data description not equal")

	version, _ := c.String("", "version")
	assert(t, version == "0.1.1", "data version not equal")

	pi, _ := c.String("", "3.14")
	assert(t, pi == "pi", "number key not equal")

	role, err := c.Pair.Pair("role")
	if err != nil {
		t.Fatal(err)
	}

	name, _ := role.String("", "name")
	assert(t, name == "role1", "role name not equal")

	level, _ := role.Int64(0, "level")
	assert(t, level == 1, "role level not equal")

	movable, _ := role.Bool(false, "movable")
	assert(t, movable == true, "role movable not equal")

	role.SetBool(false, "movable")
	movable, _ = role.Bool(true, "movable")
	assert(t, movable == false, "role setted movable not equal")

	tag1, _ := role.String("", "tags", 0)
	assert(t, tag1 == "robot", "role tag1 not equal")

	tag2, _ := role.String("", "tags", 1)
	assert(t, tag2 == "np\\c", "role tag2 not equal")

	item1Name, _ := role.String("", "items", 0, "name")
	assert(t, item1Name == "item\"1\"", "role item1 name not equal")

	item1Price, _ := role.Float64(0, "items", 0, "price")
	assert(t, item1Price == 1.5, "role item1 price not equal")

	item2Name, _ := role.String("", "items", 1, "name")
	assert(t, item2Name == "'item'2", "role item2 name not equal")

	item2Price, _ := role.Float64(0, "items", 1, "price")
	assert(t, item2Price == 12000.5, "role item2 price not equal")
}
