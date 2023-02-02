package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder provides operations to call the associateWithHubSites method.
type ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderInternal instantiates a new AssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder) {
    m := &ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes/{contentType%2Did}/microsoft.graph.associateWithHubSites";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder instantiates a new AssociateWithHubSitesRequestBuilder and sets the default values.
func NewItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post associate a [content type][contentType] with a list of hub sites.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/contenttype-associatewithhubsites?view=graph-rest-1.0
func (m *ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder) Post(ctx context.Context, body ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation associate a [content type][contentType] with a list of hub sites.
func (m *ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesPostRequestBodyable, requestConfiguration *ItemListContentTypesItemMicrosoftGraphAssociateWithHubSitesAssociateWithHubSitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
