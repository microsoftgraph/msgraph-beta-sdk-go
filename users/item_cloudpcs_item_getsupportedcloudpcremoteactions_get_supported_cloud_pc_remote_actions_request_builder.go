package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder provides operations to call the getSupportedCloudPcRemoteActions method.
type ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
type ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters struct {
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
// ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters
}
// NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal instantiates a new ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    m := &ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getSupportedCloudPcRemoteActions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder instantiates a new ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// Deprecated: This method is obsolete. Use GetAsGetSupportedCloudPcRemoteActionsGetResponse instead.
// returns a ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getsupportedcloudpcremoteactions?view=graph-rest-beta
func (m *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsResponseable), nil
}
// GetAsGetSupportedCloudPcRemoteActionsGetResponse get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// returns a ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getsupportedcloudpcremoteactions?view=graph-rest-beta
func (m *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) GetAsGetSupportedCloudPcRemoteActionsGetResponse(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsGetResponseable), nil
}
// ToGetRequestInformation get a list of supported Cloud PC remote actions for a specific Cloud PC device, including the action names and capabilities.
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
