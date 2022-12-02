package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters get contentTypes from groups
type GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetQueryParameters
}
// GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*GroupsItemDrivesItemListContentTypesItemAssociateWithHubSitesRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Base()(*GroupsItemDrivesItemListContentTypesItemBaseRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) BaseTypes()(*GroupsItemDrivesItemListContentTypesItemBaseTypesRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) BaseTypesById(id string)(*GroupsItemDrivesItemListContentTypesItemBaseTypesContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return NewGroupsItemDrivesItemListContentTypesItemBaseTypesContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnLinks()(*GroupsItemDrivesItemListContentTypesItemColumnLinksRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnLinksById(id string)(*GroupsItemDrivesItemListContentTypesItemColumnLinksColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return NewGroupsItemDrivesItemListContentTypesItemColumnLinksColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnPositions()(*GroupsItemDrivesItemListContentTypesItemColumnPositionsRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*GroupsItemDrivesItemListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewGroupsItemDrivesItemListContentTypesItemColumnPositionsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Columns()(*GroupsItemDrivesItemListContentTypesItemColumnsRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById provides operations to manage the columns property of the microsoft.graph.contentType entity.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) ColumnsById(id string)(*GroupsItemDrivesItemListContentTypesItemColumnsColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return NewGroupsItemDrivesItemListContentTypesItemColumnsColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewGroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) {
    m := &GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewGroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*GroupsItemDrivesItemListContentTypesItemCopyToDefaultContentLocationRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get contentTypes from groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get contentTypes from groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
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
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) IsPublished()(*GroupsItemDrivesItemListContentTypesItemIsPublishedRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
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
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Publish()(*GroupsItemDrivesItemListContentTypesItemPublishRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish provides operations to call the unpublish method.
func (m *GroupsItemDrivesItemListContentTypesContentTypeItemRequestBuilder) Unpublish()(*GroupsItemDrivesItemListContentTypesItemUnpublishRequestBuilder) {
    return NewGroupsItemDrivesItemListContentTypesItemUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
