package main

type TableData struct {
	Id      int
	Company string `form:"company"`
	Contact string `form:"contact"`
	Country string `form:"country"`
}

var Tabledata = []TableData{
	{1,
		"Alfreds Futterkiste",
		"Maria Anders",
		"Germany",
	},
	{2,
		"Centro comercial Moctezuma",
		"Francisco Chang",
		"Mexico",
	},
	{3,
		"Ernst Handel",
		"Roland Mendel",
		"Austria",
	},
	{4,
		"Island Trading",
		"Helen Bennett",
		"UK",
	},
	{5,
		"Laughing Bacchus Winecellars",
		"Yoshi Tannamuri",
		"Canada",
	},
	{6,
		"Magazzini Alimentari Riuniti",
		"Giovanni Rovelli",
		"Italy",
	},
}
