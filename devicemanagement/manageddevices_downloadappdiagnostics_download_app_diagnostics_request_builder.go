package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder provides operations to call the downloadAppDiagnostics method.
type ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderInternal instantiates a new ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder and sets the default values.
func NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) {
    m := &ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/downloadAppDiagnostics", pathParameters),
    }
    return m
}
// NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder instantiates a new ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder and sets the default values.
func NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action downloadAppDiagnostics
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) Post(ctx context.Context, body ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsPostRequestBodyable, requestConfiguration *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation invoke action downloadAppDiagnostics
// returns a *RequestInformation when successful
func (m *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsPostRequestBodyable, requestConfiguration *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder when successful
func (m *ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder) {
    return NewManageddevicesDownloadappdiagnosticsDownloadAppDiagnosticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
