// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder provides operations to call the retrieveCloudPcCountByStatus method.
type ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters retrieve the Cloud PC count grouped by status.
type ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters struct {
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
// ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetQueryParameters
}
// NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal instantiates a new ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder and sets the default values.
func NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    m := &ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/retrieveCloudPcCountByStatus(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder instantiates a new ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder and sets the default values.
func NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the Cloud PC count grouped by status.
// Deprecated: This method is obsolete. Use GetAsRetrieveCloudPcCountByStatusGetResponse instead.
// returns a ItemCloudPCsRetrieveCloudPcCountByStatusResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpccountbystatus?view=graph-rest-beta
func (m *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(ItemCloudPCsRetrieveCloudPcCountByStatusResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsRetrieveCloudPcCountByStatusResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsRetrieveCloudPcCountByStatusResponseable), nil
}
// GetAsRetrieveCloudPcCountByStatusGetResponse retrieve the Cloud PC count grouped by status.
// returns a ItemCloudPCsRetrieveCloudPcCountByStatusGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievecloudpccountbystatus?view=graph-rest-beta
func (m *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) GetAsRetrieveCloudPcCountByStatusGetResponse(ctx context.Context, requestConfiguration *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(ItemCloudPCsRetrieveCloudPcCountByStatusGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsRetrieveCloudPcCountByStatusGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsRetrieveCloudPcCountByStatusGetResponseable), nil
}
// ToGetRequestInformation retrieve the Cloud PC count grouped by status.
// returns a *RequestInformation when successful
func (m *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder when successful
func (m *ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder) {
    return NewItemCloudPCsRetrieveCloudPcCountByStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
