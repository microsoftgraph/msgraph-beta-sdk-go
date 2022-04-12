package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/operations"
    i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/rundiagnostics"
    i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/configuration"
    i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/updatesoftware"
    id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/health"
    idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/activity"
    iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/restart"
    i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/operations/item"
)

// TeamworkDeviceItemRequestBuilder provides operations to manage the devices property of the microsoft.graph.teamwork entity.
type TeamworkDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkDeviceItemRequestBuilderDeleteOptions options for Delete
type TeamworkDeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TeamworkDeviceItemRequestBuilderGetOptions options for Get
type TeamworkDeviceItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkDeviceItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// TeamworkDeviceItemRequestBuilderGetQueryParameters the Teams devices provisioned for the tenant.
type TeamworkDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkDeviceItemRequestBuilderPatchOptions options for Patch
type TeamworkDeviceItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkDeviceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Activity the activity property
func (m *TeamworkDeviceItemRequestBuilder) Activity()(*idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b.ActivityRequestBuilder) {
    return idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b.NewActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Configuration the configuration property
func (m *TeamworkDeviceItemRequestBuilder) Configuration()(*i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2.ConfigurationRequestBuilder) {
    return i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2.NewConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamworkDeviceItemRequestBuilderInternal instantiates a new TeamworkDeviceItemRequestBuilder and sets the default values.
func NewTeamworkDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeviceItemRequestBuilder) {
    m := &TeamworkDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/devices/{teamworkDevice%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkDeviceItemRequestBuilder instantiates a new TeamworkDeviceItemRequestBuilder and sets the default values.
func NewTeamworkDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property devices for teamwork
func (m *TeamworkDeviceItemRequestBuilder) CreateDeleteRequestInformation(options *TeamworkDeviceItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceItemRequestBuilder) CreateGetRequestInformation(options *TeamworkDeviceItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property devices in teamwork
func (m *TeamworkDeviceItemRequestBuilder) CreatePatchRequestInformation(options *TeamworkDeviceItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property devices for teamwork
func (m *TeamworkDeviceItemRequestBuilder) Delete(options *TeamworkDeviceItemRequestBuilderDeleteOptions)(error) {
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
// Get the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceItemRequestBuilder) Get(options *TeamworkDeviceItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTeamworkDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TeamworkDeviceable), nil
}
// Health the health property
func (m *TeamworkDeviceItemRequestBuilder) Health()(*id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85.HealthRequestBuilder) {
    return id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85.NewHealthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *TeamworkDeviceItemRequestBuilder) Operations()(*i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a.OperationsRequestBuilder) {
    return i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.devices.item.operations.item collection
func (m *TeamworkDeviceItemRequestBuilder) OperationsById(id string)(*i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d.TeamworkDeviceOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkDeviceOperation%2Did"] = id
    }
    return i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d.NewTeamworkDeviceOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property devices in teamwork
func (m *TeamworkDeviceItemRequestBuilder) Patch(options *TeamworkDeviceItemRequestBuilderPatchOptions)(error) {
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
// Restart the restart property
func (m *TeamworkDeviceItemRequestBuilder) Restart()(*iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c.RestartRequestBuilder) {
    return iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RunDiagnostics the runDiagnostics property
func (m *TeamworkDeviceItemRequestBuilder) RunDiagnostics()(*i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60.RunDiagnosticsRequestBuilder) {
    return i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60.NewRunDiagnosticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateSoftware the updateSoftware property
func (m *TeamworkDeviceItemRequestBuilder) UpdateSoftware()(*i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4.UpdateSoftwareRequestBuilder) {
    return i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4.NewUpdateSoftwareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
