package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContactFoldersMicrosoftGraphDeltaRequestBuilder provides operations to call the delta method.
type ItemContactFoldersMicrosoftGraphDeltaRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
type ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetQueryParameters
}
// NewItemContactFoldersMicrosoftGraphDeltaRequestBuilderInternal instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewItemContactFoldersMicrosoftGraphDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactFoldersMicrosoftGraphDeltaRequestBuilder) {
    m := &ItemContactFoldersMicrosoftGraphDeltaRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/contactFolders/microsoft.graph.delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemContactFoldersMicrosoftGraphDeltaRequestBuilder instantiates a new MicrosoftGraphDeltaRequestBuilder and sets the default values.
func NewItemContactFoldersMicrosoftGraphDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactFoldersMicrosoftGraphDeltaRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactFoldersMicrosoftGraphDeltaRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/contactfolder-delta?view=graph-rest-1.0
func (m *ItemContactFoldersMicrosoftGraphDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(ItemContactFoldersMicrosoftGraphDeltaDeltaResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemContactFoldersMicrosoftGraphDeltaDeltaResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContactFoldersMicrosoftGraphDeltaDeltaResponseable), nil
}
// ToGetRequestInformation get a set of contact folders that have been added, deleted, or removed from the user's mailbox. A **delta** function call for contact folders in a mailbox is similar to a GET request, except that by appropriately applying state tokens in one or more of these calls, you can query for incremental changes in the contact folders. This allows you to maintain and synchronize a local store of a user's contact folders without having to fetch all the contact folders of that mailbox from the server every time.
func (m *ItemContactFoldersMicrosoftGraphDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactFoldersMicrosoftGraphDeltaRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
