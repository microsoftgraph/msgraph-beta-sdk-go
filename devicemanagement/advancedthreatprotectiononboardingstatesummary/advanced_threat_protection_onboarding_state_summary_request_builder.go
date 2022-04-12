package advancedthreatprotectiononboardingstatesummary

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates"
    i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates/item"
)

// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions options for Delete
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions options for Get
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters the summary state of ATP onboarding state for this account.
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions options for Patch
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AdvancedThreatProtectionOnboardingDeviceSettingStates the advancedThreatProtectionOnboardingDeviceSettingStates property
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates()(*ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdvancedThreatProtectionOnboardingDeviceSettingStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.advancedThreatProtectionOnboardingStateSummary.advancedThreatProtectionOnboardingDeviceSettingStates.item collection
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStatesById(id string)(*i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.AdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["advancedThreatProtectionOnboardingDeviceSettingState%2Did"] = id
    }
    return i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.NewAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    m := &AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateDeleteRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the summary state of ATP onboarding state for this account.
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateGetRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreatePatchRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Delete(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the summary state of ATP onboarding state for this account.
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Get(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Patch(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
