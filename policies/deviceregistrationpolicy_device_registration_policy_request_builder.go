package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder provides operations to manage the deviceRegistrationPolicy property of the microsoft.graph.policyRoot entity.
type DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetQueryParameters read the properties and relationships of a deviceRegistrationPolicy object. Represents deviceRegistrationPolicy quota restrictions, additional authentication, and authorization policies to register device identities to your organization.
type DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetQueryParameters
}
// NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderInternal instantiates a new DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder and sets the default values.
func NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) {
    m := &DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/deviceRegistrationPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder instantiates a new DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder and sets the default values.
func NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a deviceRegistrationPolicy object. Represents deviceRegistrationPolicy quota restrictions, additional authentication, and authorization policies to register device identities to your organization.
// returns a DeviceRegistrationPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/deviceregistrationpolicy-get?view=graph-rest-beta
func (m *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceRegistrationPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceRegistrationPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceRegistrationPolicyable), nil
}
// ToGetRequestInformation read the properties and relationships of a deviceRegistrationPolicy object. Represents deviceRegistrationPolicy quota restrictions, additional authentication, and authorization policies to register device identities to your organization.
// returns a *RequestInformation when successful
func (m *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder when successful
func (m *DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) WithUrl(rawUrl string)(*DeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder) {
    return NewDeviceregistrationpolicyDeviceRegistrationPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
