package ui

const (
	rows           = 21
	cols           = 10
	plusminusrange = 42
)

var grid = make(map[int]map[int]map[int]string) // tab, row, col, name

func initGrid() {
	for t := 0; t < tabs; t++ {
		grid[t] = make(map[int]map[int]string)
		for r := 0; r < rows; r++ {
			grid[t][r] = make(map[int]string)
		}
	}

	//Tab 0
	grid[0][2][9] = "auerbach/Auerbach"

	grid[0][0][10] = "auerbach/Bergen"
	grid[0][1][10] = "auerbach/Falkenstein"

	//Tab 1
	grid[1][2][8] = "torgau-delitzsch/Dommitzsch"
	grid[1][3][8] = "torgau-delitzsch/Polbitz"

	grid[1][2][9] = "torgau-delitzsch/Drebligar"
	grid[1][3][9] = "torgau-delitzsch/Elsnig"
	grid[1][5][9] = "torgau-delitzsch/Mockritz"
	grid[1][6][9] = "torgau-delitzsch/Rosenfeld"

	grid[1][4][10] = "torgau-delitzsch/Neiden"
	grid[1][5][10] = "torgau-delitzsch/Döbern bei Neiden"
	grid[1][6][10] = "torgau-delitzsch/Zwethau"
	grid[1][8][10] = "bad-liebenwerda/Döbrichau-Löhsten"

	grid[1][5][11] = "torgau-delitzsch/Welsau"
	grid[1][6][11] = "torgau-delitzsch/Kreischau-Eulenau"
	grid[1][7][11] = "torgau-delitzsch/Beilrode"

	grid[1][5][12] = "torgau-delitzsch/Torgau St. Marien"

	grid[1][2][13] = "torgau-delitzsch/Großwig"
	grid[1][3][13] = "torgau-delitzsch/Süptitz"
	grid[1][4][13] = "torgau-delitzsch/Zinna"
	grid[1][5][13] = "torgau-delitzsch/Torgau Garnisonsgemeinde"
	grid[1][6][13] = "torgau-delitzsch/Torgau Schmerzhafte Mutter"

	grid[1][3][14] = "torgau-delitzsch/Melpitz"
	grid[1][6][14] = "torgau-delitzsch/Loßwig"
	grid[1][7][14] = "torgau-delitzsch/Triestewitz"
	grid[1][8][14] = "torgau-delitzsch/Arzberg"

	grid[1][2][15] = "torgau-delitzsch/Klitzschen"
	grid[1][7][15] = "torgau-delitzsch/Weßnig"
	grid[1][9][15] = "torgau-delitzsch/Blumberg"

	grid[1][2][16] = "torgau-delitzsch/Langenreichenbach"
	grid[1][4][16] = "torgau-delitzsch/Staupitz"
	grid[1][5][16] = "torgau-delitzsch/Beckwitz"

	grid[1][6][17] = "torgau-delitzsch/Taura"
	grid[1][8][17] = "torgau-delitzsch/Belgern"

	grid[1][2][18] = "torgau-delitzsch/Kobershain"
	grid[1][3][18] = "torgau-delitzsch/Schildau"
	grid[1][5][18] = "torgau-delitzsch/Sitzenroda"
	grid[1][7][18] = "torgau-delitzsch/Lausa"
	grid[1][8][18] = "torgau-delitzsch/Neußen"
	grid[1][9][18] = "torgau-delitzsch/Staritz"

	grid[1][9][19] = "torgau-delitzsch/Schirmenitz"

	grid[1][9][20] = "torgau-delitzsch/Paußnitz"

	//Tab 2
	grid[2][0][0] = "meissen/Bloßwitz"
	grid[2][1][0] = "meissen/Staucha"
	grid[2][2][0] = "meissen/Striegnitz"
	grid[2][3][0] = "meissen/Dörschnitz"
	grid[2][5][0] = "meissen/Großdobritz"
	grid[2][8][0] = "dresden/Großdittmannsdorf"

	grid[2][1][1] = "meissen/Lommatzsch"
	grid[2][2][1] = "meissen/Zehren"
	grid[2][3][1] = "meissen/Zadel"
	grid[2][4][1] = "meissen/Gröbern"
	grid[2][5][1] = "meissen/Oberau"
	grid[2][8][1] = "dresden/Medingen"
	grid[2][9][1] = "dresden/Ottendorf"

	grid[2][0][2] = "meissen/Neckanitz"
	grid[2][2][2] = "meissen/Meißen St. Afra"
	grid[2][3][2] = "meissen/Meißen Trinitatiskirche"
	grid[2][4][2] = "meissen/Niederau"
	grid[2][7][2] = "dresden/Grünberg"
	grid[2][8][2] = "dresden/Seifersdorf"
	grid[2][9][2] = "dresden/Wachau"

	grid[2][0][3] = "meissen/Leuben"
	grid[2][1][3] = "meissen/Planitz"
	grid[2][2][3] = "meissen/Meißen"
	grid[2][3][3] = "meissen/Meißen Frauenkirche"
	grid[2][4][3] = "meissen/Meißen Johanneskirche"
	grid[2][5][3] = "meissen/Weinböhla"
	grid[2][8][3] = "dresden/Lausa"
	grid[2][9][3] = "dresden/Schönborn"

	grid[2][1][4] = "meissen/Ziegenhain"
	grid[2][4][4] = "meissen/Brockwitz"
	grid[2][5][4] = "meissen/Coswig"
	grid[2][6][4] = "dresden/Reichenberg"
	grid[2][7][4] = "dresden/Wilschdorf"
	grid[2][8][4] = "dresden/Radeberg"
	grid[2][9][4] = "dresden/Langebrück"

	grid[2][2][5] = "meissen/Krögis"
	grid[2][3][5] = "meissen/Miltitz"
	grid[2][5][5] = "meissen/Naustadt"
	grid[2][6][5] = "dresden/Kötzschenbroda"
	grid[2][7][5] = "dresden/Klotzsche"
	grid[2][8][5] = "dresden/Dresden Neustadt"
	grid[2][9][5] = "dresden/Kleinwolmsdorf"

	grid[2][0][6] = "meissen/Rüsseina"
	grid[2][1][6] = "meissen/Raußlitz"
	grid[2][2][6] = "meissen/Heynitz"
	grid[2][3][6] = "meissen/Taubenheim"
	grid[2][4][6] = "meissen/Röhrsdorf"
	grid[2][5][6] = "meissen/Constappel"
	grid[2][6][6] = "dresden/Dresden Friedrichstadt St. Michael"
	grid[2][7][6] = "dresden/Kaditz"
	grid[2][8][6] = "dresden/Dresden Dreikönigskirche"
	grid[2][9][6] = "dresden/Großerkmannsdorf"

	grid[2][0][7] = "meissen/Wendischbora"
	grid[2][1][7] = "meissen/Rothschönberg"
	grid[2][2][7] = "meissen/Burkhardswalde"
	grid[2][3][7] = "meissen/Sora"
	grid[2][4][7] = "meissen/Weistropp"
	grid[2][5][7] = "dresden/Briesnitz"
	grid[2][6][7] = "dresden/Dresden Friedrichstadt"
	grid[2][7][7] = "dresden/Dresden Sophienkirche"
	grid[2][8][7] = "dresden/Dresden Hofkirche"
	grid[2][9][7] = "dresden/Weißig"

	grid[2][0][8] = "meissen/Nossen"
	grid[2][1][8] = "meissen/Deutschenbora"
	grid[2][2][8] = "meissen/Tanneberg"
	grid[2][3][8] = "meissen/Limbach"
	grid[2][4][8] = "meissen/Wilsdruff"
	grid[2][5][8] = "meissen/Unkersdorf"
	grid[2][6][8] = "dresden/Dresden Kreuzkirche"
	grid[2][7][8] = "dresden/Dresden Frauenkirche"
	grid[2][8][8] = "dresden/Loschwitz"

	grid[2][0][9] = "meissen/Siebenlehn"
	grid[2][1][9] = "meissen/Hirschfeld"
	grid[2][2][9] = "meissen/Neukirchen"
	grid[2][3][9] = "meissen/Blankenstein"
	grid[2][4][9] = "meissen/Grumbach"
	grid[2][5][9] = "meissen/Kesselsdorf"
	grid[2][6][9] = "dippoldiswalde/Pesterwitz"
	grid[2][7][9] = "dresden/Dresden Annenkirche"
	grid[2][8][9] = "dresden/Dresden Böhmische Exulantengemeinde"
	grid[2][9][9] = "dresden/Schönfeld"

	grid[2][0][10] = "meissen/Obergruna"
	grid[2][1][10] = "meissen/Bieberstein"
	grid[2][2][10] = "meissen/Reinsberg"
	grid[2][3][10] = "meissen/Dittmannsdorf"
	grid[2][4][10] = "meissen/Herzogswalde"
	grid[2][5][10] = "dippoldiswalde/Döhlen"
	grid[2][6][10] = "dresden/Plauen"
	grid[2][7][10] = "dresden/Leubnitz"
	grid[2][8][10] = "dresden/Leuben"
	grid[2][9][10] = "dresden/Hosterwitz"

	grid[2][0][11] = "freiberg/Großschirma"
	grid[2][1][11] = "freiberg/Krummenhennersdorf"
	grid[2][2][11] = "freiberg/Niederschöna"
	grid[2][3][11] = "meissen/Mohorn"
	grid[2][4][11] = "dippoldiswalde/Fördergersdorf"
	grid[2][5][11] = "dippoldiswalde/Tharandt"
	grid[2][6][11] = "dippoldiswalde/Deuben"
	grid[2][8][11] = "dresden/Lockwitz"

	grid[2][0][12] = "freiberg/Langhennersdorf"
	grid[2][1][12] = "freiberg/Tuttendorf"
	grid[2][2][12] = "freiberg/Conradsdorf"
	grid[2][3][12] = "freiberg/Naundorf"
	grid[2][4][12] = "dippoldiswalde/Dorfhain"
	grid[2][5][12] = "dippoldiswalde/Somsdorf"
	grid[2][6][12] = "dippoldiswalde/Rabenau"
	grid[2][7][12] = "dippoldiswalde/Possendorf"
	grid[2][8][12] = "dresden/Röhrsdorf"

	grid[2][0][13] = "freiberg/Bräunsdorf"
	grid[2][1][13] = "freiberg/Freiberg Dom St. Marien"
	grid[2][2][13] = "freiberg/Freiberg St. Nikolai"
	grid[2][4][13] = "dippoldiswalde/Klingenberg"
	grid[2][5][13] = "dippoldiswalde/Höckendorf"
	grid[2][6][13] = "dippoldiswalde/Seifersdorf"
	grid[2][8][13] = "dippoldiswalde/Kreischa"

	grid[2][0][14] = "freiberg/Kleinwaltersdorf"
	grid[2][1][14] = "freiberg/Freiberg St. Petri"
	grid[2][2][14] = "freiberg/Freiberg St. Jacobi"
	grid[2][3][14] = "freiberg/Hilbersdorf"
	grid[2][4][14] = "dippoldiswalde/Colmnitz"
	grid[2][5][14] = "dippoldiswalde/Ruppendorf"
	grid[2][6][14] = "dippoldiswalde/Dippoldiswalde"
	grid[2][7][14] = "dippoldiswalde/Reinhardtsgrimma"

	grid[2][0][15] = "freiberg/Kleinschirma"
	grid[2][1][15] = "freiberg/Freiberg St. Johannis"
	grid[2][2][15] = "freiberg/Freiberg"
	grid[2][3][15] = "freiberg/Niederbobritzsch"
	grid[2][6][15] = "dippoldiswalde/Reichstädt"

	grid[2][0][16] = "freiberg/Oberschöna"
	grid[2][1][16] = "freiberg/Erbisdorf"
	grid[2][2][16] = "freiberg/Berthelsdorf"
	grid[2][3][16] = "freiberg/Weißenborn"
	grid[2][4][16] = "freiberg/Oberbobritzsch"
	grid[2][5][16] = "dippoldiswalde/Pretzschendorf"
	grid[2][8][16] = "dippoldiswalde/Glashütte"

	grid[2][0][17] = "freiberg/Langenau"
	grid[2][1][17] = "freiberg/Weigmannsdorf"
	grid[2][2][17] = "freiberg/Lichtenberg"
	grid[2][3][17] = "dippoldiswalde/Burkersdorf"
	grid[2][4][17] = "dippoldiswalde/Hartmannsdorf"
	grid[2][5][17] = "dippoldiswalde/Hennersdorf"
	grid[2][6][17] = "dippoldiswalde/Sadisdorf"
	grid[2][7][17] = "dippoldiswalde/Schmiedeberg"
	grid[2][8][17] = "dippoldiswalde/Johnsbach"
	grid[2][9][17] = "dippoldiswalde/Dittersdorf"

	grid[2][0][18] = "freiberg/Gränitz"
	grid[2][1][18] = "freiberg/Großhartmannsdorf"
	grid[2][2][18] = "freiberg/Helbigsdorf"
	grid[2][3][18] = "freiberg/Mulda"
	grid[2][4][18] = "dippoldiswalde/Frauenstein"
	grid[2][5][18] = "dippoldiswalde/Schönfeld"
	grid[2][7][18] = "dippoldiswalde/Bärenstein"
	grid[2][8][18] = "dippoldiswalde/Lauenstein"
	grid[2][9][18] = "dippoldiswalde/Liebenau"

	grid[2][1][19] = "freiberg/Zethau"
	grid[2][2][19] = "freiberg/Dorfchemnitz"
	grid[2][3][19] = "dippoldiswalde/Dittersbach"
	grid[2][4][19] = "dippoldiswalde/Nassau"
	grid[2][5][19] = "dippoldiswalde/Hermsdorf"
	grid[2][6][19] = "dippoldiswalde/Schellerhau"
	grid[2][7][19] = "dippoldiswalde/Altenberg"
	grid[2][8][19] = "dippoldiswalde/Geising"
	grid[2][9][19] = "dippoldiswalde/Fürstenwalde"

	grid[2][1][20] = "freiberg/Voigtsdorf"
	grid[2][2][20] = "freiberg/Sayda"
	grid[2][3][20] = "freiberg/Clausnitz"
	grid[2][4][20] = "freiberg/Cämmerswalde"
	grid[2][8][20] = "dippoldiswalde/Fürstenau"

	//Tab 3
	grid[3][2][8] = "bautzen/Bautzen"
	grid[3][3][8] = "bautzen/Bautzen Dom"

	grid[3][2][9] = "bautzen/Bautzen St. Michael"
	grid[3][3][9] = "bautzen/Bautzen St. Petri"
}
