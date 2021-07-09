package lycosa

import "errors"

func NotFound(name string) error {
	return errors.New("not found: " + name)
}
