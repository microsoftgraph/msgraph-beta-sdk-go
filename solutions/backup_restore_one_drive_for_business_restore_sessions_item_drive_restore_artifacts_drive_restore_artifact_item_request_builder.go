package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder provides operations to manage the driveRestoreArtifacts property of the microsoft.graph.oneDriveForBusinessRestoreSession entity.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetQueryParameters a collection of restore points and destination details that can be used to restore a OneDrive for Business drive.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetQueryParameters
}
// BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderInternal instantiates a new BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) {
    m := &BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessRestoreSessions/{oneDriveForBusinessRestoreSession%2Did}/driveRestoreArtifacts/{driveRestoreArtifact%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder instantiates a new BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder and sets the default values.
func NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property driveRestoreArtifacts for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of restore points and destination details that can be used to restore a OneDrive for Business drive.
// returns a DriveRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable), nil
}
// Patch update the navigation property driveRestoreArtifacts in solutions
// returns a DriveRestoreArtifactable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRestoreArtifactFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable), nil
}
// RestorePoint provides operations to manage the restorePoint property of the microsoft.graph.restoreArtifactBase entity.
// returns a *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsItemRestorePointRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) RestorePoint()(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsItemRestorePointRequestBuilder) {
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsItemRestorePointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property driveRestoreArtifacts for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of restore points and destination details that can be used to restore a OneDrive for Business drive.
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property driveRestoreArtifacts in solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRestoreArtifactable, requestConfiguration *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder when successful
func (m *BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder) {
    return NewBackupRestoreOneDriveForBusinessRestoreSessionsItemDriveRestoreArtifactsDriveRestoreArtifactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
