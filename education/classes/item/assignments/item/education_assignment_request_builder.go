package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i18b62fb6d5309822100f04bac8a3161379f7816c6f4adab56845b52682eb61b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/publish"
    i2557f883ecb5587b5da5a0c0734106f82f3c5f1aa29d86f1a875b53c70b1b82b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions"
    i4875bb54e80dc288513cbd0383790b487612eadda53bf350ff9a138a5e29e1a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/resources"
    i7923129eb1b685aab5f95e3e0259dc8c04bfab2463e4d9757b2cdba960d96a0c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/setupresourcesfolder"
    iae6580328f33d9d020c6355ff035182611a1120e165b912963deb516e95364a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/categories"
    ic3d0439f8b9690e92c9beef8423b7899bca0c8873a34fca7e32f5d2e15354c00 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/rubric"
    i189527b84f3683dcbfd7a4ef6ad41af158e24950ad0928ed0e10c7445b1b1035 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/categories/item"
    i8fc04b58f02d2e1b886020ba420268c1b8b477afcae88e95f3bee06b9d6a7dcd "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item"
    id71b2ce69d7d1a2ad5737d9ff32545b32062298268db000f841be5bc7329af30 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/resources/item"
)

type EducationAssignmentRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationAssignmentRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EducationAssignmentRequestBuilder) Categories()(*iae6580328f33d9d020c6355ff035182611a1120e165b912963deb516e95364a0.CategoriesRequestBuilder) {
    return iae6580328f33d9d020c6355ff035182611a1120e165b912963deb516e95364a0.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*i189527b84f3683dcbfd7a4ef6ad41af158e24950ad0928ed0e10c7445b1b1035.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return i189527b84f3683dcbfd7a4ef6ad41af158e24950ad0928ed0e10c7445b1b1035.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/classes/{educationClass_id}/assignments/{educationAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEducationAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationAssignmentRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationAssignmentRequestBuilder) CreateGetRequestInformation(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationAssignmentRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationAssignmentRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationAssignmentRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationAssignmentRequestBuilder) Get(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationAssignment() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment), nil
}
func (m *EducationAssignmentRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationAssignmentRequestBuilder) Publish()(*i18b62fb6d5309822100f04bac8a3161379f7816c6f4adab56845b52682eb61b1.PublishRequestBuilder) {
    return i18b62fb6d5309822100f04bac8a3161379f7816c6f4adab56845b52682eb61b1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i4875bb54e80dc288513cbd0383790b487612eadda53bf350ff9a138a5e29e1a5.ResourcesRequestBuilder) {
    return i4875bb54e80dc288513cbd0383790b487612eadda53bf350ff9a138a5e29e1a5.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*id71b2ce69d7d1a2ad5737d9ff32545b32062298268db000f841be5bc7329af30.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return id71b2ce69d7d1a2ad5737d9ff32545b32062298268db000f841be5bc7329af30.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*ic3d0439f8b9690e92c9beef8423b7899bca0c8873a34fca7e32f5d2e15354c00.RubricRequestBuilder) {
    return ic3d0439f8b9690e92c9beef8423b7899bca0c8873a34fca7e32f5d2e15354c00.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i7923129eb1b685aab5f95e3e0259dc8c04bfab2463e4d9757b2cdba960d96a0c.SetUpResourcesFolderRequestBuilder) {
    return i7923129eb1b685aab5f95e3e0259dc8c04bfab2463e4d9757b2cdba960d96a0c.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*i2557f883ecb5587b5da5a0c0734106f82f3c5f1aa29d86f1a875b53c70b1b82b.SubmissionsRequestBuilder) {
    return i2557f883ecb5587b5da5a0c0734106f82f3c5f1aa29d86f1a875b53c70b1b82b.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i8fc04b58f02d2e1b886020ba420268c1b8b477afcae88e95f3bee06b9d6a7dcd.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i8fc04b58f02d2e1b886020ba420268c1b8b477afcae88e95f3bee06b9d6a7dcd.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
