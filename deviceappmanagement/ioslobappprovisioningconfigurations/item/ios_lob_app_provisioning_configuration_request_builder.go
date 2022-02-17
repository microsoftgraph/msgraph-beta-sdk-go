package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses"
    i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments"
    ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assign"
    ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments"
    if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses"
    i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses/item"
    i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments/item"
    i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses/item"
    ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments/item"
)

// IosLobAppProvisioningConfigurationRequestBuilder builds and executes requests for operations under \deviceAppManagement\iosLobAppProvisioningConfigurations\{iosLobAppProvisioningConfiguration-id}
type IosLobAppProvisioningConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IosLobAppProvisioningConfigurationRequestBuilderDeleteOptions options for Delete
type IosLobAppProvisioningConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosLobAppProvisioningConfigurationRequestBuilderGetOptions options for Get
type IosLobAppProvisioningConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IosLobAppProvisioningConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IosLobAppProvisioningConfigurationRequestBuilderPatchOptions options for Patch
type IosLobAppProvisioningConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Assign()(*ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.AssignRequestBuilder) {
    return ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Assignments()(*ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.AssignmentsRequestBuilder) {
    return ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.assignments.item collection
func (m *IosLobAppProvisioningConfigurationRequestBuilder) AssignmentsById(id string)(*ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.IosLobAppProvisioningConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment_id"] = id
    }
    return ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.NewIosLobAppProvisioningConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIosLobAppProvisioningConfigurationRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosLobAppProvisioningConfigurationRequestBuilder instantiates a new IosLobAppProvisioningConfigurationRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreateDeleteRequestInformation(options *IosLobAppProvisioningConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreateGetRequestInformation(options *IosLobAppProvisioningConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) CreatePatchRequestInformation(options *IosLobAppProvisioningConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Delete(options *IosLobAppProvisioningConfigurationRequestBuilderDeleteOptions)(error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) DeviceStatuses()(*if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.DeviceStatusesRequestBuilder) {
    return if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.deviceStatuses.item collection
func (m *IosLobAppProvisioningConfigurationRequestBuilder) DeviceStatusesById(id string)(*i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Get(options *IosLobAppProvisioningConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIosLobAppProvisioningConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfiguration), nil
}
func (m *IosLobAppProvisioningConfigurationRequestBuilder) GroupAssignments()(*i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.GroupAssignmentsRequestBuilder) {
    return i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.groupAssignments.item collection
func (m *IosLobAppProvisioningConfigurationRequestBuilder) GroupAssignmentsById(id string)(*i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.MobileAppProvisioningConfigGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment_id"] = id
    }
    return i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.NewMobileAppProvisioningConfigGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationRequestBuilder) Patch(options *IosLobAppProvisioningConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *IosLobAppProvisioningConfigurationRequestBuilder) UserStatuses()(*i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.UserStatusesRequestBuilder) {
    return i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.userStatuses.item collection
func (m *IosLobAppProvisioningConfigurationRequestBuilder) UserStatusesById(id string)(*i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
