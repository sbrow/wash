package plugin

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// Node represents a filesystem node
type Node = fs.Node

// ENOENT states the entity does not exist
const ENOENT = fuse.ENOENT

// Entry in a filesystem
type Entry struct {
	Name  string
	Isdir bool
}

// Interface that plugins are expected to model.
type Interface interface {
	Find(name string) (Node, error)
	List() ([]Entry, error)
}

// FS contains the core filesystem data.
type FS struct {
	Clients map[string]Interface
}

var _ fs.FS = (*FS)(nil)

// Dir represents a directory within the system, with the client
// necessary to represent it and the full path to the directory.
type Dir struct {
	client Interface
	name   string
}

var _ fs.Node = (*Dir)(nil)
var _ = fs.NodeRequestLookuper(&Dir{})
var _ = fs.HandleReadDirAller(&Dir{})

// File contains metadata about the file.
type File struct {
	Meta interface{}
}

var _ fs.Node = (*File)(nil)
var _ = fs.NodeOpener(&File{})