// +build pre_compile

package project

import "github.com/moisespsena/go-assetfs-example/repository"

func Main() {
	PrepareFS(repository.FileSystem)
	SetNameSpaces(repository.FileSystem)
	repository.CallCallbacks()
}