package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder provides operations to manage the definitions property of the microsoft.graph.accessReviewSet entity.
type AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters retrieve an accessReviewScheduleDefinition object by ID. This returns all properties of the scheduled access review series except for the associated accessReviewInstances. Each accessReviewScheduleDefinition has at least one instance. An instance represents a review for a specific resource (such as a particular group's members), during one occurrence (for example, March 2021) of a recurring review. To retrieve the instances of the access review series, use the list accessReviewInstance API.
type AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetQueryParameters
}
// AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal instantiates a new AccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    m := &AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder instantiates a new AccessReviewScheduleDefinitionItemRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an accessReviewScheduleDefinition object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewscheduledefinition-delete?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve an accessReviewScheduleDefinition object by ID. This returns all properties of the scheduled access review series except for the associated accessReviewInstances. Each accessReviewScheduleDefinition has at least one instance. An instance represents a review for a specific resource (such as a particular group's members), during one occurrence (for example, March 2021) of a recurring review. To retrieve the instances of the access review series, use the list accessReviewInstance API.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewscheduledefinition-get?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable), nil
}
// Instances provides operations to manage the instances property of the microsoft.graph.accessReviewScheduleDefinition entity.
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Instances()(*AccessReviewsDefinitionsItemInstancesRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update the navigation property definitions in identityGovernance
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Put(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable), nil
}
// Stop provides operations to call the stop method.
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) Stop()(*AccessReviewsDefinitionsItemStopRequestBuilder) {
    return NewAccessReviewsDefinitionsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an accessReviewScheduleDefinition object.
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve an accessReviewScheduleDefinition object by ID. This returns all properties of the scheduled access review series except for the associated accessReviewInstances. Each accessReviewScheduleDefinition has at least one instance. An instance represents a review for a specific resource (such as a particular group's members), during one occurrence (for example, March 2021) of a recurring review. To retrieve the instances of the access review series, use the list accessReviewInstance API.
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation update the navigation property definitions in identityGovernance
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, requestConfiguration *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder) {
    return NewAccessReviewsDefinitionsAccessReviewScheduleDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
