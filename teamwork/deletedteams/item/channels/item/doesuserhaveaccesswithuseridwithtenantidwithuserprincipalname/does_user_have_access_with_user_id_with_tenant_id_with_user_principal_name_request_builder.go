package doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder provides operations to call the doesUserHaveAccess method.
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters invoke function doesUserHaveAccess
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
    // Usage: tenantId='{tenantId}'
    TenantId *string
    // Usage: userId='{userId}'
    UserId *string
    // Usage: userPrincipalName='{userPrincipalName}'
    UserPrincipalName *string
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    m := &DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/microsoft.graph.doesUserHaveAccess(userId='{userId}',tenantId='{tenantId}',userPrincipalName='{userPrincipalName}'){?userId,tenantId,userPrincipalName}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable), nil
}
