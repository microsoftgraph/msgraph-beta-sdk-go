package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatussummary"
    i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses"
    ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assign"
    ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments"
    if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses"
    if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatussummary"
    i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/userstatuses/item"
    i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/assignments/item"
    i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item/devicestatuses/item"
)

// ManagedDeviceMobileAppConfigurationRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileAppConfigurations\{managedDeviceMobileAppConfiguration-id}
type ManagedDeviceMobileAppConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions options for Delete
type ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions options for Get
type ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type ManagedDeviceMobileAppConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions options for Patch
type ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assign()(*ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.AssignRequestBuilder) {
    return ibac2d0e4ca3b406649c8be63759ea1f49c5e4a6777b821a60efdd3a7c05b2b69.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Assignments()(*ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.AssignmentsRequestBuilder) {
    return ie295d58e16adbdb351de4946e4c8aace040173e000c6be8ce847cb3956139b8f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.assignments.item collection
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) AssignmentsById(id string)(*i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.ManagedDeviceMobileAppConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationAssignment_id"] = id
    }
    return i0b89b772a1c92857b752524ff6c11354157f30f3b2dd9a7d1255eb599cd446c2.NewManagedDeviceMobileAppConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedDeviceMobileAppConfigurationRequestBuilderInternal instantiates a new ManagedDeviceMobileAppConfigurationRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    m := &ManagedDeviceMobileAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceMobileAppConfigurationRequestBuilder instantiates a new ManagedDeviceMobileAppConfigurationRequestBuilder and sets the default values.
func NewManagedDeviceMobileAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceMobileAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateDeleteRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) CreatePatchRequestInformation(options *ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Delete(options *ManagedDeviceMobileAppConfigurationRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatuses()(*if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.DeviceStatusesRequestBuilder) {
    return if01c562d0b25ad35481fc3c806b08201d62dbaa863b0d4140eb330ba2cca9450.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.deviceStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusesById(id string)(*i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.ManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus_id"] = id
    }
    return i5271f84312003ab395dfea4b1460dfaa29d399e18c82ceb4ee77b6b61905f888.NewManagedDeviceMobileAppConfigurationDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) DeviceStatusSummary()(*i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.DeviceStatusSummaryRequestBuilder) {
    return i5a2cd73ff3ec9306808eb1e3300df2fc7e52f3aad1986773afba775fd8b57b8b.NewDeviceStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Get(options *ManagedDeviceMobileAppConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewManagedDeviceMobileAppConfiguration() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ManagedDeviceMobileAppConfiguration), nil
}
// Patch the Managed Device Mobile Application Configurations.
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) Patch(options *ManagedDeviceMobileAppConfigurationRequestBuilderPatchOptions)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatuses()(*i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.UserStatusesRequestBuilder) {
    return i8499825311e6b8b06993abc8855d7005b4d71a8a187427433703d1222b03e0d0.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item.userStatuses.item collection
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusesById(id string)(*i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus_id"] = id
    }
    return i0247b48acb03c0bf84935d01d7e57b2bb123143ad97e26f3e9957cc4cc1f2112.NewManagedDeviceMobileAppConfigurationUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedDeviceMobileAppConfigurationRequestBuilder) UserStatusSummary()(*if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.UserStatusSummaryRequestBuilder) {
    return if1c590ad1d01c3cd722894f058178e6da9c588aa2af587a8d25bad4064720de4.NewUserStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
