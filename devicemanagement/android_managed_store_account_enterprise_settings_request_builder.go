package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
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
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
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
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// MicrosoftGraphAddApps provides operations to call the addApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphAddApps()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphApproveApps provides operations to call the approveApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphApproveApps()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphApproveAppsApproveAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphApproveAppsApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCompleteSignup provides operations to call the completeSignup method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphCompleteSignup()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphCreateGooglePlayWebToken provides operations to call the createGooglePlayWebToken method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphCreateGooglePlayWebToken()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRequestSignupUrl provides operations to call the requestSignupUrl method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphRequestSignupUrl()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetAndroidDeviceOwnerFullyManagedEnrollmentState provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphSetAndroidDeviceOwnerFullyManagedEnrollmentState()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphSetAndroidDeviceOwnerFullyManagedEnrollmentStateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphSetAndroidDeviceOwnerFullyManagedEnrollmentStateSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSyncApps provides operations to call the syncApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphSyncApps()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphSyncAppsSyncAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphSyncAppsSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUnbind provides operations to call the unbind method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) MicrosoftGraphUnbind()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphUnbindUnbindRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphUnbindUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter)
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
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
