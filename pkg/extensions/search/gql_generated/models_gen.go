// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql_generated

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// Annotation is Key:Value pair representing custom data which is otherwise
// not available in other fields.
type Annotation struct {
	// Custom key
	Key *string `json:"Key,omitempty"`
	// Value associated with the custom key
	Value *string `json:"Value,omitempty"`
}

// Contains various details about the CVE (Common Vulnerabilities and Exposures)
// and a list of PackageInfo about the affected packages
type Cve struct {
	// CVE ID
	ID *string `json:"Id,omitempty"`
	// A short title describing the CVE
	Title *string `json:"Title,omitempty"`
	// A detailed description of the CVE
	Description *string `json:"Description,omitempty"`
	// Reference for the given CVE
	Reference *string `json:"Reference,omitempty"`
	// The impact the CVE has, one of "UNKNOWN", "LOW", "MEDIUM", "HIGH", "CRITICAL"
	Severity *string `json:"Severity,omitempty"`
	// Information on the packages in which the CVE was found
	PackageList []*PackageInfo `json:"PackageList,omitempty"`
}

// Contains the tag of the image and a list of CVEs
type CVEResultForImage struct {
	// Tag affected by the CVEs
	Tag *string `json:"Tag,omitempty"`
	// List of CVE objects which affect this specific image:tag
	CVEList []*Cve `json:"CVEList,omitempty"`
	// Counters for the different severities of the CVEs
	SeverityCounters []*SeverityCounter `json:"SeverityCounters,omitempty"`
	// The CVE pagination information, see PageInfo object for more details
	Page *PageInfo `json:"Page,omitempty"`
}

// Apply various types of filters to the queries made for repositories and images
// For example we only want to display repositories which contain images with
// a certain OS ar Architecture.
type Filter struct {
	// Only return images or repositories supporting the operating systems in the list
	// Should be values listed in the Go Language document https://go.dev/doc/install/source#environment
	Os []*string `json:"Os,omitempty"`
	// Only return images or repositories supporting the build architectures in the list
	// Should be values listed in the Go Language document https://go.dev/doc/install/source#environment
	Arch []*string `json:"Arch,omitempty"`
	// Only return images or repositories with at least one signature
	HasToBeSigned *bool `json:"HasToBeSigned,omitempty"`
	// Only returns images or repositories that are bookmarked or not bookmarked
	IsBookmarked *bool `json:"IsBookmarked,omitempty"`
	// Only returns images or repositories that are starred or not starred
	IsStarred *bool `json:"IsStarred,omitempty"`
}

// Search results, can contain images, repositories and layers
type GlobalSearchResult struct {
	// Pagination information
	Page *PageInfo `json:"Page,omitempty"`
	// List of images matching the search criteria
	Images []*ImageSummary `json:"Images,omitempty"`
	// List of repositories matching the search criteria
	Repos []*RepoSummary `json:"Repos,omitempty"`
	// List of layers matching the search criteria
	// NOTE: the actual search logic for layers is not implemented at the moment
	Layers []*LayerSummary `json:"Layers,omitempty"`
}

// Information on how a layer was created
type HistoryDescription struct {
	// Created is the time when the layer was created.
	Created *time.Time `json:"Created,omitempty"`
	// CreatedBy is the command which created the layer.
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// Author is the author of the build point.
	Author *string `json:"Author,omitempty"`
	// Comment is a custom message set when creating the layer.
	Comment *string `json:"Comment,omitempty"`
	// EmptyLayer is used to mark if the history item created a filesystem diff.
	EmptyLayer *bool `json:"EmptyLayer,omitempty"`
}

