package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeDrivesItemRootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type MeDrivesItemRootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeDrivesItemRootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeDrivesItemRootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeDrivesItemRootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.
type MeDrivesItemRootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeDrivesItemRootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeDrivesItemRootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeDrivesItemRootRequestBuilderGetQueryParameters
}
// MeDrivesItemRootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeDrivesItemRootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Activities()(*MeDrivesItemRootActivitiesRequestBuilder) {
    return NewMeDrivesItemRootActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) ActivitiesById(id string)(*MeDrivesItemRootActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewMeDrivesItemRootActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Analytics()(*MeDrivesItemRootAnalyticsRequestBuilder) {
    return NewMeDrivesItemRootAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *MeDrivesItemRootRequestBuilder) AssignSensitivityLabel()(*MeDrivesItemRootAssignSensitivityLabelRequestBuilder) {
    return NewMeDrivesItemRootAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *MeDrivesItemRootRequestBuilder) Checkin()(*MeDrivesItemRootCheckinRequestBuilder) {
    return NewMeDrivesItemRootCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *MeDrivesItemRootRequestBuilder) Checkout()(*MeDrivesItemRootCheckoutRequestBuilder) {
    return NewMeDrivesItemRootCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Children()(*MeDrivesItemRootChildrenRequestBuilder) {
    return NewMeDrivesItemRootChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) ChildrenById(id string)(*MeDrivesItemRootChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewMeDrivesItemRootChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeDrivesItemRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewMeDrivesItemRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeDrivesItemRootRequestBuilder) {
    m := &MeDrivesItemRootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeDrivesItemRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewMeDrivesItemRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeDrivesItemRootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeDrivesItemRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *MeDrivesItemRootRequestBuilder) Content()(*MeDrivesItemRootContentRequestBuilder) {
    return NewMeDrivesItemRootContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *MeDrivesItemRootRequestBuilder) Copy()(*MeDrivesItemRootCopyRequestBuilder) {
    return NewMeDrivesItemRootCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
func (m *MeDrivesItemRootRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeDrivesItemRootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *MeDrivesItemRootRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeDrivesItemRootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateLink provides operations to call the createLink method.
func (m *MeDrivesItemRootRequestBuilder) CreateLink()(*MeDrivesItemRootCreateLinkRequestBuilder) {
    return NewMeDrivesItemRootCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property root in me
func (m *MeDrivesItemRootRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *MeDrivesItemRootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *MeDrivesItemRootRequestBuilder) CreateUploadSession()(*MeDrivesItemRootCreateUploadSessionRequestBuilder) {
    return NewMeDrivesItemRootCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for me
func (m *MeDrivesItemRootRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeDrivesItemRootRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
func (m *MeDrivesItemRootRequestBuilder) Delta()(*MeDrivesItemRootDeltaRequestBuilder) {
    return NewMeDrivesItemRootDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *MeDrivesItemRootRequestBuilder) DeltaWithToken(token *string)(*MeDrivesItemRootDeltaWithTokenRequestBuilder) {
    return NewMeDrivesItemRootDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *MeDrivesItemRootRequestBuilder) ExtractSensitivityLabels()(*MeDrivesItemRootExtractSensitivityLabelsRequestBuilder) {
    return NewMeDrivesItemRootExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *MeDrivesItemRootRequestBuilder) Follow()(*MeDrivesItemRootFollowRequestBuilder) {
    return NewMeDrivesItemRootFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *MeDrivesItemRootRequestBuilder) Get(ctx context.Context, requestConfiguration *MeDrivesItemRootRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *MeDrivesItemRootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*MeDrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewMeDrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *MeDrivesItemRootRequestBuilder) Invite()(*MeDrivesItemRootInviteRequestBuilder) {
    return NewMeDrivesItemRootInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) ListItem()(*MeDrivesItemRootListItemRequestBuilder) {
    return NewMeDrivesItemRootListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
func (m *MeDrivesItemRootRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *MeDrivesItemRootRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Permissions()(*MeDrivesItemRootPermissionsRequestBuilder) {
    return NewMeDrivesItemRootPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) PermissionsById(id string)(*MeDrivesItemRootPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewMeDrivesItemRootPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *MeDrivesItemRootRequestBuilder) Preview()(*MeDrivesItemRootPreviewRequestBuilder) {
    return NewMeDrivesItemRootPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *MeDrivesItemRootRequestBuilder) Restore()(*MeDrivesItemRootRestoreRequestBuilder) {
    return NewMeDrivesItemRootRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *MeDrivesItemRootRequestBuilder) SearchWithQ(q *string)(*MeDrivesItemRootSearchWithQRequestBuilder) {
    return NewMeDrivesItemRootSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Subscriptions()(*MeDrivesItemRootSubscriptionsRequestBuilder) {
    return NewMeDrivesItemRootSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) SubscriptionsById(id string)(*MeDrivesItemRootSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewMeDrivesItemRootSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Thumbnails()(*MeDrivesItemRootThumbnailsRequestBuilder) {
    return NewMeDrivesItemRootThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) ThumbnailsById(id string)(*MeDrivesItemRootThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewMeDrivesItemRootThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *MeDrivesItemRootRequestBuilder) Unfollow()(*MeDrivesItemRootUnfollowRequestBuilder) {
    return NewMeDrivesItemRootUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *MeDrivesItemRootRequestBuilder) ValidatePermission()(*MeDrivesItemRootValidatePermissionRequestBuilder) {
    return NewMeDrivesItemRootValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) Versions()(*MeDrivesItemRootVersionsRequestBuilder) {
    return NewMeDrivesItemRootVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *MeDrivesItemRootRequestBuilder) VersionsById(id string)(*MeDrivesItemRootVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewMeDrivesItemRootVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
