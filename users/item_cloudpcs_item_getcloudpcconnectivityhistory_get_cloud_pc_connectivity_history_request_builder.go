package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder provides operations to call the getCloudPcConnectivityHistory method.
type ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters get the connectivity history of a specific Cloud PC.
type ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters struct {
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
// ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters
}
// NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderInternal instantiates a new ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    m := &ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getCloudPcConnectivityHistory(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder instantiates a new ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the connectivity history of a specific Cloud PC.
// Deprecated: This method is obsolete. Use GetAsGetCloudPcConnectivityHistoryGetResponse instead.
// returns a ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getcloudpcconnectivityhistory?view=graph-rest-beta
func (m *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryResponseable), nil
}
// GetAsGetCloudPcConnectivityHistoryGetResponse get the connectivity history of a specific Cloud PC.
// returns a ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getcloudpcconnectivityhistory?view=graph-rest-beta
func (m *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) GetAsGetCloudPcConnectivityHistoryGetResponse(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryGetResponseable), nil
}
// ToGetRequestInformation get the connectivity history of a specific Cloud PC.
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
