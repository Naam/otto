package terraform

import (
	"os"

	"github.com/hashicorp/terraform/terraform"
)

// Outputs reads the outputs from the Terraform state at the given path.
//
// This is currently done by using the Terraform API to read from the
// state file directory. In the future, this may work by shelling out to
// Terraform which might be more stable.
func Outputs(path string) (map[string]string, error) {
	// Read the state structure itself
	state, err := readState(path)
	if err != nil {
		return nil, err
	}

	// Return the outputs
	return state.RootModule().Outputs, nil
}

func readState(path string) (*terraform.State, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return terraform.ReadState(f)
}
