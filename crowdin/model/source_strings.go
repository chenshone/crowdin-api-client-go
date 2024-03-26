package model

import (
	"errors"
	"fmt"
	"net/url"
)

// SourceString represents the text units for translation.
type SourceString struct {
	ID             int     `json:"id"`
	ProjectID      int     `json:"projectId"`
	BranchID       *int    `json:"branchId,omitempty"`
	Identifier     string  `json:"identifier"`
	Text           string  `json:"text"`
	Type           string  `json:"type"`
	Context        string  `json:"context"`
	MaxLength      int     `json:"maxLength"`
	IsHidden       bool    `json:"isHidden"`
	IsDuplicate    bool    `json:"isDuplicate"`
	MasterStringID *int    `json:"masterStringId,omitempty"`
	LabelIDs       []int   `json:"labelIds"`
	WebURL         string  `json:"webUrl"`
	CreatedAt      *string `json:"createdAt,omitempty"`
	UpdatedAt      *string `json:"updatedAt,omitempty"`
	FileID         int     `json:"fileId"`
	DirectoryID    *int    `json:"directoryId,omitempty"`
	Revision       int     `json:"revision"`
}

// SourceStringsGetResponse describes the response when getting
// a source string.
type SourceStringsGetResponse struct {
	Data *SourceString `json:"data"`
}

// SourceStringsListResponse describes the response when getting
// a list of source strings.
type SourceStringsListResponse struct {
	Data []*SourceStringsGetResponse `json:"data"`
}

// SourceStringsListOptions specifies the optional parameters
// to the SourceStringsService.List method.
type SourceStringsListOptions struct {
	// Enable denormalize placeholders. Enum: 0 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`
	// Filter strings by labelIds (Label Identifiers).
	// Example: labelIds=1,2,3,4,5.
	LabelIDs string `json:"labelIds,omitempty"`
	// File Identifier.
	// Note: Can't be used with `directoryId` or `branchId` in same request.
	FileID int `json:"fileId,omitempty"`
	// Branch Identifier.
	// Note: Can't be used with `fileId` or `directoryId` in the same request.
	BranchID int `json:"branchId,omitempty"`
	// Directory Identifier.
	// Note: Can't be used with `fileId` or `branchId` in same request.
	DirectoryID int `json:"directoryId,omitempty"`
	// Filter strings by CroQL.
	// Note: Can be used only with `denormalizePlaceholders`, `offset` and
	//       `limit` in same request.
	CroQL string `json:"croql,omitempty"`
	// Filter strings by `identifier`, `text` or `context`.
	Filter string `json:"filter,omitempty"`
	// Specify field to be the target of filtering. It can be one scope or
	// a list of comma-separated scopes. Enum: "identifier" "text" "context".
	Scope string `json:"scope,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of SourceStringListOptions.
func (o *SourceStringsListOptions) Values() url.Values {
	u := o.ListOptions.Values()
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		u.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	if o.LabelIDs != "" {
		u.Add("labelIds", o.LabelIDs)
	}
	if o.FileID > 0 {
		u.Add("fileId", fmt.Sprintf("%d", o.FileID))
	}
	if o.BranchID > 0 {
		u.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}
	if o.DirectoryID > 0 {
		u.Add("directoryId", fmt.Sprintf("%d", o.DirectoryID))
	}
	if o.CroQL != "" {
		u.Add("croql", o.CroQL)
	}
	if o.Filter != "" {
		u.Add("filter", o.Filter)
	}
	if o.Scope != "" {
		u.Add("scope", o.Scope)
	}
	return u
}

// SourceStringsGetOptions specifies the optional parameters
// to the SourceStringsService.Get method.
type SourceStringsGetOptions struct {
	// Enable denormalize placeholders. Enum: 0 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`
}

// Values returns the url.Values representation of SourceStringsGetOptions.
func (o *SourceStringsGetOptions) Values() url.Values {
	u := url.Values{}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		u.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	return u
}

// SourcseStringsAddRequest defines the structure of a request
// to add a string.
type SourceStringsAddRequest struct {
	// Text for translation.
	// It can be a string or map of strings.
	// Example:
	// "text": "Not all videos are shown to users. See more"
	// or
	// "text": {
	//  "one": "string",
	//  "other": "strings"
	// }
	Text any `json:"text"`
	// File identifier.
	FileID int `json:"fileId"`
	// Defines unique string identifier.
	Identifier string `json:"identifier,omitempty"`
	// Use to provide additional information for better source text understanding.
	Context string `json:"context,omitempty"`
	// Defines whether to make string unavailable for translation. Default: false.
	IsHidden *bool `json:"isHidden,omitempty"`
	// Max. length of translated text (0 – unlimited).
	MaxLength *int `json:"maxLength,omitempty"`
	// Label Identifiers.
	LabelIDs []int `json:"labelIds,omitempty"`
	// Fields (enterprises only).
	Fields map[string]string `json:"fields,omitempty"`
}

