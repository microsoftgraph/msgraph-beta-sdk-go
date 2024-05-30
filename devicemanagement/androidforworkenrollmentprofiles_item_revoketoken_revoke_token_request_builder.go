package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder provides operations to call the revokeToken method.
type AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderInternal instantiates a new AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder and sets the default values.
func NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) {
    m := &AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkEnrollmentProfiles/{androidForWorkEnrollmentProfile%2Did}/revokeToken", pathParameters),
    }
    return m
}
// NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder instantiates a new AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder and sets the default values.
func NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action revokeToken
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action revokeToken
// returns a *RequestInformation when successful
func (m *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder when successful
func (m *AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) WithUrl(rawUrl string)(*AndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder) {
    return NewAndroidforworkenrollmentprofilesItemRevoketokenRevokeTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
