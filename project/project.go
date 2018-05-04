package project

import (
	"fmt"
	"strings"
	"github.com/moisespsena/go-assetfs"
	"github.com/moisespsena/go-assetfs/repository"
	"path/filepath"
	"github.com/moisespsena/go-path-helpers"
)

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}

func printStrings(indent int, data []string)  {
	prefix := strings.Repeat("  ", indent)
	for i, v := range data {
		fmt.Printf("%v - %02d: %v\n", prefix, i, v)
	}
}

func Run(fs assetfs.Interface) {
	PrepareFS(fs)
	SetNameSpaces(fs)
	// get asset and error
	asset, err := fs.Asset("f1.txt")
	checkError(err)
	fmt.Println("f1.txt:", asset.GetString())
	// get asset or if error, panic
	asset = fs.AssetOrPanic("f2.txt")
	fmt.Println("f2.txt:", asset.GetString())
	fmt.Println("f3.txt:", fs.AssetOrPanic("f3.txt").GetString())

	// name spaces
	ns1 := fs.NameSpace("ns1")
	fmt.Println("ns1-f1.txt:", ns1.AssetOrPanic("ns1-f1.txt").GetString())
	fmt.Println("ns1-f2.txt:", ns1.AssetOrPanic("ns1-f2.txt").GetString())
	fmt.Println("ns1-f3.txt:", ns1.AssetOrPanic("ns1-f3.txt").GetString())

	// glob
	fmt.Println("glob *.txt:")
	printStrings(0, ns1.GlobOrPanic("*.txt"))
	// glob recursive
	fmt.Println("glob recursive *.txt:")
	printStrings(0, ns1.GlobOrPanic("*.txt", true))
}

func NewAssetFS() assetfs.Interface {
	fs := assetfs.NewAssetFileSystem()
	PrepareFS(fs)
	return fs
}

func PrepareFS(fs assetfs.Interface)  {
	fs.RegisterPath("./assets1")
	fs.RegisterPath("./assets2")
}

func NewAssetFSWithNameSpaces() assetfs.Interface {
	fs := NewAssetFS()
	SetNameSpaces(fs)
	return fs
}

func SetNameSpaces(fs assetfs.Interface) {
	ns1 := fs.NameSpace("ns1")
	ns1.RegisterPath("another-ns1")
}

func NewRepository() repository.Interface {
	fs := NewAssetFSWithNameSpaces()
	repositoryPkg := filepath.Join(filepath.Dir(path_helpers.GetCalledDir()), "repository")
	fmt.Println("Repository package:", repositoryPkg)
	repo := fs.(*assetfs.AssetFileSystem).NewRepository(repositoryPkg)
	return repo
}