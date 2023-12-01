package types

// есть непонятки со стилями: https://www.plantuml.com/plantuml/umla/b8qz2iCm34PdwnGVT6de3J8rq2Fq0ewZD0R_W2Lt2RbxdPAC7Rg9nzFxxgAsQ8s1NB0Y6B2ICMkZwjq6dnGkrwHS_iQtT_OJLZkA5kjqmtVgkalg8-UgKJhy5X-rj6nEfzsHsSYah8p0ioBF6Fh-sCDb39N9F1SlY1p7XjfbK_EjyaI6quJzcOv6yQ_bZ7pg9fkS188e2EyvnGjk3GTmvLbSQqWC1s10ZGOINBjqRz8C33OuFAvGxqPMzKITHDtD-_-gEZFy2stujHDlCKCol2EXHSFWW7HkE38-nHIvPAVCRE7um4uLmpeMCOkbHONCJ8DzcN7Jf-44Z_kOOp_wORD_uNxR8RvR2SEhhiQIL2x8BM4gteKNkaO9rbSYPs5byO1QGL-Huw1RvLJVgq9pDPvCnZSNXcTZynCQuvHZn-1H0kVKEaKgQb5JR_DzGqRaJcobYDbS-iwk9L46ZmmlIAB71cuoXP2e2vc6lByjxk35NvLnxCb2hJvAXQErCfdbYvub8s4ZHawriR3kWCEnBdgwbv0XNAsDWy4bJU1_yGa0
// стили иногда можно сломать. Это не самый показательный случай, но я несколько раз их ломал, пока тыкал: https://www.plantuml.com/plantuml/umla/b8uzhi8m48Hhxob6USsBk44g8d44Bk0u6sB9FvAzfeboTsn4gA2WMus-RsVEMNIIuXqPftF6Y8sKscegrkomGM1Y2SBfv_XqZtR6fgSiIHlPyQpwRv1Opx68pps-WvUIwjiO3hedjPCAKHYE5q5SC0x3k--wsWpFVkA4UkbGIIqya-8mesdqifEzhOVUVHhL_Rc6JAumFUgcbfm4WY2wrbDKN1iEuCgpk3TiEC0CgD50WckNxarQeB4PX-8AhsxaUtRYJ6BO4EhpwdqwSu6Fc2CPlf3G8xFeaAuDHzQdQ52WRiFE3evlx3KnVCOovh6oQ57OoJW80pVzDPpWyHWB72Q3FZily7yOLrJsBi9OeODBtuuwMbAoAs5P3SxN4it1sk896nLZCuvEDByiuwZqQtVJsMnwVsDuDIw_81jHZZsBPtbRq6Q39AE63QtcFmvbKAhFLXMxHNisP6iJmnkIv6CBDmL2bToKDA5yV2uUu3LRZd4SV16XFNg5epSXKHICdgIGO2kHHFMnj4hWS4gNeVeAJ4IdTV7obY7mW_W1
// но в целом, если они нужны, то они не должны быть сильно сложными в имплементации
type Note struct {
	text      string // + можно использовать некоторые HTML теги, но я хз как сделать тип для такого
	direction string
	of        string // имя комментируемой сущности. Если нужно сделать к конкретному, например, методу, то у имени нужно добавить :: (Person::say)
	name      string
	relation  string
	onLink    bool // если false, то необходимо размещать под комментируемым классом, если true, то под комментируемой записью
}
