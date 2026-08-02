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
	grid[0][2][5] = "auerbach/Irfersgrün"

	grid[0][1][6] = "auerbach/Waldkirchen"

	grid[0][1][7] = "auerbach/Lengenfeld"
	grid[0][2][7] = "auerbach/Plohn"

	grid[0][0][8] = "auerbach/Treuen"
	grid[0][2][8] = "auerbach/Rodewisch"
	grid[0][3][8] = "auerbach/Rothenkirchen"

	grid[0][2][9] = "auerbach/Auerbach"

	grid[0][0][10] = "auerbach/Bergen"
	grid[0][1][10] = "auerbach/Falkenstein"
	grid[0][3][10] = "auerbach/Rautenkranz"

	grid[0][0][11] = "auerbach/Werda"

	grid[0][2][12] = "auerbach/Zwota"
	grid[0][3][12] = "auerbach/Klingenthal"

	//Tab 1
	grid[1][9][0] = "torgau-delitzsch/Söllichau"

	grid[1][8][1] = "torgau-delitzsch/Bad Düben Garnison"
	grid[1][9][1] = "torgau-delitzsch/Bad Düben St. Nikolai"

	grid[1][9][2] = "torgau-delitzsch/Pristäblich"

	grid[1][9][4] = "torgau-delitzsch/Gruna"

	grid[1][9][5] = "torgau-delitzsch/Mörtitz"

	grid[1][9][6] = "torgau-delitzsch/Eilenburg St. Nikolai"

	grid[1][9][7] = "torgau-delitzsch/Eilenburg St. Marien"

	//Tab 2
	grid[2][2][6] = "torgau-delitzsch/Greudnitz"

	grid[2][2][7] = "torgau-delitzsch/Wörblitz"

	grid[2][1][8] = "torgau-delitzsch/Dahlenberg"
	grid[2][3][8] = "torgau-delitzsch/Dommitzsch"
	grid[2][4][8] = "torgau-delitzsch/Polbitz"

	grid[2][1][9] = "torgau-delitzsch/Falkenberg"
	grid[2][2][9] = "torgau-delitzsch/Trossin"
	grid[2][3][9] = "torgau-delitzsch/Drebligar"
	grid[2][4][9] = "torgau-delitzsch/Elsnig"
	grid[2][5][9] = "torgau-delitzsch/Mockritz"
	grid[2][6][9] = "torgau-delitzsch/Rosenfeld"

	grid[2][0][10] = "torgau-delitzsch/Authausen"
	grid[2][2][10] = "torgau-delitzsch/Roitzsch"
	grid[2][4][10] = "torgau-delitzsch/Neiden"
	grid[2][5][10] = "torgau-delitzsch/Döbern bei Neiden"
	grid[2][6][10] = "torgau-delitzsch/Zwethau"
	grid[2][8][10] = "bad-liebenwerda/Döbrichau-Löhsten"

	grid[2][5][11] = "torgau-delitzsch/Welsau"
	grid[2][6][11] = "torgau-delitzsch/Kreischau-Eulenau"
	grid[2][7][11] = "torgau-delitzsch/Beilrode"

	grid[2][4][12] = "torgau-delitzsch/Zinna"
	grid[2][5][12] = "torgau-delitzsch/Torgau St. Marien"

	grid[2][0][13] = "torgau-delitzsch/Wöllnau"
	grid[2][2][13] = "torgau-delitzsch/Weidenhain"
	grid[2][3][13] = "torgau-delitzsch/Großwig"
	grid[2][4][13] = "torgau-delitzsch/Süptitz"
	grid[2][5][13] = "torgau-delitzsch/Torgau Garnisonsgemeinde"
	grid[2][6][13] = "torgau-delitzsch/Torgau Schmerzhafte Mutter"

	grid[2][0][14] = "torgau-delitzsch/Battaune"
	grid[2][1][14] = "torgau-delitzsch/Wildenhain"
	grid[2][4][14] = "torgau-delitzsch/Melpitz"
	grid[2][6][14] = "torgau-delitzsch/Loßwig"
	grid[2][7][14] = "torgau-delitzsch/Triestewitz"
	grid[2][8][14] = "torgau-delitzsch/Arzberg"

	grid[2][0][15] = "torgau-delitzsch/Doberschütz"
	grid[2][1][15] = "torgau-delitzsch/Mockrehna"
	grid[2][2][15] = "torgau-delitzsch/Niederaudenhain"
	grid[2][3][15] = "torgau-delitzsch/Klitzschen"
	grid[2][7][15] = "torgau-delitzsch/Weßnig"
	grid[2][9][15] = "torgau-delitzsch/Blumberg"

	grid[2][0][16] = "torgau-delitzsch/Sprotta"
	grid[2][1][16] = "torgau-delitzsch/Oberaudenhain"
	grid[2][2][16] = "torgau-delitzsch/Audenhain"
	grid[2][3][16] = "torgau-delitzsch/Langenreichenbach"
	grid[2][4][16] = "torgau-delitzsch/Staupitz"
	grid[2][5][16] = "torgau-delitzsch/Beckwitz"

	grid[2][1][17] = "torgau-delitzsch/Strelln"
	grid[2][2][17] = "torgau-delitzsch/Wildschütz"
	grid[2][6][17] = "torgau-delitzsch/Taura"
	grid[2][8][17] = "torgau-delitzsch/Belgern"

	grid[2][0][18] = "torgau-delitzsch/Paschwitz"
	grid[2][1][18] = "torgau-delitzsch/Schöna"
	grid[2][2][18] = "torgau-delitzsch/Kobershain"
	grid[2][4][18] = "torgau-delitzsch/Schildau"
	grid[2][5][18] = "torgau-delitzsch/Sitzenroda"
	grid[2][7][18] = "torgau-delitzsch/Lausa"
	grid[2][8][18] = "torgau-delitzsch/Neußen"
	grid[2][9][18] = "torgau-delitzsch/Staritz"

	grid[2][9][19] = "torgau-delitzsch/Schirmenitz"

	grid[2][9][20] = "torgau-delitzsch/Paußnitz"

	//Tab 3
	grid[3][0][0] = "meissen/Bloßwitz"
	grid[3][1][0] = "meissen/Staucha"
	grid[3][2][0] = "meissen/Striegnitz"
	grid[3][3][0] = "meissen/Dörschnitz"
	grid[3][5][0] = "meissen/Großdobritz"
	grid[3][8][0] = "dresden/Großdittmannsdorf"

	grid[3][1][1] = "meissen/Lommatzsch"
	grid[3][2][1] = "meissen/Zehren"
	grid[3][3][1] = "meissen/Zadel"
	grid[3][4][1] = "meissen/Gröbern"
	grid[3][5][1] = "meissen/Oberau"
	grid[3][8][1] = "dresden/Medingen"
	grid[3][9][1] = "dresden/Ottendorf"

	grid[3][0][2] = "meissen/Neckanitz"
	grid[3][2][2] = "meissen/Meißen St. Afra"
	grid[3][3][2] = "meissen/Meißen Trinitatiskirche"
	grid[3][4][2] = "meissen/Niederau"
	grid[3][7][2] = "dresden/Grünberg"
	grid[3][8][2] = "dresden/Seifersdorf"
	grid[3][9][2] = "dresden/Wachau"

	grid[3][0][3] = "meissen/Leuben"
	grid[3][1][3] = "meissen/Planitz"
	grid[3][2][3] = "meissen/Meißen"
	grid[3][3][3] = "meissen/Meißen Frauenkirche"
	grid[3][4][3] = "meissen/Meißen Johanneskirche"
	grid[3][5][3] = "meissen/Weinböhla"
	grid[3][8][3] = "dresden/Lausa"
	grid[3][9][3] = "dresden/Schönborn"

	grid[3][1][4] = "meissen/Ziegenhain"
	grid[3][4][4] = "meissen/Brockwitz"
	grid[3][5][4] = "meissen/Coswig"
	grid[3][6][4] = "dresden/Reichenberg"
	grid[3][7][4] = "dresden/Wilschdorf"
	grid[3][8][4] = "dresden/Radeberg"
	grid[3][9][4] = "dresden/Langebrück"

	grid[3][2][5] = "meissen/Krögis"
	grid[3][3][5] = "meissen/Miltitz"
	grid[3][5][5] = "meissen/Naustadt"
	grid[3][6][5] = "dresden/Kötzschenbroda"
	grid[3][7][5] = "dresden/Klotzsche"
	grid[3][8][5] = "dresden/Dresden Neustadt"
	grid[3][9][5] = "dresden/Kleinwolmsdorf"

	grid[3][0][6] = "meissen/Rüsseina"
	grid[3][1][6] = "meissen/Raußlitz"
	grid[3][2][6] = "meissen/Heynitz"
	grid[3][3][6] = "meissen/Taubenheim"
	grid[3][4][6] = "meissen/Röhrsdorf"
	grid[3][5][6] = "meissen/Constappel"
	grid[3][6][6] = "dresden/Dresden Friedrichstadt St. Michael"
	grid[3][7][6] = "dresden/Kaditz"
	grid[3][8][6] = "dresden/Dresden Dreikönigskirche"
	grid[3][9][6] = "dresden/Großerkmannsdorf"

	grid[3][0][7] = "meissen/Wendischbora"
	grid[3][1][7] = "meissen/Rothschönberg"
	grid[3][2][7] = "meissen/Burkhardswalde"
	grid[3][3][7] = "meissen/Sora"
	grid[3][4][7] = "meissen/Weistropp"
	grid[3][5][7] = "dresden/Briesnitz"
	grid[3][6][7] = "dresden/Dresden Friedrichstadt"
	grid[3][7][7] = "dresden/Dresden Sophienkirche"
	grid[3][8][7] = "dresden/Dresden Hofkirche"
	grid[3][9][7] = "dresden/Weißig"

	grid[3][0][8] = "meissen/Nossen"
	grid[3][1][8] = "meissen/Deutschenbora"
	grid[3][2][8] = "meissen/Tanneberg"
	grid[3][3][8] = "meissen/Limbach"
	grid[3][4][8] = "meissen/Wilsdruff"
	grid[3][5][8] = "meissen/Unkersdorf"
	grid[3][6][8] = "dresden/Dresden Kreuzkirche"
	grid[3][7][8] = "dresden/Dresden Frauenkirche"
	grid[3][8][8] = "dresden/Loschwitz"

	grid[3][0][9] = "meissen/Siebenlehn"
	grid[3][1][9] = "meissen/Hirschfeld"
	grid[3][2][9] = "meissen/Neukirchen"
	grid[3][3][9] = "meissen/Blankenstein"
	grid[3][4][9] = "meissen/Grumbach"
	grid[3][5][9] = "meissen/Kesselsdorf"
	grid[3][6][9] = "dippoldiswalde/Pesterwitz"
	grid[3][7][9] = "dresden/Dresden Annenkirche"
	grid[3][8][9] = "dresden/Dresden Böhmische Exulantengemeinde"
	grid[3][9][9] = "dresden/Schönfeld"

	grid[3][0][10] = "meissen/Obergruna"
	grid[3][1][10] = "meissen/Bieberstein"
	grid[3][2][10] = "meissen/Reinsberg"
	grid[3][3][10] = "meissen/Dittmannsdorf"
	grid[3][4][10] = "meissen/Herzogswalde"
	grid[3][5][10] = "dippoldiswalde/Döhlen"
	grid[3][6][10] = "dresden/Plauen"
	grid[3][7][10] = "dresden/Leubnitz"
	grid[3][8][10] = "dresden/Leuben"
	grid[3][9][10] = "dresden/Hosterwitz"

	grid[3][0][11] = "freiberg/Großschirma"
	grid[3][1][11] = "freiberg/Krummenhennersdorf"
	grid[3][2][11] = "freiberg/Niederschöna"
	grid[3][3][11] = "meissen/Mohorn"
	grid[3][4][11] = "dippoldiswalde/Fördergersdorf"
	grid[3][5][11] = "dippoldiswalde/Tharandt"
	grid[3][6][11] = "dippoldiswalde/Deuben"
	grid[3][8][11] = "dresden/Lockwitz"

	grid[3][0][12] = "freiberg/Langhennersdorf"
	grid[3][1][12] = "freiberg/Tuttendorf"
	grid[3][2][12] = "freiberg/Conradsdorf"
	grid[3][3][12] = "freiberg/Naundorf"
	grid[3][4][12] = "dippoldiswalde/Dorfhain"
	grid[3][5][12] = "dippoldiswalde/Somsdorf"
	grid[3][6][12] = "dippoldiswalde/Rabenau"
	grid[3][7][12] = "dippoldiswalde/Possendorf"
	grid[3][8][12] = "dresden/Röhrsdorf"

	grid[3][0][13] = "freiberg/Bräunsdorf"
	grid[3][1][13] = "freiberg/Freiberg Dom St. Marien"
	grid[3][2][13] = "freiberg/Freiberg St. Nikolai"
	grid[3][4][13] = "dippoldiswalde/Klingenberg"
	grid[3][5][13] = "dippoldiswalde/Höckendorf"
	grid[3][6][13] = "dippoldiswalde/Seifersdorf"
	grid[3][8][13] = "dippoldiswalde/Kreischa"

	grid[3][0][14] = "freiberg/Kleinwaltersdorf"
	grid[3][1][14] = "freiberg/Freiberg St. Petri"
	grid[3][2][14] = "freiberg/Freiberg St. Jacobi"
	grid[3][3][14] = "freiberg/Hilbersdorf"
	grid[3][4][14] = "dippoldiswalde/Colmnitz"
	grid[3][5][14] = "dippoldiswalde/Ruppendorf"
	grid[3][6][14] = "dippoldiswalde/Dippoldiswalde"
	grid[3][7][14] = "dippoldiswalde/Reinhardtsgrimma"

	grid[3][0][15] = "freiberg/Kleinschirma"
	grid[3][1][15] = "freiberg/Freiberg St. Johannis"
	grid[3][2][15] = "freiberg/Freiberg"
	grid[3][3][15] = "freiberg/Niederbobritzsch"
	grid[3][6][15] = "dippoldiswalde/Reichstädt"

	grid[3][0][16] = "freiberg/Oberschöna"
	grid[3][1][16] = "freiberg/Erbisdorf"
	grid[3][2][16] = "freiberg/Berthelsdorf"
	grid[3][3][16] = "freiberg/Weißenborn"
	grid[3][4][16] = "freiberg/Oberbobritzsch"
	grid[3][5][16] = "dippoldiswalde/Pretzschendorf"
	grid[3][8][16] = "dippoldiswalde/Glashütte"

	grid[3][0][17] = "freiberg/Langenau"
	grid[3][1][17] = "freiberg/Weigmannsdorf"
	grid[3][2][17] = "freiberg/Lichtenberg"
	grid[3][3][17] = "dippoldiswalde/Burkersdorf"
	grid[3][4][17] = "dippoldiswalde/Hartmannsdorf"
	grid[3][5][17] = "dippoldiswalde/Hennersdorf"
	grid[3][6][17] = "dippoldiswalde/Sadisdorf"
	grid[3][7][17] = "dippoldiswalde/Schmiedeberg"
	grid[3][8][17] = "dippoldiswalde/Johnsbach"
	grid[3][9][17] = "dippoldiswalde/Dittersdorf"

	grid[3][0][18] = "freiberg/Gränitz"
	grid[3][1][18] = "freiberg/Großhartmannsdorf"
	grid[3][2][18] = "freiberg/Helbigsdorf"
	grid[3][3][18] = "freiberg/Mulda"
	grid[3][4][18] = "dippoldiswalde/Frauenstein"
	grid[3][5][18] = "dippoldiswalde/Schönfeld"
	grid[3][7][18] = "dippoldiswalde/Bärenstein"
	grid[3][8][18] = "dippoldiswalde/Lauenstein"
	grid[3][9][18] = "dippoldiswalde/Liebenau"

	grid[3][1][19] = "freiberg/Zethau"
	grid[3][2][19] = "freiberg/Dorfchemnitz"
	grid[3][3][19] = "dippoldiswalde/Dittersbach"
	grid[3][4][19] = "dippoldiswalde/Nassau"
	grid[3][5][19] = "dippoldiswalde/Hermsdorf"
	grid[3][6][19] = "dippoldiswalde/Schellerhau"
	grid[3][7][19] = "dippoldiswalde/Altenberg"
	grid[3][8][19] = "dippoldiswalde/Geising"
	grid[3][9][19] = "dippoldiswalde/Fürstenwalde"

	grid[3][1][20] = "freiberg/Voigtsdorf"
	grid[3][2][20] = "freiberg/Sayda"
	grid[3][3][20] = "freiberg/Clausnitz"
	grid[3][4][20] = "freiberg/Cämmerswalde"
	grid[3][8][20] = "dippoldiswalde/Fürstenau"

	//Tab 4
	grid[4][2][8] = "bautzen/Bautzen"
	grid[4][3][8] = "bautzen/Bautzen Dom"

	grid[4][2][9] = "bautzen/Bautzen St. Michael"
	grid[4][3][9] = "bautzen/Bautzen St. Petri"
}
