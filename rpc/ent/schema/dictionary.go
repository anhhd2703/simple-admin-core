package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Dictionary struct {
	ent.Schema
}

func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Comment("The title shown in the ui | 展示名称 （建议配合i18n）").
			Annotations(entsql.WithComments(true)),
		field.String("name").Unique().
			Comment("The name of dictionary for search | 字典搜索名称").
			Annotations(entsql.WithComments(true)),
		field.String("desc").
			Comment("The status of dictionary (true enable | false disable) | 字典状态").
			Optional().
			Annotations(entsql.WithComments(true)),
	}
}

func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
	}
}

func (Dictionary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dictionary_details", DictionaryDetail.Type),
	}
}

func (Dictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_dictionaries"},
	}
}
