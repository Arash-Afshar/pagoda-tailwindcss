package page

type Theme string

const (
	themeLight     Theme = "light"
	themeDark      Theme = "dark"
	themeCupcake   Theme = "cupcake"
	themeBumblebee Theme = "bumblebee"
	themeEmerald   Theme = "emerald"
	themeCorporate Theme = "corporate"
	themeSynthwave Theme = "synthwave"
	themeRetro     Theme = "retro"
	themeCyberpunk Theme = "cyberpunk"
	themeValentine Theme = "valentine"
	themeHalloween Theme = "halloween"
	themeGarden    Theme = "garden"
	themeForest    Theme = "forest"
	themeAqua      Theme = "aqua"
	themeLofi      Theme = "lofi"
	themePastel    Theme = "pastel"
	themeFantasy   Theme = "fantasy"
	themeWireframe Theme = "wireframe"
	themeBlack     Theme = "black"
	themeLuxury    Theme = "luxury"
	themeDracula   Theme = "dracula"
	themeCmyk      Theme = "cmyk"
	themeAutumn    Theme = "autumn"
	themeBusiness  Theme = "business"
	themeAcid      Theme = "acid"
	themeLemonade  Theme = "lemonade"
	themeNight     Theme = "night"
	themeCoffee    Theme = "coffee"
	themeWinter    Theme = "winter"
	themeDim       Theme = "dim"
	themeNord      Theme = "nord"
	themeSunset    Theme = "sunset"
)

func GetAllThemes() []Theme {
	return []Theme{
		themeLight,
		themeDark,
		themeCupcake,
		themeBumblebee,
		themeEmerald,
		themeCorporate,
		themeSynthwave,
		themeRetro,
		themeCyberpunk,
		themeValentine,
		themeHalloween,
		themeGarden,
		themeForest,
		themeAqua,
		themeLofi,
		themePastel,
		themeFantasy,
		themeWireframe,
		themeBlack,
		themeLuxury,
		themeDracula,
		themeCmyk,
		themeAutumn,
		themeBusiness,
		themeAcid,
		themeLemonade,
		themeNight,
		themeCoffee,
		themeWinter,
		themeDim,
		themeNord,
		themeSunset,
	}
}

func defaultTheme() Theme {
	return themeWinter
}
