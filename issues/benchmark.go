package issues

import (
	"strconv"
	"time"
)

type sqlDataType struct {
	ApplicationID        int
	ApplicationName      string
	ApplicationCreatedAt time.Time
	ApplicationUpdatedAt time.Time
	ModuleID             int
	ModuleName           string
	ModuleCreatedAt      time.Time
	ModuleUpdatedAt      time.Time
	ActionID             int
	ActionName           string
	ActionCreatedAt      time.Time
	ActionUpdatedAt      time.Time
}

type DataApplicationType struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Modules   map[int]dataModuleType
}

type dataModuleType struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Actions   map[int]dataActionType
}

type dataActionType struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InitData inits data for test
func InitData() map[int]DataApplicationType {
	data := make(map[int]DataApplicationType)

	const nbApplications = 10
	const nbModules = 1000
	const nbActions = 100000

	sqlData := make([]sqlDataType, 0)
	for i := 0; i < nbActions; i++ {
		line := sqlDataType{
			ApplicationID:        (i % nbApplications) + 1,
			ApplicationName:      "Application " + strconv.Itoa((i%nbApplications)+1),
			ApplicationCreatedAt: time.Now(),
			ApplicationUpdatedAt: time.Now(),
			ModuleID:             (i % nbModules) + 1,
			ModuleName:           "Module " + strconv.Itoa((i%nbModules)+1),
			ModuleCreatedAt:      time.Now(),
			ModuleUpdatedAt:      time.Now(),
			ActionID:             i + 1,
			ActionName:           "Action " + strconv.Itoa(i+1),
			ActionCreatedAt:      time.Now(),
			ActionUpdatedAt:      time.Now(),
		}

		sqlData = append(sqlData, line)
	}

	nbData := len(sqlData)
	for i := 0; i < nbData; i++ {
		if _, ok := data[sqlData[i].ApplicationID]; !ok {
			dac := new(dataActionType)
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.Name = sqlData[i].ModuleName
			dmo.CreatedAt = sqlData[i].ModuleCreatedAt
			dmo.UpdatedAt = sqlData[i].ModuleUpdatedAt
			dmo.Actions = make(map[int]dataActionType)
			dmo.Actions[sqlData[i].ActionID] = *dac

			dap := new(DataApplicationType)
			dap.Name = sqlData[i].ApplicationName
			dap.CreatedAt = sqlData[i].ApplicationCreatedAt
			dap.UpdatedAt = sqlData[i].ApplicationUpdatedAt
			dap.Modules = make(map[int]dataModuleType)
			dap.Modules[sqlData[i].ModuleID] = *dmo

			data[sqlData[i].ApplicationID] = *dap
		}

		if _, ok := data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID]; !ok {
			dac := new(dataActionType)
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			dmo := new(dataModuleType)
			dmo.Name = sqlData[i].ModuleName
			dmo.CreatedAt = sqlData[i].ModuleCreatedAt
			dmo.UpdatedAt = sqlData[i].ModuleUpdatedAt
			dmo.Actions = make(map[int]dataActionType)
			dmo.Actions[sqlData[i].ActionID] = *dac

			data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID] = *dmo
		}

		if _, ok := data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID].Actions[sqlData[i].ActionID]; !ok {
			dac := new(dataActionType)
			dac.Name = sqlData[i].ActionName
			dac.CreatedAt = sqlData[i].ActionCreatedAt
			dac.UpdatedAt = sqlData[i].ActionUpdatedAt

			data[sqlData[i].ApplicationID].Modules[sqlData[i].ModuleID].Actions[sqlData[i].ActionID] = *dac
		}
	}

	return data
}
