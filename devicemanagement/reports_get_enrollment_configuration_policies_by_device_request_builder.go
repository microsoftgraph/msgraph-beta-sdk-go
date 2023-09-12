package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
type ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal instantiates a new GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder and sets the default values.
func NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    m := &ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getEnrollmentConfigurationPoliciesByDevice", pathParameters),
    }
    return m
}
// NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder instantiates a new GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder and sets the default values.
func NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) Post(ctx context.Context, body ReportsGetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable, requestConfiguration *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable, requestConfiguration *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) WithUrl(rawUrl string)(*ReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    return NewReportsGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
