package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2abdeba6a2120c9e857becdb50bd4401198f13618af67e4a5a7bcffdd11f7525 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/item/getsupportedproperties"
)

// DeviceAndAppManagementAssignmentFilterRequestBuilder builds and executes requests for operations under \deviceManagement\assignmentFilters\{deviceAndAppManagementAssignmentFilter-id}
type DeviceAndAppManagementAssignmentFilterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceAndAppManagementAssignmentFilterRequestBuilderDeleteOptions options for Delete
type DeviceAndAppManagementAssignmentFilterRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceAndAppManagementAssignmentFilterRequestBuilderGetOptions options for Get
type DeviceAndAppManagementAssignmentFilterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceAndAppManagementAssignmentFilterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceAndAppManagementAssignmentFilterRequestBuilderGetQueryParameters the list of assignment filters
type DeviceAndAppManagementAssignmentFilterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceAndAppManagementAssignmentFilterRequestBuilderPatchOptions options for Patch
type DeviceAndAppManagementAssignmentFilterRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilter;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceAndAppManagementAssignmentFilterRequestBuilderInternal instantiates a new DeviceAndAppManagementAssignmentFilterRequestBuilder and sets the default values.
func NewDeviceAndAppManagementAssignmentFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceAndAppManagementAssignmentFilterRequestBuilder) {
    m := &DeviceAndAppManagementAssignmentFilterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/assignmentFilters/{deviceAndAppManagementAssignmentFilter_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAndAppManagementAssignmentFilterRequestBuilder instantiates a new DeviceAndAppManagementAssignmentFilterRequestBuilder and sets the default values.
func NewDeviceAndAppManagementAssignmentFilterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceAndAppManagementAssignmentFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAndAppManagementAssignmentFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) CreateDeleteRequestInformation(options *DeviceAndAppManagementAssignmentFilterRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) CreateGetRequestInformation(options *DeviceAndAppManagementAssignmentFilterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) CreatePatchRequestInformation(options *DeviceAndAppManagementAssignmentFilterRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) Delete(options *DeviceAndAppManagementAssignmentFilterRequestBuilderDeleteOptions)(error) {
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
// Get the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) Get(options *DeviceAndAppManagementAssignmentFilterRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilter, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceAndAppManagementAssignmentFilter() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilter), nil
}
// GetSupportedProperties builds and executes requests for operations under \deviceManagement\assignmentFilters\{deviceAndAppManagementAssignmentFilter-id}\microsoft.graph.getSupportedProperties()
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) GetSupportedProperties()(*i2abdeba6a2120c9e857becdb50bd4401198f13618af67e4a5a7bcffdd11f7525.GetSupportedPropertiesRequestBuilder) {
    return i2abdeba6a2120c9e857becdb50bd4401198f13618af67e4a5a7bcffdd11f7525.NewGetSupportedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the list of assignment filters
func (m *DeviceAndAppManagementAssignmentFilterRequestBuilder) Patch(options *DeviceAndAppManagementAssignmentFilterRequestBuilderPatchOptions)(error) {
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
