package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesItemAssignRequestBuilder provides operations to call the assign method.
type DeviceCompliancePoliciesItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceCompliancePoliciesItemAssignRequestBuilderInternal instantiates a new DeviceCompliancePoliciesItemAssignRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemAssignRequestBuilder) {
    m := &DeviceCompliancePoliciesItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/assign", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesItemAssignRequestBuilder instantiates a new DeviceCompliancePoliciesItemAssignRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a DeviceCompliancePoliciesItemAssignResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesItemAssignRequestBuilder) Post(ctx context.Context, body DeviceCompliancePoliciesItemAssignPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesItemAssignRequestBuilderPostRequestConfiguration)(DeviceCompliancePoliciesItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
// returns a DeviceCompliancePoliciesItemAssignPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body DeviceCompliancePoliciesItemAssignPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesItemAssignRequestBuilderPostRequestConfiguration)(DeviceCompliancePoliciesItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
// returns a *RequestInformation when successful
func (m *DeviceCompliancePoliciesItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceCompliancePoliciesItemAssignPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceCompliancePoliciesItemAssignRequestBuilder when successful
func (m *DeviceCompliancePoliciesItemAssignRequestBuilder) WithUrl(rawUrl string)(*DeviceCompliancePoliciesItemAssignRequestBuilder) {
    return NewDeviceCompliancePoliciesItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
