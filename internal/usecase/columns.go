package usecase

type Columns map[string]interface{}

type UpdateBuilder struct {
	Columns
}

func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{Columns{}}
}

func (b *UpdateBuilder) PhotoURL(v string) *UpdateBuilder {
	b.Columns["photo_url"] = v
	return b
}

func (b *UpdateBuilder) Name(v string) *UpdateBuilder {
	b.Columns["name"] = v
	return b
}

func (b *UpdateBuilder) Age(v int) *UpdateBuilder {
	b.Columns["age"] = v
	return b
}

func (b *UpdateBuilder) City(v string) *UpdateBuilder {
	b.Columns["city"] = v
	return b
}

func (b *UpdateBuilder) InterestedIn(v string) *UpdateBuilder {
	b.Columns["interested_in"] = v
	return b
}

func (b *UpdateBuilder) Activated(v bool) *UpdateBuilder {
	b.Columns["activated"] = v
	return b
}
