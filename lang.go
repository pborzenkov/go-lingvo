//go:generate stringer -type=Lang
package lingvo

// Lang is a language to translate to/from
type Lang int

// List of supported languages
const (
	Ch Lang = 1028
	Da Lang = 1030
	De Lang = 1031
	El Lang = 1032
	En Lang = 1033
	Es Lang = 1034
	Fr Lang = 1036
	It Lang = 1040
	Pl Lang = 1045
	Ru Lang = 1049
	Uk Lang = 1058
	Kk Lang = 1087
	Tt Lang = 1092
	La Lang = 1142
)
