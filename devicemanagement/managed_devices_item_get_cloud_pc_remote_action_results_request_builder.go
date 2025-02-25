package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
type ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    m := &ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/getCloudPcRemoteActionResults(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder instantiates a new ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// Deprecated: This method is obsolete. Use GetAsGetCloudPcRemoteActionResultsGetResponse instead.
// returns a ManagedDevicesItemGetCloudPcRemoteActionResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-getcloudpcremoteactionresults?view=graph-rest-beta
func (m *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ManagedDevicesItemGetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemGetCloudPcRemoteActionResultsResponseable), nil
}
// GetAsGetCloudPcRemoteActionResultsGetResponse check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// Deprecated: Starting September 30, 2024, getCloudPcRemoteActionResults API bound to managedDevice entity will be deprecated and no longer supported. Please use retrieveCloudPcRemoteActionResults API bound to cloudpc entity instead. as of 2024-05/getCloudPcRemoteActionResults on 2024-05-08 and will be removed 2024-09-30
// returns a ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-getcloudpcremoteactionresults?view=graph-rest-beta
func (m *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) GetAsGetCloudPcRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// Deprecated: Starting September 30, 2024, getCloudPcRemoteActionResults API bound to managedDevice entity will be deprecated and no longer supported. Please use retrieveCloudPcRemoteActionResults API bound to cloudpc entity instead. as of 2024-05/getCloudPcRemoteActionResults on 2024-05-08 and will be removed 2024-09-30
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Starting September 30, 2024, getCloudPcRemoteActionResults API bound to managedDevice entity will be deprecated and no longer supported. Please use retrieveCloudPcRemoteActionResults API bound to cloudpc entity instead. as of 2024-05/getCloudPcRemoteActionResults on 2024-05-08 and will be removed 2024-09-30
// returns a *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
