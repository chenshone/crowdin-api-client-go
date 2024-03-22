package model

import (
	"fmt"
	"net/url"
)

type (
	// PreTranslation represents a pre-translation status.
	PreTranslation struct {
		Identifier string                    `json:"identifier"`
		Status     string                    `json:"status"`
		Progress   int64                     `json:"progress"`
		Attributes *PreTranslationAttributes `json:"attributes"`
		CreatedAt  string                    `json:"createdAt"`
		UpdatedAt  string                    `json:"updatedAt"`
		StartedAt  string                    `json:"startedAt,omitempty"`
		FinishedAt string                    `json:"finishedAt,omitempty"`
	}

	PreTranslationAttributes struct {
		LanguageIDs                   []string `json:"languageIds"`
		FileIDs                       []int64  `json:"fileIds"`
		Method                        *string  `json:"method,omitempty"`
		AutoApproveOption             *string  `json:"autoApproveOption,omitempty"`
		DuplicateTranslations         *bool    `json:"duplicateTranslations,omitempty"`
		SkipApprovedTranslations      *bool    `json:"skipApprovedTranslations,omitempty"`
		TranslateUntranslatedOnly     *bool    `json:"translateUntranslatedOnly,omitempty"`
		TranslateWithPerfectMatchOnly *bool    `json:"translateWithPerfectMatchOnly,omitempty"`
	}
)

// PreTranslationsResponse defines the structure of a response when
// getting a pre-translation status.
type PreTranslationsResponse struct {
	Data *PreTranslation `json:"data"`
}

// PreTranslationRequest defines the structure of a request to apply pre-translation.
type PreTranslationRequest struct {
	// Set of languages to which pre-translation should be applied.
	LanguageIDs []string `json:"languageIds"`
	// Files array that should be translated.
	FileIDs []int64 `json:"fileIds"`
	// Defines pre-translation method. Enum: "tm", "mt". Default: "tm".
	// - tm – pre-translation via Translation Memory.
	// - mt – pre-translation via Machine Translation. "mt" should be used with `engineId` parameter.
	Method string `json:"method,omitempty"`
	// Machine Translation engine Identifier. Required if `method` is set to "mt".
	EngineID int64 `json:"engineId,omitempty"`
	// Defines which translations added by TM pre-translation should be auto-approved.  Default: "none".
	// Enum: "all", "exceptAutoSubstituted", "perfectMatchApprovedOnly", "perfectMatchOnly", "none"
	// - all – all
	// - perfectMatchOnly – with perfect TM match
	// - exceptAutoSubstituted – all (skip auto-substituted suggestions)
	// - perfectMatchApprovedOnly - with perfect TM match (approved previously)
	// - none – no auto-approve
	AutoApproveOption string `json:"autoApproveOption,omitempty"`
	// Adds translations even if the same translation already exists. Default is false.
	// Note: Works only with TM pre-translation method.
	DuplicateTranslations bool `json:"duplicateTranslations,omitempty"`
	// Skip approved translations. Default is false.
	// Note: Works only with TM pre-translation method.
	SkipApprovedTranslations bool `json:"skipApprovedTranslations,omitempty"`
	// Applies pre-translation for untranslated strings only. Default is true.
	// Note: Works only with TM pre-translation method.
	TranslateUntranslatedOnly bool `json:"translateUntranslatedOnly,omitempty"`
	// Applies pre-translation only for the strings with perfect match
	// (source text and contextual information are identical). Default is false.
	// Note: Works only with TM pre-translation method.
	TranslateWithPerfectMatchOnly bool `json:"translateWithPerfectMatchOnly,omitempty"`
	// Defines fallback languages mapping. The passed value should contain a map of
	// language ID as a key and an array of fallback language IDs as a value.
	// Note: Available only for TM Pre-Translation.
	FallbackLanguages map[string][]string `json:"fallbackLanguages,omitempty"`
	// Label Identifiers.
	LabelIDs []int64 `json:"labelIds,omitempty"`
	// Exclude Label Identifiers.
	ExcludeLabelIDs []int64 `json:"excludeLabelIds,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PreTranslationRequest) Validate() error {
	if len(r.LanguageIDs) == 0 {
		return fmt.Errorf("languageIds is required")
	}
	if len(r.FileIDs) == 0 {
		return fmt.Errorf("fileIds is required")
	}
	return nil
}

// BuildProjectDirectoryTranslationRequest defines the structure of a request
// to build project directory translation.
type BuildProjectDirectoryTranslationRequest struct {
	// Specify target languages for build.
	// Leave this field empty to build all target languages.
	TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
	// Defines whether to export only translated strings. Default: false.
	// Note: true value can't be used with `skipUntranslatedFiles=true` in same request .
	SkipUntranslatedStrings bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated file. Default: false.
	// Note: true value can't be used with `skipUntranslatedStrings=true`` in same request.
	SkipUntranslatedFiles bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default: false.
	ExportApprovedOnly bool `json:"exportApprovedOnly,omitempty"`
	// Preserve folder hierarchy. Default: false.
	PreserveFolderHierarchy bool `json:"preserveFolderHierarchy,omitempty"`
}

