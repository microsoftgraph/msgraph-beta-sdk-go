package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/operations"
    i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/rundiagnostics"
    i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/configuration"
    i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/updatesoftware"
    id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/health"
    idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/activity"
    iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/restart"
    i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/devices/item/operations/item"
)

// TeamworkDeviceRequestBuilder builds and executes requests for operations under \teamwork\devices\{teamworkDevice-id}
type TeamworkDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamworkDeviceRequestBuilderDeleteOptions options for Delete
type TeamworkDeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamworkDeviceRequestBuilderGetOptions options for Get
type TeamworkDeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamworkDeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamworkDeviceRequestBuilderGetQueryParameters the Teams devices provisioned for the tenant.
type TeamworkDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamworkDeviceRequestBuilderPatchOptions options for Patch
type TeamworkDeviceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkDevice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TeamworkDeviceRequestBuilder) Activity()(*idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b.ActivityRequestBuilder) {
    return idcb97c68de16b48e5f40ae7c33010c8ab78995f83aafa559ed280e2cd6db993b.NewActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamworkDeviceRequestBuilder) Configuration()(*i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2.ConfigurationRequestBuilder) {
    return i5991b16c8216191b7b79a2649c2cbd2785a5237f266a4b1a2b3caaf26d95c5d2.NewConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamworkDeviceRequestBuilderInternal instantiates a new TeamworkDeviceRequestBuilder and sets the default values.
func NewTeamworkDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkDeviceRequestBuilder) {
    m := &TeamworkDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/devices/{teamworkDevice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkDeviceRequestBuilder instantiates a new TeamworkDeviceRequestBuilder and sets the default values.
func NewTeamworkDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamworkDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) CreateDeleteRequestInformation(options *TeamworkDeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) CreateGetRequestInformation(options *TeamworkDeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) CreatePatchRequestInformation(options *TeamworkDeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) Delete(options *TeamworkDeviceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) Get(options *TeamworkDeviceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkDevice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTeamworkDevice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkDevice), nil
}
func (m *TeamworkDeviceRequestBuilder) Health()(*id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85.HealthRequestBuilder) {
    return id65a906da2d62457b975164936dd205db649d8f0fcf34067fb93cdc441739c85.NewHealthRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamworkDeviceRequestBuilder) Operations()(*i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a.OperationsRequestBuilder) {
    return i370f825d40713fc0eb3f34cb88138713e23e2630f08716d8e4f4315992444b1a.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.devices.item.operations.item collection
func (m *TeamworkDeviceRequestBuilder) OperationsById(id string)(*i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d.TeamworkDeviceOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkDeviceOperation_id"] = id
    }
    return i9322b599ad80a79ed0450750e6a1607ac4158d6981a51a70253206f7041ccd6d.NewTeamworkDeviceOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the Teams devices provisioned for the tenant.
func (m *TeamworkDeviceRequestBuilder) Patch(options *TeamworkDeviceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamworkDeviceRequestBuilder) Restart()(*iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c.RestartRequestBuilder) {
    return iea334c08ce97640ab2365649a261444b2335f9441fceb47ed8150a28e7a66e0c.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamworkDeviceRequestBuilder) RunDiagnostics()(*i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60.RunDiagnosticsRequestBuilder) {
    return i45ce7cab530b471e7b705e982738e8c6ce9013eadd24b6a9491480dd05821d60.NewRunDiagnosticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamworkDeviceRequestBuilder) UpdateSoftware()(*i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4.UpdateSoftwareRequestBuilder) {
    return i885ba2537ef1730ed6c80f571dac9a76afac394788b86766b383cd7ed03da0c4.NewUpdateSoftwareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
