package response

type AggregateDTO struct {
	AreaProvinsi string `json:"area_provinsi"`
	Tahun        string `json:"tahun"`
	MingguKe     string `json:"minggu_ke"`
	Min          string `json:"min"`
	Max          string `json:"max"`
	Median       string `json:"median"`
	Avg          string `json:"avg"`
}
