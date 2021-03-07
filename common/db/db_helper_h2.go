package db

type h2DbHelper struct{}

func (helper h2DbHelper) Init() {
}

func (helper h2DbHelper) Insert(tablename string, entity Entity) int {
	return 1
}

func (helper h2DbHelper) Select() interface{} {
	return 1
}