// Validate checks if the build project directory translation request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectDirectoryTranslationRequest) Validate() error {
	if r.SkipUntranslatedStrings && r.SkipUntranslatedFiles {
		return fmt.Errorf("skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request")
	}
	return nil
}

// BuildProjectDirectoryTranslation represents a project directory build.
type BuildProjectDirectoryTranslation struct {
	ID         int64  `json:"id"`
	ProjectID  int64  `json:"projectId"`
	Status     string `json:"status"`
	Progress   int64  `json:"progress"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	FinishedAt string `json:"finishedAt,omitempty"`
}

// BuildProjectFileTranslationRequest defines the structure of a request
// to build project file translation.
type BuildProjectFileTranslationRequest struct {
	// Target Language Identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// Defines whether to export only translated strings. Default: false.
	// Note: true value can't be used with `skipUntranslatedFiles=true`` in same request.
	SkipUntranslatedStrings bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated file. Default: false.
	// Note: true value can't be used with skipUntranslatedStrings=true in same request.
	SkipUntranslatedFiles bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default: false.
	ExportApprovedOnly bool `json:"exportApprovedOnly,omitempty"`
}

// Validate checks if the build project file translation request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectFileTranslationRequest) Validate() error {
	if len(r.TargetLanguageID) == 0 {
		return fmt.Errorf("targetLanguageId is required")
	}
	if r.SkipUntranslatedStrings && r.SkipUntranslatedFiles {
		return fmt.Errorf("skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request")
	}
	return nil
}

// TranslationsBuildsListOptions specifies the optional parameters to the
type TranslationsBuildsListOptions struct {
	ListOptions

	// Branch Identifier. Filter builds by branchId.
	BranchID int64 `url:"branchId,omitempty"`
}

// Values returns the url.Values representation of the query options.
func (p *TranslationsBuildsListOptions) Values() url.Values {
	v := p.ListOptions.Values()
	if p.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", p.BranchID))
	}
	return v
}

