package credentialuserregistrationssummaries

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    i24f4eb51f717c3034d4688f6ba15526e13c45a6c6e98a9701f1318efcd704d4b "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/credentialuserregistrationssummaries/count"
)

// CredentialUserRegistrationsSummariesRequestBuilder provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type CredentialUserRegistrationsSummariesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CredentialUserRegistrationsSummariesRequestBuilderGetOptions options for Get
type CredentialUserRegistrationsSummariesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
type CredentialUserRegistrationsSummariesRequestBuilderGetQueryParameters struct {
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
// CredentialUserRegistrationsSummariesRequestBuilderPostOptions options for Post
type CredentialUserRegistrationsSummariesRequestBuilderPostOptions struct {
    // 
    Body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewCredentialUserRegistrationsSummariesRequestBuilderInternal instantiates a new CredentialUserRegistrationsSummariesRequestBuilder and sets the default values.
func NewCredentialUserRegistrationsSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CredentialUserRegistrationsSummariesRequestBuilder) {
    m := &CredentialUserRegistrationsSummariesRequestBuilder{
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
// NewCredentialUserRegistrationsSummariesRequestBuilder instantiates a new CredentialUserRegistrationsSummariesRequestBuilder and sets the default values.
func NewCredentialUserRegistrationsSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CredentialUserRegistrationsSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCredentialUserRegistrationsSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *CredentialUserRegistrationsSummariesRequestBuilder) Count()(*i24f4eb51f717c3034d4688f6ba15526e13c45a6c6e98a9701f1318efcd704d4b.CountRequestBuilder) {
    return i24f4eb51f717c3034d4688f6ba15526e13c45a6c6e98a9701f1318efcd704d4b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *CredentialUserRegistrationsSummariesRequestBuilder) CreateGetRequestInformation(options *CredentialUserRegistrationsSummariesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to credentialUserRegistrationsSummaries for tenantRelationships
func (m *CredentialUserRegistrationsSummariesRequestBuilder) CreatePostRequestInformation(options *CredentialUserRegistrationsSummariesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Get summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *CredentialUserRegistrationsSummariesRequestBuilder) Get(options *CredentialUserRegistrationsSummariesRequestBuilderGetOptions)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateCredentialUserRegistrationsSummaryCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryCollectionResponseable), nil
}
// Post create new navigation property to credentialUserRegistrationsSummaries for tenantRelationships
func (m *CredentialUserRegistrationsSummariesRequestBuilder) Post(options *CredentialUserRegistrationsSummariesRequestBuilderPostOptions)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CredentialUserRegistrationsSummaryable), nil
}
