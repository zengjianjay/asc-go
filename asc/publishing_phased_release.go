package asc

import (
	"context"
	"fmt"
	"time"
)

// PhasedReleaseState defines model for PhasedReleaseState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/phasedreleasestate
type PhasedReleaseState string

const (
	// PhasedReleaseStateInactive is a representation of the INACTIVE state
	PhasedReleaseStateInactive PhasedReleaseState = "INACTIVE"
	// PhasedReleaseStateActive is a representation of the ACTIVE state
	PhasedReleaseStateActive PhasedReleaseState = "ACTIVE"
	// PhasedReleaseStatePaused is a representation of the PAUSED state
	PhasedReleaseStatePaused PhasedReleaseState = "PAUSED"
	// PhasedReleaseStateComplete is a representation of the COMPLETE state
	PhasedReleaseStateComplete PhasedReleaseState = "COMPLETE"
)

// AppStoreVersionPhasedRelease defines model for AppStoreVersionPhasedRelease.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedrelease
type AppStoreVersionPhasedRelease struct {
	Attributes *AppStoreVersionPhasedReleaseAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Links      ResourceLinks                           `json:"links"`
	Type       string                                  `json:"type"`
}

// AppStoreVersionPhasedReleaseAttributes defines model for AppStoreVersionPhasedRelease.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedrelease/attributes
type AppStoreVersionPhasedReleaseAttributes struct {
	CurrentDayNumber   *int                `json:"currentDayNumber,omitempty"`
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
	StartDate          *time.Time          `json:"startDate,omitempty"`
	TotalPauseDuration *int                `json:"totalPauseDuration,omitempty"`
}

// appStoreVersionPhasedReleaseCreateRequest defines model for appStoreVersionPhasedReleaseCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest
type appStoreVersionPhasedReleaseCreateRequest struct {
	Attributes    *appStoreVersionPhasedReleaseCreateRequestAttributes   `json:"attributes,omitempty"`
	Relationships appStoreVersionPhasedReleaseCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// AppStoreVersionPhasedReleaseCreateRequestAttributes are attributes for AppStoreVersionPhasedReleaseCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest/data/attributes
type appStoreVersionPhasedReleaseCreateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// AppStoreVersionPhasedReleaseCreateRequestRelationships are relationships for AppStoreVersionPhasedReleaseCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleasecreaterequest/data/relationships
type appStoreVersionPhasedReleaseCreateRequestRelationships struct {
	AppStoreVersion RelationshipDeclaration `json:"appStoreVersion"`
}

// AppStoreVersionPhasedReleaseResponse defines model for AppStoreVersionPhasedReleaseResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseresponse
type AppStoreVersionPhasedReleaseResponse struct {
	Data  AppStoreVersionPhasedRelease `json:"data"`
	Links DocumentLinks                `json:"links"`
}

// AppStoreVersionPhasedReleaseUpdateRequest defines model for AppStoreVersionPhasedReleaseUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseupdaterequest
type AppStoreVersionPhasedReleaseUpdateRequest struct {
	Attributes *AppStoreVersionPhasedReleaseUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                               `json:"id"`
	Type       string                                               `json:"type"`
}

// AppStoreVersionPhasedReleaseUpdateRequestAttributes are attributes for AppStoreVersionPhasedReleaseUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionphasedreleaseupdaterequest/data/attributes
type AppStoreVersionPhasedReleaseUpdateRequestAttributes struct {
	PhasedReleaseState *PhasedReleaseState `json:"phasedReleaseState,omitempty"`
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery are query options for GetAppStoreVersionPhasedReleaseForAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_phased_release_information_of_an_app_store_version
type GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery struct {
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
}

// CreatePhasedRelease enables phased release for an App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version_phased_release
func (s *PublishingService) CreatePhasedRelease(ctx context.Context, phasedReleaseState *PhasedReleaseState, appStoreVersionID string) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	req := appStoreVersionPhasedReleaseCreateRequest{
		Relationships: appStoreVersionPhasedReleaseCreateRequestRelationships{
			AppStoreVersion: RelationshipDeclaration{
				Data: &RelationshipData{
					ID:   appStoreVersionID,
					Type: "appStoreVersions",
				},
			},
		},
		Type: "appStoreVersionPhasedReleases",
	}
	if phasedReleaseState != nil {
		req.Attributes = &appStoreVersionPhasedReleaseCreateRequestAttributes{
			PhasedReleaseState: phasedReleaseState,
		}
	}
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.post(ctx, "appStoreVersionPhasedReleases", req, res)
	return res, resp, err
}

// UpdatePhasedRelease pauses or resumes a phased release, or immediately release an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_version_phased_release
func (s *PublishingService) UpdatePhasedRelease(ctx context.Context, id string, body AppStoreVersionPhasedReleaseUpdateRequest) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.patch(ctx, url, body, res)
	return res, resp, err
}

// DeletePhasedRelease cancels a planned phased release that has not been started.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version_phased_release
func (s *PublishingService) DeletePhasedRelease(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersionPhasedReleases/%s", id)
	return s.client.delete(ctx, url, nil)
}

// GetAppStoreVersionPhasedReleaseForAppStoreVersion reads the phased release status and configuration for a version with phased release enabled.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_app_store_version_phased_release_information_of_an_app_store_version
func (s *PublishingService) GetAppStoreVersionPhasedReleaseForAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionPhasedReleaseForAppStoreVersionQuery) (*AppStoreVersionPhasedReleaseResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/appStoreVersionPhasedRelease", id)
	res := new(AppStoreVersionPhasedReleaseResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
