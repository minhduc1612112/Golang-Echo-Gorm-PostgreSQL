package models

import (
	"gorm.io/gorm"

	database "echo-demo/db"
	form "echo-demo/forms"
)

func GetBuildingsList() (buildings []form.Building, err error) {
	var db *gorm.DB

	db, err = database.Connect()
	if err == nil {
		db.AutoMigrate(&form.Building{})
		err = db.Model(&buildings).Find(&buildings).Error
	}
	return buildings, err
}

func GetBuildingsList2() (buildings []form.Building, err error) {
	var db *gorm.DB
	db, err = database.Connect()
	if err == nil {
		// err = db.Table("buildings").Find(&buildings).Error
		// if err == nil {
		// 	for i := range buildings {
		// 		db.Table("rooms").Where("buildingid = ?", buildings[i].ID).Find(&buildings[i].Rooms)
		// 	}
		// }
		err = db.Raw(`select *,
									array_to_json(array(select json_build_object('id', id, 'createdAt',created_at,
																							'updatedAt', updated_at, 'deletedAt', deleted_at,
																							'name', name, 'description', description,
																							'numberOfPeople', number_of_people,
																							'area', area, 'buildingId', building_id)
																			from rooms where building_id = 1))
										as rooms
									from buildings`).Scan(&buildings).Error
	}
	return buildings, err
}

func GetBuilding(id uint) (building form.Building, err error) {
	var db *gorm.DB

	db, err = database.Connect()
	if err == nil {
		db.AutoMigrate(&form.Building{})
		err = db.Model(&building).First(&building, id).Error
	}
	return building, err
}