// Details about a specific image, it is used by queries returning a list of images
// We define an image as a pairing or a repository and a tag belonging to that repository
type ImageSummary struct {
	// Name of the repository where the image is found
	RepoName *string `json:"RepoName,omitempty"`
	// Tag identifying the image within the repository
	Tag *string `json:"Tag,omitempty"`
	// The digest of the descriptor of this image
	Digest *string `json:"Digest,omitempty"`
	// The media type of the descriptor of this image
	MediaType *string `json:"MediaType,omitempty"`
	// List of manifests for all supported versions of the image for different operating systems and architectures
	Manifests []*ManifestSummary `json:"Manifests,omitempty"`
	// Total size of the files associated with all images (manifest, config, layers)
	Size *string `json:"Size,omitempty"`
	// Number of downloads of the manifest of this image
	DownloadCount *int `json:"DownloadCount,omitempty"`
	// Timestamp of the last modification done to the image (from config or the last updated layer)
	LastUpdated *time.Time `json:"LastUpdated,omitempty"`
	// Human-readable description of the software packaged in the image
	Description *string `json:"Description,omitempty"`
	// True if the image has a signature associated with it, false otherwise
	IsSigned *bool `json:"IsSigned,omitempty"`
	// Info about signature validity
	SignatureInfo []*SignatureSummary `json:"SignatureInfo,omitempty"`
	// License(s) under which contained software is distributed as an SPDX License Expression
	Licenses *string `json:"Licenses,omitempty"`
	// Labels associated with this image
	// NOTE: currently this field is unused
	Labels *string `json:"Labels,omitempty"`
	// Human-readable title of the image
	Title *string `json:"Title,omitempty"`
	// URL to get source code for building the image
	Source *string `json:"Source,omitempty"`
	// URL to get documentation on the image
	Documentation *string `json:"Documentation,omitempty"`
	// Vendor associated with this image, the distributing entity, organization or individual
	Vendor *string `json:"Vendor,omitempty"`
	// Contact details of the people or organization responsible for the image
	Authors *string `json:"Authors,omitempty"`
	// Short summary of the identified CVEs
	Vulnerabilities *ImageVulnerabilitySummary `json:"Vulnerabilities,omitempty"`
	// Information about objects that reference this image
	Referrers []*Referrer `json:"Referrers,omitempty"`
}

// Contains summary of vulnerabilities found in a specific image
type ImageVulnerabilitySummary struct {
	// Maximum severity of all CVEs found in this image
	MaxSeverity *string `json:"MaxSeverity,omitempty"`
	// Count of all CVEs found in this image
	Count *int `json:"Count,omitempty"`
}

// Information about how/when a layer was built
type LayerHistory struct {
	// Information specific to the layer such as size and digest.
	Layer *LayerSummary `json:"Layer,omitempty"`
	// Additional information about how the layer was created.
	HistoryDescription *HistoryDescription `json:"HistoryDescription,omitempty"`
}

// Contains details about a specific layer which is part of an image
type LayerSummary struct {
	// The size of the layer in bytes
	Size *string `json:"Size,omitempty"`
	// Digest of the layer content
	Digest *string `json:"Digest,omitempty"`
}

// Details about a specific version of an image for a certain operating system and architecture.
type ManifestSummary struct {
	// Digest of the manifest file associated with this image
	Digest *string `json:"Digest,omitempty"`
	// Digest of the config file associated with this image
	ConfigDigest *string `json:"ConfigDigest,omitempty"`
	// Timestamp of the last update to an image inside this repository
	LastUpdated *time.Time `json:"LastUpdated,omitempty"`
	// Total size of the files associated with this manifest (manifest, config, layers)
	Size *string `json:"Size,omitempty"`
	// True if the manifest has a signature associated with it, false otherwise
	IsSigned *bool `json:"IsSigned,omitempty"`
	// Info about signature validity
	SignatureInfo []*SignatureSummary `json:"SignatureInfo,omitempty"`
	// OS and architecture supported by this image
	Platform *Platform `json:"Platform,omitempty"`
	// Total number of image manifest downloads from this repository
	DownloadCount *int `json:"DownloadCount,omitempty"`
	// List of layers matching the search criteria
	// NOTE: the actual search logic for layers is not implemented at the moment
	Layers []*LayerSummary `json:"Layers,omitempty"`
	// Information about the history of the specific image, see LayerHistory
	History []*LayerHistory `json:"History,omitempty"`
	// Short summary of the identified CVEs
	Vulnerabilities *ImageVulnerabilitySummary `json:"Vulnerabilities,omitempty"`
	// Information about objects that reference this image
	Referrers []*Referrer `json:"Referrers,omitempty"`
	// Value of the artifactType field if present else the value of the config media type
	ArtifactType *string `json:"ArtifactType,omitempty"`
}

