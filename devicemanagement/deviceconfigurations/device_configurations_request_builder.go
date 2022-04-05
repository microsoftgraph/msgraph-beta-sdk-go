package deviceconfigurations

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/gettargetedusersanddevices"
    i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/getiosavailableupdateversions"
    id4f8b95e9a4f408fff016b0499932369a6013687b3c089be2d425060978e03b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/count"
    ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/haspayloadlinks"
)

// DeviceConfigurationsRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationsRequestBuilderGetOptions options for Get
type DeviceConfigurationsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceConfigurationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceConfigurationsRequestBuilderGetQueryParameters the device configurations.
type DeviceConfigurationsRequestBuilderGetQueryParameters struct {
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
// DeviceConfigurationsRequestBuilderPostOptions options for Post
type DeviceConfigurationsRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeviceConfigurationsRequestBuilderInternal instantiates a new DeviceConfigurationsRequestBuilder and sets the default values.
func NewDeviceConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsRequestBuilder) {
    m := &DeviceConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationsRequestBuilder instantiates a new DeviceConfigurationsRequestBuilder and sets the default values.
func NewDeviceConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeviceConfigurationsRequestBuilder) Count()(*id4f8b95e9a4f408fff016b0499932369a6013687b3c089be2d425060978e03b0.CountRequestBuilder) {
    return id4f8b95e9a4f408fff016b0499932369a6013687b3c089be2d425060978e03b0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the device configurations.
func (m *DeviceConfigurationsRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deviceConfigurations for deviceManagement
func (m *DeviceConfigurationsRequestBuilder) CreatePostRequestInformation(options *DeviceConfigurationsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the device configurations.
func (m *DeviceConfigurationsRequestBuilder) Get(options *DeviceConfigurationsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationCollectionResponseable), nil
}
// GetIosAvailableUpdateVersions provides operations to call the getIosAvailableUpdateVersions method.
func (m *DeviceConfigurationsRequestBuilder) GetIosAvailableUpdateVersions()(*i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a.GetIosAvailableUpdateVersionsRequestBuilder) {
    return i91812d422ff32b6087dbdb578859075772e6e50438228306a697e5aa7c00dc0a.NewGetIosAvailableUpdateVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetTargetedUsersAndDevices the getTargetedUsersAndDevices property
func (m *DeviceConfigurationsRequestBuilder) GetTargetedUsersAndDevices()(*i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a.GetTargetedUsersAndDevicesRequestBuilder) {
    return i7c8fc03ac714b8a0b8442431c0a15160448a9c352c4c71e08219dd57f3689a4a.NewGetTargetedUsersAndDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HasPayloadLinks the hasPayloadLinks property
func (m *DeviceConfigurationsRequestBuilder) HasPayloadLinks()(*ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f.HasPayloadLinksRequestBuilder) {
    return ie1dbe7768034fc9acbabff177deec668725e3e1b963712ee72a4c861c0c6347f.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deviceConfigurations for deviceManagement
func (m *DeviceConfigurationsRequestBuilder) Post(options *DeviceConfigurationsRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
