package charts

type BarGroupTheme string

const (
	Walden      BarGroupTheme = "walden"
	Essos       BarGroupTheme = "essos"
	Infographic BarGroupTheme = "infographic"
	Shine       BarGroupTheme = "shine"
	Macarons    BarGroupTheme = "macarons"
	Roma        BarGroupTheme = "roma"
)

const defaultBarGroupTheme = Walden

var barGroupThemes = map[BarGroupTheme][]string{
	Walden: {
		"#3fb1e3",
		"#6be6c1",
		"#626c91",
		"#a0a7e6",
		"#c4ebad",
		"#96dee8",
	},
	Essos: {
		"#893448",
		"#d95850",
		"#eb8146",
		"#ffb248",
		"#f2d643",
		"#ebdba4",
	},
	Infographic: {
		"#c1232b",
		"#27727b",
		"#fcce10",
		"#e87c25",
		"#b5c334",
		"#fe8463",
		"#9bca63",
		"#fad860",
		"#f3a43b",
		"#60c0dd",
		"#d7504b",
		"#c6e579",
		"#f4e001",
		"#f0805a",
		"#26c0c0",
	},
	Shine: {
		"#c12e34",
		"#e6b600",
		"#0098d9",
		"#2b821d",
		"#005eaa",
		"#339ca8",
		"#cda819",
		"#32a487",
	},
	Macarons: {
		"#2ec7c9",
		"#b6a2de",
		"#5ab1ef",
		"#ffb980",
		"#d87a80",
		"#8d98b3",
		"#e5cf0d",
		"#97b552",
		"#95706d",
		"#dc69aa",
		"#07a2a4",
		"#9a7fd1",
		"#588dd5",
		"#f5994e",
		"#c05050",
		"#59678c",
		"#c9ab00",
		"#7eb00a",
		"#6f5553",
		"#c14089",
	},
	Roma: {
		"#e01f54",
		"#001852",
		"#f5e8c8",
		"#b8d2c7",
		"#c6b38e",
		"#a4d8c2",
		"#f3d999",
		"#d3758f",
		"#dcc392",
		"#2e4783",
		"#82b6e9",
		"#ff6347",
		"#a092f1",
		"#0a915d",
		"#eaf889",
		"#6699FF",
		"#ff6666",
		"#3cb371",
		"#d5b158",
		"#38b6b6",
	},
}
