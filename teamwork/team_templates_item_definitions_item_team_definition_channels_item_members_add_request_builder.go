package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder provides operations to call the add method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderInternal instantiates a new AddRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/{channel%2Did}/members/microsoft.graph.add";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder instantiates a new AddRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/conversationmembers-add?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder) Post(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddResponseable), nil
}
// ToPostRequestInformation add multiple members in a single request to a team. The response provides details about which memberships could and couldn't be created.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsItemMembersAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}