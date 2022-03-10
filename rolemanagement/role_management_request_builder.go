package rolemanagement

import (
    i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement"
    i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement"
    i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/directory"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// RoleManagementRequestBuilder provides operations to manage the roleManagement singleton.
type RoleManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleManagementRequestBuilderGetOptions options for Get
type RoleManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleManagementRequestBuilderGetQueryParameters get roleManagement
type RoleManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RoleManagementRequestBuilderPatchOptions options for Patch
type RoleManagementRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagementable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RoleManagementRequestBuilder) CloudPC()(*ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.CloudPCRequestBuilder) {
    return ie179728a78650b7ea21a75747c7c0c71a4b3eeffeb7edcfa3f337c75f028ffa7.NewCloudPCRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRoleManagementRequestBuilderInternal instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleManagementRequestBuilder) {
    m := &RoleManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementRequestBuilder instantiates a new RoleManagementRequestBuilder and sets the default values.
func NewRoleManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get roleManagement
func (m *RoleManagementRequestBuilder) CreateGetRequestInformation(options *RoleManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update roleManagement
func (m *RoleManagementRequestBuilder) CreatePatchRequestInformation(options *RoleManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *RoleManagementRequestBuilder) DeviceManagement()(*i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.DeviceManagementRequestBuilder) {
    return i45ee08ad7f18888228aa301bf50b8ba0f15edd802a1aab237f60b36937784346.NewDeviceManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleManagementRequestBuilder) Directory()(*i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.DirectoryRequestBuilder) {
    return i8a93901c4a3aed17badfa3c9c54168d3e7506fa6cc5ed2d60a7be06761c00759.NewDirectoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RoleManagementRequestBuilder) EntitlementManagement()(*i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.EntitlementManagementRequestBuilder) {
    return i10372b71511cc47bf754e90b050fc97fee639d10b03f642f5ac8321f1d6eee9e.NewEntitlementManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleManagement
func (m *RoleManagementRequestBuilder) Get(options *RoleManagementRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateRoleManagementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RoleManagementable), nil
}
// Patch update roleManagement
func (m *RoleManagementRequestBuilder) Patch(options *RoleManagementRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
