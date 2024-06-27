package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetQueryParameters
}
// FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemAssociateWithHubSitesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*FileStorageDeletedContainersItemDriveListContentTypesItemAssociateWithHubSitesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemAssociateWithHubSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemBaseRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Base()(*FileStorageDeletedContainersItemDriveListContentTypesItemBaseRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemBaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemBaseTypesRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) BaseTypes()(*FileStorageDeletedContainersItemDriveListContentTypesItemBaseTypesRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemBaseTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemColumnLinksRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) ColumnLinks()(*FileStorageDeletedContainersItemDriveListContentTypesItemColumnLinksRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemColumnLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) ColumnPositions()(*FileStorageDeletedContainersItemDriveListContentTypesItemColumnPositionsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemColumnPositionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemColumnsRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Columns()(*FileStorageDeletedContainersItemDriveListContentTypesItemColumnsRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderInternal instantiates a new FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) {
    m := &FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/deletedContainers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder instantiates a new FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder and sets the default values.
func NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemCopyToDefaultContentLocationRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*FileStorageDeletedContainersItemDriveListContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentTypes for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of content types present in this list.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) IsPublished()(*FileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemIsPublishedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property contentTypes in storage
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// Publish provides operations to call the publish method.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemPublishRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Publish()(*FileStorageDeletedContainersItemDriveListContentTypesItemPublishRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property contentTypes for storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of content types present in this list.
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in storage
// returns a *RequestInformation when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unpublish provides operations to call the unpublish method.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesItemUnpublishRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) Unpublish()(*FileStorageDeletedContainersItemDriveListContentTypesItemUnpublishRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesItemUnpublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder when successful
func (m *FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder) {
    return NewFileStorageDeletedContainersItemDriveListContentTypesContentTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
