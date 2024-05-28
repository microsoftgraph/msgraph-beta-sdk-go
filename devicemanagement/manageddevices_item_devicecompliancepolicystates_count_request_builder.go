package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder provides operations to count the resources in the collection.
type ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetQueryParameters get the number of the resource
type ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetQueryParameters
}
// NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderInternal instantiates a new ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) {
    m := &ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceCompliancePolicyStates/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder instantiates a new ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder and sets the default values.
func NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder when successful
func (m *ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder) {
    return NewManageddevicesItemDevicecompliancepolicystatesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