// TranslationsProjectBuild represents a project build.
type TranslationsProjectBuild struct {
	ID         int64  `json:"id"`
	ProjectID  int64  `json:"projectId"`
	Status     string `json:"status"`
	Progress   int64  `json:"progress"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	FinishedAt string `json:"finishedAt,omitempty"`
	Attributes struct {
		BranchID                *int64   `json:"branchId,omitempty"`
		DirectoryID             *int64   `json:"directoryId,omitempty"`
		TargetLanguageIDs       []string `json:"targetLanguageIds,omitempty"`
		SkipUntranslatedStrings bool     `json:"skipUntranslatedStrings,omitempty"`
		SkipUntranslatedFiles   bool     `json:"skipUntranslatedFiles,omitempty"`
		ExportApprovedOnly      bool     `json:"exportApprovedOnly,omitempty"`
	} `json:"attributes,omitempty"`
}

// TranslationsProjectBuildResponse defines the structure of a response when
// getting a project build.
type TranslationsProjectBuildResponse struct {
	Data *TranslationsProjectBuild `json:"data"`
}

// TranslationsProjectBuildsListResponse defines the structure of a response when
// getting a list of project builds.
type TranslationsProjectBuildsListResponse struct {
	Data       []*TranslationsProjectBuildResponse `json:"data"`
	Pagination *Pagination                         `json:"pagination"`
}

type (
	// BuildProjectTranslationRequest interface that allows accepting
	// BuildProjectRequest and PseudoBuildProjectRequest types.
	BuildProjectTranslationRequest interface {
		ValidateBuildRequest() error
	}

	// BuildProjectRequest defines the structure of a request to build a project.
	BuildProjectRequest struct {
		// Branch Identifier.
		BranchID int64 `json:"branchId,omitempty"`
		// Specify target languages for build.
		// Leave this field empty to build all target languages
		TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
		// Defines whether to export only translated strings.
		// Note: true value can't be used with `skipUntranslatedFiles=true` in same request.
		SkipUntranslatedStrings bool `json:"skipUntranslatedStrings,omitempty"`
		// Defines whether to export only translated files.
		// Note: true value can't be used with `skipUntranslatedStrings=true` in same request.
		SkipUntranslatedFiles bool `json:"skipUntranslatedFiles,omitempty"`
		// Defines whether to export only approved strings.
		ExportApprovedOnly bool `json:"exportApprovedOnly,omitempty"`
	}

	// PsuedoBuildProjectRequest defines the structure of a request to build a project
	// with pseudo translations.
	PseudoBuildProjectRequest struct {
		// Flag for detecting pseudo translation. Default: false.
		Pseudo bool `json:"pseudo"`
		// Branch Identifier.
		BranchID int64 `json:"branchId,omitempty"`
		// Add special characters at the beginning of each string to show
		// where messages have been concatenated together.
		Prefix string `json:"prefix,omitempty"`
		// Add special characters at the end of each string to show where
		// messages have been concatenated together.
		Suffix string `json:"suffix,omitempty"`
		// Make string larger or shorter.
		// Acceptable values must be from -50 to 100. Default is 0.
		LengthTransformation int64 `json:"lengthTransformation,omitempty"`
		// Transforms characters to other languages.
		// Enum: "asian", "cyrillic", "european", "arabic".
		CharTransformation string `json:"charTransformation,omitempty"`
	}
)

// ValidateBuildRequest implements the BuildProjectTranslationRequest interface.
func (r *BuildProjectRequest) ValidateBuildRequest() error {
	if r.SkipUntranslatedStrings && r.SkipUntranslatedFiles {
		return fmt.Errorf("`skipUntranslatedStrings` and `skipUntranslatedFiles` must not be true at the same request")
	}
	return nil
}

// Validate checks if the build project request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PseudoBuildProjectRequest) Validate() error {
	return r.ValidateBuildRequest()
}

// PseudoBuildProjectRequest implements the BuildProjectTranslationRequest interface.
func (r *PseudoBuildProjectRequest) ValidateBuildRequest() error {
	if r.LengthTransformation < -50 || r.LengthTransformation > 100 {
		return fmt.Errorf("lengthTransformation must be from -50 to 100")
	}
	return nil
}

// Validate checks if the build project request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectRequest) Validate() error {
	return r.ValidateBuildRequest()
}

// UploadTranslationsRequest defines the structure of a request to upload translations.
type UploadTranslationsRequest struct {
	// Storage Identifier.
	StorageID int64 `json:"storageId"`
	// File Identifier for import.
	// Note: Required for content in all formats except XLIFF.
	FileID int64 `json:"fileId,omitempty"`
	// Branch Identifier for import.
	// Note: Required for string based API.
	BranchID int64 `json:"branchId,omitempty"`
	// Defines whether to add translation if it's the same as the source string.
	// Default: false.
	ImportEqSuggestions bool `json:"importEqSuggestions,omitempty"`
	// Mark uploaded translations as approved. Default: false.
	AutoApproveImported bool `json:"autoApproveImported,omitempty"`
	// Allow translations upload to hidden source strings. Default: false.
	TranslateHidden bool `json:"translateHidden,omitempty"`
}

// Validate checks if the upload translations request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *UploadTranslationsRequest) Validate() error {
	if r.StorageID == 0 {
		return fmt.Errorf("storageId is required")
	}
	if r.FileID > 0 && r.BranchID > 0 {
		return fmt.Errorf("fileId and branchId can not be used at the same request")
	}
	return nil
}

type (
	// UploadTranslations represents the uploaded translations.
	UploadTranslations struct {
		ProjectID  int64  `json:"projectId"`
		StorageID  int64  `json:"storageId"`
		LanguageID string `json:"languageId"`
		FileID     int64  `json:"fileId"`
	}

	// UploadTranslationsResponse defines the structure of a response when
	// uploading translations.
	UploadTranslationsResponse struct {
		Data UploadTranslations `json:"data"`
	}
)

type ExportTranslationRequest struct {
	// Specify target language for export.
	TargetLanguageID string `json:"targetLanguageId"`
	// Defines export file format. Use API Type feature specified at the
	// corresponding file format from Crowdin Store.
	// Note: the `format` parameter is required in all cases except when you'd like
	// to export translations for a single file in its original format.
	Format string `json:"format,omitempty"`
	// Label Identifiers.
	LabelIDs []int64 `json:"labelIds,omitempty"`
	// Branch Identifiers.
	// Note: Can't be used with `directoryIds` or `fileIds` in same request.
	BranchIDs []int64 `json:"branchIds,omitempty"`
	// Directory Identifiers.
	// Note: Can't be used with `branchIds` or `fileIds` in same request.
	DirectoryIDs []int64 `json:"directoryIds,omitempty"`
	// File Identifiers.
	// Note: Can't be used with `branchIds` or `directoryIds` in same request.
	FileIds []int64 `json:"fileIds,omitempty"`
	// Defines whether to export only translated strings. Default is false.
	// Note: Can't be used with `skipUntranslatedFiles` in same request.
	SkipUntranslatedStrings bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only approved strings. Default is false.
	ExportApprovedOnly bool `json:"exportApprovedOnly,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ExportTranslationRequest) Validate() error {
	if r.TargetLanguageID == "" {
		return fmt.Errorf("targetLanguageId is required")
	}
	return nil
}

type (
	// DownloadLink represents a download link.
	DownloadLink struct {
		URL      string `json:"url"`
		ExpireIn string `json:"expireIn"`

		Etag *string `json:"etag,omitempty"`
	}

	// DownloadLinkResponse defines the structure of a response when
	// getting a download URL with its expiration time.
	DownloadLinkResponse struct {
		Data *DownloadLink `json:"data"`
	}
)