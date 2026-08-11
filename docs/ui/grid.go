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
	grid[0][1][2] = "torgau-delitzsch/*"
	grid[0][2][2] = "bad-liebenwerda/*"

	grid[0][3][4] = "meissen/*"
	grid[0][4][4] = "dresden/*"
	grid[0][5][4] = "bautzen/*"

	grid[0][3][5] = "freiberg/*"
	grid[0][4][5] = "dippoldiswalde/*"

	grid[0][0][7] = "auerbach/*"

	//Tab 1
	grid[1][2][5] = "auerbach/Irfersgrün"

	grid[1][1][6] = "auerbach/Waldkirchen"

	grid[1][1][7] = "auerbach/Lengenfeld"
	grid[1][2][7] = "auerbach/Plohn"

	grid[1][0][8] = "auerbach/Treuen"
	grid[1][2][8] = "auerbach/Rodewisch"
	grid[1][3][8] = "auerbach/Rothenkirchen"

	grid[1][2][9] = "auerbach/Auerbach"

	grid[1][0][10] = "auerbach/Bergen"
	grid[1][1][10] = "auerbach/Falkenstein"
	grid[1][3][10] = "auerbach/Rautenkranz"

	grid[1][0][11] = "auerbach/Werda"

	grid[1][2][12] = "auerbach/Zwota"
	grid[1][3][12] = "auerbach/Klingenthal"

	//Tab 2
	grid[2][9][0] = "torgau-delitzsch/Söllichau"

	grid[2][8][1] = "torgau-delitzsch/Bad Düben Garnison"
	grid[2][9][1] = "torgau-delitzsch/Bad Düben St. Nikolai"

	grid[2][9][2] = "torgau-delitzsch/Pristäblich"

	grid[2][9][4] = "torgau-delitzsch/Gruna"

	grid[2][9][5] = "torgau-delitzsch/Mörtitz"

	grid[2][9][6] = "torgau-delitzsch/Eilenburg St. Nikolai"

	grid[2][9][7] = "torgau-delitzsch/Eilenburg St. Marien"

	grid[2][9][8] = "torgau-delitzsch/Eilenburg St. Franziskus Xaver"

	//Tab 3
	grid[3][2][6] = "torgau-delitzsch/Greudnitz"

	grid[3][2][7] = "torgau-delitzsch/Wörblitz"

	grid[3][1][8] = "torgau-delitzsch/Dahlenberg"
	grid[3][3][8] = "torgau-delitzsch/Dommitzsch"
	grid[3][4][8] = "torgau-delitzsch/Polbitz"

	grid[3][1][9] = "torgau-delitzsch/Falkenberg"
	grid[3][2][9] = "torgau-delitzsch/Trossin"
	grid[3][3][9] = "torgau-delitzsch/Drebligar"
	grid[3][4][9] = "torgau-delitzsch/Elsnig"
	grid[3][5][9] = "torgau-delitzsch/Mockritz"
	grid[3][6][9] = "torgau-delitzsch/Rosenfeld"

	grid[3][0][10] = "torgau-delitzsch/Authausen"
	grid[3][2][10] = "torgau-delitzsch/Roitzsch"
	grid[3][4][10] = "torgau-delitzsch/Neiden"
	grid[3][5][10] = "torgau-delitzsch/Döbern bei Neiden"
	grid[3][6][10] = "torgau-delitzsch/Zwethau"
	grid[3][8][10] = "bad-liebenwerda/Döbrichau-Löhsten"

	grid[3][5][11] = "torgau-delitzsch/Welsau"
	grid[3][6][11] = "torgau-delitzsch/Kreischau-Eulenau"
	grid[3][7][11] = "torgau-delitzsch/Beilrode"

	grid[3][4][12] = "torgau-delitzsch/Zinna"
	grid[3][5][12] = "torgau-delitzsch/Torgau St. Marien"

	grid[3][0][13] = "torgau-delitzsch/Wöllnau"
	grid[3][2][13] = "torgau-delitzsch/Weidenhain"
	grid[3][3][13] = "torgau-delitzsch/Großwig"
	grid[3][4][13] = "torgau-delitzsch/Süptitz"
	grid[3][5][13] = "torgau-delitzsch/Torgau Garnisonsgemeinde"
	grid[3][6][13] = "torgau-delitzsch/Torgau Schmerzhafte Mutter"

	grid[3][0][14] = "torgau-delitzsch/Battaune"
	grid[3][1][14] = "torgau-delitzsch/Wildenhain"
	grid[3][4][14] = "torgau-delitzsch/Melpitz"
	grid[3][6][14] = "torgau-delitzsch/Loßwig"
	grid[3][7][14] = "torgau-delitzsch/Triestewitz"
	grid[3][8][14] = "torgau-delitzsch/Arzberg"

	grid[3][0][15] = "torgau-delitzsch/Doberschütz"
	grid[3][1][15] = "torgau-delitzsch/Mockrehna"
	grid[3][2][15] = "torgau-delitzsch/Niederaudenhain"
	grid[3][3][15] = "torgau-delitzsch/Klitzschen"
	grid[3][7][15] = "torgau-delitzsch/Weßnig"
	grid[3][9][15] = "torgau-delitzsch/Blumberg"

	grid[3][0][16] = "torgau-delitzsch/Sprotta"
	grid[3][1][16] = "torgau-delitzsch/Oberaudenhain"
	grid[3][2][16] = "torgau-delitzsch/Audenhain"
	grid[3][3][16] = "torgau-delitzsch/Langenreichenbach"
	grid[3][4][16] = "torgau-delitzsch/Staupitz"
	grid[3][5][16] = "torgau-delitzsch/Beckwitz"

	grid[3][1][17] = "torgau-delitzsch/Strelln"
	grid[3][2][17] = "torgau-delitzsch/Wildschütz"
	grid[3][6][17] = "torgau-delitzsch/Taura"
	grid[3][8][17] = "torgau-delitzsch/Belgern"

	grid[3][0][18] = "torgau-delitzsch/Paschwitz"
	grid[3][1][18] = "torgau-delitzsch/Schöna"
	grid[3][2][18] = "torgau-delitzsch/Kobershain"
	grid[3][4][18] = "torgau-delitzsch/Schildau"
	grid[3][5][18] = "torgau-delitzsch/Sitzenroda"
	grid[3][7][18] = "torgau-delitzsch/Lausa"
	grid[3][8][18] = "torgau-delitzsch/Neußen"
	grid[3][9][18] = "torgau-delitzsch/Staritz"

	grid[3][9][19] = "torgau-delitzsch/Schirmenitz"

	grid[3][9][20] = "torgau-delitzsch/Paußnitz"

	//Tab 4
	grid[4][0][0] = "meissen/Bloßwitz"
	grid[4][1][0] = "meissen/Staucha"
	grid[4][2][0] = "meissen/Striegnitz"
	grid[4][3][0] = "meissen/Dörschnitz"
	grid[4][5][0] = "meissen/Großdobritz"
	grid[4][8][0] = "dresden/Großdittmannsdorf"

	grid[4][1][1] = "meissen/Lommatzsch"
	grid[4][2][1] = "meissen/Zehren"
	grid[4][3][1] = "meissen/Zadel"
	grid[4][4][1] = "meissen/Gröbern"
	grid[4][5][1] = "meissen/Oberau"
	grid[4][8][1] = "dresden/Medingen"
	grid[4][9][1] = "dresden/Ottendorf"

	grid[4][0][2] = "meissen/Neckanitz"
	grid[4][2][2] = "meissen/Meißen St. Afra"
	grid[4][3][2] = "meissen/Meißen Trinitatiskirche"
	grid[4][4][2] = "meissen/Niederau"
	grid[4][7][2] = "dresden/Grünberg"
	grid[4][8][2] = "dresden/Seifersdorf"
	grid[4][9][2] = "dresden/Wachau"

	grid[4][0][3] = "meissen/Leuben"
	grid[4][1][3] = "meissen/Planitz"
	grid[4][2][3] = "meissen/Meißen"
	grid[4][3][3] = "meissen/Meißen Frauenkirche"
	grid[4][4][3] = "meissen/Meißen Johanneskirche"
	grid[4][5][3] = "meissen/Weinböhla"
	grid[4][8][3] = "dresden/Lausa"
	grid[4][9][3] = "dresden/Schönborn"

	grid[4][1][4] = "meissen/Ziegenhain"
	grid[4][4][4] = "meissen/Brockwitz"
	grid[4][5][4] = "meissen/Coswig"
	grid[4][6][4] = "dresden/Reichenberg"
	grid[4][7][4] = "dresden/Wilschdorf"
	grid[4][8][4] = "dresden/Radeberg"
	grid[4][9][4] = "dresden/Langebrück"

	grid[4][2][5] = "meissen/Krögis"
	grid[4][3][5] = "meissen/Miltitz"
	grid[4][5][5] = "meissen/Naustadt"
	grid[4][6][5] = "dresden/Kötzschenbroda"
	grid[4][7][5] = "dresden/Klotzsche"
	grid[4][8][5] = "dresden/Dresden Neustadt"
	grid[4][9][5] = "dresden/Kleinwolmsdorf"

	grid[4][0][6] = "meissen/Rüsseina"
	grid[4][1][6] = "meissen/Raußlitz"
	grid[4][2][6] = "meissen/Heynitz"
	grid[4][3][6] = "meissen/Taubenheim"
	grid[4][4][6] = "meissen/Röhrsdorf"
	grid[4][5][6] = "meissen/Constappel"
	grid[4][6][6] = "dresden/Dresden Friedrichstadt St. Michael"
	grid[4][7][6] = "dresden/Kaditz"
	grid[4][8][6] = "dresden/Dresden Dreikönigskirche"
	grid[4][9][6] = "dresden/Großerkmannsdorf"

	grid[4][0][7] = "meissen/Wendischbora"
	grid[4][1][7] = "meissen/Rothschönberg"
	grid[4][2][7] = "meissen/Burkhardswalde"
	grid[4][3][7] = "meissen/Sora"
	grid[4][4][7] = "meissen/Weistropp"
	grid[4][5][7] = "dresden/Briesnitz"
	grid[4][6][7] = "dresden/Dresden Friedrichstadt"
	grid[4][7][7] = "dresden/Dresden Sophienkirche"
	grid[4][8][7] = "dresden/Dresden Hofkirche"
	grid[4][9][7] = "dresden/Weißig"

	grid[4][0][8] = "meissen/Nossen"
	grid[4][1][8] = "meissen/Deutschenbora"
	grid[4][2][8] = "meissen/Tanneberg"
	grid[4][3][8] = "meissen/Limbach"
	grid[4][4][8] = "meissen/Wilsdruff"
	grid[4][5][8] = "meissen/Unkersdorf"
	grid[4][6][8] = "dresden/Dresden Kreuzkirche"
	grid[4][7][8] = "dresden/Dresden Frauenkirche"
	grid[4][8][8] = "dresden/Loschwitz"

	grid[4][0][9] = "meissen/Siebenlehn"
	grid[4][1][9] = "meissen/Hirschfeld"
	grid[4][2][9] = "meissen/Neukirchen"
	grid[4][3][9] = "meissen/Blankenstein"
	grid[4][4][9] = "meissen/Grumbach"
	grid[4][5][9] = "meissen/Kesselsdorf"
	grid[4][6][9] = "dippoldiswalde/Pesterwitz"
	grid[4][7][9] = "dresden/Dresden Annenkirche"
	grid[4][8][9] = "dresden/Dresden Böhmische Exulantengemeinde"
	grid[4][9][9] = "dresden/Schönfeld"

	grid[4][0][10] = "meissen/Obergruna"
	grid[4][1][10] = "meissen/Bieberstein"
	grid[4][2][10] = "meissen/Reinsberg"
	grid[4][3][10] = "meissen/Dittmannsdorf"
	grid[4][4][10] = "meissen/Herzogswalde"
	grid[4][5][10] = "dippoldiswalde/Döhlen"
	grid[4][6][10] = "dresden/Plauen"
	grid[4][7][10] = "dresden/Leubnitz"
	grid[4][8][10] = "dresden/Leuben"
	grid[4][9][10] = "dresden/Hosterwitz"

	grid[4][0][11] = "freiberg/Großschirma"
	grid[4][1][11] = "freiberg/Krummenhennersdorf"
	grid[4][2][11] = "freiberg/Niederschöna"
	grid[4][3][11] = "meissen/Mohorn"
	grid[4][4][11] = "dippoldiswalde/Fördergersdorf"
	grid[4][5][11] = "dippoldiswalde/Tharandt"
	grid[4][6][11] = "dippoldiswalde/Deuben"
	grid[4][8][11] = "dresden/Lockwitz"

	grid[4][0][12] = "freiberg/Langhennersdorf"
	grid[4][1][12] = "freiberg/Tuttendorf"
	grid[4][2][12] = "freiberg/Conradsdorf"
	grid[4][3][12] = "freiberg/Naundorf"
	grid[4][4][12] = "dippoldiswalde/Dorfhain"
	grid[4][5][12] = "dippoldiswalde/Somsdorf"
	grid[4][6][12] = "dippoldiswalde/Rabenau"
	grid[4][7][12] = "dippoldiswalde/Possendorf"
	grid[4][8][12] = "dresden/Röhrsdorf"

	grid[4][0][13] = "freiberg/Bräunsdorf"
	grid[4][1][13] = "freiberg/Freiberg Dom St. Marien"
	grid[4][2][13] = "freiberg/Freiberg St. Nikolai"
	grid[4][4][13] = "dippoldiswalde/Klingenberg"
	grid[4][5][13] = "dippoldiswalde/Höckendorf"
	grid[4][6][13] = "dippoldiswalde/Seifersdorf"
	grid[4][8][13] = "dippoldiswalde/Kreischa"

	grid[4][0][14] = "freiberg/Kleinwaltersdorf"
	grid[4][1][14] = "freiberg/Freiberg St. Petri"
	grid[4][2][14] = "freiberg/Freiberg St. Jacobi"
	grid[4][3][14] = "freiberg/Hilbersdorf"
	grid[4][4][14] = "dippoldiswalde/Colmnitz"
	grid[4][5][14] = "dippoldiswalde/Ruppendorf"
	grid[4][6][14] = "dippoldiswalde/Dippoldiswalde"
	grid[4][7][14] = "dippoldiswalde/Reinhardtsgrimma"

	grid[4][0][15] = "freiberg/Kleinschirma"
	grid[4][1][15] = "freiberg/Freiberg St. Johannis"
	grid[4][2][15] = "freiberg/Freiberg"
	grid[4][3][15] = "freiberg/Niederbobritzsch"
	grid[4][6][15] = "dippoldiswalde/Reichstädt"

	grid[4][0][16] = "freiberg/Oberschöna"
	grid[4][1][16] = "freiberg/Erbisdorf"
	grid[4][2][16] = "freiberg/Berthelsdorf"
	grid[4][3][16] = "freiberg/Weißenborn"
	grid[4][4][16] = "freiberg/Oberbobritzsch"
	grid[4][5][16] = "dippoldiswalde/Pretzschendorf"
	grid[4][8][16] = "dippoldiswalde/Glashütte"

	grid[4][0][17] = "freiberg/Langenau"
	grid[4][1][17] = "freiberg/Weigmannsdorf"
	grid[4][2][17] = "freiberg/Lichtenberg"
	grid[4][3][17] = "dippoldiswalde/Burkersdorf"
	grid[4][4][17] = "dippoldiswalde/Hartmannsdorf"
	grid[4][5][17] = "dippoldiswalde/Hennersdorf"
	grid[4][6][17] = "dippoldiswalde/Sadisdorf"
	grid[4][7][17] = "dippoldiswalde/Schmiedeberg"
	grid[4][8][17] = "dippoldiswalde/Johnsbach"
	grid[4][9][17] = "dippoldiswalde/Dittersdorf"

	grid[4][0][18] = "freiberg/Gränitz"
	grid[4][1][18] = "freiberg/Großhartmannsdorf"
	grid[4][2][18] = "freiberg/Helbigsdorf"
	grid[4][3][18] = "freiberg/Mulda"
	grid[4][4][18] = "dippoldiswalde/Frauenstein"
	grid[4][5][18] = "dippoldiswalde/Schönfeld"
	grid[4][7][18] = "dippoldiswalde/Bärenstein"
	grid[4][8][18] = "dippoldiswalde/Lauenstein"
	grid[4][9][18] = "dippoldiswalde/Liebenau"

	grid[4][1][19] = "freiberg/Zethau"
	grid[4][2][19] = "freiberg/Dorfchemnitz"
	grid[4][3][19] = "dippoldiswalde/Dittersbach"
	grid[4][4][19] = "dippoldiswalde/Nassau"
	grid[4][5][19] = "dippoldiswalde/Hermsdorf"
	grid[4][6][19] = "dippoldiswalde/Schellerhau"
	grid[4][7][19] = "dippoldiswalde/Altenberg"
	grid[4][8][19] = "dippoldiswalde/Geising"
	grid[4][9][19] = "dippoldiswalde/Fürstenwalde"

	grid[4][1][20] = "freiberg/Voigtsdorf"
	grid[4][2][20] = "freiberg/Sayda"
	grid[4][3][20] = "freiberg/Clausnitz"
	grid[4][4][20] = "freiberg/Cämmerswalde"
	grid[4][8][20] = "dippoldiswalde/Fürstenau"

	//Tab 5
	grid[5][2][8] = "bautzen/Bautzen"
	grid[5][3][8] = "bautzen/Bautzen Dom"

	grid[5][2][9] = "bautzen/Bautzen St. Michael"
	grid[5][3][9] = "bautzen/Bautzen St. Petri"
}
