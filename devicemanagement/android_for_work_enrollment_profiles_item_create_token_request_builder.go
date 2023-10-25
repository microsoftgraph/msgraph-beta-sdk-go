package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder provides operations to call the createToken method.
type AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderInternal instantiates a new CreateTokenRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) {
    m := &AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkEnrollmentProfiles/{androidForWorkEnrollmentProfile%2Did}/createToken", pathParameters),
    }
    return m
}
// NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder instantiates a new CreateTokenRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createToken
func (m *AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) Post(ctx context.Context, body AndroidForWorkEnrollmentProfilesItemCreateTokenPostRequestBodyable, requestConfiguration *AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action createToken
func (m *AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidForWorkEnrollmentProfilesItemCreateTokenPostRequestBodyable, requestConfiguration *AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) WithUrl(rawUrl string)(*AndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder) {
    return NewAndroidForWorkEnrollmentProfilesItemCreateTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