// Validate checks if the add request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *SourceStringsAddRequest) Validate() error {
	if r == nil {
		return errors.New("request cannot be nil")
	}

	// check if `text` is a string or map of strings
	switch text := r.Text.(type) {
	case string:
		if r.Text == "" {
			return errors.New("text cannot be empty")
		}
	case map[string]string:
		if len(text) == 0 {
			return errors.New("text cannot be empty")
		}
	default:
		return errors.New("text must be a string or map of strings")
	}

	if r.FileID == 0 {
		return errors.New("fileId is required")
	}
	return nil
}

// SourceStringsService represents the upload strings status.
type SourceStringsUpload struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
	Attributes struct {
		BranchID      int    `json:"branchId"`
		SotrageID     int    `json:"storageId"`
		FileType      string `json:"fileType"`
		ParserVersion int    `json:"parserVersion"`
		LabelIDs      []int  `json:"labelIds"`
		ImportOptions struct {
			FirstLineContainsHeader bool           `json:"firstLineContainsHeader"`
			ImportTranslations      bool           `json:"importTranslations"`
			Scheme                  map[string]int `json:"scheme"`
		} `json:"importOptions"`
		UpdateStrings bool `json:"updateStrings"`
		CleanupMode   bool `json:"cleanupMode"`
	} `json:"attributes"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
}

// SourceStringsUploadResponse defines the response when
// uploading strings.
type SourceStringsUploadResponse struct {
	Data *SourceStringsUpload `json:"data"`
}

// SourceStringsUploadRequest defines the structure of a request
// to upload strings.
type SourceStringsUploadRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// Branch Identifier.
	// Defines branch to which file will be added.
	BranchID int `json:"branchId"`
	// Default: auto
	// Enum: auto, android, macosx, arb, csv, json, xlsx, xliff, xliff_two
	// - empty value or `auto` — Try to detect file type by extension or MIME type
	// - `android` — Android (*.xml)
	// - `macosx` — Mac OS X / iOS (*.strings)
	// - `arb` — Application Resource Bundle (*.arb)
	// - `csv` — Comma Separated Values (*.csv)
	// - `json` — Generic JSON (*.json)
	// - `xliff` — XLIFF (*.xliff, *.xlf)
	// - `xliff_two` — XLIFF 2.0 (*.xliff, *.xlf)
	// - `xlsx` — Microsoft Excel (*.xlsx)
	Type string `json:"type,omitempty"`
	// Using latest parser version by default.
	// Note: Must be used together with `type`.
	ParserVersion int `json:"parserVersion,omitempty"`
	// Attach labels to strings.
	LabelIDs []int `json:"labelIds,omitempty"`
	// Update strings that have the same keys. Default: false.
	UpdateStrings *bool `json:"updateStrings,omitempty"`
	// If true, all strings with a system label that do not exist in the file
	// will be deleted. Default: false.
	CleanupMode *bool `json:"cleanupMode,omitempty"`
	// Options for importing strings.
	ImportOptions *SourceStringsImportOptions `json:"importOptions,omitempty"`
}

// SourceStringsImportOptions defines the options for importing strings.
type SourceStringsImportOptions struct {
	// Defines whether the file includes a first-row header that should
	// not be imported. Default: false.
	FirstLineContainsHeader *bool `json:"firstLineContainsHeader,omitempty"`
	// Defines whether to import translations from the file. Default: false.
	ImportTranslations *bool `json:"importTranslations,omitempty"`
	// Defines data columns mapping. The key is the column name and the value
	// is the column index. The column numbering starts at 0.
	Scheme map[string]int `json:"scheme,omitempty"`
}

// Validate checks if the upload request is valid.
// It implements the crowdin.RequestValidator interface.
func (o *SourceStringsUploadRequest) Validate() error {
	if o == nil {
		return errors.New("request cannot be nil")
	}
	if o.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if o.BranchID == 0 {
		return errors.New("branchId is required")
	}
	return nil
}