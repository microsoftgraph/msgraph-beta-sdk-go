package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder provides operations to manage the mailboxInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetQueryParameters the list of mailbox inclusion rules applied to the tenant.
type BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetQueryParameters
}
// BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderInternal instantiates a new BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder and sets the default values.
func NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) {
    m := &BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/mailboxInclusionRules/{mailboxProtectionRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder instantiates a new BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder and sets the default values.
func NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mailboxInclusionRules for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of mailbox inclusion rules applied to the tenant.
// returns a MailboxProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxProtectionRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable), nil
}
// Patch update the navigation property mailboxInclusionRules in solutions
// returns a MailboxProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailboxProtectionRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable), nil
}
// ToDeleteRequestInformation delete navigation property mailboxInclusionRules for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of mailbox inclusion rules applied to the tenant.
// returns a *RequestInformation when successful
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mailboxInclusionRules in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailboxProtectionRuleable, requestConfiguration *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder when successful
func (m *BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder) {
    return NewBackupRestoreMailboxInclusionRulesMailboxProtectionRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
