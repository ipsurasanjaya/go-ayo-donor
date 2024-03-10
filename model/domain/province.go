package domain

type GetProvinceOut struct {
	ID   int64
	Name string
}

type GetProvinceIn struct {
	Limit  int64
	Search string
}
