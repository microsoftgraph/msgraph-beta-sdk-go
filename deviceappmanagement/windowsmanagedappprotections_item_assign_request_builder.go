package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsmanagedappprotectionsItemAssignRequestBuilder provides operations to call the assign method.
type WindowsmanagedappprotectionsItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsmanagedappprotectionsItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsmanagedappprotectionsItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsmanagedappprotectionsItemAssignRequestBuilderInternal instantiates a new WindowsmanagedappprotectionsItemAssignRequestBuilder and sets the default values.
func NewWindowsmanagedappprotectionsItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagedappprotectionsItemAssignRequestBuilder) {
    m := &WindowsmanagedappprotectionsItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}/assign", pathParameters),
    }
    return m
}
// NewWindowsmanagedappprotectionsItemAssignRequestBuilder instantiates a new WindowsmanagedappprotectionsItemAssignRequestBuilder and sets the default values.
func NewWindowsmanagedappprotectionsItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagedappprotectionsItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsmanagedappprotectionsItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsmanagedappprotectionsItemAssignRequestBuilder) Post(ctx context.Context, body WindowsmanagedappprotectionsItemAssignPostRequestBodyable, requestConfiguration *WindowsmanagedappprotectionsItemAssignRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action assign
// returns a *RequestInformation when successful
func (m *WindowsmanagedappprotectionsItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsmanagedappprotectionsItemAssignPostRequestBodyable, requestConfiguration *WindowsmanagedappprotectionsItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsmanagedappprotectionsItemAssignRequestBuilder when successful
func (m *WindowsmanagedappprotectionsItemAssignRequestBuilder) WithUrl(rawUrl string)(*WindowsmanagedappprotectionsItemAssignRequestBuilder) {
    return NewWindowsmanagedappprotectionsItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
