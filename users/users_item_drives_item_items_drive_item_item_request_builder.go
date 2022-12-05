package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDrivesItemItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type UsersItemDrivesItemItemsDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetQueryParameters
}
// UsersItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Activities()(*UsersItemDrivesItemItemsItemActivitiesRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ActivitiesById(id string)(*UsersItemDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Analytics()(*UsersItemDrivesItemItemsItemAnalyticsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*UsersItemDrivesItemItemsItemAssignSensitivityLabelRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Checkin()(*UsersItemDrivesItemItemsItemCheckinRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Checkout()(*UsersItemDrivesItemItemsItemCheckoutRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Children()(*UsersItemDrivesItemItemsItemChildrenRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ChildrenById(id string)(*UsersItemDrivesItemItemsItemChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return NewUsersItemDrivesItemItemsItemChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemDrivesItemItemsDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemItemsDriveItemItemRequestBuilder) {
    m := &UsersItemDrivesItemItemsDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDrivesItemItemsDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDrivesItemItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Content()(*UsersItemDrivesItemItemsItemContentRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Copy()(*UsersItemDrivesItemItemsItemCopyRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property items for users
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) CreateLink()(*UsersItemDrivesItemItemsItemCreateLinkRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in users
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) CreateUploadSession()(*UsersItemDrivesItemItemsItemCreateUploadSessionRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property items for users
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Delta()(*UsersItemDrivesItemItemsItemDeltaRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*UsersItemDrivesItemItemsItemDeltaWithTokenRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*UsersItemDrivesItemItemsItemExtractSensitivityLabelsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Follow()(*UsersItemDrivesItemItemsItemFollowRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*UsersItemDrivesItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Invite()(*UsersItemDrivesItemItemsItemInviteRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ListItem()(*UsersItemDrivesItemItemsItemListItemRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property items in users
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *UsersItemDrivesItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Permissions()(*UsersItemDrivesItemItemsItemPermissionsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) PermissionsById(id string)(*UsersItemDrivesItemItemsItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Preview()(*UsersItemDrivesItemItemsItemPreviewRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Restore()(*UsersItemDrivesItemItemsItemRestoreRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*UsersItemDrivesItemItemsItemSearchWithQRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Subscriptions()(*UsersItemDrivesItemItemsItemSubscriptionsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) SubscriptionsById(id string)(*UsersItemDrivesItemItemsItemSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsItemSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Thumbnails()(*UsersItemDrivesItemItemsItemThumbnailsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ThumbnailsById(id string)(*UsersItemDrivesItemItemsItemThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsItemThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unfollow provides operations to call the unfollow method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Unfollow()(*UsersItemDrivesItemItemsItemUnfollowRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) ValidatePermission()(*UsersItemDrivesItemItemsItemValidatePermissionRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) Versions()(*UsersItemDrivesItemItemsItemVersionsRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *UsersItemDrivesItemItemsDriveItemItemRequestBuilder) VersionsById(id string)(*UsersItemDrivesItemItemsItemVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsItemVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
