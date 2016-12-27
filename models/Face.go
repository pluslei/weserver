package models

import (
	"github.com/astaxie/beego/orm"
)

// 表情库
type Face struct {
	Id        int64  // 标识
	Title     string // 标题
	Url       string // 路径
	Group     int64  // 分组
	GroupFace string // 分组表情
	// GroupUrl  string `orm:"-"` // 无实际意义
}

func (this *Face) TableName() string {
	return "face"
}

func init() {
	orm.RegisterModel(new(Face))
}

// 查询表情列表
func GetFaceList(page int64, page_size int64, sort string) (faces []orm.Params, count int64) {
	o := orm.NewOrm()
	face := new(Face)
	qs := o.QueryTable(face)
	qs.Limit(page_size, page).OrderBy(sort).Values(&faces)
	count, _ = qs.Count()
	return faces, count
}

// 添加表情
func AddFace(f *Face) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(f)
	return id, err
}

// 修改表情
func EditFace(f *Face) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(f)
	return id, err
}

// 删除表情
func DelFace(id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Face{Id: id})
	return status, err
}

// 按照Id查询表情
func GetFaceById(id int64) (Face, error) {
	o := orm.NewOrm()
	face := Face{Id: id}
	err := o.Read(&face)
	if err != nil {
		return face, err
	}
	return face, nil
}

// 查询所有表情
func GetFaces() (faces []orm.Params) {
	o := orm.NewOrm()
	face := new(Face)
	qs := o.QueryTable(face)
	qs.Values(&faces)
	return faces
}

// 按照 Group 查询表情
func GetFaceByGroup(group int64) (faces []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("face").Filter("group", group).Values(&faces)
	return faces, err
}

// 查询所有的分组
func GetGroupList() (faces []Face, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("face").Distinct().All(&faces, "group", "group_face")
	return faces, err
}

// 根据分组查询表情
func GetGroupFace(group int64) (face Face, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("face").Filter("group", group).Distinct().One(&face, "group", "group_face")
	return face, err
}

func GetMaxGroup() (max []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT MAX(`group`) AS Max FROM face").Values(&max)
	return max, err
}
