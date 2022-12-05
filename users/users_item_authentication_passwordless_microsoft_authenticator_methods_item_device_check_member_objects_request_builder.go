package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderInternal instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder) {
    m := &UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/microsoft.graph.checkMemberObjects";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action checkMemberObjects
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsPostRequestBodyable, requestConfiguration *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action checkMemberObjects
func (m *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsPostRequestBodyable, requestConfiguration *UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceCheckMemberObjectsResponseable), nil
}
