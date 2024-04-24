package form_file

import (
	"io"
	"net/http"
	"os"
	"sync"
)

type files map[string]IFileContainer

type IFiles interface {
	Get(name string) IFileContainer
	Len() int
}

type IFormFile interface {
	Get(formKey string) IFileContainer
	Files() IFiles
}

type IFileInfo interface {
	Size() int64
	Name() string
	ContentType() string
	Extension() string
}

type IFile interface {
	Read(writer io.Writer) IFile
	Store(dir string, storagePath *string) IFile
	Rollback() IFile
	RandomFileName() IFile
	Info() IFileInfo
	SetNewName(name string) IFile
	GetNewName() string
	SetPrem(perm os.FileMode) IFile
	IsValid() bool
	Error() error
}

type IFileContainer interface {
	Files() []IFile
	GetFirst() IFile
	Errors() []error
	RollbackAll() IFileContainer
	StoreAll(dir string, storagePaths *[]string) IFileContainer
	Count() int
	Has() bool
	HasMultiple() bool
}

func NewFormFile(w http.ResponseWriter, r *http.Request) IFormFile {
	return &formFile{
		w:     w,
		r:     r,
		once:  &sync.Once{},
		files: files{},
	}
}

type formFile struct {
	w     http.ResponseWriter
	r     *http.Request
	once  *sync.Once
	files files
}

func (f *formFile) init() {
	f.once.Do(func() {

		f.r.FormValue("")

		var multipartForm = f.r.MultipartForm

		if multipartForm == nil || len(multipartForm.File) < 1 {
			return
		}

		var multipartFormFiles = multipartForm.File

		for key, headers := range multipartFormFiles {

			if len(headers) < 1 {
				continue
			}

			var _files = make([]IFile, 0, len(headers))

			for _, header := range headers {
				_files = append(_files, &file{
					w:          f.w,
					r:          f.r,
					FileHeader: header,
					perm:       0666,
					Once:       &sync.Once{},
				})
			}

			f.files[key] = &fileContainer{
				files: _files,
			}

		}

	})
}

func (f *formFile) Get(name string) IFileContainer {
	f.init()
	return f.files.Get(name)
}

func (f *formFile) Files() IFiles {
	f.init()
	return f.files
}

func (fs files) Get(name string) IFileContainer {
	if f := fs[name]; f != nil {
		return f
	}
	return &fileContainer{}
}

func (fs files) Len() int {
	return len(fs)
}
