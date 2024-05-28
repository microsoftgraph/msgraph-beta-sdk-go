package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder provides operations to call the remediate method.
type CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderInternal instantiates a new CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) {
    m := &CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/collaboration/analyzedEmails/microsoft.graph.security.remediate", pathParameters),
    }
    return m
}
// NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder instantiates a new CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder and sets the default values.
func NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove a potential threat from end users' mailboxes. Remediation means to take prescribed action against a threat. This API can trigger email purge actions like move to junk, move to deleted items, soft delete, hard delete, or move to Inbox. This API enables scenarios and use cases such as SOAR integration, playbooks, and automations. For more information read email remediation, trigger action and track actions. If there is false positives admins can take move to inbox action.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-analyzedemail-remediate?view=graph-rest-beta
func (m *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) Post(ctx context.Context, body CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateRemediatePostRequestBodyable, requestConfiguration *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation remove a potential threat from end users' mailboxes. Remediation means to take prescribed action against a threat. This API can trigger email purge actions like move to junk, move to deleted items, soft delete, hard delete, or move to Inbox. This API enables scenarios and use cases such as SOAR integration, playbooks, and automations. For more information read email remediation, trigger action and track actions. If there is false positives admins can take move to inbox action.
// returns a *RequestInformation when successful
func (m *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) ToPostRequestInformation(ctx context.Context, body CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateRemediatePostRequestBodyable, requestConfiguration *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder when successful
func (m *CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) WithUrl(rawUrl string)(*CollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder) {
    return NewCollaborationAnalyzedemailsMicrosoftgraphsecurityremediateMicrosoftGraphSecurityRemediateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
