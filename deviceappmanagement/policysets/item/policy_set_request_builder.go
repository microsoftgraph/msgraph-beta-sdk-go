package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/items"
    i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/update"
    ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/assignments"
    i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/assignments/item"
    i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/items/item"
)

// PolicySetRequestBuilder builds and executes requests for operations under \deviceAppManagement\policySets\{policySet-id}
type PolicySetRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PolicySetRequestBuilderDeleteOptions options for Delete
type PolicySetRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PolicySetRequestBuilderGetOptions options for Get
type PolicySetRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PolicySetRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PolicySetRequestBuilderGetQueryParameters the PolicySet of Policies and Applications
type PolicySetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PolicySetRequestBuilderPatchOptions options for Patch
type PolicySetRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicySet;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PolicySetRequestBuilder) Assignments()(*ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a.AssignmentsRequestBuilder) {
    return ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.policySets.item.assignments.item collection
func (m *PolicySetRequestBuilder) AssignmentsById(id string)(*i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4.PolicySetAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["policySetAssignment_id"] = id
    }
    return i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4.NewPolicySetAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPolicySetRequestBuilderInternal instantiates a new PolicySetRequestBuilder and sets the default values.
func NewPolicySetRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PolicySetRequestBuilder) {
    m := &PolicySetRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/policySets/{policySet_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPolicySetRequestBuilder instantiates a new PolicySetRequestBuilder and sets the default values.
func NewPolicySetRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PolicySetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicySetRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) CreateDeleteRequestInformation(options *PolicySetRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) CreateGetRequestInformation(options *PolicySetRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) CreatePatchRequestInformation(options *PolicySetRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) Delete(options *PolicySetRequestBuilderDeleteOptions)(error) {
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
// Get the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) Get(options *PolicySetRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicySet, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPolicySet() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicySet), nil
}
func (m *PolicySetRequestBuilder) Items()(*i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d.ItemsRequestBuilder) {
    return i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.policySets.item.items.item collection
func (m *PolicySetRequestBuilder) ItemsById(id string)(*i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1.PolicySetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["policySetItem_id"] = id
    }
    return i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1.NewPolicySetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the PolicySet of Policies and Applications
func (m *PolicySetRequestBuilder) Patch(options *PolicySetRequestBuilderPatchOptions)(error) {
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
func (m *PolicySetRequestBuilder) Update()(*i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8.UpdateRequestBuilder) {
    return i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8.NewUpdateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
