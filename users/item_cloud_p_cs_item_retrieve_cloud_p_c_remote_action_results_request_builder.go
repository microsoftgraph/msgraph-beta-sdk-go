package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder provides operations to call the retrieveCloudPCRemoteActionResults method.
type ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
type ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal instantiates a new ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    m := &ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/retrieveCloudPCRemoteActionResults(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder instantiates a new ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// Deprecated: This method is obsolete. Use GetAsRetrieveCloudPCRemoteActionResultsGetResponse instead.
// returns a ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpcremoteactionresults?view=graph-rest-beta
func (m *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsResponseable), nil
}
// GetAsRetrieveCloudPCRemoteActionResultsGetResponse retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// returns a ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpcremoteactionresults?view=graph-rest-beta
func (m *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) GetAsRetrieveCloudPCRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation retrieve remote action results and check the status of a specific remote action performed on the associated Cloud PC device.
// returns a *RequestInformation when successful
func (m *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder when successful
func (m *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    return NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
