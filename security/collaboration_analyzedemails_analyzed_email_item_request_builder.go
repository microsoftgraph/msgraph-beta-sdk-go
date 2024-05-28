package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder provides operations to manage the analyzedEmails property of the microsoft.graph.security.collaborationRoot entity.
type CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetQueryParameters read the properties and relationships of an analyzedEmail object.
type CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetQueryParameters
}
// CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderInternal instantiates a new CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) {
    m := &CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/collaboration/analyzedEmails/{analyzedEmail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder instantiates a new CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property analyzedEmails for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an analyzedEmail object.
// returns a AnalyzedEmailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-analyzedemail-get?view=graph-rest-beta
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property analyzedEmails in security
// returns a AnalyzedEmailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property analyzedEmails for security
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an analyzedEmail object.
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property analyzedEmails in security
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.AnalyzedEmailable, requestConfiguration *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder when successful
func (m *CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) WithUrl(rawUrl string)(*CollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder) {
    return NewCollaborationAnalyzedemailsAnalyzedEmailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
