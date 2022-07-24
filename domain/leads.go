package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func leadInputToLead(ctx context.Context, input *models.LeadInput) *models.Lead {
	return &models.Lead{
		Description: input.Description,
		Name:        input.Name,
		NextVisitID: input.NextVisitID,
		Phone:       input.Phone,
		Published:   input.Published,
		Source:      input.Source,
		Status:      input.Status,
		StudentIds:  input.StudentIds,
		TeamID:      input.TeamID,
		YearBorn:    input.YearBorn,
	}
}
func leadInputWithIdToLead(ctx context.Context, lead *models.Lead, input *models.LeadInput) *models.Lead {
	return &models.Lead{
		ID:          lead.ID,
		Description: input.Description,
		Name:        input.Name,
		NextVisitID: input.NextVisitID,
		Phone:       input.Phone,
		Published:   input.Published,
		Source:      input.Source,
		Status:      input.Status,
		StudentIds:  input.StudentIds,
		TeamID:      input.TeamID,
		YearBorn:    input.YearBorn,
	}
}

func (d *Domain) LeadSave(ctx context.Context, input models.LeadInput) (*models.LeadPayload, error) {
	return d.LeadsRepo.LeadsSave(leadInputToLead(ctx, &input))
}

func (d *Domain) LeadUpdate(ctx context.Context, leadInput models.LeadInputWithID) (*models.LeadPayload, error) {

	lead, err := d.LeadsRepo.GetLeadsByID(leadInput.ID)
	if err != nil || lead == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	lead = leadInputWithIdToLead(ctx, lead, leadInput.Input)
	leadPayload, err := d.LeadsRepo.LeadsUpdate(lead)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return leadPayload, nil
}

func (d *Domain) LeadPublish(id string) (*models.LeadPayload, error) {
	lead, err := d.LeadsRepo.GetLeadsByID(id)
	if err != nil || lead == nil {
		return nil, errors.New("филиала не существует")
	}

	leadPayload, err := d.LeadsRepo.LeadsPublish(lead)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return leadPayload, nil
}

func (d *Domain) LeadRestore(id string) (*models.LeadPayload, error) {
	lead, err := d.LeadsRepo.GetLeadsByID(id)
	if err != nil || lead == nil {
		return nil, errors.New("филиала не существует")
	}

	leadPayload, err := d.LeadsRepo.LeadsRestore(lead)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return leadPayload, nil
}

// LeadDelete is the resolver for the leadDelete field.
func (d *Domain) LeadDelete(id string) (bool, error) {
	lead, err := d.LeadsRepo.GetLeadsByID(id)
	if err != nil || lead == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.LeadsRepo.LeadsDelete(lead)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
