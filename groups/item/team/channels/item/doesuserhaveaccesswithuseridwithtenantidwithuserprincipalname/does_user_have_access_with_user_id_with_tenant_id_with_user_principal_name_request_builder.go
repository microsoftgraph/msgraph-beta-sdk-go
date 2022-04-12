package doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions options for Get
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    m := &DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/microsoft.graph.doesUserHaveAccess(userId='{userId}',tenantId='{tenantId}',userPrincipalName='{userPrincipalName}'){?userId,tenantId,userPrincipalName}";
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
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformation(options *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) Get(options *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions)(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable), nil
}
