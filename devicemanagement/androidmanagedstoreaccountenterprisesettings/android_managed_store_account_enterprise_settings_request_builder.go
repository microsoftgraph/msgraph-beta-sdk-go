package androidmanagedstoreaccountenterprisesettings

import (
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
// ApproveApps the approveApps property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) ApproveApps()(*ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.ApproveAppsRequestBuilder) {
    return ia5c047b600e80aaec8a495894dd576800c93609eea95fb0910bcad732a992e54.NewApproveAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteSignup the completeSignup property
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGooglePlayWebToken the createGooglePlayWebToken property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreateGooglePlayWebToken()(*i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.CreateGooglePlayWebTokenRequestBuilder) {
    return i0f4c5b6e577471df0b671fb66a33d7c548c770695d5d464721827e7069ad5d5c.NewCreateGooglePlayWebTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property androidManagedStoreAccountEnterpriseSettings for deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) DeleteWithResponseHandler(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) GetWithResponseHandler(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the singleton Android managed store account enterprise settings entity.
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) GetWithResponseHandler(requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable), nil
}
// PatchWithResponseHandler update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property androidManagedStoreAccountEnterpriseSettings in deviceManagement
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAccountEnterpriseSettingsable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RequestSignupUrl the requestSignupUrl property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) RequestSignupUrl()(*ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.RequestSignupUrlRequestBuilder) {
    return ia312e54dde3f35887f50da33001441eb2d48646e6fd640a4c7387523b4934a7e.NewRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetAndroidDeviceOwnerFullyManagedEnrollmentState the setAndroidDeviceOwnerFullyManagedEnrollmentState property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SetAndroidDeviceOwnerFullyManagedEnrollmentState()(*i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilder) {
    return i9821f431564656dfb670bc3758898c81b836fe062cdacc51bf1462209f22c694.NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncApps the syncApps property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) SyncApps()(*i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.SyncAppsRequestBuilder) {
    return i5570da22a9e8ab3a3faa339899d50bf58d9d37e83641e9d383b921ee28101901.NewSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unbind the unbind property
func (m *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) Unbind()(*i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.UnbindRequestBuilder) {
    return i573858361497d031356ed7937a66002f0e820977a38c57afb805332721bfbd14.NewUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
