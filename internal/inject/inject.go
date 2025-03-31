package inject

// InjectorFunc describes an injector function
type InjectorFunc struct {

	// ImportPath is the import path of the injector function.
	ImportPath string

	// FuncName is the name of the injector function.
	FuncName string

	// ReturnTypes is the number of return values. This does not include the cleanup return value, or error return value, if either exists.
	ReturnTypes int

	// HasCleanup reports whether the injector function has a cleanup return value.
	HasCleanup bool

	// HasError reports whether the injector function has an error return value.
	HasError bool
}

// Injectors is a list of all known injector functions.
var Injectors = []InjectorFunc{
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject",
		ReturnTypes: 1,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "InjectE",
		ReturnTypes: 1,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "InjectC",
		ReturnTypes: 1,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "InjectCE",
		ReturnTypes: 1,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject2",
		ReturnTypes: 2,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject2E",
		ReturnTypes: 2,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject2C",
		ReturnTypes: 2,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject2CE",
		ReturnTypes: 2,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject3",
		ReturnTypes: 3,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject3E",
		ReturnTypes: 3,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject3C",
		ReturnTypes: 3,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject3CE",
		ReturnTypes: 3,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject4",
		ReturnTypes: 4,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject4E",
		ReturnTypes: 4,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject4C",
		ReturnTypes: 4,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject4CE",
		ReturnTypes: 4,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject5",
		ReturnTypes: 5,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject5E",
		ReturnTypes: 5,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject5C",
		ReturnTypes: 5,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject5CE",
		ReturnTypes: 5,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject6",
		ReturnTypes: 6,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject6E",
		ReturnTypes: 6,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject6C",
		ReturnTypes: 6,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject6CE",
		ReturnTypes: 6,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject7",
		ReturnTypes: 7,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject7E",
		ReturnTypes: 7,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject7C",
		ReturnTypes: 7,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject7CE",
		ReturnTypes: 7,
		HasCleanup:  true,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject8",
		ReturnTypes: 8,
		HasCleanup:  false,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject8E",
		ReturnTypes: 8,
		HasCleanup:  false,
		HasError:    true,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject8C",
		ReturnTypes: 8,
		HasCleanup:  true,
		HasError:    false,
	},
	{
		ImportPath:  "github.com/a-jentleman/inject",
		FuncName:    "Inject8CE",
		ReturnTypes: 8,
		HasCleanup:  true,
		HasError:    true,
	},
}

// InjectorsByName is a map of known injector functions by indexed by their (1) import path and (2) function name.
var InjectorsByName = map[string]map[string]*InjectorFunc{
	"github.com/a-jentleman/inject": {
		"Inject":    &Injectors[0],
		"InjectE":   &Injectors[1],
		"InjectC":   &Injectors[2],
		"InjectCE":  &Injectors[3],
		"Inject2":   &Injectors[4],
		"Inject2E":  &Injectors[5],
		"Inject2C":  &Injectors[6],
		"Inject2CE": &Injectors[7],
		"Inject3":   &Injectors[8],
		"Inject3E":  &Injectors[9],
		"Inject3C":  &Injectors[10],
		"Inject3CE": &Injectors[11],
		"Inject4":   &Injectors[12],
		"Inject4E":  &Injectors[13],
		"Inject4C":  &Injectors[14],
		"Inject4CE": &Injectors[15],
		"Inject5":   &Injectors[16],
		"Inject5E":  &Injectors[17],
		"Inject5C":  &Injectors[18],
		"Inject5CE": &Injectors[19],
		"Inject6":   &Injectors[20],
		"Inject6E":  &Injectors[21],
		"Inject6C":  &Injectors[22],
		"Inject6CE": &Injectors[23],
		"Inject7":   &Injectors[24],
		"Inject7E":  &Injectors[25],
		"Inject7C":  &Injectors[26],
		"Inject7CE": &Injectors[27],
		"Inject8":   &Injectors[28],
		"Inject8E":  &Injectors[29],
		"Inject8C":  &Injectors[30],
		"Inject8CE": &Injectors[31],
	},
}
