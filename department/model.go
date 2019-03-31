package department

type Container struct {
	Id       int         `json:"id" db:"id"`
	Name     string      `json:"name" db:"name"`
	ParentId int         `json:"_" db:"parent_id"`
	URL      string      `json:"url" db:"url"`
	Children []Container `json:"children"`
}
