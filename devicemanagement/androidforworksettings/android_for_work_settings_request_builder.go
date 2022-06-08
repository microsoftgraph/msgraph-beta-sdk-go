package androidforworksettings

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/completesignup"
    i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/requestsignupurl"
    i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/syncapps"
    ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androidforworksettings/unbind"
)

// AndroidForWorkSettingsRequestBuilder provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
type AndroidForWorkSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidForWorkSettingsRequestBuilderGetQueryParameters the singleton Android for Work settings entity.
type AndroidForWorkSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidForWorkSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidForWorkSettingsRequestBuilderGetQueryParameters
}
// AndroidForWorkSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteSignup the completeSignup property
func (m *AndroidForWorkSettingsRequestBuilder) CompleteSignup()(*i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359.CompleteSignupRequestBuilder) {
    return i280e19d80d7a529d2b9933cf3044f19490e2f99605f9859eeb0710b5d4c6b359.NewCompleteSignupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAndroidForWorkSettingsRequestBuilderInternal instantiates a new AndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsRequestBuilder) {
    m := &AndroidForWorkSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidForWorkSettings{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidForWorkSettingsRequestBuilder instantiates a new AndroidForWorkSettingsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property androidForWorkSettings for deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property androidForWorkSettings for deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AndroidForWorkSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property androidForWorkSettings in deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property androidForWorkSettings in deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, requestConfiguration *AndroidForWorkSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property androidForWorkSettings for deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property androidForWorkSettings for deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *AndroidForWorkSettingsRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the singleton Android for Work settings entity.
func (m *AndroidForWorkSettingsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AndroidForWorkSettingsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkSettingsFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable), nil
}
// Patch update the navigation property androidForWorkSettings in deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property androidForWorkSettings in deviceManagement
func (m *AndroidForWorkSettingsRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkSettingsable, requestConfiguration *AndroidForWorkSettingsRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *AndroidForWorkSettingsRequestBuilder) RequestSignupUrl()(*i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff.RequestSignupUrlRequestBuilder) {
    return i61c16d732624a212d9246ecf43fd01db8d4bc43debd10e8907b417795c13f8ff.NewRequestSignupUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncApps the syncApps property
func (m *AndroidForWorkSettingsRequestBuilder) SyncApps()(*i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817.SyncAppsRequestBuilder) {
    return i97a07e45ebea436c88e03068b5e337121e7e6f6d82dd6f519d86bd92618a0817.NewSyncAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unbind the unbind property
func (m *AndroidForWorkSettingsRequestBuilder) Unbind()(*ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8.UnbindRequestBuilder) {
    return ibc8cdcda9d0880d123b6a234fd0980335960dc7e2f066cbce430ed5b5bd54ef8.NewUnbindRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
