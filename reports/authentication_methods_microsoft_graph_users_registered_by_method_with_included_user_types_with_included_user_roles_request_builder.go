package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder provides operations to call the usersRegisteredByMethod method.
type AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal instantiates a new MicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewAuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, includedUserRoles *string, includedUserTypes *string)(*AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    m := &AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/authenticationMethods/microsoft.graph.usersRegisteredByMethod(includedUserTypes='{includedUserTypes}',includedUserRoles='{includedUserRoles}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if includedUserRoles != nil {
        urlTplParams["includedUserRoles"] = *includedUserRoles
    }
    if includedUserTypes != nil {
        urlTplParams["includedUserTypes"] = *includedUserTypes
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder instantiates a new MicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewAuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function usersRegisteredByMethod
func (m *AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationMethodSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRegistrationMethodSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationMethodSummaryable), nil
}
// ToGetRequestInformation invoke function usersRegisteredByMethod
func (m *AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsMicrosoftGraphUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
