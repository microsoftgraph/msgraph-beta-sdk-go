package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder provides operations to call the addCopyFromContentTypeHub method.
type ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) {
    m := &ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/addCopyFromContentTypeHub", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder instantiates a new ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-addcopyfromcontenttypehub?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) Post(ctx context.Context, body ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder when successful
func (m *ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder) {
    return NewItemSitesItemListsItemContentTypesAddCopyFromContentTypeHubRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
