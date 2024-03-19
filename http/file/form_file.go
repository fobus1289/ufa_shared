package form_file

import "net/http"

func NewForm(r *http.Request) IFormFile {
	return &form{
		Request: r,
	}
}

type IFormFile interface {
	GetFile(formKey string) IFiles
	Files() IFileMap
}

type form struct {
	*http.Request
	FormFiles fileMap
}

func (r *form) GetFile(formKey string) IFiles {
	r.ParseFormFiles()
	return r.FormFiles.Get(formKey)
}

func (r *form) Files() IFileMap {
	r.ParseFormFiles()
	return r.FormFiles
}

func (r *form) ParseFormFiles() {

	if r.FormFiles != nil {
		return
	}

	r.FormFiles = fileMap{}

	r.FormValue("")

	var multipartForm = r.MultipartForm

	if multipartForm == nil || len(multipartForm.File) == 0 {
		return
	}

	var multipartFormFiles = multipartForm.File

	for key, headers := range multipartFormFiles {

		if len(headers) == 0 {
			continue
		}

		for _, fileHeader := range headers {
			if fileHeader == nil {
				continue
			}
			r.FormFiles[key] = append(r.FormFiles[key], &file{
				FileHeader: fileHeader,
			})
		}

	}

}
