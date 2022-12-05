package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters the singleton Android managed store account enterprise settings entity.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApproveApps provides operations to call the approveApps method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsApproveAppsRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteSignup provides operations to call the completeSignup method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    m := &DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the singleton Android managed store account enterprise settings entity.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGooglePlayWebToken provides operations to call the createGooglePlayWebToken method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the singleton Android managed store account enterprise settings entity.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// Patch update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// RequestSignupUrl provides operations to call the requestSignupUrl method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestSignupUrlRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentState provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncApps provides operations to call the syncApps method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unbind provides operations to call the unbind method.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsUnbindRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
