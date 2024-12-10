package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemRetrieveSnapshotsRequestBuilder provides operations to call the retrieveSnapshots method.
type ItemCloudPCsItemRetrieveSnapshotsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters list all cloudPcSnapshot resources for a Cloud PC.
type ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetQueryParameters
}
// NewItemCloudPCsItemRetrieveSnapshotsRequestBuilderInternal instantiates a new ItemCloudPCsItemRetrieveSnapshotsRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetrieveSnapshotsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    m := &ItemCloudPCsItemRetrieveSnapshotsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/retrieveSnapshots(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemRetrieveSnapshotsRequestBuilder instantiates a new ItemCloudPCsItemRetrieveSnapshotsRequestBuilder and sets the default values.
func NewItemCloudPCsItemRetrieveSnapshotsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemRetrieveSnapshotsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all cloudPcSnapshot resources for a Cloud PC.
// Deprecated: This method is obsolete. Use GetAsRetrieveSnapshotsGetResponse instead.
// returns a ItemCloudPCsItemRetrieveSnapshotsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievesnapshots?view=graph-rest-beta
func (m *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemRetrieveSnapshotsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemRetrieveSnapshotsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemRetrieveSnapshotsResponseable), nil
}
// GetAsRetrieveSnapshotsGetResponse list all cloudPcSnapshot resources for a Cloud PC.
// returns a ItemCloudPCsItemRetrieveSnapshotsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-retrievesnapshots?view=graph-rest-beta
func (m *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) GetAsRetrieveSnapshotsGetResponse(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(ItemCloudPCsItemRetrieveSnapshotsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCloudPCsItemRetrieveSnapshotsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCloudPCsItemRetrieveSnapshotsGetResponseable), nil
}
// ToGetRequestInformation list all cloudPcSnapshot resources for a Cloud PC.
// returns a *RequestInformation when successful
func (m *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemRetrieveSnapshotsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder when successful
func (m *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    return NewItemCloudPCsItemRetrieveSnapshotsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
