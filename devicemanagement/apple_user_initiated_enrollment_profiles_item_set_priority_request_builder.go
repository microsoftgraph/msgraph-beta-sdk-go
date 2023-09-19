package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder provides operations to call the setPriority method.
type AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderInternal instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) {
    m := &AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}/setPriority", pathParameters),
    }
    return m
}
// NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder instantiates a new SetPriorityRequestBuilder and sets the default values.
func NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setPriority
func (m *AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) Post(ctx context.Context, body AppleUserInitiatedEnrollmentProfilesItemSetPriorityPostRequestBodyable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action setPriority
func (m *AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppleUserInitiatedEnrollmentProfilesItemSetPriorityPostRequestBodyable, requestConfiguration *AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) WithUrl(rawUrl string)(*AppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
