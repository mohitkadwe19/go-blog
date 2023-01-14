package service

import (
	"Blog/config"
	"Blog/ent"
	"Blog/ent/blog"
	"Blog/ent/user"
	"context"
)

type BlogOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewBlogOps(ctx context.Context) *BlogOps {
	return &BlogOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *BlogOps) BlogGetByID(id int) (*ent.Blog, error) {

	blog, err := r.client.Blog.Query().Where(blog.ID(id)).Only(r.ctx)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (r *BlogOps) BlogCreate(newBlog ent.Blog) (*ent.Blog, error) {

	user, err := r.client.User.Query().Where(
		user.ID(newBlog.Edges.Owner.ID),
	).Only(r.ctx)
	if err != nil {
		return nil, err
	}

	blog, err := r.client.Blog.Create().
		SetTitle(newBlog.Title).
		SetContent(newBlog.Content).
		SetAuthor(newBlog.Author).
		SetOwner(user).
		Save(r.ctx)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (r *BlogOps) BlogUpdate(blog ent.Blog) (*ent.Blog, error) {

	updatedBlog, err := r.client.Blog.UpdateOneID(blog.ID).
		SetTitle(blog.Title).
		SetContent(blog.Content).
		SetAuthor(blog.Author).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func (r *BlogOps) BlogDelete(id int) (int, error) {

	err := r.client.Blog.
		DeleteOneID(id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BlogOps) BlogGetAll() ([]*ent.Blog, error) {

	blogs, err := r.client.Blog.Query().All(r.ctx)

	if err != nil {
		return nil, err
	}

	return blogs, nil
}
