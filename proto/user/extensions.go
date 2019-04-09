package go_micro_srv_user
//
// import (
// 	"github.com/jinzhu/gorm"
// 	uuid "github.com/satori/go.uuid"
// )
//
// //
// // import "github.com/jinzhu/gorm"
// // import "github.com/satori/go.uuid"
// //
// // func (model *User) BeforeCreate(scope *gorm.Scope) error {
// // 	uuid, err := uuid.NewV4()
// // 	if err != nil {
// // 		return err
// // 	}
// // 	return scope.SetColumn("Id", uuid.String())
// // }
//
//
//
// func (model *User) BeforeCreate(scope *gorm.Scope) error {
// 	uuid, _ := uuid.NewV4()
// 	return scope.SetColumn("Id", uuid.String())
// }