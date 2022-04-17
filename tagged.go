package errutil

func NewTagged(err error, tags ...Tag) error {
	return Tagged{
		wrappedError: wrappedError{error: err},
		tags:         tags,
	}
}

type Tagged struct {
	wrappedError
	tags []Tag
}

func (t Tagged) Tags() []Tag {
	return t.tags
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
