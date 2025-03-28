// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	io "io"
	time "time"
)

type FilesGetRequest struct {
	IdIn []*string `json:"-" url:"id__in,omitempty"`
}

type AllowedFileTypes string

const (
	AllowedFileTypesOcrResults                  AllowedFileTypes = "ocr_results"
	AllowedFileTypesOcrFiles                    AllowedFileTypes = "ocr_files"
	AllowedFileTypesPayables                    AllowedFileTypes = "payables"
	AllowedFileTypesReceivables                 AllowedFileTypes = "receivables"
	AllowedFileTypesReceipts                    AllowedFileTypes = "receipts"
	AllowedFileTypesUserpics                    AllowedFileTypes = "userpics"
	AllowedFileTypesEntityLogo                  AllowedFileTypes = "entity_logo"
	AllowedFileTypesCompaniesLogo               AllowedFileTypes = "companies_logo"
	AllowedFileTypesZip                         AllowedFileTypes = "zip"
	AllowedFileTypesIdentityDocuments           AllowedFileTypes = "identity_documents"
	AllowedFileTypesAdditionalIdentityDocuments AllowedFileTypes = "additional_identity_documents"
	AllowedFileTypesReceivableSignatures        AllowedFileTypes = "receivable_signatures"
)

func NewAllowedFileTypesFromString(s string) (AllowedFileTypes, error) {
	switch s {
	case "ocr_results":
		return AllowedFileTypesOcrResults, nil
	case "ocr_files":
		return AllowedFileTypesOcrFiles, nil
	case "payables":
		return AllowedFileTypesPayables, nil
	case "receivables":
		return AllowedFileTypesReceivables, nil
	case "receipts":
		return AllowedFileTypesReceipts, nil
	case "userpics":
		return AllowedFileTypesUserpics, nil
	case "entity_logo":
		return AllowedFileTypesEntityLogo, nil
	case "companies_logo":
		return AllowedFileTypesCompaniesLogo, nil
	case "zip":
		return AllowedFileTypesZip, nil
	case "identity_documents":
		return AllowedFileTypesIdentityDocuments, nil
	case "additional_identity_documents":
		return AllowedFileTypesAdditionalIdentityDocuments, nil
	case "receivable_signatures":
		return AllowedFileTypesReceivableSignatures, nil
	}
	var t AllowedFileTypes
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AllowedFileTypes) Ptr() *AllowedFileTypes {
	return &a
}

type FileResponse struct {
	Id         string    `json:"id" url:"id"`
	CreatedAt  time.Time `json:"created_at" url:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" url:"updated_at"`
	FileType   string    `json:"file_type" url:"file_type"`
	Md5        string    `json:"md5" url:"md5"`
	Mimetype   string    `json:"mimetype" url:"mimetype"`
	Name       string    `json:"name" url:"name"`
	Region     string    `json:"region" url:"region"`
	S3Bucket   string    `json:"s3_bucket" url:"s3_bucket"`
	S3FilePath string    `json:"s3_file_path" url:"s3_file_path"`
	Size       int       `json:"size" url:"size"`
	Url        string    `json:"url" url:"url"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *FileResponse) GetId() string {
	if f == nil {
		return ""
	}
	return f.Id
}

func (f *FileResponse) GetCreatedAt() time.Time {
	if f == nil {
		return time.Time{}
	}
	return f.CreatedAt
}

func (f *FileResponse) GetUpdatedAt() time.Time {
	if f == nil {
		return time.Time{}
	}
	return f.UpdatedAt
}

func (f *FileResponse) GetFileType() string {
	if f == nil {
		return ""
	}
	return f.FileType
}

func (f *FileResponse) GetMd5() string {
	if f == nil {
		return ""
	}
	return f.Md5
}

func (f *FileResponse) GetMimetype() string {
	if f == nil {
		return ""
	}
	return f.Mimetype
}

func (f *FileResponse) GetName() string {
	if f == nil {
		return ""
	}
	return f.Name
}

func (f *FileResponse) GetRegion() string {
	if f == nil {
		return ""
	}
	return f.Region
}

func (f *FileResponse) GetS3Bucket() string {
	if f == nil {
		return ""
	}
	return f.S3Bucket
}

func (f *FileResponse) GetS3FilePath() string {
	if f == nil {
		return ""
	}
	return f.S3FilePath
}

func (f *FileResponse) GetSize() int {
	if f == nil {
		return 0
	}
	return f.Size
}

func (f *FileResponse) GetUrl() string {
	if f == nil {
		return ""
	}
	return f.Url
}

func (f *FileResponse) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FileResponse) UnmarshalJSON(data []byte) error {
	type embed FileResponse
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*f),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*f = FileResponse(unmarshaler.embed)
	f.CreatedAt = unmarshaler.CreatedAt.Time()
	f.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FileResponse) MarshalJSON() ([]byte, error) {
	type embed FileResponse
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*f),
		CreatedAt: internal.NewDateTime(f.CreatedAt),
		UpdatedAt: internal.NewDateTime(f.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (f *FileResponse) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FilesResponse struct {
	Data []*FileResponse `json:"data" url:"data"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (f *FilesResponse) GetData() []*FileResponse {
	if f == nil {
		return nil
	}
	return f.Data
}

func (f *FilesResponse) GetExtraProperties() map[string]interface{} {
	return f.extraProperties
}

func (f *FilesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler FilesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FilesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	f.rawJSON = json.RawMessage(data)
	return nil
}

func (f *FilesResponse) String() string {
	if len(f.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(f.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type UploadFile struct {
	File     io.Reader        `json:"-" url:"-"`
	FileType AllowedFileTypes `json:"file_type" url:"-"`
}
