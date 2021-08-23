package main

import (
	"context"

	"github.com/dishbreak/terraform-cloud-launcher/lib"
	"github.com/dishbreak/terraform-cloud-launcher/models"
	"github.com/hashicorp/go-tfe"
)

type WorkspaceListCmd struct{}

func (w *WorkspaceListCmd) Help() string {
	return `
Return a list of workspaces in Alfred JSON format.
	`
}

func paginate(nextPage int) tfe.WorkspaceListOptions {
	return tfe.WorkspaceListOptions{
		ListOptions: tfe.ListOptions{
			PageNumber: nextPage,
		},
	}
}

func (w *WorkspaceListCmd) Run(ctx *Context) error {
	resp := models.NewScriptResponse()

	client, err := lib.NewTfeClient()
	if err != nil {
		resp.SetError(err)
		return err
	}

	for nextPage := 1; nextPage != 0; {
		workspaceList, err := client.Workspaces.List(context.Background(), "nerdwallet", paginate(nextPage))
		if err != nil {
			resp.SetError(err)
			return err
		}
		for _, workspace := range workspaceList.Items {
			resp.AddItem(models.ListItem{
				Title:    workspace.Name,
				Subtitle: workspace.Description,
				Arg:      workspace.Name,
				Valid:    true,
			})
		}
		nextPage = workspaceList.NextPage
	}

	resp.SendFeedback()

	return nil
}
