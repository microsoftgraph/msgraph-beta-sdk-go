package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveRequestBuilder provides operations to manage the drive property of the microsoft.graph.fileStorageContainer entity.
type FileStorageDeletedContainersItemDriveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageDeletedContainersItemDriveRequestBuilderGetQueryParameters the drive of the resource fileStorageContainer. Read-only.
type FileStorageDeletedContainersItemDriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersItemDriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveRequestBuilderGetQueryParameters
}
// FileStorageDeletedContainersItemDriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveActivitiesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Activities()(*FileStorageDeletedContainersItemDriveActivitiesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveBundlesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Bundles()(*FileStorageDeletedContainersItemDriveBundlesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveBundlesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersItemDriveRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageDeletedContainersItemDriveCreatedByUserRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) CreatedByUser()(*FileStorageDeletedContainersItemDriveCreatedByUserRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property drive for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderDeleteRequestConfiguration)(error) {
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
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveFollowingRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Following()(*FileStorageDeletedContainersItemDriveFollowingRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveFollowingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the drive of the resource fileStorageContainer. Read-only.
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveItemsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Items()(*FileStorageDeletedContainersItemDriveItemsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *FileStorageDeletedContainersItemDriveLastModifiedByUserRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) LastModifiedByUser()(*FileStorageDeletedContainersItemDriveLastModifiedByUserRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveListRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) List()(*FileStorageDeletedContainersItemDriveListRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property drive in storage
// returns a Driveable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Recent provides operations to call the recent method.
// returns a *FileStorageDeletedContainersItemDriveRecentRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Recent()(*FileStorageDeletedContainersItemDriveRecentRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveRecentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveRootRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Root()(*FileStorageDeletedContainersItemDriveRootRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveRootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SearchWithQ provides operations to call the search method.
// returns a *FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) SearchWithQ(q *string)(*FileStorageDeletedContainersItemDriveSearchWithQRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveSearchWithQRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, q)
}
// SharedWithMe provides operations to call the sharedWithMe method.
// returns a *FileStorageDeletedContainersItemDriveSharedWithMeRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) SharedWithMe()(*FileStorageDeletedContainersItemDriveSharedWithMeRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveSharedWithMeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
// returns a *FileStorageDeletedContainersItemDriveSpecialRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) Special()(*FileStorageDeletedContainersItemDriveSpecialRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveSpecialRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property drive for storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the drive of the resource fileStorageContainer. Read-only.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property drive in storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *FileStorageDeletedContainersItemDriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageDeletedContainersItemDriveRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
