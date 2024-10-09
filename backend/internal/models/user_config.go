package models

type UserConfig struct {
	Id          uint64   `json:"id"`
	Name        string   `json:"name"`
	Phones      []string `json:"phones"`
	Description string   `json:"description,omitempty"`
}

func (c *UserConfig) IsValid() bool {
	if len(c.Name) == 0 || len(c.Phones) == 0 {
		return false
	}
	return true
}

func (newUserConfig *UserConfig) Update(oldUserConfig *UserConfig) {
	newUserConfig.Id = oldUserConfig.Id

	if len(newUserConfig.Name) == 0 {
		newUserConfig.Name = oldUserConfig.Name
	}

	if len(newUserConfig.Phones) == 0 {
		newUserConfig.Phones = oldUserConfig.Phones
	}

	if len(newUserConfig.Description) == 0 {
		newUserConfig.Description = oldUserConfig.Description
	}
}
