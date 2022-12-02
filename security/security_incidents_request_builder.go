package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityIncidentsRequestBuilder provides operations to manage the incidents property of the microsoft.graph.security entity.
type SecurityIncidentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityIncidentsRequestBuilderGetQueryParameters get a list of incident objects that Microsoft 365 Defender has created to track attacks in an organization. Attacks are typically inflicted on different types of entities, such as devices, users, and mailboxes, resulting in multiple alert objects. Microsoft 365 Defender correlates alerts with the same attack techniques or the same attacker into an **incident**.  This operation allows you to filter and sort through incidents to create an informed cyber security response. It exposes a collection of incidents that were flagged in your network, within the time range you specified in your environment retention policy. The most recent incidents are displayed at the top of the list.
type SecurityIncidentsRequestBuilderGetQueryParameters struct {
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
// SecurityIncidentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityIncidentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityIncidentsRequestBuilderGetQueryParameters
}
// SecurityIncidentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityIncidentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityIncidentsRequestBuilderInternal instantiates a new IncidentsRequestBuilder and sets the default values.
func NewSecurityIncidentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityIncidentsRequestBuilder) {
    m := &SecurityIncidentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/incidents{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityIncidentsRequestBuilder instantiates a new IncidentsRequestBuilder and sets the default values.
func NewSecurityIncidentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityIncidentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityIncidentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *SecurityIncidentsRequestBuilder) Count()(*SecurityIncidentsCountRequestBuilder) {
    return NewSecurityIncidentsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get a list of incident objects that Microsoft 365 Defender has created to track attacks in an organization. Attacks are typically inflicted on different types of entities, such as devices, users, and mailboxes, resulting in multiple alert objects. Microsoft 365 Defender correlates alerts with the same attack techniques or the same attacker into an **incident**.  This operation allows you to filter and sort through incidents to create an informed cyber security response. It exposes a collection of incidents that were flagged in your network, within the time range you specified in your environment retention policy. The most recent incidents are displayed at the top of the list.
func (m *SecurityIncidentsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityIncidentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to incidents for security
func (m *SecurityIncidentsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Incidentable, requestConfiguration *SecurityIncidentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get a list of incident objects that Microsoft 365 Defender has created to track attacks in an organization. Attacks are typically inflicted on different types of entities, such as devices, users, and mailboxes, resulting in multiple alert objects. Microsoft 365 Defender correlates alerts with the same attack techniques or the same attacker into an **incident**.  This operation allows you to filter and sort through incidents to create an informed cyber security response. It exposes a collection of incidents that were flagged in your network, within the time range you specified in your environment retention policy. The most recent incidents are displayed at the top of the list.
func (m *SecurityIncidentsRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityIncidentsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IncidentCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateIncidentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.IncidentCollectionResponseable), nil
}
// Post create new navigation property to incidents for security
func (m *SecurityIncidentsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Incidentable, requestConfiguration *SecurityIncidentsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Incidentable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateIncidentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Incidentable), nil
}
