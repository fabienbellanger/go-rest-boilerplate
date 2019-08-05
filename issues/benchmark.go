package issues

import (
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
)

type sqlDataType struct {
	ApplicationID        uint64
	ApplicationName      string
	ApplicationCreatedAt time.Time
	ApplicationUpdatedAt time.Time
	ModuleID             uint64
	ModuleName           string
	ModuleCreatedAt      time.Time
	ModuleUpdatedAt      time.Time
	ActionID             uint64
	ActionName           string
	ActionCreatedAt      time.Time
	ActionUpdatedAt      time.Time
}

type dataApplicationType struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Modules   map[uint64]dataModuleType
}

type dataModuleType struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Actions   map[uint64]dataActionType
}

type dataActionType struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InitData inits data for test
func InitData() *map[uint64]dataApplicationType {
	// sqlData := constructSQLArray()
	// data := constructFinalArray(sqlData)

	// fmt.Printf("%+v", data)

	data := getSQLData()

	return data
}

func getSQLData() *map[uint64]dataApplicationType {
	query := `
		SELECT 
			applications.id,
			applications.name,
			applications.created_at,
			applications.updated_at,
			modules.id,
			modules.name,
			modules.created_at,
			modules.updated_at,
			actions.id,
			actions.name,
			actions.created_at,
			actions.updated_at
		FROM applications
			INNER JOIN modules ON applications.id = modules.application_id
			INNER JOIN actions ON modules.id = actions.module_id`
	rows, _ := database.Select(query)

	data := make(map[uint64]dataApplicationType)

	var line sqlDataType
	for rows.Next() {
		rows.Scan(
			&line.ApplicationID,
			&line.ApplicationName,
			&line.ApplicationCreatedAt,
			&line.ApplicationUpdatedAt,
			&line.ModuleID,
			&line.ModuleName,
			&line.ModuleCreatedAt,
			&line.ModuleUpdatedAt,
			&line.ActionID,
			&line.ActionName,
			&line.ActionCreatedAt,
			&line.ActionUpdatedAt)

		if _, ok := data[line.ApplicationID]; !ok {
			dac := new(dataActionType)
			dac.ID = line.ActionID
			dac.Name = line.ActionName
			dac.CreatedAt = line.ActionCreatedAt
			dac.UpdatedAt = line.ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.ID = line.ModuleID
			dmo.Name = line.ModuleName
			dmo.CreatedAt = line.ModuleCreatedAt
			dmo.UpdatedAt = line.ModuleUpdatedAt
			dmo.Actions = make(map[uint64]dataActionType)
			dmo.Actions[line.ActionID] = *dac

			dap := new(dataApplicationType)
			dap.ID = line.ApplicationID
			dap.Name = line.ApplicationName
			dap.CreatedAt = line.ApplicationCreatedAt
			dap.UpdatedAt = line.ApplicationUpdatedAt
			dap.Modules = make(map[uint64]dataModuleType)
			dap.Modules[line.ModuleID] = *dmo

			data[line.ApplicationID] = *dap
		}

		if _, ok := data[line.ApplicationID].Modules[line.ModuleID]; !ok {
			dac := new(dataActionType)
			dac.ID = line.ActionID
			dac.Name = line.ActionName
			dac.CreatedAt = line.ActionCreatedAt
			dac.UpdatedAt = line.ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.ID = line.ModuleID
			dmo.Name = line.ModuleName
			dmo.CreatedAt = line.ModuleCreatedAt
			dmo.UpdatedAt = line.ModuleUpdatedAt
			dmo.Actions = make(map[uint64]dataActionType)
			dmo.Actions[line.ActionID] = *dac

			data[line.ApplicationID].Modules[line.ModuleID] = *dmo
		}

		if _, ok := data[line.ApplicationID].Modules[line.ModuleID].Actions[line.ActionID]; !ok {
			dac := new(dataActionType)
			dac.ID = line.ActionID
			dac.Name = line.ActionName
			dac.CreatedAt = line.ActionCreatedAt
			dac.UpdatedAt = line.ActionUpdatedAt

			data[line.ApplicationID].Modules[line.ModuleID].Actions[line.ActionID] = *dac
		}
	}

	return &data
}

func constructSQLArray() []sqlDataType {
	const nbApplications = 10
	const nbModules = 1000
	const nbActions = 100000

	sqlData := make([]sqlDataType, 0)
	for i := 0; i < nbActions; i++ {
		line := sqlDataType{
			ApplicationID:        uint64(i%nbApplications) + 1,
			ApplicationName:      "Application " + strconv.Itoa((i%nbApplications)+1),
			ApplicationCreatedAt: time.Now(),
			ApplicationUpdatedAt: time.Now(),
			ModuleID:             uint64(i%nbModules) + 1,
			ModuleName:           "Module " + strconv.Itoa((i%nbModules)+1),
			ModuleCreatedAt:      time.Now(),
			ModuleUpdatedAt:      time.Now(),
			ActionID:             uint64(i) + 1,
			ActionName:           "Action " + strconv.Itoa(i+1),
			ActionCreatedAt:      time.Now(),
			ActionUpdatedAt:      time.Now(),
		}

		sqlData = append(sqlData, line)
	}

	return sqlData
}

func constructFinalArray(sqlData []sqlDataType) *map[uint64]dataApplicationType {
	data := make(map[uint64]dataApplicationType)

	nbData := uint64(len(sqlData))
	var i uint64
	for i = 0; i < nbData; i++ {
		if _, ok := data[sqlData[i].ApplicationID]; !ok {
			dac := new(dataActionType)
			dac.ID = sqlData[i].ActionID
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.ID = sqlData[i].ModuleID
			dmo.Name = sqlData[i].ModuleName
			dmo.CreatedAt = sqlData[i].ModuleCreatedAt
			dmo.UpdatedAt = sqlData[i].ModuleUpdatedAt
			dmo.Actions = make(map[uint64]dataActionType)
			dmo.Actions[sqlData[i].ActionID] = *dac

			dap := new(dataApplicationType)
			dap.ID = sqlData[i].ApplicationID
			dap.Name = sqlData[i].ApplicationName
			dap.CreatedAt = sqlData[i].ApplicationCreatedAt
			dap.UpdatedAt = sqlData[i].ApplicationUpdatedAt
			dap.Modules = make(map[uint64]dataModuleType)
			dap.Modules[sqlData[i].ModuleID] = *dmo

			data[sqlData[i].ApplicationID] = *dap
		}

		if _, ok := data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID]; !ok {
			dac := new(dataActionType)
			dac.ID = sqlData[i].ActionID
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.ID = sqlData[i].ModuleID
			dmo.Name = sqlData[i].ModuleName
			dmo.CreatedAt = sqlData[i].ModuleCreatedAt
			dmo.UpdatedAt = sqlData[i].ModuleUpdatedAt
			dmo.Actions = make(map[uint64]dataActionType)
			dmo.Actions[sqlData[i].ActionID] = *dac

			data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID] = *dmo
		}

		if _, ok := data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID].Actions[sqlData[i].ActionID]; !ok {
			dac := new(dataActionType)
			dac.ID = sqlData[i].ActionID
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID].Actions[sqlData[i].ActionID] = *dac
		}
	}

	return &data
}
