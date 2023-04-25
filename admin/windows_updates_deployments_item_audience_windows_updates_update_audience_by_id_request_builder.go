package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder provides operations to call the updateAudienceById method.
type WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal instantiates a new WindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/windowsUpdates.updateAudienceById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder instantiates a new WindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateAudienceById
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action updateAudienceById
func (m *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
