package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageDeletedcontainersItemDriveListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type FilestorageDeletedcontainersItemDriveListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageDeletedcontainersItemDriveListRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveListRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageDeletedcontainersItemDriveListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type FilestorageDeletedcontainersItemDriveListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageDeletedcontainersItemDriveListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageDeletedcontainersItemDriveListRequestBuilderGetQueryParameters
}
// FilestorageDeletedcontainersItemDriveListRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageDeletedcontainersItemDriveListRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListActivitiesRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Activities()(*FilestorageDeletedcontainersItemDriveListActivitiesRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListColumnsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Columns()(*FilestorageDeletedcontainersItemDriveListColumnsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageDeletedcontainersItemDriveListRequestBuilderInternal instantiates a new FilestorageDeletedcontainersItemDriveListRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListRequestBuilder) {
    m := &FilestorageDeletedcontainersItemDriveListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageDeletedcontainersItemDriveListRequestBuilder instantiates a new FilestorageDeletedcontainersItemDriveListRequestBuilder and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageDeletedcontainersItemDriveListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageDeletedcontainersItemDriveListRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListContenttypesContentTypesRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) ContentTypes()(*FilestorageDeletedcontainersItemDriveListContenttypesContentTypesRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageDeletedcontainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) CreatedByUser()(*FilestorageDeletedcontainersItemDriveListCreatedbyuserCreatedByUserRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property list for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderDeleteRequestConfiguration)(error) {
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
// Drive provides operations to manage the drive property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListDriveRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Drive()(*FilestorageDeletedcontainersItemDriveListDriveRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
// returns a Listable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListItemsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Items()(*FilestorageDeletedcontainersItemDriveListItemsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FilestorageDeletedcontainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) LastModifiedByUser()(*FilestorageDeletedcontainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListOperationsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Operations()(*FilestorageDeletedcontainersItemDriveListOperationsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property list in storage
// returns a Listable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListPermissionsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Permissions()(*FilestorageDeletedcontainersItemDriveListPermissionsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.list entity.
// returns a *FilestorageDeletedcontainersItemDriveListSubscriptionsRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) Subscriptions()(*FilestorageDeletedcontainersItemDriveListSubscriptionsRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property list for storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property list in storage
// returns a *RequestInformation when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *FilestorageDeletedcontainersItemDriveListRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageDeletedcontainersItemDriveListRequestBuilder when successful
func (m *FilestorageDeletedcontainersItemDriveListRequestBuilder) WithUrl(rawUrl string)(*FilestorageDeletedcontainersItemDriveListRequestBuilder) {
    return NewFilestorageDeletedcontainersItemDriveListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
