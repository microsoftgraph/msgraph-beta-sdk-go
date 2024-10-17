package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder provides operations to manage the exchangeProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetQueryParameters the list of Exchange protection policies in the tenant.
type BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetQueryParameters
}
// BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderInternal instantiates a new BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) {
    m := &BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/exchangeProtectionPolicies/{exchangeProtectionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder instantiates a new BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exchangeProtectionPolicies for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of Exchange protection policies in the tenant.
// returns a ExchangeProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExchangeProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable), nil
}
// MailboxInclusionRules provides operations to manage the mailboxInclusionRules property of the microsoft.graph.exchangeProtectionPolicy entity.
// returns a *BackupRestoreExchangeProtectionPoliciesItemMailboxInclusionRulesRequestBuilder when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) MailboxInclusionRules()(*BackupRestoreExchangeProtectionPoliciesItemMailboxInclusionRulesRequestBuilder) {
    return NewBackupRestoreExchangeProtectionPoliciesItemMailboxInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MailboxProtectionUnits provides operations to manage the mailboxProtectionUnits property of the microsoft.graph.exchangeProtectionPolicy entity.
// returns a *BackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsRequestBuilder when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) MailboxProtectionUnits()(*BackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsRequestBuilder) {
    return NewBackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MailboxProtectionUnitsBulkAdditionJobs provides operations to manage the mailboxProtectionUnitsBulkAdditionJobs property of the microsoft.graph.exchangeProtectionPolicy entity.
// returns a *BackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsBulkAdditionJobsRequestBuilder when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) MailboxProtectionUnitsBulkAdditionJobs()(*BackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsBulkAdditionJobsRequestBuilder) {
    return NewBackupRestoreExchangeProtectionPoliciesItemMailboxProtectionUnitsBulkAdditionJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update an Exchange protection policy. This method adds a mailboxprotectionunit to or removes it from the protection policy.
// returns a ExchangeProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/exchangeprotectionpolicy-update?view=graph-rest-beta
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExchangeProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property exchangeProtectionPolicies for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of Exchange protection policies in the tenant.
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update an Exchange protection policy. This method adds a mailboxprotectionunit to or removes it from the protection policy.
// returns a *RequestInformation when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeProtectionPolicyable, requestConfiguration *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder when successful
func (m *BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder) {
    return NewBackupRestoreExchangeProtectionPoliciesExchangeProtectionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
