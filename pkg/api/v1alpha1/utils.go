package v1alpha1

import (
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

// NamespacedNameToString converts a NamespacedName to a string.
func NamespacedNameToString(resource *types.NamespacedName) string {
	return resource.Namespace + "/" + resource.Name
}

// StringToNamespacedName converts a string to a NamespacedName.
func StringToNamespacedName(resource string) types.NamespacedName {
	return types.NamespacedName{
		Namespace: resource[:strings.Index(resource, "/")],
		Name:      resource[strings.Index(resource, "/")+1:],
	}
}
