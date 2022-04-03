package errutil

func NewTagged(err error, tags ...Tag) error {
	return Tagged{
		error: err,
		tags:  tags,
	}
}

type Tagged struct {
	error
	tags []Tag
}

func (t Tagged) Tags() []Tag {
	return t.tags
}

func (t Tagged) Unwrap() error {
	return t.error
}

func NewTag(key, value string) Tag {
	return Tag{
		Key:   key,
		Value: value,
	}
}

type Tag struct {
	Key   string
	Value string
}
