package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder provides operations to call the import method.
type ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImportedwindowsautopilotdeviceidentitiesImportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedwindowsautopilotdeviceidentitiesImportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilderInternal instantiates a new ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) {
    m := &ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities/import", pathParameters),
    }
    return m
}
// NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilder instantiates a new ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action import
// Deprecated: This method is obsolete. Use PostAsImportPostResponse instead.
// returns a ImportedwindowsautopilotdeviceidentitiesImportResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) Post(ctx context.Context, body ImportedwindowsautopilotdeviceidentitiesImportPostRequestBodyable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilderPostRequestConfiguration)(ImportedwindowsautopilotdeviceidentitiesImportResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedwindowsautopilotdeviceidentitiesImportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedwindowsautopilotdeviceidentitiesImportResponseable), nil
}
// PostAsImportPostResponse invoke action import
// returns a ImportedwindowsautopilotdeviceidentitiesImportPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) PostAsImportPostResponse(ctx context.Context, body ImportedwindowsautopilotdeviceidentitiesImportPostRequestBodyable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilderPostRequestConfiguration)(ImportedwindowsautopilotdeviceidentitiesImportPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImportedwindowsautopilotdeviceidentitiesImportPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportedwindowsautopilotdeviceidentitiesImportPostResponseable), nil
}
// ToPostRequestInformation invoke action import
// returns a *RequestInformation when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImportedwindowsautopilotdeviceidentitiesImportPostRequestBodyable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) WithUrl(rawUrl string)(*ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