// Contains the name of the package, the current installed version and the version where the CVE was fixed
type PackageInfo struct {
	// Name of the package affected by a CVE
	Name *string `json:"Name,omitempty"`
	// Current version of the package, typically affected by the CVE
	InstalledVersion *string `json:"InstalledVersion,omitempty"`
	// Minimum version of the package in which the CVE is fixed
	FixedVersion *string `json:"FixedVersion,omitempty"`
}

// Information on current page returned by the API
type PageInfo struct {
	// The total number of objects on all pages
	TotalCount int `json:"TotalCount"`
	// The number of objects in this page
	ItemCount int `json:"ItemCount"`
}

// Pagination parameters
// If PageInput is empty, the request should return all objects.
type PageInput struct {
	// The maximum amount of results to return for this page
	// Negative values are not allowed
	Limit *int `json:"limit,omitempty"`
	// The results page number you want to receive
	// Negative values are not allowed
	Offset *int `json:"offset,omitempty"`
	// The criteria used to sort the results on the page
	SortBy *SortCriteria `json:"sortBy,omitempty"`
}

// Paginated list of ImageSummary objects
type PaginatedImagesResult struct {
	// Information on the returned page
	Page *PageInfo `json:"Page,omitempty"`
	// List of images
	Results []*ImageSummary `json:"Results"`
}

// Paginated list of RepoSummary objects
type PaginatedReposResult struct {
	// Information on the returned page
	Page *PageInfo `json:"Page,omitempty"`
	// List of repositories
	Results []*RepoSummary `json:"Results"`
}

// Contains details about the OS and architecture of the image
type Platform struct {
	// The name of the operating system which the image is built to run on,
	// Should be values listed in the Go Language document https://go.dev/doc/install/source#environment
	Os *string `json:"Os,omitempty"`
	// The name of the compilation architecture which the image is built to run on,
	// Should be values listed in the Go Language document https://go.dev/doc/install/source#environment
	Arch *string `json:"Arch,omitempty"`
}

// A referrer is an object which has a reference to a another object
type Referrer struct {
	// Referrer MediaType
	// See https://github.com/opencontainers/artifacts for more details
	MediaType *string `json:"MediaType,omitempty"`
	// Referrer ArtifactType
	// See https://github.com/opencontainers/artifacts for more details
	ArtifactType *string `json:"ArtifactType,omitempty"`
	// Total size of the referrer files in bytes
	Size *int `json:"Size,omitempty"`
	// Digest of the manifest file of the referrer
	Digest *string `json:"Digest,omitempty"`
	// A list of annotations associated with this referrer
	Annotations []*Annotation `json:"Annotations"`
}

// Contains details about the repo: both general information on the repo, and the list of images
type RepoInfo struct {
	// List of images in the repo
	Images []*ImageSummary `json:"Images,omitempty"`
	// Details about the repository itself
	Summary *RepoSummary `json:"Summary,omitempty"`
}

