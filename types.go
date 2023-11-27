package core

// Class body
// By default, methods and fields are automatically regrouped by PlantUML.
// You can use separators to define your own way of ordering fields and methods.
// The following separators are possible: "--" ".." "==" "__".
//
// You can also use titles within the separators:
//
// @startuml
// class Foo1 {
//   You can use
//   several lines
//   ..
//   as you want
//   and group
//   ==
//   things together.
//   __
//   You can have as many groups
//   as you want
//   --
//   End of class
// }
//
// class User {
//   .. Simple Getter ..
//   + getName()
//   + getAddress()
//   .. Some setter ..
//   + setName()
//   __ private data __
//   int age
//   -- encrypted --
//   String password
// }
//
// @enduml

type BodyGroupKind int

const (
	BodyGroupInvalid BodyGroupKind = iota
	BodyGroupSolid                 // --
	BodyGroupDotty                 // ..
	BodyGroupWeak                  // __
	BodyGroupDouble                // ==
)

type BodyGroup struct {
	Ident string
	Text  string
	Kind  BodyGroupKind
}
