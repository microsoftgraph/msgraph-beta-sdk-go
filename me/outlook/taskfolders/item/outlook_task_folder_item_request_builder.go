package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks"
    i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties"
    i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties"
    i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/tasks/item"
    i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/multivalueextendedproperties/item"
    i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/outlook/taskfolders/item/singlevalueextendedproperties/item"
)

// OutlookTaskFolderItemRequestBuilder provides operations to manage the taskFolders property of the microsoft.graph.outlookUser entity.
type OutlookTaskFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OutlookTaskFolderItemRequestBuilderDeleteOptions options for Delete
type OutlookTaskFolderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskFolderItemRequestBuilderGetOptions options for Get
type OutlookTaskFolderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OutlookTaskFolderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OutlookTaskFolderItemRequestBuilderGetQueryParameters get taskFolders from me
type OutlookTaskFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// OutlookTaskFolderItemRequestBuilderPatchOptions options for Patch
type OutlookTaskFolderItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOutlookTaskFolderItemRequestBuilderInternal instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    m := &OutlookTaskFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/outlook/taskFolders/{outlookTaskFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOutlookTaskFolderItemRequestBuilder instantiates a new OutlookTaskFolderItemRequestBuilder and sets the default values.
func NewOutlookTaskFolderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OutlookTaskFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutlookTaskFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) CreateDeleteRequestInformation(options *OutlookTaskFolderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) CreateGetRequestInformation(options *OutlookTaskFolderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) CreatePatchRequestInformation(options *OutlookTaskFolderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property taskFolders for me
func (m *OutlookTaskFolderItemRequestBuilder) Delete(options *OutlookTaskFolderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get taskFolders from me
func (m *OutlookTaskFolderItemRequestBuilder) Get(options *OutlookTaskFolderItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOutlookTaskFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OutlookTaskFolderable), nil
}
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedProperties()(*i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.MultiValueExtendedPropertiesRequestBuilder) {
    return i99568de587864405cbb7ec5179efa3e533e234570ca91e410212ec1d04fe089c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.multiValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i645e01811e5e9cf150ad7755ded1c4cca58e15173c57613b642c0c878c5c4224.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property taskFolders in me
func (m *OutlookTaskFolderItemRequestBuilder) Patch(options *OutlookTaskFolderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedProperties()(*i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.SingleValueExtendedPropertiesRequestBuilder) {
    return i44c0c086502e5c42f8c0b70140fda98b421c68cee2ad57dce93ec3e17b4a1a9a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.singleValueExtendedProperties.item collection
func (m *OutlookTaskFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7f227a30fb4ec3965178563ee0f1c3b874776270d21828095647c10a35379e25.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OutlookTaskFolderItemRequestBuilder) Tasks()(*i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.TasksRequestBuilder) {
    return i43e25b1a2e7db024aa2c8195c093b9f95136385066d5a6b0e1a1718ccbb4941a.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.outlook.taskFolders.item.tasks.item collection
func (m *OutlookTaskFolderItemRequestBuilder) TasksById(id string)(*i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.OutlookTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["outlookTask_id"] = id
    }
    return i4c149284b5e78d79fdca30504030b7dd8696f7b3a114850f0c0c50352668aeb9.NewOutlookTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
