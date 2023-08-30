package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder provides operations to call the revokeToken method.
type AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderInternal instantiates a new RevokeTokenRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) {
    m := &AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkEnrollmentProfiles/{androidForWorkEnrollmentProfile%2Did}/revokeToken", pathParameters),
    }
    return m
}
// NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder instantiates a new RevokeTokenRequestBuilder and sets the default values.
func NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action revokeToken
func (m *AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) Post(ctx context.Context, requestConfiguration *AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action revokeToken
func (m *AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) WithUrl(rawUrl string)(*AndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder) {
    return NewAndroidForWorkEnrollmentProfilesItemRevokeTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
