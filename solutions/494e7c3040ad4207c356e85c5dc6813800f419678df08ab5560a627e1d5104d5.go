package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder provides operations to manage the driveProtectionUnits property of the microsoft.graph.oneDriveForBusinessProtectionPolicy entity.
type BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetQueryParameters contains the protection units associated with a  OneDrive for Work or School protection policy.
type BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetQueryParameters
}
// NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderInternal instantiates a new BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) {
    m := &BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessProtectionPolicies/{oneDriveForBusinessProtectionPolicy%2Did}/driveProtectionUnits/{driveProtectionUnit%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder instantiates a new BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get contains the protection units associated with a  OneDrive for Work or School protection policy.
// returns a DriveProtectionUnitable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionUnitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionUnitable), nil
}
// ToGetRequestInformation contains the protection units associated with a  OneDrive for Work or School protection policy.
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder) {
    return NewBackupRestoreOneDriveForBusinessProtectionPoliciesItemDriveProtectionUnitsDriveProtectionUnitItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
