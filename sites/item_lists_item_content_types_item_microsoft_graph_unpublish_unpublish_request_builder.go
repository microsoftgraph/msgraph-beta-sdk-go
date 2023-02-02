package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder provides operations to call the unpublish method.
type ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderInternal instantiates a new UnpublishRequestBuilder and sets the default values.
func NewItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder) {
    m := &ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/microsoft.graph.unpublish";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder instantiates a new UnpublishRequestBuilder and sets the default values.
func NewItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unpublish a [contentType][] from a content type hub site.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-unpublish?view=graph-rest-1.0
func (m *ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation unpublish a [contentType][] from a content type hub site.
func (m *ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContentTypesItemMicrosoftGraphUnpublishUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
