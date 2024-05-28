package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.drive entity.
type FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters all items contained in the drive. Read-only. Nullable.
type FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetQueryParameters
}
// FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemActivitiesRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Activities()(*FilestorageDeletedcontainersItemDriveItemsItemActivitiesRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemAnalyticsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Analytics()(*FilestorageDeletedcontainersItemDriveItemsItemAnalyticsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) AssignSensitivityLabel()(*FilestorageDeletedcontainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemAssignsensitivitylabelAssignSensitivityLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkin provides operations to call the checkin method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCheckinRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Checkin()(*FilestorageDeletedcontainersItemDriveItemsItemCheckinRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCheckinRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Checkout provides operations to call the checkout method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCheckoutRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Checkout()(*FilestorageDeletedcontainersItemDriveItemsItemCheckoutRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCheckoutRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemChildrenRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Children()(*FilestorageDeletedcontainersItemDriveItemsItemChildrenRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the storage entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemContentRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Content()(*FilestorageDeletedcontainersItemDriveItemsItemContentRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ContentStream provides operations to manage the media for the storage entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemContentstreamContentStreamRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ContentStream()(*FilestorageDeletedcontainersItemDriveItemsItemContentstreamContentStreamRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemContentstreamContentStreamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Copy provides operations to call the copy method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCopyRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Copy()(*FilestorageDeletedcontainersItemDriveItemsItemCopyRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCopyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) CreatedByUser()(*FilestorageDeletedcontainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) CreateLink()(*FilestorageDeletedcontainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCreatelinkCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) CreateUploadSession()(*FilestorageDeletedcontainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property items for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Delta()(*FilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeltaWithToken provides operations to call the delta method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) DeltaWithToken(token *string)(*FilestorageDeletedcontainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemDeltawithtokenDeltaWithTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, token)
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ExtractSensitivityLabels()(*FilestorageDeletedcontainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Follow provides operations to call the follow method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemFollowRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Follow()(*FilestorageDeletedcontainersItemDriveItemsItemFollowRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemFollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get all items contained in the drive. Read-only. Nullable.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*FilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// Invite provides operations to call the invite method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemInviteRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Invite()(*FilestorageDeletedcontainersItemDriveItemsItemInviteRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemInviteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) LastModifiedByUser()(*FilestorageDeletedcontainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemListitemListItemRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ListItem()(*FilestorageDeletedcontainersItemDriveItemsItemListitemListItemRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemListitemListItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in storage
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) PermanentDelete()(*FilestorageDeletedcontainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemPermanentdeletePermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemPermissionsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Permissions()(*FilestorageDeletedcontainersItemDriveItemsItemPermissionsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Preview provides operations to call the preview method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemPreviewRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Preview()(*FilestorageDeletedcontainersItemDriveItemsItemPreviewRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemRestoreRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Restore()(*FilestorageDeletedcontainersItemDriveItemsItemRestoreRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetentionLabel provides operations to manage the retentionLabel property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) RetentionLabel()(*FilestorageDeletedcontainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemRetentionlabelRetentionLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) SearchWithQ(q *string)(*FilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemSearchwithqSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemSubscriptionsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Subscriptions()(*FilestorageDeletedcontainersItemDriveItemsItemSubscriptionsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemThumbnailsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Thumbnails()(*FilestorageDeletedcontainersItemDriveItemsItemThumbnailsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemThumbnailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property items for storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageDeletedcontainersItemDriveItemsItemUnfollowRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Unfollow()(*FilestorageDeletedcontainersItemDriveItemsItemUnfollowRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemUnfollowRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ValidatePermission provides operations to call the validatePermission method.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) ValidatePermission()(*FilestorageDeletedcontainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemVersionsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Versions()(*FilestorageDeletedcontainersItemDriveItemsItemVersionsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workbook provides operations to manage the workbook property of the microsoft.graph.driveItem entity.
// returns a *FilestorageDeletedcontainersItemDriveItemsItemWorkbookRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveItemsDriveItemItemRequestBuilder) Workbook()(*FilestorageDeletedcontainersItemDriveItemsItemWorkbookRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemWorkbookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
