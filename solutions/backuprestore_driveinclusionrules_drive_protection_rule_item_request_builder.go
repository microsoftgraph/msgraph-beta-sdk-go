package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder provides operations to manage the driveInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters the list of drive inclusion rules applied to the tenant.
type BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters
}
// BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal instantiates a new BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder and sets the default values.
func NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    m := &BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/driveInclusionRules/{driveProtectionRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder instantiates a new BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder and sets the default values.
func NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property driveInclusionRules for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of drive inclusion rules applied to the tenant.
// returns a DriveProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property driveInclusionRules in solutions
// returns a DriveProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property driveInclusionRules for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of drive inclusion rules applied to the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property driveInclusionRules in solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, requestConfiguration *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder when successful
func (m *BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    return NewBackuprestoreDriveinclusionrulesDriveProtectionRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
