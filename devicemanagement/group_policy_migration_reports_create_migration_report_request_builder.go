package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder provides operations to call the createMigrationReport method.
type GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyMigrationReportsCreateMigrationReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsCreateMigrationReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilderInternal instantiates a new CreateMigrationReportRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) {
    m := &GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/createMigrationReport", pathParameters),
    }
    return m
}
// NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilder instantiates a new CreateMigrationReportRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createMigrationReport
// Deprecated: This method is obsolete. Use PostAsCreateMigrationReportPostResponse instead.
func (m *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) Post(ctx context.Context, body GroupPolicyMigrationReportsCreateMigrationReportPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsCreateMigrationReportResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsCreateMigrationReportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsCreateMigrationReportResponseable), nil
}
// PostAsCreateMigrationReportPostResponse invoke action createMigrationReport
func (m *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) PostAsCreateMigrationReportPostResponse(ctx context.Context, body GroupPolicyMigrationReportsCreateMigrationReportPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsCreateMigrationReportPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsCreateMigrationReportPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsCreateMigrationReportPostResponseable), nil
}
// ToPostRequestInformation invoke action createMigrationReport
func (m *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body GroupPolicyMigrationReportsCreateMigrationReportPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyMigrationReportsCreateMigrationReportRequestBuilder) {
    return NewGroupPolicyMigrationReportsCreateMigrationReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
