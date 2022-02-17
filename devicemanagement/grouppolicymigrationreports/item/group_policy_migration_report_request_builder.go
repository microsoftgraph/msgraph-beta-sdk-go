package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings"
    ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions"
    i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings/item"
    i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions/item"
)

// GroupPolicyMigrationReportRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyMigrationReports\{groupPolicyMigrationReport-id}
type GroupPolicyMigrationReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyMigrationReportRequestBuilderDeleteOptions options for Delete
type GroupPolicyMigrationReportRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyMigrationReportRequestBuilderGetOptions options for Get
type GroupPolicyMigrationReportRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyMigrationReportRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyMigrationReportRequestBuilderGetQueryParameters a list of Group Policy migration reports.
type GroupPolicyMigrationReportRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyMigrationReportRequestBuilderPatchOptions options for Patch
type GroupPolicyMigrationReportRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupPolicyMigrationReportRequestBuilderInternal instantiates a new GroupPolicyMigrationReportRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportRequestBuilder) {
    m := &GroupPolicyMigrationReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyMigrationReportRequestBuilder instantiates a new GroupPolicyMigrationReportRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyMigrationReportRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) CreateGetRequestInformation(options *GroupPolicyMigrationReportRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyMigrationReportRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) Delete(options *GroupPolicyMigrationReportRequestBuilderDeleteOptions)(error) {
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
// Get a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) Get(options *GroupPolicyMigrationReportRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyMigrationReport() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReport), nil
}
func (m *GroupPolicyMigrationReportRequestBuilder) GroupPolicySettingMappings()(*i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.GroupPolicySettingMappingsRequestBuilder) {
    return i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.NewGroupPolicySettingMappingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicySettingMappingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyMigrationReports.item.groupPolicySettingMappings.item collection
func (m *GroupPolicyMigrationReportRequestBuilder) GroupPolicySettingMappingsById(id string)(*i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.GroupPolicySettingMappingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicySettingMapping_id"] = id
    }
    return i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.NewGroupPolicySettingMappingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportRequestBuilder) Patch(options *GroupPolicyMigrationReportRequestBuilderPatchOptions)(error) {
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
func (m *GroupPolicyMigrationReportRequestBuilder) UnsupportedGroupPolicyExtensions()(*ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.UnsupportedGroupPolicyExtensionsRequestBuilder) {
    return ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.NewUnsupportedGroupPolicyExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsupportedGroupPolicyExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyMigrationReports.item.unsupportedGroupPolicyExtensions.item collection
func (m *GroupPolicyMigrationReportRequestBuilder) UnsupportedGroupPolicyExtensionsById(id string)(*i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.UnsupportedGroupPolicyExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unsupportedGroupPolicyExtension_id"] = id
    }
    return i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.NewUnsupportedGroupPolicyExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
