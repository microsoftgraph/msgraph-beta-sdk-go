package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type ItemItemsDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type ItemItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsDriveItemItemRequestBuilderGetQueryParameters
}
// ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Activities()(*ItemItemsItemActivitiesRequestBuilder) {
    return NewItemItemsItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ActivitiesById(id string)(*ItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Analytics()(*ItemItemsItemAnalyticsRequestBuilder) {
    return NewItemItemsItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *ItemItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*ItemItemsItemAssignSensitivityLabelRequestBuilder) {
    return NewItemItemsItemAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Checkin provides operations to call the checkin method.
func (m *ItemItemsDriveItemItemRequestBuilder) Checkin()(*ItemItemsItemCheckinRequestBuilder) {
    return NewItemItemsItemCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Checkout provides operations to call the checkout method.
func (m *ItemItemsDriveItemItemRequestBuilder) Checkout()(*ItemItemsItemCheckoutRequestBuilder) {
    return NewItemItemsItemCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Children()(*ItemItemsItemChildrenRequestBuilder) {
    return NewItemItemsItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ChildrenById(id string)(*ItemItemsItemChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did1"] = id
    }
    return NewItemItemsItemChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewItemItemsDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    m := &ItemItemsDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemItemsDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewItemItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Content()(*ItemItemsItemContentRequestBuilder) {
    return NewItemItemsItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Copy provides operations to call the copy method.
func (m *ItemItemsDriveItemItemRequestBuilder) Copy()(*ItemItemsItemCopyRequestBuilder) {
    return NewItemItemsItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreateLink provides operations to call the createLink method.
func (m *ItemItemsDriveItemItemRequestBuilder) CreateLink()(*ItemItemsItemCreateLinkRequestBuilder) {
    return NewItemItemsItemCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *ItemItemsDriveItemItemRequestBuilder) CreateUploadSession()(*ItemItemsItemCreateUploadSessionRequestBuilder) {
    return NewItemItemsItemCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property items for drives
func (m *ItemItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
func (m *ItemItemsDriveItemItemRequestBuilder) Delta()(*ItemItemsItemDeltaRequestBuilder) {
    return NewItemItemsItemDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
func (m *ItemItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*ItemItemsItemDeltaWithTokenRequestBuilder) {
    return NewItemItemsItemDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *ItemItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*ItemItemsItemExtractSensitivityLabelsRequestBuilder) {
    return NewItemItemsItemExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Follow provides operations to call the follow method.
func (m *ItemItemsDriveItemItemRequestBuilder) Follow()(*ItemItemsItemFollowRequestBuilder) {
    return NewItemItemsItemFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
func (m *ItemItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
func (m *ItemItemsDriveItemItemRequestBuilder) Invite()(*ItemItemsItemInviteRequestBuilder) {
    return NewItemItemsItemInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ListItem()(*ItemItemsItemListItemRequestBuilder) {
    return NewItemItemsItemListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property items in drives
func (m *ItemItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Permissions()(*ItemItemsItemPermissionsRequestBuilder) {
    return NewItemItemsItemPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) PermissionsById(id string)(*ItemItemsItemPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewItemItemsItemPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Preview provides operations to call the preview method.
func (m *ItemItemsDriveItemItemRequestBuilder) Preview()(*ItemItemsItemPreviewRequestBuilder) {
    return NewItemItemsItemPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Restore provides operations to call the restore method.
func (m *ItemItemsDriveItemItemRequestBuilder) Restore()(*ItemItemsItemRestoreRequestBuilder) {
    return NewItemItemsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SearchWithQ provides operations to call the search method.
func (m *ItemItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*ItemItemsItemSearchWithQRequestBuilder) {
    return NewItemItemsItemSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Subscriptions()(*ItemItemsItemSubscriptionsRequestBuilder) {
    return NewItemItemsItemSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) SubscriptionsById(id string)(*ItemItemsItemSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewItemItemsItemSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Thumbnails()(*ItemItemsItemThumbnailsRequestBuilder) {
    return NewItemItemsItemThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) ThumbnailsById(id string)(*ItemItemsItemThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewItemItemsItemThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for drives
func (m *ItemItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
func (m *ItemItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property items in drives
func (m *ItemItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *ItemItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Unfollow provides operations to call the unfollow method.
func (m *ItemItemsDriveItemItemRequestBuilder) Unfollow()(*ItemItemsItemUnfollowRequestBuilder) {
    return NewItemItemsItemUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *ItemItemsDriveItemItemRequestBuilder) ValidatePermission()(*ItemItemsItemValidatePermissionRequestBuilder) {
    return NewItemItemsItemValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Versions()(*ItemItemsItemVersionsRequestBuilder) {
    return NewItemItemsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) VersionsById(id string)(*ItemItemsItemVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewItemItemsItemVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
func (m *ItemItemsDriveItemItemRequestBuilder) Workbook()(*ItemItemsItemWorkbookRequestBuilder) {
    return NewItemItemsItemWorkbookRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
