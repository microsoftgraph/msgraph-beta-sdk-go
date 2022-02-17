package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3f4546266ed71e9c91b44a543a0beffa807c5b7d55c61515df51c44e98a53d73 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/targetedmanagedappprotection"
    i98f27ff67c0e5b948335041a3bf29d931ab7be58573ae29bfd1554d78924873d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/windowsinformationprotection"
    ia5b743276ffdb496eb66220a0a53ba3adf25d50fba689c1c1a6dbe1ac9c91d34 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/managedappprotection"
    ie7f5606c139c22faaa101d1eb10801cba4ca1282d26fa67419451925a809558b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/targetapps"
)

// ManagedAppPolicyRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\appliedPolicies\{managedAppPolicy-id}
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
// ManagedAppPolicyRequestBuilderGetQueryParameters zero or more policys already applied on the registered app when it last synchronized with managment service.
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
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/appliedPolicies/{managedAppPolicy_id}{?select,expand}";
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
// CreateDeleteRequestInformation zero or more policys already applied on the registered app when it last synchronized with managment service.
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
// CreateGetRequestInformation zero or more policys already applied on the registered app when it last synchronized with managment service.
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
// CreatePatchRequestInformation zero or more policys already applied on the registered app when it last synchronized with managment service.
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
// Delete zero or more policys already applied on the registered app when it last synchronized with managment service.
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
// Get zero or more policys already applied on the registered app when it last synchronized with managment service.
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
func (m *ManagedAppPolicyRequestBuilder) ManagedAppProtection()(*ia5b743276ffdb496eb66220a0a53ba3adf25d50fba689c1c1a6dbe1ac9c91d34.ManagedAppProtectionRequestBuilder) {
    return ia5b743276ffdb496eb66220a0a53ba3adf25d50fba689c1c1a6dbe1ac9c91d34.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch zero or more policys already applied on the registered app when it last synchronized with managment service.
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
func (m *ManagedAppPolicyRequestBuilder) TargetApps()(*ie7f5606c139c22faaa101d1eb10801cba4ca1282d26fa67419451925a809558b.TargetAppsRequestBuilder) {
    return ie7f5606c139c22faaa101d1eb10801cba4ca1282d26fa67419451925a809558b.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) TargetedManagedAppProtection()(*i3f4546266ed71e9c91b44a543a0beffa807c5b7d55c61515df51c44e98a53d73.TargetedManagedAppProtectionRequestBuilder) {
    return i3f4546266ed71e9c91b44a543a0beffa807c5b7d55c61515df51c44e98a53d73.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyRequestBuilder) WindowsInformationProtection()(*i98f27ff67c0e5b948335041a3bf29d931ab7be58573ae29bfd1554d78924873d.WindowsInformationProtectionRequestBuilder) {
    return i98f27ff67c0e5b948335041a3bf29d931ab7be58573ae29bfd1554d78924873d.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
