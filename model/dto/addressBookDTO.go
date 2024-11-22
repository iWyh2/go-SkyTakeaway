package dto

// AddressBookDTO 地址簿传输数据模型
type AddressBookDTO struct {
	Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	// 用户id
	UserId int `json:"userId"`
	// 收货人
	Consignee string `json:"consignee"`
	// 手机号
	Phone string `json:"phone"`
	// 性别 0 女 1 男
	Sex string `json:"sex"`
	// 省级区划编号
	ProvinceCode string `json:"provinceCode"`
	// 省级名称
	ProvinceName string `json:"provinceName"`
	// 市级区划编号
	CityCode string `json:"cityCode"`
	// 市级名称
	CityName string `json:"cityName"`
	// 区级区划编号
	DistrictCode string `json:"districtCode"`
	// 区级名称
	DistrictName string `json:"districtName"`
	// 详细地址
	Detail string `json:"detail"`
	// 标签
	Label int `json:"label"`
	// 是否默认 0否 1是
	IsDefault int `json:"isDefault"`
}
