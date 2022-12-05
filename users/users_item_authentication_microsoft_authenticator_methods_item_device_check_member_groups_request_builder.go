package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder provides operations to call the checkMemberGroups method.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderInternal instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder) {
    m := &UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/microsoft.graph.checkMemberGroups";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Azure AD. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyable, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Azure AD. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsPostRequestBodyable, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceCheckMemberGroupsResponseable), nil
}
