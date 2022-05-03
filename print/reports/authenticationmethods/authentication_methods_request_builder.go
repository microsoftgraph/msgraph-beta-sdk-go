package authenticationmethods

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbymethod"
    i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbymethodwithincludedusertypeswithincludeduserroles"
    i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbyfeaturewithincludedusertypeswithincludeduserroles"
    ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/userregistrationdetails"
    ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/usersregisteredbyfeature"
    i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods/userregistrationdetails/item"
)

// AuthenticationMethodsRequestBuilder provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
type AuthenticationMethodsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMethodsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationMethodsRequestBuilderGetQueryParameters container for navigation properties for Azure AD authentication methods resources.
type AuthenticationMethodsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodsRequestBuilderGetQueryParameters
}
// AuthenticationMethodsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationMethodsRequestBuilderInternal instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationMethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    m := &AuthenticationMethodsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/authenticationMethods{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMethodsRequestBuilder instantiates a new AuthenticationMethodsRequestBuilder and sets the default values.
func NewAuthenticationMethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property authenticationMethods for print
func (m *AuthenticationMethodsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property authenticationMethods for print
func (m *AuthenticationMethodsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationMethodsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformationWithRequestConfiguration container for navigation properties for Azure AD authentication methods resources.
func (m *AuthenticationMethodsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration container for navigation properties for Azure AD authentication methods resources.
func (m *AuthenticationMethodsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationMethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property authenticationMethods in print
func (m *AuthenticationMethodsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property authenticationMethods in print
func (m *AuthenticationMethodsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, requestConfiguration *AuthenticationMethodsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property authenticationMethods for print
func (m *AuthenticationMethodsRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AuthenticationMethodsRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property authenticationMethods for print
func (m *AuthenticationMethodsRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AuthenticationMethodsRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler container for navigation properties for Azure AD authentication methods resources.
func (m *AuthenticationMethodsRequestBuilder) GetWithResponseHandler(requestConfiguration *AuthenticationMethodsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler container for navigation properties for Azure AD authentication methods resources.
func (m *AuthenticationMethodsRequestBuilder) GetWithResponseHandler(requestConfiguration *AuthenticationMethodsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationMethodsRootFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable), nil
}
// PatchWithResponseHandler update the navigation property authenticationMethods in print
func (m *AuthenticationMethodsRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, requestConfiguration *AuthenticationMethodsRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property authenticationMethods in print
func (m *AuthenticationMethodsRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationMethodsRootable, requestConfiguration *AuthenticationMethodsRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserRegistrationDetails the userRegistrationDetails property
func (m *AuthenticationMethodsRequestBuilder) UserRegistrationDetails()(*ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31.UserRegistrationDetailsRequestBuilder) {
    return ie54d2504e9f581782cfc818a41a8b995a582cdd5f7976792d36e04e314a0ca31.NewUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.authenticationMethods.userRegistrationDetails.item collection
func (m *AuthenticationMethodsRequestBuilder) UserRegistrationDetailsById(id string)(*i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0.UserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userRegistrationDetails%2Did"] = id
    }
    return i8465f6f7d0d944a33802781ad4947a28b996037208b7f788a58b7c1946465bf0.NewUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UsersRegisteredByFeature provides operations to call the usersRegisteredByFeature method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeature()(*ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487.UsersRegisteredByFeatureRequestBuilder) {
    return ieccbed6a2bbaa40cf04b762541a9ab50dd6ed3158a44574052a51b40a95f8487.NewUsersRegisteredByFeatureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByFeature method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0.UsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return i47f70708cebf48f687242c96efc5507b89cca5b077a72c1d685b5c61a39d46c0.NewUsersRegisteredByFeatureWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserRoles, includedUserTypes);
}
// UsersRegisteredByMethod provides operations to call the usersRegisteredByMethod method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethod()(*i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134.UsersRegisteredByMethodRequestBuilder) {
    return i0dd7ffa9993603b98e542fb6765c2bc58b69867a61d8a29faf5aa5bd28dc5134.NewUsersRegisteredByMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles provides operations to call the usersRegisteredByMethod method.
func (m *AuthenticationMethodsRequestBuilder) UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRoles(includedUserRoles *string, includedUserTypes *string)(*i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439.UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    return i1bba8e4b27ab9c120aa0ff2d39a209b233e51379ecc5ae53660f9bc4e1dd4439.NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter, includedUserRoles, includedUserTypes);
}
