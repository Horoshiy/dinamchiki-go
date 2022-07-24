package domain

import (
	"context"
	"errors"
	"fmt"
	models "gitlab.com/dinamchiki/go-graphql/graph/model"
)

func teamInputToTeam(ctx context.Context, input *models.TeamInput) *models.Team {
	return &models.Team{
		Ages:        input.Ages,
		CoachIds:    input.CoachIds,
		HeadCoachID: input.HeadCoachID,
		Name:        input.Name,
		PlaceID:     input.PlaceID,
		Published:   input.Published,
		Writable:    input.Writable,
	}
}
func teamInputWithIdToTeam(ctx context.Context, team *models.Team, input *models.TeamInput) *models.Team {
	return &models.Team{
		ID:          team.ID,
		Ages:        input.Ages,
		CoachIds:    input.CoachIds,
		HeadCoachID: input.HeadCoachID,
		Name:        input.Name,
		PlaceID:     input.PlaceID,
		Published:   input.Published,
		Writable:    input.Writable,
	}
}

func (d *Domain) TeamSave(ctx context.Context, input models.TeamInput) (*models.TeamPayload, error) {

	_, err := d.TeamsRepo.GetTeamsByName(input.Name)
	if err == nil {
		return nil, errors.New("такое название уже есть у стадиона")
	}

	return d.TeamsRepo.TeamsSave(teamInputToTeam(ctx, &input))
}

func (d *Domain) TeamUpdate(ctx context.Context, teamInput models.TeamInputWithID) (*models.TeamPayload, error) {

	team, err := d.TeamsRepo.GetTeamsByID(teamInput.ID)
	if err != nil || team == nil {
		return nil, errors.New("такого стадиона не существует")
	}

	if len(teamInput.Input.Name) < 3 {
		return nil, errors.New("кол-во символов в названии мало")
	}

	team = teamInputWithIdToTeam(ctx, team, teamInput.Input)
	teamPayload, err := d.TeamsRepo.TeamsUpdate(team)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamPayload, nil
}

func (d *Domain) TeamPublish(id string) (*models.TeamPayload, error) {
	team, err := d.TeamsRepo.GetTeamsByID(id)
	if err != nil || team == nil {
		return nil, errors.New("филиала не существует")
	}

	teamPayload, err := d.TeamsRepo.TeamsPublish(team)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamPayload, nil
}

func (d *Domain) TeamRestore(id string) (*models.TeamPayload, error) {
	team, err := d.TeamsRepo.GetTeamsByID(id)
	if err != nil || team == nil {
		return nil, errors.New("филиала не существует")
	}

	teamPayload, err := d.TeamsRepo.TeamsRestore(team)
	if err != nil {
		return nil, fmt.Errorf("ошибка при обновлении филиала: %v", err)
	}

	return teamPayload, nil
}

// TeamDelete is the resolver for the teamDelete field.
func (d *Domain) TeamDelete(id string) (bool, error) {
	team, err := d.TeamsRepo.GetTeamsByID(id)
	if err != nil || team == nil {
		return false, errors.New("филиала не существует")
	}

	err = d.TeamsRepo.TeamsDelete(team)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup %v", err)
	}

	return true, nil
}
