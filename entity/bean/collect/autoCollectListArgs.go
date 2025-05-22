package collect

// AutoCollectListArgs struct
type AutoCollectListArgs struct {
	MaxId string `json:"max_id"`
}

func (p AutoCollectListArgs) Validate() bool {
	return p.MaxId == ""
}
