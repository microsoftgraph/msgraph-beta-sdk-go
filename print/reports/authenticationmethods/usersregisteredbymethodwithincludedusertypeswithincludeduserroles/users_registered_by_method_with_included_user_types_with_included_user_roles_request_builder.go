package usersregisteredbymethodwithincludedusertypeswithincludeduserroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder provides operations to call the usersRegisteredByMethod method.
type UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal instantiates a new UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, includedUserRoles *string, includedUserTypes *string)(*UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    m := &UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/authenticationMethods/microsoft.graph.usersRegisteredByMethod(includedUserTypes='{includedUserTypes}',includedUserRoles='{includedUserRoles}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if includedUserRoles != nil {
        urlTplParams[""] = *includedUserRoles
    }
    if includedUserTypes != nil {
        urlTplParams[""] = *includedUserTypes
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder instantiates a new UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder and sets the default values.
func NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformationWithRequestConfiguration invoke function usersRegisteredByMethod
func (m *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function usersRegisteredByMethod
func (m *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// GetWithResponseHandler invoke function usersRegisteredByMethod
func (m *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) GetWithResponseHandler(requestConfiguration *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationMethodSummaryable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler invoke function usersRegisteredByMethod
func (m *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilder) GetWithResponseHandler(requestConfiguration *UsersRegisteredByMethodWithIncludedUserTypesWithIncludedUserRolesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationMethodSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserRegistrationMethodSummaryFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserRegistrationMethodSummaryable), nil
}
