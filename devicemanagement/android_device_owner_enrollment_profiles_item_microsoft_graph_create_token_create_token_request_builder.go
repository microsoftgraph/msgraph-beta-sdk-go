package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder provides operations to call the createToken method.
type AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderInternal instantiates a new CreateTokenRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder) {
    m := &AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile%2Did}/microsoft.graph.createToken";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder instantiates a new CreateTokenRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createToken
func (m *AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder) Post(ctx context.Context, body AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBodyable, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action createToken
func (m *AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenPostRequestBodyable, requestConfiguration *AndroidDeviceOwnerEnrollmentProfilesItemMicrosoftGraphCreateTokenCreateTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