// Details of a specific repo, it is used by queries returning a list of repos
type RepoSummary struct {
	// Name of the repository
	Name *string `json:"Name,omitempty"`
	// Timestamp of the last update to an image inside this repository
	LastUpdated *time.Time `json:"LastUpdated,omitempty"`
	// Total size of the files within this repository
	Size *string `json:"Size,omitempty"`
	// List of platforms supported by this repository
	Platforms []*Platform `json:"Platforms,omitempty"`
	// Vendors associated with this image, the distributing entities, organizations or individuals
	Vendors []*string `json:"Vendors,omitempty"`
	// Details of the newest image inside the repository
	// NOTE: not the image with the `latest` tag, the one with the most recent created timestamp
	NewestImage *ImageSummary `json:"NewestImage,omitempty"`
	// Total number of image manifest downloads from this repository
	DownloadCount *int `json:"DownloadCount,omitempty"`
	// Number of stars attributed to this repository by users
	StarCount *int `json:"StarCount,omitempty"`
	// True if the repository is bookmarked by the current user, false otherwise
	IsBookmarked *bool `json:"IsBookmarked,omitempty"`
	// True if the repository is starred by the current user, false otherwise
	IsStarred *bool `json:"IsStarred,omitempty"`
	// Rank represents how good the match was between the queried repo name and this repo summary.
	Rank *int `json:"Rank,omitempty"`
}

// Define how many vulnerabilities for the given severity exist
type SeverityCounter struct {
	// The impact the CVE has, one of "UNKNOWN", "LOW", "MEDIUM", "HIGH", "CRITICAL"
	Severity string `json:"Severity"`
	// Count of all CVEs found in this image with the given severity
	Count int `json:"Count"`
}

// Contains details about the signature
type SignatureSummary struct {
	// Tool is the tool used for signing image
	Tool *string `json:"Tool,omitempty"`
	// True if the signature is trusted, false otherwise
	IsTrusted *bool `json:"IsTrusted,omitempty"`
	// Author is the author of the signature
	Author *string `json:"Author,omitempty"`
}

// All sort criteria usable with pagination, some of these criteria applies only
// to certain queries. For example sort by severity is available for CVEs but not
// for repositories
type SortCriteria string

const (
	// How relevant the result is based on the user input used while searching
	// Applies to: images and repositories
	SortCriteriaRelevance SortCriteria = "RELEVANCE"
	// Sort by the most recently created timestamp of the images
	// Applies to: images and repositories
	SortCriteriaUpdateTime SortCriteria = "UPDATE_TIME"
	// Sort alphabetically ascending
	// Applies to: images, repositories and CVEs
	SortCriteriaAlphabeticAsc SortCriteria = "ALPHABETIC_ASC"
	// Sort alphabetically descending
	// Applies to: images, repositories and CVEs
	SortCriteriaAlphabeticDsc SortCriteria = "ALPHABETIC_DSC"
	// Sort from the most severe to the least severe
	// Applies to: CVEs
	SortCriteriaSeverity SortCriteria = "SEVERITY"
	// Sort by the total number of stars given by users
	// Applies to: repositories
	SortCriteriaStars SortCriteria = "STARS"
	// Sort by the total download count
	// Applies to: repositories and images
	SortCriteriaDownloads SortCriteria = "DOWNLOADS"
)

var AllSortCriteria = []SortCriteria{
	SortCriteriaRelevance,
	SortCriteriaUpdateTime,
	SortCriteriaAlphabeticAsc,
	SortCriteriaAlphabeticDsc,
	SortCriteriaSeverity,
	SortCriteriaStars,
	SortCriteriaDownloads,
}

func (e SortCriteria) IsValid() bool {
	switch e {
	case SortCriteriaRelevance, SortCriteriaUpdateTime, SortCriteriaAlphabeticAsc, SortCriteriaAlphabeticDsc, SortCriteriaSeverity, SortCriteriaStars, SortCriteriaDownloads:
		return true
	}
	return false
}

func (e SortCriteria) String() string {
	return string(e)
}

func (e *SortCriteria) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortCriteria(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortCriteria", str)
	}
	return nil
}

func (e SortCriteria) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
