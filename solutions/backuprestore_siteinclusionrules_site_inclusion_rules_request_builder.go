package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder provides operations to manage the siteInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetQueryParameters the list of site inclusion rules applied to the tenant.
type BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetQueryParameters struct {
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
// BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetQueryParameters
}
// BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySiteProtectionRuleId provides operations to manage the siteInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSiteinclusionrulesSiteProtectionRuleItemRequestBuilder when successful
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) BySiteProtectionRuleId(siteProtectionRuleId string)(*BackuprestoreSiteinclusionrulesSiteProtectionRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if siteProtectionRuleId != "" {
        urlTplParams["siteProtectionRule%2Did"] = siteProtectionRuleId
    }
    return NewBackuprestoreSiteinclusionrulesSiteProtectionRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderInternal instantiates a new BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder and sets the default values.
func NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) {
    m := &BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/siteInclusionRules{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder instantiates a new BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder and sets the default values.
func NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreSiteinclusionrulesCountRequestBuilder when successful
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) Count()(*BackuprestoreSiteinclusionrulesCountRequestBuilder) {
    return NewBackuprestoreSiteinclusionrulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of site inclusion rules applied to the tenant.
// returns a SiteProtectionRuleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleCollectionResponseable), nil
}
// Post create new navigation property to siteInclusionRules for solutions
// returns a SiteProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleable, requestConfiguration *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteProtectionRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleable), nil
}
// ToGetRequestInformation the list of site inclusion rules applied to the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to siteInclusionRules for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SiteProtectionRuleable, requestConfiguration *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder when successful
func (m *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) {
    return NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
