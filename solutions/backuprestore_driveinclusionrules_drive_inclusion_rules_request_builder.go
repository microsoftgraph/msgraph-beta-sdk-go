package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder provides operations to manage the driveInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetQueryParameters the list of drive inclusion rules applied to the tenant.
type BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetQueryParameters struct {
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
// BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetQueryParameters
}
// BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDriveProtectionRuleId provides operations to manage the driveInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder when successful
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) ByDriveProtectionRuleId(driveProtectionRuleId string)(*BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if driveProtectionRuleId != "" {
        urlTplParams["driveProtectionRule%2Did"] = driveProtectionRuleId
    }
    return NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderInternal instantiates a new BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder and sets the default values.
func NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) {
    m := &BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/driveInclusionRules{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder instantiates a new BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder and sets the default values.
func NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreDriveinclusionrulesCountRequestBuilder when successful
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) Count()(*BackuprestoreDriveinclusionrulesCountRequestBuilder) {
    return NewBackuprestoreDriveinclusionrulesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of drive inclusion rules applied to the tenant.
// returns a DriveProtectionRuleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionRuleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleCollectionResponseable), nil
}
// Post create new navigation property to driveInclusionRules for solutions
// returns a DriveProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, requestConfiguration *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable), nil
}
// ToGetRequestInformation the list of drive inclusion rules applied to the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to driveInclusionRules for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, requestConfiguration *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder when successful
func (m *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) {
    return NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
