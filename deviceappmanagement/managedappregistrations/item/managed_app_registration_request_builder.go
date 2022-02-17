package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies"
    id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations"
    if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies"
    i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item"
    i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/operations/item"
    id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item"
)

// ManagedAppRegistrationRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}
type ManagedAppRegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppRegistrationRequestBuilderDeleteOptions options for Delete
type ManagedAppRegistrationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppRegistrationRequestBuilderGetOptions options for Get
type ManagedAppRegistrationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppRegistrationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppRegistrationRequestBuilderGetQueryParameters the managed app registrations.
type ManagedAppRegistrationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedAppRegistrationRequestBuilderPatchOptions options for Patch
type ManagedAppRegistrationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedAppRegistrationRequestBuilder) AppliedPolicies()(*i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.AppliedPoliciesRequestBuilder) {
    return i56070fa6720bc0b45941309da76f8b33eaeab41361bc4d8deae7cf9b20365caa.NewAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.appliedPolicies.item collection
func (m *ManagedAppRegistrationRequestBuilder) AppliedPoliciesById(id string)(*i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i1c13dfabdff5dc012fd7d12bf25406e90dff95d5a2f3f2db4980cf2baccac50e.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedAppRegistrationRequestBuilderInternal instantiates a new ManagedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    m := &ManagedAppRegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationRequestBuilder instantiates a new ManagedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppRegistrationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreateGetRequestInformation(options *ManagedAppRegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreatePatchRequestInformation(options *ManagedAppRegistrationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Delete(options *ManagedAppRegistrationRequestBuilderDeleteOptions)(error) {
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
// Get the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Get(options *ManagedAppRegistrationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedAppRegistration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedAppRegistration), nil
}
func (m *ManagedAppRegistrationRequestBuilder) IntendedPolicies()(*if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.IntendedPoliciesRequestBuilder) {
    return if8c50fd03afc54f53eb9adc01674971598fcc5e91f81ad3c1573605a4ba9a43c.NewIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.intendedPolicies.item collection
func (m *ManagedAppRegistrationRequestBuilder) IntendedPoliciesById(id string)(*id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return id0b04ce4d5e56adfc6f008d1080876eacead82cead17447895ea3eb8353b96ac.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) Operations()(*id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.OperationsRequestBuilder) {
    return id957f7cf90d5e3cc168823ce7a8346428ac2a3905e6406e2ee0dc61e8312d433.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item.operations.item collection
func (m *ManagedAppRegistrationRequestBuilder) OperationsById(id string)(*i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.ManagedAppOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation_id"] = id
    }
    return i4188f43081529892f83979f26fa2b9dca74e78bf1260dae80736bc3d6815fb36.NewManagedAppOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Patch(options *ManagedAppRegistrationRequestBuilderPatchOptions)(error) {
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
