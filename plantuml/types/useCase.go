package types

type useCasePrecedent struct {
	title    string
	business bool //тупо добавляет черточку справа снизу
}

type useCaseActorStyle string

const (
	Admin   useCaseActorStyle = "Admin"
	awesome useCaseActorStyle = "awesome"
	hollow  useCaseActorStyle = "hollow"
)

type useCaseActor struct {
	name     string
	style    useCaseActorStyle // стиль меняется на всех человечках
	business bool              //тупо добавляет черточку справа снизу
}
