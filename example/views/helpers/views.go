package helpers

import "path"

func AssetURL(file string) string {
	return path.Join("/assets/", file)
}
