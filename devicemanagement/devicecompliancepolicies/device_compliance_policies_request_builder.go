package devicecompliancepolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/validatecompliancescript"
    i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/setscheduledretirestate"
    i66a18fc3065359dba5afba474db05606a4e59242a39b023a6c78e234799dcef5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/count"
    i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/haspayloadlinks"
    ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/refreshdevicecompliancereportsummarization"
    ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/getdevicesscheduledtoretire"
)

// DeviceCompliancePoliciesRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceCompliancePoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCompliancePoliciesRequestBuilderGetOptions options for Get
type DeviceCompliancePoliciesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceCompliancePoliciesRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// DeviceCompliancePoliciesRequestBuilderPostOptions options for Post
type DeviceCompliancePoliciesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeviceCompliancePoliciesRequestBuilderInternal instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesRequestBuilder) {
    m := &DeviceCompliancePoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCompliancePoliciesRequestBuilder instantiates a new DeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeviceCompliancePoliciesRequestBuilder) Count()(*i66a18fc3065359dba5afba474db05606a4e59242a39b023a6c78e234799dcef5.CountRequestBuilder) {
    return i66a18fc3065359dba5afba474db05606a4e59242a39b023a6c78e234799dcef5.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) CreateGetRequestInformation(options *DeviceCompliancePoliciesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesRequestBuilder) CreatePostRequestInformation(options *DeviceCompliancePoliciesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Get the device compliance policies.
func (m *DeviceCompliancePoliciesRequestBuilder) Get(options *DeviceCompliancePoliciesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable), nil
}
// GetDevicesScheduledToRetire the getDevicesScheduledToRetire property
func (m *DeviceCompliancePoliciesRequestBuilder) GetDevicesScheduledToRetire()(*ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5.GetDevicesScheduledToRetireRequestBuilder) {
    return ice8333de4faa45aa6e283d6de5cec5fb2d91ddb4641c608a190a6c8ebda92bc5.NewGetDevicesScheduledToRetireRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasPayloadLinks the hasPayloadLinks property
func (m *DeviceCompliancePoliciesRequestBuilder) HasPayloadLinks()(*i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a.HasPayloadLinksRequestBuilder) {
    return i8c922b0106272c33024960c0cdcce96f4a55b9e9d3d906cb19b2281dbb71993a.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesRequestBuilder) Post(options *DeviceCompliancePoliciesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// RefreshDeviceComplianceReportSummarization the refreshDeviceComplianceReportSummarization property
func (m *DeviceCompliancePoliciesRequestBuilder) RefreshDeviceComplianceReportSummarization()(*ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51.RefreshDeviceComplianceReportSummarizationRequestBuilder) {
    return ib90740a26262712164b9160ce4822a569344cfc69adafea017f8563b8e220b51.NewRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetScheduledRetireState the setScheduledRetireState property
func (m *DeviceCompliancePoliciesRequestBuilder) SetScheduledRetireState()(*i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc.SetScheduledRetireStateRequestBuilder) {
    return i5f664252c457e99b8d5db6b834754f966442a3164d883ca219d4dbf2362b3cdc.NewSetScheduledRetireStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateComplianceScript the validateComplianceScript property
func (m *DeviceCompliancePoliciesRequestBuilder) ValidateComplianceScript()(*i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093.ValidateComplianceScriptRequestBuilder) {
    return i58651a263a770247eef91ec5783125ddffea19f0485aeac13f756043009a7093.NewValidateComplianceScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
