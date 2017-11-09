package teamcity

import (
	"errors"
	"fmt"
	"github.com/Cardfree/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllProjectParameters(projectID string, parameters *types.Parameters) error {
	path := fmt.Sprintf("/httpAuth/app/rest/projects/id:%s/parameters", projectID)
	var parametersReturn *types.Parameters

	err := c.doRetryRequest("PUT", path, parameters, &parametersReturn)
	if err != nil {
		return err
	}

	if parametersReturn == nil {
		return errors.New("project not updated")
	}
	*parameters = *parametersReturn

	return nil
}