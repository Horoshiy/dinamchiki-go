package postgres

import (
	"encoding/base64"
	"fmt"
	"github.com/go-pg/pg/v10"
	"gitlab.com/dinamchiki/go-graphql/graph/model"
)

type ArticlesRepo struct {
	DB *pg.DB
}

func (r *ArticlesRepo) GetArticles(filter *models.ArticleFilter, first, last *int, after, before *string) (*models.ArticleConnection, error) {
	var items []*models.Article
	query := r.DB.Model(&items)

	var decodedCursor string
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}
	if before != nil && after == nil {
		b, err := base64.StdEncoding.DecodeString(*before)
		if err != nil {
			return nil, err
		}
		decodedCursor = string(b)
	}
	var edges []*models.ArticleEdge
	countElems, err := query.Count()
	if err != nil {
		return nil, err
	}

	edges = make([]*models.ArticleEdge, countElems)

	if first != nil {
		query.Order("id ASC")
		edges = make([]*models.ArticleEdge, *first)
	}
	if last != nil && first == nil && before == nil && after == nil {
		query.Order("id DESC")
		edges = make([]*models.ArticleEdge, *last)
	}
	if (last != nil && first == nil) || (before != nil && after == nil) {
		query.Order("id DESC")
		edges = make([]*models.ArticleEdge, countElems)
	}

	count := 0
	currentPage := false

	if decodedCursor == "" {
		currentPage = true
	}
	hasNextPage := false

	if filter != nil {
		if filter.Title != nil && *filter.Title != "" {
			query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Title))
		}
	}

	err = query.Select()
	if err != nil {
		return nil, err
	}

	for i, v := range items {
		if v.ID == decodedCursor {
			currentPage = true
		}
		if err != nil {
			return nil, err
		}
		if first != nil {
			if currentPage && count < *first {
				edges[count] = &models.ArticleEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *first && i < len(items) {
				hasNextPage = true
			}
		}
		if last != nil && first == nil {
			if currentPage && count < *last {
				edges[count] = &models.ArticleEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}

			if count == *last && i < len(items) {
				hasNextPage = true
			}
		}
		if last == nil && first == nil {
			if currentPage && count < countElems {
				edges[count] = &models.ArticleEdge{
					Cursor: base64.StdEncoding.EncodeToString([]byte(v.ID)),
					Node:   v,
				}
				count++
			}
		}
	}

	pageInfo := models.PageInfo{
		StartCursor: base64.StdEncoding.EncodeToString([]byte(edges[0].Node.ID)),
		EndCursor:   base64.StdEncoding.EncodeToString([]byte(edges[count-1].Node.ID)),
		HasNextPage: &hasNextPage,
	}

	pc := &models.ArticleConnection{
		Edges:    edges[:count],
		PageInfo: &pageInfo,
	}

	return pc, nil
}

func (r *ArticlesRepo) GetArticleByFiled(field, value string) (*models.Article, error) {
	var item models.Article
	err := r.DB.Model(&item).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &item, err
}

func (r *ArticlesRepo) GetArticleByID(id string) (*models.Article, error) {
	return r.GetArticleByFiled("id", id)
}

func (r *ArticlesRepo) GetArticleByTitle(title string) (*models.Article, error) {
	return r.GetArticleByFiled("title", title)
}

func (r *ArticlesRepo) ArticleSave(place *models.Article) (*models.ArticlePayload, error) {
	_, err := r.DB.Model(place).Returning("*").Insert()
	fmt.Println(err)
	placePayload := &models.ArticlePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ArticlesRepo) ArticleUpdate(place *models.Article) (*models.ArticlePayload, error) {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Update()
	placePayload := &models.ArticlePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ArticlesRepo) ArticlePublish(place *models.Article) (*models.ArticlePayload, error) {
	_, err := r.DB.Model(place).Set("published = NOT published").Where("id = ?", place.ID).Update()
	placePayload := &models.ArticlePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ArticlesRepo) ArticleRestore(place *models.Article) (*models.ArticlePayload, error) {
	_, err := r.DB.Model(place).Set("deletedAt = NULL").Where("id = ?", place.ID).Update()
	placePayload := &models.ArticlePayload{
		Record:   place,
		RecordID: place.ID,
	}
	return placePayload, err
}

func (r *ArticlesRepo) ArticleDelete(place *models.Article) error {
	_, err := r.DB.Model(place).Where("id = ?", place.ID).Delete()
	return err
}
