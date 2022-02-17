package devicemanagement

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces"
    ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roledefinitions"
    i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item"
    ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item"
    id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roledefinitions/item"
)

// DeviceManagementRequestBuilder builds and executes requests for operations under \roleManagement\deviceManagement
type DeviceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementRequestBuilderDeleteOptions options for Delete
type DeviceManagementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementRequestBuilderGetOptions options for Get
type DeviceManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementRequestBuilderGetQueryParameters the RbacApplication for Device Management
type DeviceManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementRequestBuilderPatchOptions options for Patch
type DeviceManagementRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) CreateGetRequestInformation(options *DeviceManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) Delete(options *DeviceManagementRequestBuilderDeleteOptions)(error) {
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
// Get the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) Get(options *DeviceManagementRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRbacApplicationMultiple() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RbacApplicationMultiple), nil
}
// Patch the RbacApplication for Device Management
func (m *DeviceManagementRequestBuilder) Patch(options *DeviceManagementRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementRequestBuilder) ResourceNamespaces()(*i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163.ResourceNamespacesRequestBuilder) {
    return i74d924764342fa6022f3fd5e5f9706420912c7cdd1a2bb0b50c8bf67c140c163.NewResourceNamespacesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceNamespacesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.resourceNamespaces.item collection
func (m *DeviceManagementRequestBuilder) ResourceNamespacesById(id string)(*i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15.UnifiedRbacResourceNamespaceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceNamespace_id"] = id
    }
    return i2f6986ad54e9019a27955e607e7a5e6a3900d7d2a186631d40a82c7420f91a15.NewUnifiedRbacResourceNamespaceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9.RoleAssignmentsRequestBuilder) {
    return i16ba9d37411ef84c105365cabee4cb8ed289e66e837ce2bc6d3458b7cdcfb9e9.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleAssignments.item collection
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265.UnifiedRoleAssignmentMultipleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleAssignmentMultiple_id"] = id
    }
    return ibf3ec838b277e1140e5e57824ced75b7c6eed602c42fa7b33e6783984c251265.NewUnifiedRoleAssignmentMultipleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e.RoleDefinitionsRequestBuilder) {
    return ic7f2d26945adadc926e1a8a84b9142418a67273d0f9ba3ad2fda23bc4c71665e.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.roleDefinitions.item collection
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c.UnifiedRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleDefinition_id"] = id
    }
    return id1e540118321507d6c775d35c43257337ecb59533ab19347820eefca0118bb2c.NewUnifiedRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
