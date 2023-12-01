package types

type RelationModel struct {
	first  string
	second string
	// Type   classRealtionType // с типами связями есть проблема:многие связи поддерживают только определенные сущности
	label string
}
