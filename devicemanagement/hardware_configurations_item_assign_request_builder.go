package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareConfigurationsItemAssignRequestBuilder provides operations to call the assign method.
type HardwareConfigurationsItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareConfigurationsItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHardwareConfigurationsItemAssignRequestBuilderInternal instantiates a new HardwareConfigurationsItemAssignRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemAssignRequestBuilder) {
    m := &HardwareConfigurationsItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}/assign", pathParameters),
    }
    return m
}
// NewHardwareConfigurationsItemAssignRequestBuilder instantiates a new HardwareConfigurationsItemAssignRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareConfigurationsItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
// returns a HardwareConfigurationsItemAssignResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemAssignRequestBuilder) Post(ctx context.Context, body HardwareConfigurationsItemAssignPostRequestBodyable, requestConfiguration *HardwareConfigurationsItemAssignRequestBuilderPostRequestConfiguration)(HardwareConfigurationsItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHardwareConfigurationsItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HardwareConfigurationsItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
// returns a HardwareConfigurationsItemAssignPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body HardwareConfigurationsItemAssignPostRequestBodyable, requestConfiguration *HardwareConfigurationsItemAssignRequestBuilderPostRequestConfiguration)(HardwareConfigurationsItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHardwareConfigurationsItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(HardwareConfigurationsItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body HardwareConfigurationsItemAssignPostRequestBodyable, requestConfiguration *HardwareConfigurationsItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwareConfigurationsItemAssignRequestBuilder when successful
func (m *HardwareConfigurationsItemAssignRequestBuilder) WithUrl(rawUrl string)(*HardwareConfigurationsItemAssignRequestBuilder) {
    return NewHardwareConfigurationsItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
