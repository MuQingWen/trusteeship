package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateClass
//@description: 创建Class记录
//@param: class model.Class
//@return: err error

func CreateClass(class model.Class) (err error) {
	err = global.GVA_DB.Create(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteClass
//@description: 删除Class记录
//@param: class model.Class
//@return: err error

func DeleteClass(class model.Class) (err error) {
	err = global.GVA_DB.Delete(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteClassByIds
//@description: 批量删除Class记录
//@param: ids request.IdsReq
//@return: err error

func DeleteClassByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Class{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateClass
//@description: 更新Class记录
//@param: class *model.Class
//@return: err error

func UpdateClass(class model.Class) (err error) {
	err = global.GVA_DB.Save(&class).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetClass
//@description: 根据id获取Class记录
//@param: id uint
//@return: err error, class model.Class

func GetClass(id uint) (err error, class model.Class) {
	err = global.GVA_DB.Where("id = ?", id).First(&class).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetClassInfoList
//@description: 分页获取Class记录
//@param: info request.ClassSearch
//@return: err error, list interface{}, total int64

func GetClassInfoList(info request.ClassSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Class{})
    var classs []model.Class
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Cno != "" {
        db = db.Where("`cno` = ?",info.Cno)
    }
    if info.Cname != "" {
        db = db.Where("`cname` = ?",info.Cname)
    }
    if info.Chours != 0 {
        db = db.Where("`chours` = ?",info.Chours)
    }
    if info.Ccycle != "" {
        db = db.Where("`ccycle` = ?",info.Ccycle)
    }
    if info.Ctime != "" {
        db = db.Where("`ctime` = ?",info.Ctime)
    }
    if info.Cplace != "" {
        db = db.Where("`cplace` = ?",info.Cplace)
    }
    if info.Tno != "" {
        db = db.Where("`tno` = ?",info.Tno)
    }
    if info.Tgrade != 0 {
        db = db.Where("`tgrade` = ?",info.Tgrade)
    }
    if info.Caccount != 0 {
        db = db.Where("`caccount` = ?",info.Caccount)
    }
    if info.Cneed != 0 {
        db = db.Where("`cneed` = ?",info.Cneed)
    }
    if info.Cadd != 0 {
        db = db.Where("`cadd` = ?",info.Cadd)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&classs).Error
	return err, classs, total
}