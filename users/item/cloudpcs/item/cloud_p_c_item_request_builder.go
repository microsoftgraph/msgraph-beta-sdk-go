package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/troubleshoot"
    i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/changeuseraccounttype"
    i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/reprovision"
    iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/endgraceperiod"
    iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/rename"
    ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/reboot"
    if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/getcloudpclaunchinfo"
)

// CloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type CloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPCItemRequestBuilderDeleteOptions options for Delete
type CloudPCItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCItemRequestBuilderGetOptions options for Get
type CloudPCItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPCItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCItemRequestBuilderGetQueryParameters get cloudPCs from users
type CloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPCItemRequestBuilderPatchOptions options for Patch
type CloudPCItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CloudPCItemRequestBuilder) ChangeUserAccountType()(*i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f.ChangeUserAccountTypeRequestBuilder) {
    return i97a56d161f68b1fd3ad12ad108342214d2bfeedae4b18c9748dc91bc1e04df9f.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCItemRequestBuilder) {
    m := &CloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/cloudPCs/{cloudPC_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for users
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPCItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get cloudPCs from users
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformation(options *CloudPCItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPCs in users
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformation(options *CloudPCItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property cloudPCs for users
func (m *CloudPCItemRequestBuilder) Delete(options *CloudPCItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
func (m *CloudPCItemRequestBuilder) EndGracePeriod()(*iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b.EndGracePeriodRequestBuilder) {
    return iabd9a09efde7cff52110da1822c7742e59aeb1e5915ccfe038d929a6b7ad268b.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from users
func (m *CloudPCItemRequestBuilder) Get(options *CloudPCItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCloudPCFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable), nil
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0.GetCloudPcLaunchInfoRequestBuilder) {
    return if4dc1d7e15dcd2713c85cab00a5a53d94f288c3f116521764c65b1ad462b0eb0.NewGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in users
func (m *CloudPCItemRequestBuilder) Patch(options *CloudPCItemRequestBuilderPatchOptions)(error) {
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
func (m *CloudPCItemRequestBuilder) Reboot()(*ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e.RebootRequestBuilder) {
    return ief59a775175af2e50c2ba2215af92a5bdc70b7dceef0de8704dccb1033d7ed7e.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Rename()(*iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022.RenameRequestBuilder) {
    return iecadcd2b05f98b70514e4bc46dddb17b2c7871f504fd101f5951b02df655c022.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Reprovision()(*i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0.ReprovisionRequestBuilder) {
    return i97b00529155a730b3ecc9a3f5590a17060250dc44e1d660aed1789b4345f43c0.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Troubleshoot()(*i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e.TroubleshootRequestBuilder) {
    return i07d40583df1b4aa1b0a3e7e32084dc8b4ab30be70327b272f2ad24d3ccd4095e.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
