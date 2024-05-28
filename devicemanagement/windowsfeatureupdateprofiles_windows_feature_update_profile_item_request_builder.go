package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows feature update profiles
type WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetQueryParameters
}
// WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *WindowsfeatureupdateprofilesItemAssignRequestBuilder when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) Assign()(*WindowsfeatureupdateprofilesItemAssignRequestBuilder) {
    return NewWindowsfeatureupdateprofilesItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsFeatureUpdateProfile entity.
// returns a *WindowsfeatureupdateprofilesItemAssignmentsRequestBuilder when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) Assignments()(*WindowsfeatureupdateprofilesItemAssignmentsRequestBuilder) {
    return NewWindowsfeatureupdateprofilesItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderInternal instantiates a new WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) {
    m := &WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsFeatureUpdateProfiles/{windowsFeatureUpdateProfile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder instantiates a new WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder and sets the default values.
func NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsFeatureUpdateProfiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of windows feature update profiles
// returns a WindowsFeatureUpdateProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsFeatureUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable), nil
}
// Patch update the navigation property windowsFeatureUpdateProfiles in deviceManagement
// returns a WindowsFeatureUpdateProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsFeatureUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable), nil
}
// ToDeleteRequestInformation delete navigation property windowsFeatureUpdateProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of windows feature update profiles
// returns a *RequestInformation when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsFeatureUpdateProfiles in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsFeatureUpdateProfileable, requestConfiguration *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder when successful
func (m *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) WithUrl(rawUrl string)(*WindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder) {
    return NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
