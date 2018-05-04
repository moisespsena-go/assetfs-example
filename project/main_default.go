// +build !bindata,!init_repository,!pre_compile

package project

import "github.com/moisespsena/go-assetfs-example/repository"

func Main() {
	Run(repository.AssetFS)
}