package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters windows managed app policies.
type WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.windowsManagedAppProtection entity.
// returns a *WindowsManagedAppProtectionsItemAppsRequestBuilder when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Apps()(*WindowsManagedAppProtectionsItemAppsRequestBuilder) {
    return NewWindowsManagedAppProtectionsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assign provides operations to call the assign method.
// returns a *WindowsManagedAppProtectionsItemAssignRequestBuilder when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Assign()(*WindowsManagedAppProtectionsItemAssignRequestBuilder) {
    return NewWindowsManagedAppProtectionsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsManagedAppProtection entity.
// returns a *WindowsManagedAppProtectionsItemAssignmentsRequestBuilder when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Assignments()(*WindowsManagedAppProtectionsItemAssignmentsRequestBuilder) {
    return NewWindowsManagedAppProtectionsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal instantiates a new WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    m := &WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder instantiates a new WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsManagedAppProtections for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windows managed app policies.
// returns a WindowsManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// Patch update the navigation property windowsManagedAppProtections in deviceAppManagement
// returns a WindowsManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// TargetApps provides operations to call the targetApps method.
// returns a *WindowsManagedAppProtectionsItemTargetAppsRequestBuilder when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) TargetApps()(*WindowsManagedAppProtectionsItemTargetAppsRequestBuilder) {
    return NewWindowsManagedAppProtectionsItemTargetAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property windowsManagedAppProtections for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation windows managed app policies.
// returns a *RequestInformation when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsManagedAppProtections in deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder when successful
func (m *WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) WithUrl(rawUrl string)(*WindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder) {
    return NewWindowsManagedAppProtectionsWindowsManagedAppProtectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
