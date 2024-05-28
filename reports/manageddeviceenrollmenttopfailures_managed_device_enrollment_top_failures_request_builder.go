package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder provides operations to call the managedDeviceEnrollmentTopFailures method.
type ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderInternal instantiates a new ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder and sets the default values.
func NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    m := &ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/managedDeviceEnrollmentTopFailures()", pathParameters),
    }
    return m
}
// NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder instantiates a new ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder and sets the default values.
func NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function managedDeviceEnrollmentTopFailures
// returns a Reportable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Reportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Reportable), nil
}
// ToGetRequestInformation invoke function managedDeviceEnrollmentTopFailures
// returns a *RequestInformation when successful
func (m *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder when successful
func (m *ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) WithUrl(rawUrl string)(*ManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewManageddeviceenrollmenttopfailuresManagedDeviceEnrollmentTopFailuresRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
