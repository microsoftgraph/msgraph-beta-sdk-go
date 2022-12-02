package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters get contentTypes from users
type UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters
}
// UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*UsersItemDrivesItemListContentTypesItemAssociateWithHubSitesRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Base()(*UsersItemDrivesItemListContentTypesItemBaseRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) BaseTypes()(*UsersItemDrivesItemListContentTypesItemBaseTypesRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) BaseTypesById(id string)(*UsersItemDrivesItemListContentTypesItemBaseTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return NewUsersItemDrivesItemListContentTypesItemBaseTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnLinks()(*UsersItemDrivesItemListContentTypesItemColumnLinksRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnLinksById(id string)(*UsersItemDrivesItemListContentTypesItemColumnLinksColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return NewUsersItemDrivesItemListContentTypesItemColumnLinksColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnPositions()(*UsersItemDrivesItemListContentTypesItemColumnPositionsRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*UsersItemDrivesItemListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewUsersItemDrivesItemListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Columns()(*UsersItemDrivesItemListContentTypesItemColumnsRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnsById(id string)(*UsersItemDrivesItemListContentTypesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewUsersItemDrivesItemListContentTypesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) {
    m := &UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*UsersItemDrivesItemListContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get contentTypes from users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get contentTypes from users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) IsPublished()(*UsersItemDrivesItemListContentTypesItemIsPublishedRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in users
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// Publish provides operations to call the publish method.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Publish()(*UsersItemDrivesItemListContentTypesItemPublishRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *UsersItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Unpublish()(*UsersItemDrivesItemListContentTypesItemUnpublishRequestBuilder) {
    return NewUsersItemDrivesItemListContentTypesItemUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
