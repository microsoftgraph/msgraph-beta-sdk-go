package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assign"
    i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assignments"
    if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assignments/item"
)

// CloudPcProvisioningPolicyRequestBuilder builds and executes requests for operations under \deviceManagement\virtualEndpoint\provisioningPolicies\{cloudPcProvisioningPolicy-id}
type CloudPcProvisioningPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPcProvisioningPolicyRequestBuilderDeleteOptions options for Delete
type CloudPcProvisioningPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcProvisioningPolicyRequestBuilderGetOptions options for Get
type CloudPcProvisioningPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPcProvisioningPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcProvisioningPolicyRequestBuilderGetQueryParameters cloud PC provisioning policy.
type CloudPcProvisioningPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPcProvisioningPolicyRequestBuilderPatchOptions options for Patch
type CloudPcProvisioningPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CloudPcProvisioningPolicyRequestBuilder) Assign()(*i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104.AssignRequestBuilder) {
    return i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPcProvisioningPolicyRequestBuilder) Assignments()(*i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c.AssignmentsRequestBuilder) {
    return i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.provisioningPolicies.item.assignments.item collection
func (m *CloudPcProvisioningPolicyRequestBuilder) AssignmentsById(id string)(*if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848.CloudPcProvisioningPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicyAssignment_id"] = id
    }
    return if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848.NewCloudPcProvisioningPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCloudPcProvisioningPolicyRequestBuilderInternal instantiates a new CloudPcProvisioningPolicyRequestBuilder and sets the default values.
func NewCloudPcProvisioningPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcProvisioningPolicyRequestBuilder) {
    m := &CloudPcProvisioningPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPcProvisioningPolicyRequestBuilder instantiates a new CloudPcProvisioningPolicyRequestBuilder and sets the default values.
func NewCloudPcProvisioningPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcProvisioningPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPcProvisioningPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) CreateDeleteRequestInformation(options *CloudPcProvisioningPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) CreateGetRequestInformation(options *CloudPcProvisioningPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) CreatePatchRequestInformation(options *CloudPcProvisioningPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) Delete(options *CloudPcProvisioningPolicyRequestBuilderDeleteOptions)(error) {
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
// Get cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) Get(options *CloudPcProvisioningPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCloudPcProvisioningPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicy), nil
}
// Patch cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyRequestBuilder) Patch(options *CloudPcProvisioningPolicyRequestBuilderPatchOptions)(error) {
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
