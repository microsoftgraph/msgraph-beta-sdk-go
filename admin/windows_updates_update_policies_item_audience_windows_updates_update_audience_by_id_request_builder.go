package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder provides operations to call the updateAudienceById method.
type WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal instantiates a new WindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/windowsUpdates.updateAudienceById", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder instantiates a new WindowsUpdatesUpdateAudienceByIdRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateAudienceById
func (m *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) Post(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdUpdateAudienceByIdPostRequestBodyable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceWindowsUpdatesUpdateAudienceByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
