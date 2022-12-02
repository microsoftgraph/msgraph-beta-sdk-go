package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters get a list of the credentialUserRegistrationsSummary objects and their properties.
type TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal instantiates a new CredentialUserRegistrationsSummariesRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/credentialUserRegistrationsSummaries{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder instantiates a new CredentialUserRegistrationsSummariesRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) Count()(*TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesCountRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of the credentialUserRegistrationsSummary objects and their properties.
func (m *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to credentialUserRegistrationsSummaries for tenantRelationships
func (m *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable, requestConfiguration *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get a list of the credentialUserRegistrationsSummary objects and their properties.
func (m *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateCredentialUserRegistrationsSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryCollectionResponseable), nil
}
// Post create new navigation property to credentialUserRegistrationsSummaries for tenantRelationships
func (m *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable, requestConfiguration *TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable), nil
}
