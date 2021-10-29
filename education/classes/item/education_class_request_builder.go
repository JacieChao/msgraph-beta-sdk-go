package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentsettings"
    i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/schools"
    i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/members"
    i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/teachers"
    i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments"
    ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentdefaults"
    ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentcategories"
    if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/group"
    i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item"
    id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentcategories/item"
)

type EducationClassRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationClassRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *EducationClassRequestBuilder) AssignmentCategories()(*ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1.AssignmentCategoriesRequestBuilder) {
    return ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1.NewAssignmentCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentCategoriesById(id string)(*id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentDefaults()(*ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.AssignmentDefaultsRequestBuilder) {
    return ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Assignments()(*i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9.AssignmentsRequestBuilder) {
    return i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentsById(id string)(*i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613.EducationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613.NewEducationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentSettings()(*i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.AssignmentSettingsRequestBuilder) {
    return i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEducationClassRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassRequestBuilder) {
    m := &EducationClassRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/classes/{educationClass_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEducationClassRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationClassRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationClassRequestBuilder) CreateGetRequestInformation(q func (value *EducationClassRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationClassRequestBuilderGetQueryParameters)
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
func (m *EducationClassRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClass, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationClassRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationClassRequestBuilder) Get(q func (value *EducationClassRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClass, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationClass() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClass), nil
}
func (m *EducationClassRequestBuilder) Group()(*if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.GroupRequestBuilder) {
    return if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Members()(*i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782.MembersRequestBuilder) {
    return i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClass, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationClassRequestBuilder) Schools()(*i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f.SchoolsRequestBuilder) {
    return i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Teachers()(*i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c.TeachersRequestBuilder) {
    return i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c.NewTeachersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}