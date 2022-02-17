package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i001eb6cb139a524dd351cd98bb0c439290700a903f23d19a2f27b36ffe531777 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/presentationvalues/item/definitionvalue"
    i0c67a157507f229b79bb5a42860b892b4a5eafc70b612c8f218a33421b0a7dfc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/presentationvalues/item/presentation"
)

// GroupPolicyPresentationValueRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyConfigurations\{groupPolicyConfiguration-id}\definitionValues\{groupPolicyDefinitionValue-id}\presentationValues\{groupPolicyPresentationValue-id}
type GroupPolicyPresentationValueRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyPresentationValueRequestBuilderDeleteOptions options for Delete
type GroupPolicyPresentationValueRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyPresentationValueRequestBuilderGetOptions options for Get
type GroupPolicyPresentationValueRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyPresentationValueRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyPresentationValueRequestBuilderGetQueryParameters the associated group policy presentation values with the definition value.
type GroupPolicyPresentationValueRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyPresentationValueRequestBuilderPatchOptions options for Patch
type GroupPolicyPresentationValueRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyPresentationValue;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupPolicyPresentationValueRequestBuilderInternal instantiates a new GroupPolicyPresentationValueRequestBuilder and sets the default values.
func NewGroupPolicyPresentationValueRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyPresentationValueRequestBuilder) {
    m := &GroupPolicyPresentationValueRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration_id}/definitionValues/{groupPolicyDefinitionValue_id}/presentationValues/{groupPolicyPresentationValue_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyPresentationValueRequestBuilder instantiates a new GroupPolicyPresentationValueRequestBuilder and sets the default values.
func NewGroupPolicyPresentationValueRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyPresentationValueRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyPresentationValueRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyPresentationValueRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) CreateGetRequestInformation(options *GroupPolicyPresentationValueRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyPresentationValueRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyPresentationValueRequestBuilder) DefinitionValue()(*i001eb6cb139a524dd351cd98bb0c439290700a903f23d19a2f27b36ffe531777.DefinitionValueRequestBuilder) {
    return i001eb6cb139a524dd351cd98bb0c439290700a903f23d19a2f27b36ffe531777.NewDefinitionValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) Delete(options *GroupPolicyPresentationValueRequestBuilderDeleteOptions)(error) {
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
// Get the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) Get(options *GroupPolicyPresentationValueRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyPresentationValue, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyPresentationValue() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyPresentationValue), nil
}
// Patch the associated group policy presentation values with the definition value.
func (m *GroupPolicyPresentationValueRequestBuilder) Patch(options *GroupPolicyPresentationValueRequestBuilderPatchOptions)(error) {
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
func (m *GroupPolicyPresentationValueRequestBuilder) Presentation()(*i0c67a157507f229b79bb5a42860b892b4a5eafc70b612c8f218a33421b0a7dfc.PresentationRequestBuilder) {
    return i0c67a157507f229b79bb5a42860b892b4a5eafc70b612c8f218a33421b0a7dfc.NewPresentationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
