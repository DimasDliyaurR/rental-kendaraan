package models

import (
	"github.com/goravel/framework/database/orm"
)

type Kendaraan struct {
	orm.Model
	NamaKendaraan string
	Plat          string
	Kapasitas     int
	orm.SoftDeletes
}
