package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
type ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters struct {
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
// ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    m := &ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/getCloudPcRemoteActionResults(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder instantiates a new ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// Deprecated: This method is obsolete. Use GetAsGetCloudPcRemoteActionResultsGetResponse instead.
// returns a ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-getcloudpcremoteactionresults?view=graph-rest-beta
func (m *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsResponseable), nil
}
// GetAsGetCloudPcRemoteActionResultsGetResponse check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// returns a ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/manageddevice-getcloudpcremoteactionresults?view=graph-rest-beta
func (m *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) GetAsGetCloudPcRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation check the Cloud PC-specified remote action results for a Cloud PC device. Cloud PC supports reprovision and resize remote actions.
// returns a *RequestInformation when successful
func (m *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewManageddevicesItemGetcloudpcremoteactionresultsGetCloudPcRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
