package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func (v *GetAnArticleRequest) Validate() error {
	return validate.Struct(v)
}

func (v *GetArticlesRequest) Validate() error {
	return validate.Struct(v)
}

func (v *NewArticleRequest) Validate() error {
	return validate.Struct(v)
}

func (v *UpdateArticleRequest) Validate() error {
	return validate.Struct(v)
}

func (v *DeleteArticleRequest) Validate() error {
	return validate.Struct(v)
}

func (v *LikeArticleRequest) Validate() error {
	return validate.Struct(v)
}

func (v *GetTopAuthorsRequest) Validate() error {
	return validate.Struct(v)
}

func (v *GetPopularArticlesRequest) Validate() error {
	return validate.Struct(v)
}
