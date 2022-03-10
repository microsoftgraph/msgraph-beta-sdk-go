package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments"
    ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assign"
    ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments"
    if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses"
    i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/devicestatuses/item"
    i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/groupassignments/item"
    i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/userstatuses/item"
    ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item/assignments/item"
)

// IosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type IosLobAppProvisioningConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IosLobAppProvisioningConfigurationItemRequestBuilderDeleteOptions options for Delete
type IosLobAppProvisioningConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosLobAppProvisioningConfigurationItemRequestBuilderGetOptions options for Get
type IosLobAppProvisioningConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IosLobAppProvisioningConfigurationItemRequestBuilderPatchOptions options for Patch
type IosLobAppProvisioningConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.AssignRequestBuilder) {
    return ia3ce539b9f79e09109b9926ad281d97d40b25705ecefa910e2bfd8db0c875c85.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.AssignmentsRequestBuilder) {
    return ie78372adfad521ab316f12bda3bf095efac22a79d88a6727c101e09421a58567.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.assignments.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) AssignmentsById(id string)(*ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.IosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfigurationAssignment_id"] = id
    }
    return ia1a84f1e276568c2aeb23fcbd50f8de951fa5878159c134fa0228aec984ff3c0.NewIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *IosLobAppProvisioningConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreateGetRequestInformation(options *IosLobAppProvisioningConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *IosLobAppProvisioningConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Delete(options *IosLobAppProvisioningConfigurationItemRequestBuilderDeleteOptions)(error) {
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
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.DeviceStatusesRequestBuilder) {
    return if9007b7e9e190f805a646bbd1f3e6ae76781048bd465090c2016244c10d7e0e5.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.deviceStatuses.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.ManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i072136c315e8186c0ff2a10126695a2fb03cdd14dfc2acf41de66f1f1caba3d8.NewManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the IOS Lob App Provisioning Configurations.
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Get(options *IosLobAppProvisioningConfigurationItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosLobAppProvisioningConfigurationable), nil
}
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.GroupAssignmentsRequestBuilder) {
    return i5131d1ed5fbac3ef96a626337be97ffb49d1383128e9d493a6940ceb8833b3a0.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.groupAssignments.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.MobileAppProvisioningConfigGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppProvisioningConfigGroupAssignment_id"] = id
    }
    return i11301c42fede611eafd37db9196e74ba08d3d72e93b62bd910672a1181aa6c19.NewMobileAppProvisioningConfigGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) Patch(options *IosLobAppProvisioningConfigurationItemRequestBuilderPatchOptions)(error) {
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
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.UserStatusesRequestBuilder) {
    return i36927486fc4ece174ed7ee7d23b6d06925d565134b4116881a5a149b17667929.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item.userStatuses.item collection
func (m *IosLobAppProvisioningConfigurationItemRequestBuilder) UserStatusesById(id string)(*i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.ManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i6393c179ab957b9b86c9cd839a5d7820fb3ce75696832f0ef436bd70a8a4196d.NewManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
