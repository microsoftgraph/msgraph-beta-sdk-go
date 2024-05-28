package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder provides operations to call the getStorageAccounts method.
type VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters list all storage accounts cloudPcForensicStorageAccount that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
type VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters struct {
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
// VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters
}
// NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderInternal instantiates a new VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, subscriptionId *string)(*VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    m := &VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/getStorageAccounts(subscriptionId='{subscriptionId}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if subscriptionId != nil {
        m.BaseRequestBuilder.PathParameters["subscriptionId"] = *subscriptionId
    }
    return m
}
// NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder instantiates a new VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get list all storage accounts cloudPcForensicStorageAccount that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// Deprecated: This method is obsolete. Use GetAsGetStorageAccountsWithSubscriptionIdGetResponse instead.
// returns a VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcsnapshot-getstorageaccounts?view=graph-rest-beta
func (m *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdResponseable), nil
}
// GetAsGetStorageAccountsWithSubscriptionIdGetResponse list all storage accounts cloudPcForensicStorageAccount that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// returns a VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcsnapshot-getstorageaccounts?view=graph-rest-beta
func (m *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) GetAsGetStorageAccountsWithSubscriptionIdGetResponse(ctx context.Context, requestConfiguration *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdGetResponseable), nil
}
// ToGetRequestInformation list all storage accounts cloudPcForensicStorageAccount that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// returns a *RequestInformation when successful
func (m *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder when successful
func (m *VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    return NewVirtualendpointSnapshotsGetstorageaccountswithsubscriptionidGetStorageAccountsWithSubscriptionIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
