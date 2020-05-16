package request

type FetchingMapperDTO struct {
	Uuid         *string `json:"uuid"`
	AreaKota     *string `json:"area_kota"`
	AreaProvinsi *string `json:"area_provinsi"`
	Komoditas    *string `json:"komoditas"`
	Price        *string `json:"price"`
	Size         *string `json:"size"`
	TglParsed    *string `json:"tgl_parsed"`
	Timestamp    *string `json:"timestamp"`
}
