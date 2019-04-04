// +build bindata

package project

import "github.com/moisespsena-go/assetfs-example/repository"

func Main() {
	Run(repository.AssetFS)
}
