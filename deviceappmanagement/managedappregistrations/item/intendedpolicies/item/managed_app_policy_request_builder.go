package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i281e1c8a748086adf38b7084418b11f03ab91a590efb0860d3a49a31529931c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection"
    i3d87d99714147c990c532d86358aed917a008540bf0e82de3546748456448bb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/managedappprotection"
    i5cd6d74ba596318f265190129cfc42fd3fcb93764fb956aabda82f2de1438925 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetapps"
    ia5e8ce427876d3bb4743896822b16f0365266b66560b2c73e1b2fec2b4625524 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/windowsinformationprotection"
)

// ManagedAppPolicyRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\intendedPolicies\{managedAppPolicy-id}
type ManagedAppPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppPolicyRequestBuilderDeleteOptions options for Delete
type ManagedAppPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyRequestBuilderGetOptions options for Get
type ManagedAppPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyRequestBuilderGetQueryParameters zero or more policies admin intended for the app as of now.
type ManagedAppPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedAppPolicyRequestBuilderPatchOptions options for Patch
type ManagedAppPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewManagedAppPolicyRequestBuilderInternal instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
func NewManagedAppPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    m := &ManagedAppPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/intendedPolicies/{managedAppPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppPolicyRequestBuilder instantiates a new ManagedAppPolicyRequestBuilder and sets the default values.
func NewManagedAppPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) CreateGetRequestInformation(options *ManagedAppPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) CreatePatchRequestInformation(options *ManagedAppPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) Delete(options *ManagedAppPolicyRequestBuilderDeleteOptions)(error) {
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
// Get zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) Get(options *ManagedAppPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedAppPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppPolicy), nil
}
func (m *ManagedAppPolicyRequestBuilder) ManagedAppProtection()(*i3d87d99714147c990c532d86358aed917a008540bf0e82de3546748456448bb1.ManagedAppProtectionRequestBuilder) {
    return i3d87d99714147c990c532d86358aed917a008540bf0e82de3546748456448bb1.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyRequestBuilder) Patch(options *ManagedAppPolicyRequestBuilderPatchOptions)(error) {
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
func (m *ManagedAppPolicyRequestBuilder) TargetApps()(*i5cd6d74ba596318f265190129cfc42fd3fcb93764fb956aabda82f2de1438925.TargetAppsRequestBuilder) {
    return i5cd6d74ba596318f265190129cfc42fd3fcb93764fb956aabda82f2de1438925.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) TargetedManagedAppProtection()(*i281e1c8a748086adf38b7084418b11f03ab91a590efb0860d3a49a31529931c3.TargetedManagedAppProtectionRequestBuilder) {
    return i281e1c8a748086adf38b7084418b11f03ab91a590efb0860d3a49a31529931c3.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) WindowsInformationProtection()(*ia5e8ce427876d3bb4743896822b16f0365266b66560b2c73e1b2fec2b4625524.WindowsInformationProtectionRequestBuilder) {
    return ia5e8ce427876d3bb4743896822b16f0365266b66560b2c73e1b2fec2b4625524.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
