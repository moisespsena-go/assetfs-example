// +build init_repository

package project

func Main() {
	repo := NewRepository()
	repo.Init()
}