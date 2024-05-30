package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder provides operations to manage the analyzedEmails property of the microsoft.graph.security.collaborationRoot entity.
type CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetQueryParameters read the properties and relationships of an analyzedEmail object.
type CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetQueryParameters struct {
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
// CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetQueryParameters
}
// CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAnalyzedEmailId provides operations to manage the analyzedEmails property of the microsoft.graph.security.collaborationRoot entity.
// returns a *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) ByAnalyzedEmailId(analyzedEmailId string)(*CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if analyzedEmailId != "" {
        urlTplParams["analyzedEmail%2Did"] = analyzedEmailId
    }
    return NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderInternal instantiates a new CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) {
    m := &CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/collaboration/analyzedEmails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder instantiates a new CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CollaborationAnalyzedemailsCountRequestBuilder when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) Count()(*CollaborationAnalyzedemailsCountRequestBuilder) {
    return NewCollaborationAnalyzedemailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an analyzedEmail object.
// returns a AnalyzedEmailCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailCollectionResponseable, error) {
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
// returns a *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) MicrosoftGraphSecurityRemediate()(*CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) {
    return NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to analyzedEmails for security
// returns a AnalyzedEmailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, error) {
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
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) WithUrl(rawUrl string)(*CollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder) {
    return NewCollaborationAnalyzedemailsAnalyzedEmailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
