package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters
}
// FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemActivitiesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Activities()(*FileStorageDeletedContainersItemDriveItemsItemActivitiesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemAnalyticsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Analytics()(*FileStorageDeletedContainersItemDriveItemsItemAnalyticsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*FileStorageDeletedContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkin provides operations to call the checkin method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCheckinRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Checkin()(*FileStorageDeletedContainersItemDriveItemsItemCheckinRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCheckinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkout provides operations to call the checkout method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCheckoutRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Checkout()(*FileStorageDeletedContainersItemDriveItemsItemCheckoutRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCheckoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemChildrenRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Children()(*FileStorageDeletedContainersItemDriveItemsItemChildrenRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemContentRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Content()(*FileStorageDeletedContainersItemDriveItemsItemContentRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ContentStream provides operations to manage the media for the storage entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemContentStreamRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ContentStream()(*FileStorageDeletedContainersItemDriveItemsItemContentStreamRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemContentStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCopyRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Copy()(*FileStorageDeletedContainersItemDriveItemsItemCopyRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCreatedByUserRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) CreatedByUser()(*FileStorageDeletedContainersItemDriveItemsItemCreatedByUserRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCreateLinkRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) CreateLink()(*FileStorageDeletedContainersItemDriveItemsItemCreateLinkRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemCreateUploadSessionRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) CreateUploadSession()(*FileStorageDeletedContainersItemDriveItemsItemCreateUploadSessionRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Delta provides operations to call the delta method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemDeltaRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Delta()(*FileStorageDeletedContainersItemDriveItemsItemDeltaRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemDeltaWithTokenRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*FileStorageDeletedContainersItemDriveItemsItemDeltaWithTokenRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemDeltaWithTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*FileStorageDeletedContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow provides operations to call the follow method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemFollowRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Follow()(*FileStorageDeletedContainersItemDriveItemsItemFollowRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*FileStorageDeletedContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Invite()(*FileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemLastModifiedByUserRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) LastModifiedByUser()(*FileStorageDeletedContainersItemDriveItemsItemLastModifiedByUserRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemListItemRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ListItem()(*FileStorageDeletedContainersItemDriveItemsItemListItemRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemListItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemPermanentDeleteRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) PermanentDelete()(*FileStorageDeletedContainersItemDriveItemsItemPermanentDeleteRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemPermissionsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Permissions()(*FileStorageDeletedContainersItemDriveItemsItemPermissionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview provides operations to call the preview method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemPreviewRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Preview()(*FileStorageDeletedContainersItemDriveItemsItemPreviewRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemRestoreRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Restore()(*FileStorageDeletedContainersItemDriveItemsItemRestoreRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetentionLabel provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemRetentionLabelRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) RetentionLabel()(*FileStorageDeletedContainersItemDriveItemsItemRetentionLabelRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemRetentionLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemSearchWithQRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*FileStorageDeletedContainersItemDriveItemsItemSearchWithQRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemSubscriptionsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Subscriptions()(*FileStorageDeletedContainersItemDriveItemsItemSubscriptionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemThumbnailsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Thumbnails()(*FileStorageDeletedContainersItemDriveItemsItemThumbnailsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemThumbnailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation all items contained in the drive. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property items in storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Unfollow provides operations to call the unfollow method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemUnfollowRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Unfollow()(*FileStorageDeletedContainersItemDriveItemsItemUnfollowRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
// returns a *FileStorageDeletedContainersItemDriveItemsItemValidatePermissionRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) ValidatePermission()(*FileStorageDeletedContainersItemDriveItemsItemValidatePermissionRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemValidatePermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemVersionsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Versions()(*FileStorageDeletedContainersItemDriveItemsItemVersionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
// returns a *FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveItemsDriveItemItemRequestBuilder) Workbook()(*FileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsItemWorkbookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
