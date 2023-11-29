package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters the singleton Android managed store account enterprise settings entity.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters
}
// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddApps provides operations to call the addApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) AddApps()(*AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApproveApps provides operations to call the approveApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*AndroidManagedStoreAccountEnterpriseSettingsApproveAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsApproveAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteSignup provides operations to call the completeSignup method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*AndroidManagedStoreAccountEnterpriseSettingsCompleteSignupRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGooglePlayWebToken provides operations to call the createGooglePlayWebToken method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// Patch update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// RequestSignupUrl provides operations to call the requestSignupUrl method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*AndroidManagedStoreAccountEnterpriseSettingsRequestSignupUrlRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestSignupUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentState provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*AndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SyncApps provides operations to call the syncApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*AndroidManagedStoreAccountEnterpriseSettingsUnbindRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsUnbindRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) WithUrl(rawUrl string)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
