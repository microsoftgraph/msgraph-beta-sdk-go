package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

// TargetedManagedAppConfigurationRequestBuilder builds and executes requests for operations under \deviceAppManagement\targetedManagedAppConfigurations\{targetedManagedAppConfiguration-id}
type TargetedManagedAppConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TargetedManagedAppConfigurationRequestBuilderDeleteOptions options for Delete
type TargetedManagedAppConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TargetedManagedAppConfigurationRequestBuilderGetOptions options for Get
type TargetedManagedAppConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TargetedManagedAppConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TargetedManagedAppConfigurationRequestBuilderGetQueryParameters targeted managed app configurations.
type TargetedManagedAppConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TargetedManagedAppConfigurationRequestBuilderPatchOptions options for Patch
type TargetedManagedAppConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Apps()(*i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.AppsRequestBuilder) {
    return i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.apps.item collection
func (m *TargetedManagedAppConfigurationRequestBuilder) AppsById(id string)(*i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.ManagedMobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.NewManagedMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assign()(*iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.AssignRequestBuilder) {
    return iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assignments()(*i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.AssignmentsRequestBuilder) {
    return i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.assignments.item collection
func (m *TargetedManagedAppConfigurationRequestBuilder) AssignmentsById(id string)(*ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.TargetedManagedAppPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment_id"] = id
    }
    return ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.NewTargetedManagedAppPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTargetedManagedAppConfigurationRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    m := &TargetedManagedAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppConfigurationRequestBuilder instantiates a new TargetedManagedAppConfigurationRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateDeleteRequestInformation(options *TargetedManagedAppConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateGetRequestInformation(options *TargetedManagedAppConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) CreatePatchRequestInformation(options *TargetedManagedAppConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) Delete(options *TargetedManagedAppConfigurationRequestBuilderDeleteOptions)(error) {
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
func (m *TargetedManagedAppConfigurationRequestBuilder) DeploymentSummary()(*ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.DeploymentSummaryRequestBuilder) {
    return ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) Get(options *TargetedManagedAppConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTargetedManagedAppConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration), nil
}
// Patch targeted managed app configurations.
func (m *TargetedManagedAppConfigurationRequestBuilder) Patch(options *TargetedManagedAppConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *TargetedManagedAppConfigurationRequestBuilder) TargetApps()(*i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.TargetAppsRequestBuilder) {
    return i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
