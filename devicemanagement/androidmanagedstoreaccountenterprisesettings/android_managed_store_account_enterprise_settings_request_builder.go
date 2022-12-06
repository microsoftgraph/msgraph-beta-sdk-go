package androidmanagedstoreaccountenterprisesettings

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/creategoogleplaywebtoken"
    i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/completesignup"
    i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/syncapps"
    i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/unbind"
    i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/setandroiddeviceownerfullymanagedenrollmentstate"
    ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/requestsignupurl"
    ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidmanagedstoreaccountenterprisesettings/approveapps"
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
    Headers map[string]string
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetQueryParameters
}
// AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApproveApps provides operations to call the approveApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.ApproveAppsRequestBuilder) {
    return ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.NewApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteSignup provides operations to call the completeSignup method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CompleteSignup()(*i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.CompleteSignupRequestBuilder) {
    return i314c76ed635b4004a342dd0c841e00e8f471d83c6cb71a6344ebcca4508e202a.NewCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder instantiates a new AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.CreateGooglePlayWebTokenRequestBuilder) {
    return i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.NewCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.RequestSignupUrlRequestBuilder) {
    return ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.NewRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentState provides operations to call the setAndroidDeviceOwnerFullyManagedEnrollmentState method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncApps provides operations to call the syncApps method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.SyncAppsRequestBuilder) {
    return i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.NewSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unbind provides operations to call the unbind method.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.UnbindRequestBuilder) {
    return i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.NewUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
