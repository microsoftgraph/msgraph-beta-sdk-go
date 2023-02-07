package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderInternal instantiates a new MicrosoftGraphRestoreVersionRequestBuilder and sets the default values.
func NewItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder) {
    m := &ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/versions/{listItemVersion%2Did}/microsoft.graph.restoreVersion";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder instantiates a new MicrosoftGraphRestoreVersionRequestBuilder and sets the default values.
func NewItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action restoreVersion
func (m *ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action restoreVersion
func (m *ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemVersionsItemMicrosoftGraphRestoreVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
