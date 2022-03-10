package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assign"
    i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assignments"
    if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item/assignments/item"
)

// CloudPcProvisioningPolicyItemRequestBuilder provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
type CloudPcProvisioningPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPcProvisioningPolicyItemRequestBuilderDeleteOptions options for Delete
type CloudPcProvisioningPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcProvisioningPolicyItemRequestBuilderGetOptions options for Get
type CloudPcProvisioningPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters cloud PC provisioning policy.
type CloudPcProvisioningPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPcProvisioningPolicyItemRequestBuilderPatchOptions options for Patch
type CloudPcProvisioningPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CloudPcProvisioningPolicyItemRequestBuilder) Assign()(*i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104.AssignRequestBuilder) {
    return i21d65c3429cff028859a2c620580cf0b27cf65a897ffa7c6f9bd17a4506e4104.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPcProvisioningPolicyItemRequestBuilder) Assignments()(*i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c.AssignmentsRequestBuilder) {
    return i8efaf7f9224f8bf3ba72fe2c8e64b5a0019684376d8d47f60bfc6702485abc3c.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.provisioningPolicies.item.assignments.item collection
func (m *CloudPcProvisioningPolicyItemRequestBuilder) AssignmentsById(id string)(*if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848.CloudPcProvisioningPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicyAssignment_id"] = id
    }
    return if91535156d5443297583fcd65d99597809e9e5b6115e4948cfc32c6c18b25848.NewCloudPcProvisioningPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCloudPcProvisioningPolicyItemRequestBuilderInternal instantiates a new CloudPcProvisioningPolicyItemRequestBuilder and sets the default values.
func NewCloudPcProvisioningPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcProvisioningPolicyItemRequestBuilder) {
    m := &CloudPcProvisioningPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPcProvisioningPolicyItemRequestBuilder instantiates a new CloudPcProvisioningPolicyItemRequestBuilder and sets the default values.
func NewCloudPcProvisioningPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcProvisioningPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPcProvisioningPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property provisioningPolicies for deviceManagement
func (m *CloudPcProvisioningPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPcProvisioningPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CloudPcProvisioningPolicyItemRequestBuilder) CreateGetRequestInformation(options *CloudPcProvisioningPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property provisioningPolicies in deviceManagement
func (m *CloudPcProvisioningPolicyItemRequestBuilder) CreatePatchRequestInformation(options *CloudPcProvisioningPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property provisioningPolicies for deviceManagement
func (m *CloudPcProvisioningPolicyItemRequestBuilder) Delete(options *CloudPcProvisioningPolicyItemRequestBuilderDeleteOptions)(error) {
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
// Get cloud PC provisioning policy.
func (m *CloudPcProvisioningPolicyItemRequestBuilder) Get(options *CloudPcProvisioningPolicyItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCloudPcProvisioningPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcProvisioningPolicyable), nil
}
// Patch update the navigation property provisioningPolicies in deviceManagement
func (m *CloudPcProvisioningPolicyItemRequestBuilder) Patch(options *CloudPcProvisioningPolicyItemRequestBuilderPatchOptions)(error) {
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
