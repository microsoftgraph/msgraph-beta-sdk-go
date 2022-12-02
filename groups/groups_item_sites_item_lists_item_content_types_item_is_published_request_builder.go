package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder provides operations to call the isPublished method.
type GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderInternal instantiates a new IsPublishedRequestBuilder and sets the default values.
func NewGroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder) {
    m := &GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.isPublished()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder instantiates a new IsPublishedRequestBuilder and sets the default values.
func NewGroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function isPublished
func (m *GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function isPublished
func (m *GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemSitesItemListsItemContentTypesItemIsPublishedRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupsItemSitesItemListsItemContentTypesItemIsPublishedResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupsItemSitesItemListsItemContentTypesItemIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupsItemSitesItemListsItemContentTypesItemIsPublishedResponseable), nil
}
