package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
	"gitlab.com/dinamchiki/go-graphql/middleware"
)

func articleInputToArticle(ctx context.Context, input *models.ArticleInput) *models.Article {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil
	}
	return &models.Article{
		AuthorID:    currentUser.ID,
		Description: input.Description,
		FileName:    input.FileName,
		Published:   input.Published,
		Tags:        input.Tags,
		Title:       input.Title,
	}
}
func articleInputWithIdToArticle(ctx context.Context, article *models.Article, input *models.ArticleInput) *models.Article {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil
	}
	return &models.Article{
		ID:          article.ID,
		AuthorID:    currentUser.ID,
		Description: input.Description,
		FileName:    input.FileName,
		Published:   input.Published,
		Tags:        input.Tags,
		Title:       input.Title,
	}
}

func (d *Domain) ArticleSave(ctx context.Context, input models.ArticleInput) (*models.ArticlePayload, error) {

	_, err := d.ArticlesRepo.GetArticleByTitle(input.Title)
	if err == nil {
		return nil, errors.New("такое название уже есть у статьи")
	}

	return d.ArticlesRepo.ArticleSave(articleInputToArticle(ctx, &input))
}

func (d *Domain) ArticleUpdate(ctx context.Context, articleInput models.ArticleInputWithID) (*models.ArticlePayload, error) {

	article, err := d.ArticlesRepo.GetArticleByID(articleInput.ID)
	if err != nil || article == nil {
		return nil, errors.New("такого филиала не существует")
	}

	if len(articleInput.Input.Title) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}
	if len(articleInput.Input.Description) < 3 {
		return nil, errors.New("кол-во символов в описании мало")
	}

	article = articleInputWithIdToArticle(ctx, article, articleInput.Input)
	articlePayload, err := d.ArticlesRepo.ArticleUpdate(article)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return articlePayload, nil
}

func (d *Domain) ArticlePublish(id string) (*models.ArticlePayload, error) {
	article, err := d.ArticlesRepo.GetArticleByID(id)
	if err != nil || article == nil {
		return nil, errors.New("филиала не существует")
	}

	articlePayload, err := d.ArticlesRepo.ArticlePublish(article)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return articlePayload, nil
}

func (d *Domain) ArticleRestore(id string) (*models.ArticlePayload, error) {
	article, err := d.ArticlesRepo.GetArticleByID(id)
	if err != nil || article == nil {
		return nil, errors.New("филиала не существует")
	}

	articlePayload, err := d.ArticlesRepo.ArticleRestore(article)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return articlePayload, nil
}

// ArticleDelete is the resolver for the articleDelete field.
func (d *Domain) ArticleDelete(id string) (bool, error) {
	article, err := d.ArticlesRepo.GetArticleByID(id)
	if err != nil || article == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.ArticlesRepo.ArticleDelete(article)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
