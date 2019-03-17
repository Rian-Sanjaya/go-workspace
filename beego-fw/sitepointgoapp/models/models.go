package models

type Article struct {
	// form:"-" means id is not required, in sqlite id is an auto-increment field
	Id int `form:"-"`
	// "name,test,name"
	// from the form field with name: name
	// the field will be a text field
	// the label will be set to 'name'
	// valid:"MinSize(5);MaxSize(20)" - value supplied must be at least 5 characters but not longer than 20
	Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
	Client string `form:"client,text,client:"`
	Url    string `form:"url,text,url:"`
}

// TableName function
// The reason why I added it is because the article table name doesn’t match the struct’s name.
// Instead, it’s called articles.
//If these were both the same, then it would be automatically found by Beego’s ORM.
func (a *Article) TableName() string {
	return "articles"
}
