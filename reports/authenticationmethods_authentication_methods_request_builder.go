package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationmethodsAuthenticationMethodsRequestBuilder provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
type AuthenticationmethodsAuthenticationMethodsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationmethodsAuthenticationMethodsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsAuthenticationMethodsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationmethodsAuthenticationMethodsRequestBuilderGetQueryParameters container for navigation properties for Microsoft Entra authentication methods resources.
type AuthenticationmethodsAuthenticationMethodsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationmethodsAuthenticationMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsAuthenticationMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationmethodsAuthenticationMethodsRequestBuilderGetQueryParameters
}
// AuthenticationmethodsAuthenticationMethodsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsAuthenticationMethodsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationmethodsAuthenticationMethodsRequestBuilderInternal instantiates a new AuthenticationmethodsAuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationmethodsAuthenticationMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsAuthenticationMethodsRequestBuilder) {
    m := &AuthenticationmethodsAuthenticationMethodsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/authenticationMethods{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationmethodsAuthenticationMethodsRequestBuilder instantiates a new AuthenticationmethodsAuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationmethodsAuthenticationMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsAuthenticationMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationmethodsAuthenticationMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authenticationMethods for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get container for navigation properties for Microsoft Entra authentication methods resources.
// returns a AuthenticationMethodsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationMethodsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable), nil
}
// Patch update the navigation property authenticationMethods in reports
// returns a AuthenticationMethodsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationMethodsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable), nil
}
// ToDeleteRequestInformation delete navigation property authenticationMethods for reports
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation container for navigation properties for Microsoft Entra authentication methods resources.
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property authenticationMethods in reports
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, requestConfiguration *AuthenticationmethodsAuthenticationMethodsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UserRegistrationDetails provides operations to manage the userRegistrationDetails property of the microsoft.graph.authenticationMethodsRoot entity.
// returns a *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) UserRegistrationDetails()(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilder) {
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsersRegisteredByFeature provides operations to call the usersRegisteredByFeature method.
// returns a *AuthenticationmethodsUsersregisteredbyfeatureUsersRegisteredByFeatureRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) UsersRegisteredByFeature()(*AuthenticationmethodsUsersregisteredbyfeatureUsersRegisteredByFeatureRequestBuilder) {
    return NewAuthenticationmethodsUsersregisteredbyfeatureUsersRegisteredByFeatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByFeature method.
// returns a *AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*AuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return NewAuthenticationmethodsUsersregisteredbyfeaturewithincludedusertypeswithincludeduserrolesUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, includedUserRoles, includedUserTypes)
}
// UsersRegisteredByMethod provides operations to call the usersRegisteredByMethod method.
// returns a *AuthenticationmethodsUsersregisteredbymethodUsersRegisteredByMethodRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) UsersRegisteredByMethod()(*AuthenticationmethodsUsersregisteredbymethodUsersRegisteredByMethodRequestBuilder) {
    return NewAuthenticationmethodsUsersregisteredbymethodUsersRegisteredByMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByMethod method.
// returns a *AuthenticationmethodsUsersregisteredbymethodwithincludedusertypeswithincludeduserrolesUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*AuthenticationmethodsUsersregisteredbymethodwithincludedusertypeswithincludeduserrolesUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return NewAuthenticationmethodsUsersregisteredbymethodwithincludedusertypeswithincludeduserrolesUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, includedUserRoles, includedUserTypes)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AuthenticationmethodsAuthenticationMethodsRequestBuilder when successful
func (m *AuthenticationmethodsAuthenticationMethodsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationmethodsAuthenticationMethodsRequestBuilder) {
    return NewAuthenticationmethodsAuthenticationMethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
