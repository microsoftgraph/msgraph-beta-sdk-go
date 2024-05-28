package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidforworksettingsAndroidForWorkSettingsRequestBuilder provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
type AndroidforworksettingsAndroidForWorkSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidforworksettingsAndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworksettingsAndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetQueryParameters the singleton Android for Work settings entity.
type AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetQueryParameters
}
// AndroidforworksettingsAndroidForWorkSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworksettingsAndroidForWorkSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteSignup provides operations to call the completeSignup method.
// returns a *AndroidforworksettingsCompletesignupCompleteSignupRequestBuilder when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) CompleteSignup()(*AndroidforworksettingsCompletesignupCompleteSignupRequestBuilder) {
    return NewAndroidforworksettingsCompletesignupCompleteSignupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilderInternal instantiates a new AndroidforworksettingsAndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) {
    m := &AndroidforworksettingsAndroidForWorkSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilder instantiates a new AndroidforworksettingsAndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property androidForWorkSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the singleton Android for Work settings entity.
// returns a AndroidForWorkSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable), nil
}
// Patch update the navigation property androidForWorkSettings in deviceManagement
// returns a AndroidForWorkSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable), nil
}
// RequestSignupUrl provides operations to call the requestSignupUrl method.
// returns a *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) RequestSignupUrl()(*AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    return NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncApps provides operations to call the syncApps method.
// returns a *AndroidforworksettingsSyncappsSyncAppsRequestBuilder when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) SyncApps()(*AndroidforworksettingsSyncappsSyncAppsRequestBuilder) {
    return NewAndroidforworksettingsSyncappsSyncAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidForWorkSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the singleton Android for Work settings entity.
// returns a *RequestInformation when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property androidForWorkSettings in deviceManagement
// returns a *RequestInformation when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, requestConfiguration *AndroidforworksettingsAndroidForWorkSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unbind provides operations to call the unbind method.
// returns a *AndroidforworksettingsUnbindRequestBuilder when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) Unbind()(*AndroidforworksettingsUnbindRequestBuilder) {
    return NewAndroidforworksettingsUnbindRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder when successful
func (m *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) WithUrl(rawUrl string)(*AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) {
    return NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
