package gostruct

import "reflect"

type Builder struct {
	field []reflect.StructField
	tags  map[string]*reflect.StructField
}

func NewBuilder() *Builder {
	return &Builder{
		tags: map[string]*reflect.StructField{},
	}
}

func (b *Builder) SetTagForField(fieldName, tagName string) *Builder {
	if !b.containsField(fieldName) {
		return b
	}
	b.tags[fieldName].Tag = reflect.StructTag(tagName)
	return b
}

func (b *Builder) containsField(field string) bool {
	_, ok := b.tags[field]
	return ok
}

func (b *Builder) AddField(name string, ftype reflect.Type) *Builder {
	if !b.containsField(name) {
		b.field = append(b.field, reflect.StructField{
			Name: name,
			Type: ftype,
		})
		b.tags[name] = &b.field[len(b.field)-1]
	}
	return b
}

func (b *Builder) AddString(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(""))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(true))
}

func (b *Builder) AddInt64(name string) *Builder {

	return b.AddField(name, reflect.TypeOf(int64(0)))
}

func (b *Builder) AddInt32(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(int32(0)))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.AddField(name, reflect.TypeOf(float64(1.2)))
}

func (b *Builder) Build() *Struct {
	strct := reflect.StructOf(b.field)
	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		f := strct.Field(i)
		index[f.Name] = i
	}

	return &Struct{strct, index}
}

type Struct struct {
	strct reflect.Type
	index map[string]int
}

func (s *Struct) New() *Instance {
	instance := reflect.New(s.strct).Elem()

	return &Instance{instance, s.index}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt64(name string, value int64) {
	i.Field(name).SetInt(value)
}

func (i *Instance) SetInt32(name string, value int32) {
	i.Field(name).SetInt(int64(value))
}

func (i *Instance) SetFloat64(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) Interface() any {
	return i.internal.Interface()
}

func (i *Instance) Addr() any {
	return i.internal.Addr().Interface()
}
