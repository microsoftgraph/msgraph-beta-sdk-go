package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CollaborationAnalyzedEmailsRequestBuilder provides operations to manage the analyzedEmails property of the microsoft.graph.security.collaborationRoot entity.
type CollaborationAnalyzedEmailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CollaborationAnalyzedEmailsRequestBuilderGetQueryParameters read the properties and relationships of an analyzedEmail object.
type CollaborationAnalyzedEmailsRequestBuilderGetQueryParameters struct {
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
// CollaborationAnalyzedEmailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedEmailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CollaborationAnalyzedEmailsRequestBuilderGetQueryParameters
}
// CollaborationAnalyzedEmailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedEmailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAnalyzedEmailId provides operations to manage the analyzedEmails property of the microsoft.graph.security.collaborationRoot entity.
// returns a *CollaborationAnalyzedEmailsAnalyzedEmailItemRequestBuilder when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) ByAnalyzedEmailId(analyzedEmailId string)(*CollaborationAnalyzedEmailsAnalyzedEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if analyzedEmailId != "" {
        urlTplParams["analyzedEmail%2Did"] = analyzedEmailId
    }
    return NewCollaborationAnalyzedEmailsAnalyzedEmailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCollaborationAnalyzedEmailsRequestBuilderInternal instantiates a new CollaborationAnalyzedEmailsRequestBuilder and sets the default values.
func NewCollaborationAnalyzedEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedEmailsRequestBuilder) {
    m := &CollaborationAnalyzedEmailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/collaboration/analyzedEmails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCollaborationAnalyzedEmailsRequestBuilder instantiates a new CollaborationAnalyzedEmailsRequestBuilder and sets the default values.
func NewCollaborationAnalyzedEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedEmailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCollaborationAnalyzedEmailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CollaborationAnalyzedEmailsCountRequestBuilder when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) Count()(*CollaborationAnalyzedEmailsCountRequestBuilder) {
    return NewCollaborationAnalyzedEmailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an analyzedEmail object.
// returns a AnalyzedEmailCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedEmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *CollaborationAnalyzedEmailsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateAnalyzedEmailCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailCollectionResponseable), nil
}
// MicrosoftGraphSecurityRemediate provides operations to call the remediate method.
// returns a *CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRequestBuilder when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) MicrosoftGraphSecurityRemediate()(*CollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRequestBuilder) {
    return NewCollaborationAnalyzedEmailsMicrosoftGraphSecurityRemediateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to analyzedEmails for security
// returns a AnalyzedEmailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedEmailsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedEmailsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateAnalyzedEmailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable), nil
}
// ToGetRequestInformation read the properties and relationships of an analyzedEmail object.
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CollaborationAnalyzedEmailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to analyzedEmails for security
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedEmailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CollaborationAnalyzedEmailsRequestBuilder when successful
func (m *CollaborationAnalyzedEmailsRequestBuilder) WithUrl(rawUrl string)(*CollaborationAnalyzedEmailsRequestBuilder) {
    return NewCollaborationAnalyzedEmailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
